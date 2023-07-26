package featureflagenvironment


type FeatureFlagEnvironmentFallthrough struct {
	// Group percentage rollout by a custom attribute. This argument is only valid if rollout_weights is also specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/resources/feature_flag_environment#bucket_by FeatureFlagEnvironment#bucket_by}
	BucketBy *string `field:"optional" json:"bucketBy" yaml:"bucketBy"`
	// The context kind associated with the specified rollout.
	//
	// This argument is only valid if rollout_weights is also specified. If omitted, defaults to 'user'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/resources/feature_flag_environment#context_kind FeatureFlagEnvironment#context_kind}
	ContextKind *string `field:"optional" json:"contextKind" yaml:"contextKind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/resources/feature_flag_environment#rollout_weights FeatureFlagEnvironment#rollout_weights}.
	RolloutWeights *[]*float64 `field:"optional" json:"rolloutWeights" yaml:"rolloutWeights"`
	// The integer variation index to serve in case of fallthrough.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/resources/feature_flag_environment#variation FeatureFlagEnvironment#variation}
	Variation *float64 `field:"optional" json:"variation" yaml:"variation"`
}

