/*
Copyright 2020 The Kubernetes Authors.

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
// This file was generated by "./hack/update-codegen.sh". Do not edit directly.
// directly.

package composite

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kr/pretty"
	computealpha "google.golang.org/api/compute/v0.alpha"
	computebeta "google.golang.org/api/compute/v0.beta"
	compute "google.golang.org/api/compute/v1"
)

func TestAuthenticationPolicy(t *testing.T) {
	compositeType := reflect.TypeOf(AuthenticationPolicy{})
	alphaType := reflect.TypeOf(computealpha.AuthenticationPolicy{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestAuthorizationConfig(t *testing.T) {
	compositeType := reflect.TypeOf(AuthorizationConfig{})
	alphaType := reflect.TypeOf(computealpha.AuthorizationConfig{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestBackend(t *testing.T) {
	compositeType := reflect.TypeOf(Backend{})
	alphaType := reflect.TypeOf(computealpha.Backend{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}
func TestBackendService(t *testing.T) {
	// Use reflection to verify that our composite type contains all the
	// same fields as the alpha type.
	compositeType := reflect.TypeOf(BackendService{})
	alphaType := reflect.TypeOf(computealpha.BackendService{})
	betaType := reflect.TypeOf(computebeta.BackendService{})
	gaType := reflect.TypeOf(compute.BackendService{})

	// For the composite type, remove the Version field from consideration
	compositeTypeNumFields := compositeType.NumField() - 2
	if compositeTypeNumFields != alphaType.NumField() {
		t.Fatalf("%v should contain %v fields. Got %v", alphaType.Name(), alphaType.NumField(), compositeTypeNumFields)
	}

	// Compare all the fields by doing a lookup since we can't guarantee that they'll be in the same order
	// Make sure that composite type is strictly alpha fields + internal bookkeeping
	for i := 2; i < compositeType.NumField(); i++ {
		lookupField, found := alphaType.FieldByName(compositeType.Field(i).Name)
		if !found {
			t.Fatal(fmt.Errorf("Field %v not present in alpha type %v", compositeType.Field(i), alphaType))
		}
		if err := compareFields(compositeType.Field(i), lookupField); err != nil {
			t.Fatal(err)
		}
	}

	// Verify that all beta fields are in composite type
	if err := typeEquality(betaType, compositeType, false); err != nil {
		t.Fatal(err)
	}

	// Verify that all GA fields are in composite type
	if err := typeEquality(gaType, compositeType, false); err != nil {
		t.Fatal(err)
	}
}

func TestToBackendService(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected *BackendService
	}{
		{
			computealpha.BackendService{},
			&BackendService{},
		},
		{
			computebeta.BackendService{},
			&BackendService{},
		},
		{
			compute.BackendService{},
			&BackendService{},
		},
	}
	for _, testCase := range testCases {
		result, _ := ToBackendService(testCase.input)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Fatalf("ToBackendService(input) = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(testCase.input), pretty.Sprint(result), pretty.Sprint(testCase.expected))
		}
	}
}

func TestBackendServiceToAlpha(t *testing.T) {
	composite := BackendService{}
	expected := &computealpha.BackendService{}
	result, err := composite.ToAlpha()
	if err != nil {
		t.Fatalf("BackendService.ToAlpha() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("BackendService.ToAlpha() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}
func TestBackendServiceToBeta(t *testing.T) {
	composite := BackendService{}
	expected := &computebeta.BackendService{}
	result, err := composite.ToBeta()
	if err != nil {
		t.Fatalf("BackendService.ToBeta() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("BackendService.ToBeta() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}
func TestBackendServiceToGA(t *testing.T) {
	composite := BackendService{}
	expected := &compute.BackendService{}
	result, err := composite.ToGA()
	if err != nil {
		t.Fatalf("BackendService.ToGA() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("BackendService.ToGA() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}

func TestBackendServiceCdnPolicy(t *testing.T) {
	compositeType := reflect.TypeOf(BackendServiceCdnPolicy{})
	alphaType := reflect.TypeOf(computealpha.BackendServiceCdnPolicy{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestBackendServiceFailoverPolicy(t *testing.T) {
	compositeType := reflect.TypeOf(BackendServiceFailoverPolicy{})
	alphaType := reflect.TypeOf(computealpha.BackendServiceFailoverPolicy{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestBackendServiceIAP(t *testing.T) {
	compositeType := reflect.TypeOf(BackendServiceIAP{})
	alphaType := reflect.TypeOf(computealpha.BackendServiceIAP{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestBackendServiceIAPOAuth2ClientInfo(t *testing.T) {
	compositeType := reflect.TypeOf(BackendServiceIAPOAuth2ClientInfo{})
	alphaType := reflect.TypeOf(computealpha.BackendServiceIAPOAuth2ClientInfo{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestBackendServiceLogConfig(t *testing.T) {
	compositeType := reflect.TypeOf(BackendServiceLogConfig{})
	alphaType := reflect.TypeOf(computealpha.BackendServiceLogConfig{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestCacheKeyPolicy(t *testing.T) {
	compositeType := reflect.TypeOf(CacheKeyPolicy{})
	alphaType := reflect.TypeOf(computealpha.CacheKeyPolicy{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestCallCredentials(t *testing.T) {
	compositeType := reflect.TypeOf(CallCredentials{})
	alphaType := reflect.TypeOf(computealpha.CallCredentials{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestChannelCredentials(t *testing.T) {
	compositeType := reflect.TypeOf(ChannelCredentials{})
	alphaType := reflect.TypeOf(computealpha.ChannelCredentials{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestCircuitBreakers(t *testing.T) {
	compositeType := reflect.TypeOf(CircuitBreakers{})
	alphaType := reflect.TypeOf(computealpha.CircuitBreakers{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestClientTlsSettings(t *testing.T) {
	compositeType := reflect.TypeOf(ClientTlsSettings{})
	alphaType := reflect.TypeOf(computealpha.ClientTlsSettings{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestConnectionDraining(t *testing.T) {
	compositeType := reflect.TypeOf(ConnectionDraining{})
	alphaType := reflect.TypeOf(computealpha.ConnectionDraining{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestConsistentHashLoadBalancerSettings(t *testing.T) {
	compositeType := reflect.TypeOf(ConsistentHashLoadBalancerSettings{})
	alphaType := reflect.TypeOf(computealpha.ConsistentHashLoadBalancerSettings{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestConsistentHashLoadBalancerSettingsHttpCookie(t *testing.T) {
	compositeType := reflect.TypeOf(ConsistentHashLoadBalancerSettingsHttpCookie{})
	alphaType := reflect.TypeOf(computealpha.ConsistentHashLoadBalancerSettingsHttpCookie{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestCorsPolicy(t *testing.T) {
	compositeType := reflect.TypeOf(CorsPolicy{})
	alphaType := reflect.TypeOf(computealpha.CorsPolicy{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestDuration(t *testing.T) {
	compositeType := reflect.TypeOf(Duration{})
	alphaType := reflect.TypeOf(computealpha.Duration{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}
func TestForwardingRule(t *testing.T) {
	// Use reflection to verify that our composite type contains all the
	// same fields as the alpha type.
	compositeType := reflect.TypeOf(ForwardingRule{})
	alphaType := reflect.TypeOf(computealpha.ForwardingRule{})
	betaType := reflect.TypeOf(computebeta.ForwardingRule{})
	gaType := reflect.TypeOf(compute.ForwardingRule{})

	// For the composite type, remove the Version field from consideration
	compositeTypeNumFields := compositeType.NumField() - 2
	if compositeTypeNumFields != alphaType.NumField() {
		t.Fatalf("%v should contain %v fields. Got %v", alphaType.Name(), alphaType.NumField(), compositeTypeNumFields)
	}

	// Compare all the fields by doing a lookup since we can't guarantee that they'll be in the same order
	// Make sure that composite type is strictly alpha fields + internal bookkeeping
	for i := 2; i < compositeType.NumField(); i++ {
		lookupField, found := alphaType.FieldByName(compositeType.Field(i).Name)
		if !found {
			t.Fatal(fmt.Errorf("Field %v not present in alpha type %v", compositeType.Field(i), alphaType))
		}
		if err := compareFields(compositeType.Field(i), lookupField); err != nil {
			t.Fatal(err)
		}
	}

	// Verify that all beta fields are in composite type
	if err := typeEquality(betaType, compositeType, false); err != nil {
		t.Fatal(err)
	}

	// Verify that all GA fields are in composite type
	if err := typeEquality(gaType, compositeType, false); err != nil {
		t.Fatal(err)
	}
}

func TestToForwardingRule(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected *ForwardingRule
	}{
		{
			computealpha.ForwardingRule{},
			&ForwardingRule{},
		},
		{
			computebeta.ForwardingRule{},
			&ForwardingRule{},
		},
		{
			compute.ForwardingRule{},
			&ForwardingRule{},
		},
	}
	for _, testCase := range testCases {
		result, _ := ToForwardingRule(testCase.input)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Fatalf("ToForwardingRule(input) = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(testCase.input), pretty.Sprint(result), pretty.Sprint(testCase.expected))
		}
	}
}

func TestForwardingRuleToAlpha(t *testing.T) {
	composite := ForwardingRule{}
	expected := &computealpha.ForwardingRule{}
	result, err := composite.ToAlpha()
	if err != nil {
		t.Fatalf("ForwardingRule.ToAlpha() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("ForwardingRule.ToAlpha() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}
func TestForwardingRuleToBeta(t *testing.T) {
	composite := ForwardingRule{}
	expected := &computebeta.ForwardingRule{}
	result, err := composite.ToBeta()
	if err != nil {
		t.Fatalf("ForwardingRule.ToBeta() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("ForwardingRule.ToBeta() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}
func TestForwardingRuleToGA(t *testing.T) {
	composite := ForwardingRule{}
	expected := &compute.ForwardingRule{}
	result, err := composite.ToGA()
	if err != nil {
		t.Fatalf("ForwardingRule.ToGA() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("ForwardingRule.ToGA() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}

func TestGrpcServiceConfig(t *testing.T) {
	compositeType := reflect.TypeOf(GrpcServiceConfig{})
	alphaType := reflect.TypeOf(computealpha.GrpcServiceConfig{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestHTTP2HealthCheck(t *testing.T) {
	compositeType := reflect.TypeOf(HTTP2HealthCheck{})
	alphaType := reflect.TypeOf(computealpha.HTTP2HealthCheck{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestHTTPHealthCheck(t *testing.T) {
	compositeType := reflect.TypeOf(HTTPHealthCheck{})
	alphaType := reflect.TypeOf(computealpha.HTTPHealthCheck{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestHTTPSHealthCheck(t *testing.T) {
	compositeType := reflect.TypeOf(HTTPSHealthCheck{})
	alphaType := reflect.TypeOf(computealpha.HTTPSHealthCheck{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}
func TestHealthCheck(t *testing.T) {
	// Use reflection to verify that our composite type contains all the
	// same fields as the alpha type.
	compositeType := reflect.TypeOf(HealthCheck{})
	alphaType := reflect.TypeOf(computealpha.HealthCheck{})
	betaType := reflect.TypeOf(computebeta.HealthCheck{})
	gaType := reflect.TypeOf(compute.HealthCheck{})

	// For the composite type, remove the Version field from consideration
	compositeTypeNumFields := compositeType.NumField() - 2
	if compositeTypeNumFields != alphaType.NumField() {
		t.Fatalf("%v should contain %v fields. Got %v", alphaType.Name(), alphaType.NumField(), compositeTypeNumFields)
	}

	// Compare all the fields by doing a lookup since we can't guarantee that they'll be in the same order
	// Make sure that composite type is strictly alpha fields + internal bookkeeping
	for i := 2; i < compositeType.NumField(); i++ {
		lookupField, found := alphaType.FieldByName(compositeType.Field(i).Name)
		if !found {
			t.Fatal(fmt.Errorf("Field %v not present in alpha type %v", compositeType.Field(i), alphaType))
		}
		if err := compareFields(compositeType.Field(i), lookupField); err != nil {
			t.Fatal(err)
		}
	}

	// Verify that all beta fields are in composite type
	if err := typeEquality(betaType, compositeType, false); err != nil {
		t.Fatal(err)
	}

	// Verify that all GA fields are in composite type
	if err := typeEquality(gaType, compositeType, false); err != nil {
		t.Fatal(err)
	}
}

func TestToHealthCheck(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected *HealthCheck
	}{
		{
			computealpha.HealthCheck{},
			&HealthCheck{},
		},
		{
			computebeta.HealthCheck{},
			&HealthCheck{},
		},
		{
			compute.HealthCheck{},
			&HealthCheck{},
		},
	}
	for _, testCase := range testCases {
		result, _ := ToHealthCheck(testCase.input)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Fatalf("ToHealthCheck(input) = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(testCase.input), pretty.Sprint(result), pretty.Sprint(testCase.expected))
		}
	}
}

func TestHealthCheckToAlpha(t *testing.T) {
	composite := HealthCheck{}
	expected := &computealpha.HealthCheck{}
	result, err := composite.ToAlpha()
	if err != nil {
		t.Fatalf("HealthCheck.ToAlpha() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("HealthCheck.ToAlpha() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}
func TestHealthCheckToBeta(t *testing.T) {
	composite := HealthCheck{}
	expected := &computebeta.HealthCheck{}
	result, err := composite.ToBeta()
	if err != nil {
		t.Fatalf("HealthCheck.ToBeta() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("HealthCheck.ToBeta() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}
func TestHealthCheckToGA(t *testing.T) {
	composite := HealthCheck{}
	expected := &compute.HealthCheck{}
	result, err := composite.ToGA()
	if err != nil {
		t.Fatalf("HealthCheck.ToGA() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("HealthCheck.ToGA() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}

func TestHealthCheckLogConfig(t *testing.T) {
	compositeType := reflect.TypeOf(HealthCheckLogConfig{})
	alphaType := reflect.TypeOf(computealpha.HealthCheckLogConfig{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestHostRule(t *testing.T) {
	compositeType := reflect.TypeOf(HostRule{})
	alphaType := reflect.TypeOf(computealpha.HostRule{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestHttpFaultAbort(t *testing.T) {
	compositeType := reflect.TypeOf(HttpFaultAbort{})
	alphaType := reflect.TypeOf(computealpha.HttpFaultAbort{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestHttpFaultDelay(t *testing.T) {
	compositeType := reflect.TypeOf(HttpFaultDelay{})
	alphaType := reflect.TypeOf(computealpha.HttpFaultDelay{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestHttpFaultInjection(t *testing.T) {
	compositeType := reflect.TypeOf(HttpFaultInjection{})
	alphaType := reflect.TypeOf(computealpha.HttpFaultInjection{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestHttpHeaderAction(t *testing.T) {
	compositeType := reflect.TypeOf(HttpHeaderAction{})
	alphaType := reflect.TypeOf(computealpha.HttpHeaderAction{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestHttpHeaderMatch(t *testing.T) {
	compositeType := reflect.TypeOf(HttpHeaderMatch{})
	alphaType := reflect.TypeOf(computealpha.HttpHeaderMatch{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestHttpHeaderOption(t *testing.T) {
	compositeType := reflect.TypeOf(HttpHeaderOption{})
	alphaType := reflect.TypeOf(computealpha.HttpHeaderOption{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestHttpQueryParameterMatch(t *testing.T) {
	compositeType := reflect.TypeOf(HttpQueryParameterMatch{})
	alphaType := reflect.TypeOf(computealpha.HttpQueryParameterMatch{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestHttpRedirectAction(t *testing.T) {
	compositeType := reflect.TypeOf(HttpRedirectAction{})
	alphaType := reflect.TypeOf(computealpha.HttpRedirectAction{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestHttpRetryPolicy(t *testing.T) {
	compositeType := reflect.TypeOf(HttpRetryPolicy{})
	alphaType := reflect.TypeOf(computealpha.HttpRetryPolicy{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestHttpRouteAction(t *testing.T) {
	compositeType := reflect.TypeOf(HttpRouteAction{})
	alphaType := reflect.TypeOf(computealpha.HttpRouteAction{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestHttpRouteRule(t *testing.T) {
	compositeType := reflect.TypeOf(HttpRouteRule{})
	alphaType := reflect.TypeOf(computealpha.HttpRouteRule{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestHttpRouteRuleMatch(t *testing.T) {
	compositeType := reflect.TypeOf(HttpRouteRuleMatch{})
	alphaType := reflect.TypeOf(computealpha.HttpRouteRuleMatch{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestInt64RangeMatch(t *testing.T) {
	compositeType := reflect.TypeOf(Int64RangeMatch{})
	alphaType := reflect.TypeOf(computealpha.Int64RangeMatch{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestJwt(t *testing.T) {
	compositeType := reflect.TypeOf(Jwt{})
	alphaType := reflect.TypeOf(computealpha.Jwt{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestJwtHeader(t *testing.T) {
	compositeType := reflect.TypeOf(JwtHeader{})
	alphaType := reflect.TypeOf(computealpha.JwtHeader{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestMetadataCredentialsFromPlugin(t *testing.T) {
	compositeType := reflect.TypeOf(MetadataCredentialsFromPlugin{})
	alphaType := reflect.TypeOf(computealpha.MetadataCredentialsFromPlugin{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestMetadataFilter(t *testing.T) {
	compositeType := reflect.TypeOf(MetadataFilter{})
	alphaType := reflect.TypeOf(computealpha.MetadataFilter{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestMetadataFilterLabelMatch(t *testing.T) {
	compositeType := reflect.TypeOf(MetadataFilterLabelMatch{})
	alphaType := reflect.TypeOf(computealpha.MetadataFilterLabelMatch{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestMutualTls(t *testing.T) {
	compositeType := reflect.TypeOf(MutualTls{})
	alphaType := reflect.TypeOf(computealpha.MutualTls{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestOriginAuthenticationMethod(t *testing.T) {
	compositeType := reflect.TypeOf(OriginAuthenticationMethod{})
	alphaType := reflect.TypeOf(computealpha.OriginAuthenticationMethod{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestOutlierDetection(t *testing.T) {
	compositeType := reflect.TypeOf(OutlierDetection{})
	alphaType := reflect.TypeOf(computealpha.OutlierDetection{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestPathMatcher(t *testing.T) {
	compositeType := reflect.TypeOf(PathMatcher{})
	alphaType := reflect.TypeOf(computealpha.PathMatcher{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestPathRule(t *testing.T) {
	compositeType := reflect.TypeOf(PathRule{})
	alphaType := reflect.TypeOf(computealpha.PathRule{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestPeerAuthenticationMethod(t *testing.T) {
	compositeType := reflect.TypeOf(PeerAuthenticationMethod{})
	alphaType := reflect.TypeOf(computealpha.PeerAuthenticationMethod{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestPermission(t *testing.T) {
	compositeType := reflect.TypeOf(Permission{})
	alphaType := reflect.TypeOf(computealpha.Permission{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestPermissionConstraint(t *testing.T) {
	compositeType := reflect.TypeOf(PermissionConstraint{})
	alphaType := reflect.TypeOf(computealpha.PermissionConstraint{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestPrincipal(t *testing.T) {
	compositeType := reflect.TypeOf(Principal{})
	alphaType := reflect.TypeOf(computealpha.Principal{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestRbacPolicy(t *testing.T) {
	compositeType := reflect.TypeOf(RbacPolicy{})
	alphaType := reflect.TypeOf(computealpha.RbacPolicy{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestRequestMirrorPolicy(t *testing.T) {
	compositeType := reflect.TypeOf(RequestMirrorPolicy{})
	alphaType := reflect.TypeOf(computealpha.RequestMirrorPolicy{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestSSLHealthCheck(t *testing.T) {
	compositeType := reflect.TypeOf(SSLHealthCheck{})
	alphaType := reflect.TypeOf(computealpha.SSLHealthCheck{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestSdsConfig(t *testing.T) {
	compositeType := reflect.TypeOf(SdsConfig{})
	alphaType := reflect.TypeOf(computealpha.SdsConfig{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestSecuritySettings(t *testing.T) {
	compositeType := reflect.TypeOf(SecuritySettings{})
	alphaType := reflect.TypeOf(computealpha.SecuritySettings{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestTCPHealthCheck(t *testing.T) {
	compositeType := reflect.TypeOf(TCPHealthCheck{})
	alphaType := reflect.TypeOf(computealpha.TCPHealthCheck{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}
func TestTargetHttpProxy(t *testing.T) {
	// Use reflection to verify that our composite type contains all the
	// same fields as the alpha type.
	compositeType := reflect.TypeOf(TargetHttpProxy{})
	alphaType := reflect.TypeOf(computealpha.TargetHttpProxy{})
	betaType := reflect.TypeOf(computebeta.TargetHttpProxy{})
	gaType := reflect.TypeOf(compute.TargetHttpProxy{})

	// For the composite type, remove the Version field from consideration
	compositeTypeNumFields := compositeType.NumField() - 2
	if compositeTypeNumFields != alphaType.NumField() {
		t.Fatalf("%v should contain %v fields. Got %v", alphaType.Name(), alphaType.NumField(), compositeTypeNumFields)
	}

	// Compare all the fields by doing a lookup since we can't guarantee that they'll be in the same order
	// Make sure that composite type is strictly alpha fields + internal bookkeeping
	for i := 2; i < compositeType.NumField(); i++ {
		lookupField, found := alphaType.FieldByName(compositeType.Field(i).Name)
		if !found {
			t.Fatal(fmt.Errorf("Field %v not present in alpha type %v", compositeType.Field(i), alphaType))
		}
		if err := compareFields(compositeType.Field(i), lookupField); err != nil {
			t.Fatal(err)
		}
	}

	// Verify that all beta fields are in composite type
	if err := typeEquality(betaType, compositeType, false); err != nil {
		t.Fatal(err)
	}

	// Verify that all GA fields are in composite type
	if err := typeEquality(gaType, compositeType, false); err != nil {
		t.Fatal(err)
	}
}

func TestToTargetHttpProxy(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected *TargetHttpProxy
	}{
		{
			computealpha.TargetHttpProxy{},
			&TargetHttpProxy{},
		},
		{
			computebeta.TargetHttpProxy{},
			&TargetHttpProxy{},
		},
		{
			compute.TargetHttpProxy{},
			&TargetHttpProxy{},
		},
	}
	for _, testCase := range testCases {
		result, _ := ToTargetHttpProxy(testCase.input)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Fatalf("ToTargetHttpProxy(input) = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(testCase.input), pretty.Sprint(result), pretty.Sprint(testCase.expected))
		}
	}
}

func TestTargetHttpProxyToAlpha(t *testing.T) {
	composite := TargetHttpProxy{}
	expected := &computealpha.TargetHttpProxy{}
	result, err := composite.ToAlpha()
	if err != nil {
		t.Fatalf("TargetHttpProxy.ToAlpha() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("TargetHttpProxy.ToAlpha() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}
func TestTargetHttpProxyToBeta(t *testing.T) {
	composite := TargetHttpProxy{}
	expected := &computebeta.TargetHttpProxy{}
	result, err := composite.ToBeta()
	if err != nil {
		t.Fatalf("TargetHttpProxy.ToBeta() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("TargetHttpProxy.ToBeta() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}
func TestTargetHttpProxyToGA(t *testing.T) {
	composite := TargetHttpProxy{}
	expected := &compute.TargetHttpProxy{}
	result, err := composite.ToGA()
	if err != nil {
		t.Fatalf("TargetHttpProxy.ToGA() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("TargetHttpProxy.ToGA() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}
func TestTargetHttpsProxy(t *testing.T) {
	// Use reflection to verify that our composite type contains all the
	// same fields as the alpha type.
	compositeType := reflect.TypeOf(TargetHttpsProxy{})
	alphaType := reflect.TypeOf(computealpha.TargetHttpsProxy{})
	betaType := reflect.TypeOf(computebeta.TargetHttpsProxy{})
	gaType := reflect.TypeOf(compute.TargetHttpsProxy{})

	// For the composite type, remove the Version field from consideration
	compositeTypeNumFields := compositeType.NumField() - 2
	if compositeTypeNumFields != alphaType.NumField() {
		t.Fatalf("%v should contain %v fields. Got %v", alphaType.Name(), alphaType.NumField(), compositeTypeNumFields)
	}

	// Compare all the fields by doing a lookup since we can't guarantee that they'll be in the same order
	// Make sure that composite type is strictly alpha fields + internal bookkeeping
	for i := 2; i < compositeType.NumField(); i++ {
		lookupField, found := alphaType.FieldByName(compositeType.Field(i).Name)
		if !found {
			t.Fatal(fmt.Errorf("Field %v not present in alpha type %v", compositeType.Field(i), alphaType))
		}
		if err := compareFields(compositeType.Field(i), lookupField); err != nil {
			t.Fatal(err)
		}
	}

	// Verify that all beta fields are in composite type
	if err := typeEquality(betaType, compositeType, false); err != nil {
		t.Fatal(err)
	}

	// Verify that all GA fields are in composite type
	if err := typeEquality(gaType, compositeType, false); err != nil {
		t.Fatal(err)
	}
}

func TestToTargetHttpsProxy(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected *TargetHttpsProxy
	}{
		{
			computealpha.TargetHttpsProxy{},
			&TargetHttpsProxy{},
		},
		{
			computebeta.TargetHttpsProxy{},
			&TargetHttpsProxy{},
		},
		{
			compute.TargetHttpsProxy{},
			&TargetHttpsProxy{},
		},
	}
	for _, testCase := range testCases {
		result, _ := ToTargetHttpsProxy(testCase.input)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Fatalf("ToTargetHttpsProxy(input) = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(testCase.input), pretty.Sprint(result), pretty.Sprint(testCase.expected))
		}
	}
}

func TestTargetHttpsProxyToAlpha(t *testing.T) {
	composite := TargetHttpsProxy{}
	expected := &computealpha.TargetHttpsProxy{}
	result, err := composite.ToAlpha()
	if err != nil {
		t.Fatalf("TargetHttpsProxy.ToAlpha() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("TargetHttpsProxy.ToAlpha() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}
func TestTargetHttpsProxyToBeta(t *testing.T) {
	composite := TargetHttpsProxy{}
	expected := &computebeta.TargetHttpsProxy{}
	result, err := composite.ToBeta()
	if err != nil {
		t.Fatalf("TargetHttpsProxy.ToBeta() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("TargetHttpsProxy.ToBeta() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}
func TestTargetHttpsProxyToGA(t *testing.T) {
	composite := TargetHttpsProxy{}
	expected := &compute.TargetHttpsProxy{}
	result, err := composite.ToGA()
	if err != nil {
		t.Fatalf("TargetHttpsProxy.ToGA() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("TargetHttpsProxy.ToGA() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}

func TestTlsCertificateContext(t *testing.T) {
	compositeType := reflect.TypeOf(TlsCertificateContext{})
	alphaType := reflect.TypeOf(computealpha.TlsCertificateContext{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestTlsCertificatePaths(t *testing.T) {
	compositeType := reflect.TypeOf(TlsCertificatePaths{})
	alphaType := reflect.TypeOf(computealpha.TlsCertificatePaths{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestTlsContext(t *testing.T) {
	compositeType := reflect.TypeOf(TlsContext{})
	alphaType := reflect.TypeOf(computealpha.TlsContext{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestTlsValidationContext(t *testing.T) {
	compositeType := reflect.TypeOf(TlsValidationContext{})
	alphaType := reflect.TypeOf(computealpha.TlsValidationContext{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestUDPHealthCheck(t *testing.T) {
	compositeType := reflect.TypeOf(UDPHealthCheck{})
	alphaType := reflect.TypeOf(computealpha.UDPHealthCheck{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}
func TestUrlMap(t *testing.T) {
	// Use reflection to verify that our composite type contains all the
	// same fields as the alpha type.
	compositeType := reflect.TypeOf(UrlMap{})
	alphaType := reflect.TypeOf(computealpha.UrlMap{})
	betaType := reflect.TypeOf(computebeta.UrlMap{})
	gaType := reflect.TypeOf(compute.UrlMap{})

	// For the composite type, remove the Version field from consideration
	compositeTypeNumFields := compositeType.NumField() - 2
	if compositeTypeNumFields != alphaType.NumField() {
		t.Fatalf("%v should contain %v fields. Got %v", alphaType.Name(), alphaType.NumField(), compositeTypeNumFields)
	}

	// Compare all the fields by doing a lookup since we can't guarantee that they'll be in the same order
	// Make sure that composite type is strictly alpha fields + internal bookkeeping
	for i := 2; i < compositeType.NumField(); i++ {
		lookupField, found := alphaType.FieldByName(compositeType.Field(i).Name)
		if !found {
			t.Fatal(fmt.Errorf("Field %v not present in alpha type %v", compositeType.Field(i), alphaType))
		}
		if err := compareFields(compositeType.Field(i), lookupField); err != nil {
			t.Fatal(err)
		}
	}

	// Verify that all beta fields are in composite type
	if err := typeEquality(betaType, compositeType, false); err != nil {
		t.Fatal(err)
	}

	// Verify that all GA fields are in composite type
	if err := typeEquality(gaType, compositeType, false); err != nil {
		t.Fatal(err)
	}
}

func TestToUrlMap(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected *UrlMap
	}{
		{
			computealpha.UrlMap{},
			&UrlMap{},
		},
		{
			computebeta.UrlMap{},
			&UrlMap{},
		},
		{
			compute.UrlMap{},
			&UrlMap{},
		},
	}
	for _, testCase := range testCases {
		result, _ := ToUrlMap(testCase.input)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Fatalf("ToUrlMap(input) = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(testCase.input), pretty.Sprint(result), pretty.Sprint(testCase.expected))
		}
	}
}

func TestUrlMapToAlpha(t *testing.T) {
	composite := UrlMap{}
	expected := &computealpha.UrlMap{}
	result, err := composite.ToAlpha()
	if err != nil {
		t.Fatalf("UrlMap.ToAlpha() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("UrlMap.ToAlpha() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}
func TestUrlMapToBeta(t *testing.T) {
	composite := UrlMap{}
	expected := &computebeta.UrlMap{}
	result, err := composite.ToBeta()
	if err != nil {
		t.Fatalf("UrlMap.ToBeta() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("UrlMap.ToBeta() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}
func TestUrlMapToGA(t *testing.T) {
	composite := UrlMap{}
	expected := &compute.UrlMap{}
	result, err := composite.ToGA()
	if err != nil {
		t.Fatalf("UrlMap.ToGA() error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("UrlMap.ToGA() = \ninput = %s\n%s\nwant = \n%s", pretty.Sprint(composite), pretty.Sprint(result), pretty.Sprint(expected))
	}
}

func TestUrlMapTest(t *testing.T) {
	compositeType := reflect.TypeOf(UrlMapTest{})
	alphaType := reflect.TypeOf(computealpha.UrlMapTest{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestUrlRewrite(t *testing.T) {
	compositeType := reflect.TypeOf(UrlRewrite{})
	alphaType := reflect.TypeOf(computealpha.UrlRewrite{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}

func TestWeightedBackendService(t *testing.T) {
	compositeType := reflect.TypeOf(WeightedBackendService{})
	alphaType := reflect.TypeOf(computealpha.WeightedBackendService{})
	if err := typeEquality(compositeType, alphaType, true); err != nil {
		t.Fatal(err)
	}
}
