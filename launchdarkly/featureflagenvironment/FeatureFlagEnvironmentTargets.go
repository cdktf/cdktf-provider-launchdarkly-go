// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureflagenvironment


type FeatureFlagEnvironmentTargets struct {
	// List of `user` strings to target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.0/docs/resources/feature_flag_environment#values FeatureFlagEnvironment#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// The index of the variation to serve if a user target value is matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.0/docs/resources/feature_flag_environment#variation FeatureFlagEnvironment#variation}
	Variation *float64 `field:"required" json:"variation" yaml:"variation"`
}

