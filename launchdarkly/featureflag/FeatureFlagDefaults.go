// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureflag


type FeatureFlagDefaults struct {
	// The index of the variation the flag will default to in all new environments when off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.26.0/docs/resources/feature_flag#off_variation FeatureFlag#off_variation}
	OffVariation *float64 `field:"required" json:"offVariation" yaml:"offVariation"`
	// The index of the variation the flag will default to in all new environments when on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.26.0/docs/resources/feature_flag#on_variation FeatureFlag#on_variation}
	OnVariation *float64 `field:"required" json:"onVariation" yaml:"onVariation"`
}

