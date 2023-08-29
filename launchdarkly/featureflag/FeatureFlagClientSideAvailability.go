// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureflag


type FeatureFlagClientSideAvailability struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#using_environment_id FeatureFlag#using_environment_id}.
	UsingEnvironmentId interface{} `field:"optional" json:"usingEnvironmentId" yaml:"usingEnvironmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#using_mobile_key FeatureFlag#using_mobile_key}.
	UsingMobileKey interface{} `field:"optional" json:"usingMobileKey" yaml:"usingMobileKey"`
}

