// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureflagenvironment


type FeatureFlagEnvironmentPrerequisites struct {
	// The prerequisite feature flag's `key`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.4/docs/resources/feature_flag_environment#flag_key FeatureFlagEnvironment#flag_key}
	FlagKey *string `field:"required" json:"flagKey" yaml:"flagKey"`
	// The index of the prerequisite feature flag's variation to target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.4/docs/resources/feature_flag_environment#variation FeatureFlagEnvironment#variation}
	Variation *float64 `field:"required" json:"variation" yaml:"variation"`
}

