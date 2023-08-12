package featureflag


type FeatureFlagVariations struct {
	// The value of the flag for this variation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#value FeatureFlag#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// A description for the variation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#description FeatureFlag#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the variation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#name FeatureFlag#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

