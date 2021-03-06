/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package loadbalancers

import (
	"k8s.io/ingress-gce/pkg/composite"
	"k8s.io/ingress-gce/pkg/flags"
	"k8s.io/ingress-gce/pkg/translator"
	"k8s.io/ingress-gce/pkg/utils"
	"k8s.io/ingress-gce/pkg/utils/namer"
	"k8s.io/klog"
)

const (
	// Every target https proxy accepts upto 10 ssl certificates.
	TargetProxyCertLimit = 10
)

// checkProxy ensures the correct TargetHttpProxy for a loadbalancer
func (l *L7) checkProxy() (err error) {
	isL7ILB := flags.F.EnableL7Ilb && utils.IsGCEL7ILBIngress(l.runtimeInfo.Ingress)
	tr := translator.NewTranslator(isL7ILB, l.namer)

	description, err := l.description()
	if err != nil {
		return err
	}
	urlMapKey, err := l.CreateKey(l.um.Name)
	if err != nil {
		return err
	}

	version := l.Versions().TargetHttpProxy
	proxy := tr.ToCompositeTargetHttpProxy(description, version, urlMapKey)

	key, err := l.CreateKey(proxy.Name)
	if err != nil {
		return err
	}

	currentProxy, _ := composite.GetTargetHttpProxy(l.cloud, key, version)
	if currentProxy == nil {
		klog.V(3).Infof("Creating new http proxy for urlmap %v", l.um.Name)

		key, err := l.CreateKey(proxy.Name)
		if err != nil {
			return err
		}
		if err = composite.CreateTargetHttpProxy(l.cloud, key, proxy); err != nil {
			return err
		}
		key, err = l.CreateKey(proxy.Name)
		if err != nil {
			return err
		}
		currentProxy, err = composite.GetTargetHttpProxy(l.cloud, key, version)
		if err != nil {
			return err
		}
		l.tp = currentProxy
		return nil
	}
	if !utils.EqualResourcePaths(currentProxy.UrlMap, proxy.UrlMap) {
		klog.V(3).Infof("Proxy %v has the wrong url map, setting %v overwriting %v",
			currentProxy.Name, proxy.UrlMap, currentProxy.UrlMap)
		key, err := l.CreateKey(currentProxy.Name)
		if err != nil {
			return err
		}
		if err := composite.SetUrlMapForTargetHttpProxy(l.cloud, key, currentProxy, proxy.UrlMap); err != nil {
			return err
		}
	}
	l.tp = currentProxy
	return nil
}

func (l *L7) checkHttpsProxy() (err error) {
	isL7ILB := flags.F.EnableL7Ilb && utils.IsGCEL7ILBIngress(l.runtimeInfo.Ingress)
	tr := translator.NewTranslator(isL7ILB, l.namer)
	env := &translator.Env{FrontendConfig: l.runtimeInfo.FrontendConfig}

	if len(l.sslCerts) == 0 {
		klog.V(3).Infof("No SSL certificates for %q, will not create HTTPS Proxy.", l)
		return nil
	}

	urlMapKey, err := l.CreateKey(l.um.Name)
	if err != nil {
		return err
	}
	description, err := l.description()
	version := l.Versions().TargetHttpProxy
	proxy, sslPolicySet, err := tr.ToCompositeTargetHttpsProxy(env, description, version, urlMapKey, l.sslCerts)
	if err != nil {
		return err
	}

	key, err := l.CreateKey(proxy.Name)
	if err != nil {
		return err
	}

	currentProxy, _ := composite.GetTargetHttpsProxy(l.cloud, key, version)
	if err != nil {
		return err
	}

	if currentProxy == nil {
		klog.V(3).Infof("Creating new https Proxy for urlmap %q", l.um.Name)

		if err = composite.CreateTargetHttpsProxy(l.cloud, key, proxy); err != nil {
			return err
		}

		key, err = l.CreateKey(proxy.Name)
		if err != nil {
			return err
		}
		currentProxy, err = composite.GetTargetHttpsProxy(l.cloud, key, version)
		if err != nil {
			return err
		}

		l.tps = currentProxy
		return nil
	}
	if !utils.EqualResourcePaths(currentProxy.UrlMap, proxy.UrlMap) {
		klog.V(3).Infof("Https Proxy %v has the wrong url map, setting %v overwriting %v",
			currentProxy.Name, proxy.UrlMap, currentProxy.UrlMap)
		key, err := l.CreateKey(currentProxy.Name)
		if err != nil {
			return err
		}
		if err := composite.SetUrlMapForTargetHttpsProxy(l.cloud, key, currentProxy, proxy.UrlMap); err != nil {
			return err
		}
	}

	if !l.compareCerts(currentProxy.SslCertificates) {
		klog.V(3).Infof("Https Proxy %q has the wrong ssl certs, setting %v overwriting %v",
			currentProxy.Name, toCertNames(l.sslCerts), currentProxy.SslCertificates)
		var sslCertURLs []string
		for _, cert := range l.sslCerts {
			sslCertURLs = append(sslCertURLs, cert.SelfLink)
		}
		key, err := l.CreateKey(currentProxy.Name)
		if err != nil {
			return err
		}
		if err := composite.SetSslCertificateForTargetHttpsProxy(l.cloud, key, currentProxy, sslCertURLs); err != nil {
			return err
		}
	}

	if flags.F.EnableFrontendConfig && sslPolicySet {
		if err := l.ensureSslPolicy(env, currentProxy, proxy.SslPolicy); err != nil {
			return err
		}
	}

	l.tps = currentProxy
	return nil
}

func (l *L7) getSslCertLinkInUse() ([]string, error) {
	proxyName := l.namer.TargetProxy(namer.HTTPSProtocol)
	key, err := l.CreateKey(proxyName)
	if err != nil {
		return nil, err
	}
	proxy, err := composite.GetTargetHttpsProxy(l.cloud, key, l.Versions().TargetHttpsProxy)
	if err != nil {
		return nil, err
	}

	return proxy.SslCertificates, nil
}

// ensureSslPolicy ensures that the SslPolicy described in the frontendconfig is
// properly applied to the proxy.
func (l *L7) ensureSslPolicy(env *translator.Env, currentProxy *composite.TargetHttpsProxy, policyLink string) error {
	if !utils.EqualResourceIDs(policyLink, currentProxy.SslPolicy) {
		key, err := l.CreateKey(currentProxy.Name)
		if err != nil {
			return err
		}
		if err := composite.SetSslPolicyForTargetHttpsProxy(l.cloud, key, currentProxy, policyLink); err != nil {
			return err
		}
	}
	return nil
}
