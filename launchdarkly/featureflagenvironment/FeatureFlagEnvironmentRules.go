// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureflagenvironment


type FeatureFlagEnvironmentRules struct {
	// Group percentage rollout by a custom attribute. This argument is only valid if `rollout_weights` is also specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag_environment#bucket_by FeatureFlagEnvironment#bucket_by}
	BucketBy *string `field:"optional" json:"bucketBy" yaml:"bucketBy"`
	// clauses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag_environment#clauses FeatureFlagEnvironment#clauses}
	Clauses interface{} `field:"optional" json:"clauses" yaml:"clauses"`
	// The context kind associated with the specified rollout.
	//
	// This argument is only valid if `rollout_weights` is also specified. Defaults to `user` if omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag_environment#context_kind FeatureFlagEnvironment#context_kind}
	ContextKind *string `field:"optional" json:"contextKind" yaml:"contextKind"`
	// A human-readable description of the targeting rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag_environment#description FeatureFlagEnvironment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of integer percentage rollout weights (in thousandths of a percent) to apply to each variation if the rule clauses evaluates to `true`.
	//
	// The sum of the `rollout_weights` must equal 100000 and the number of rollout weights specified in the array must match the number of flag variations. You must specify either `variation` or `rollout_weights`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag_environment#rollout_weights FeatureFlagEnvironment#rollout_weights}
	RolloutWeights *[]*float64 `field:"optional" json:"rolloutWeights" yaml:"rolloutWeights"`
	// The integer variation index to serve if the rule clauses evaluate to `true`.
	//
	// You must specify either `variation` or `rollout_weights`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag_environment#variation FeatureFlagEnvironment#variation}
	Variation *float64 `field:"optional" json:"variation" yaml:"variation"`
}

