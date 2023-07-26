package datalaunchdarklyfeatureflag


type DataLaunchdarklyFeatureFlagDefaults struct {
	// The index of the variation served when the flag is off for new environments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/feature_flag#off_variation DataLaunchdarklyFeatureFlag#off_variation}
	OffVariation *float64 `field:"required" json:"offVariation" yaml:"offVariation"`
	// The index of the variation served when the flag is on for new environments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/feature_flag#on_variation DataLaunchdarklyFeatureFlag#on_variation}
	OnVariation *float64 `field:"required" json:"onVariation" yaml:"onVariation"`
}

