// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureflagenvironment


type FeatureFlagEnvironmentContextTargets struct {
	// The context kind on which the flag should target in this environment.
	//
	// User (`user`) targets should be specified as `targets` attribute blocks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.5/docs/resources/feature_flag_environment#context_kind FeatureFlagEnvironment#context_kind}
	ContextKind *string `field:"required" json:"contextKind" yaml:"contextKind"`
	// List of `user` strings to target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.5/docs/resources/feature_flag_environment#values FeatureFlagEnvironment#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// The index of the variation to serve if a user target value is matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.5/docs/resources/feature_flag_environment#variation FeatureFlagEnvironment#variation}
	Variation *float64 `field:"required" json:"variation" yaml:"variation"`
}

