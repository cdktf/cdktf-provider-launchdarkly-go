package datalaunchdarklyfeatureflagenvironment


type DataLaunchdarklyFeatureFlagEnvironmentRules struct {
	// Group percentage rollout by a custom attribute. This argument is only valid if rollout_weights is also specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/feature_flag_environment#bucket_by DataLaunchdarklyFeatureFlagEnvironment#bucket_by}
	BucketBy *string `field:"optional" json:"bucketBy" yaml:"bucketBy"`
	// clauses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/feature_flag_environment#clauses DataLaunchdarklyFeatureFlagEnvironment#clauses}
	Clauses interface{} `field:"optional" json:"clauses" yaml:"clauses"`
	// A human-readable description of the targeting rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/feature_flag_environment#description DataLaunchdarklyFeatureFlagEnvironment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/feature_flag_environment#rollout_weights DataLaunchdarklyFeatureFlagEnvironment#rollout_weights}.
	RolloutWeights *[]*float64 `field:"optional" json:"rolloutWeights" yaml:"rolloutWeights"`
	// The integer variation index to serve if the rule clauses evaluate to true.
	//
	// This argument is only valid if clauses are also specified
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/feature_flag_environment#variation DataLaunchdarklyFeatureFlagEnvironment#variation}
	Variation *float64 `field:"optional" json:"variation" yaml:"variation"`
}

