package datalaunchdarklyfeatureflagenvironment


type DataLaunchdarklyFeatureFlagEnvironmentContextTargets struct {
	// The context kind on which the flag should target in this environment.
	//
	// If the context_kind has not been previously specified at the project level, it will be automatically created on the project. User targets should be specified as targets attribute blocks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.3/docs/data-sources/feature_flag_environment#context_kind DataLaunchdarklyFeatureFlagEnvironment#context_kind}
	ContextKind *string `field:"required" json:"contextKind" yaml:"contextKind"`
	// List of user strings to target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.3/docs/data-sources/feature_flag_environment#values DataLaunchdarklyFeatureFlagEnvironment#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Index of the variation to serve if a user_target is matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.3/docs/data-sources/feature_flag_environment#variation DataLaunchdarklyFeatureFlagEnvironment#variation}
	Variation *float64 `field:"required" json:"variation" yaml:"variation"`
}

