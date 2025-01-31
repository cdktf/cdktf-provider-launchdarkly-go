// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureflag


type FeatureFlagCustomProperties struct {
	// The unique custom property key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.5/docs/resources/feature_flag#key FeatureFlag#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The name of the custom property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.5/docs/resources/feature_flag#name FeatureFlag#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of custom property value strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.5/docs/resources/feature_flag#value FeatureFlag#value}
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

