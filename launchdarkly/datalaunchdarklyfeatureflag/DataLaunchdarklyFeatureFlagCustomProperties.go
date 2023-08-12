package datalaunchdarklyfeatureflag


type DataLaunchdarklyFeatureFlagCustomProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/data-sources/feature_flag#key DataLaunchdarklyFeatureFlag#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/data-sources/feature_flag#name DataLaunchdarklyFeatureFlag#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/data-sources/feature_flag#value DataLaunchdarklyFeatureFlag#value}.
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

