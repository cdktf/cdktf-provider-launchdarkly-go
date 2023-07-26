package featureflagenvironment


type FeatureFlagEnvironmentTargets struct {
	// List of user strings to target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.3/docs/resources/feature_flag_environment#values FeatureFlagEnvironment#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Index of the variation to serve if a user_target is matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.3/docs/resources/feature_flag_environment#variation FeatureFlagEnvironment#variation}
	Variation *float64 `field:"required" json:"variation" yaml:"variation"`
}

