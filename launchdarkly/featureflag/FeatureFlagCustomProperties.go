package featureflag


type FeatureFlagCustomProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#key FeatureFlag#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#name FeatureFlag#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#value FeatureFlag#value}.
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

