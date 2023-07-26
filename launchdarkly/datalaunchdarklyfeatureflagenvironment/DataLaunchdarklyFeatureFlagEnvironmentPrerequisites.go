package datalaunchdarklyfeatureflagenvironment


type DataLaunchdarklyFeatureFlagEnvironmentPrerequisites struct {
	// The prerequisite feature flag's key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.3/docs/data-sources/feature_flag_environment#flag_key DataLaunchdarklyFeatureFlagEnvironment#flag_key}
	FlagKey *string `field:"required" json:"flagKey" yaml:"flagKey"`
	// The index of the prerequisite feature flag's variation to target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.3/docs/data-sources/feature_flag_environment#variation DataLaunchdarklyFeatureFlagEnvironment#variation}
	Variation *float64 `field:"required" json:"variation" yaml:"variation"`
}

