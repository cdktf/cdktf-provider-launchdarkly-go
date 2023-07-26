package datalaunchdarklyfeatureflagenvironment


type DataLaunchdarklyFeatureFlagEnvironmentRulesClauses struct {
	// The user attribute to operate on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.3/docs/data-sources/feature_flag_environment#attribute DataLaunchdarklyFeatureFlagEnvironment#attribute}
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// The operator associated with the rule clause.
	//
	// Available options are in, endsWith, startsWith, matches, contains, lessThan, lessThanOrEqual, greaterThanOrEqual, before, after, segmentMatch, semVerEqual, semVerLessThan, and semVerGreaterThan
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.3/docs/data-sources/feature_flag_environment#op DataLaunchdarklyFeatureFlagEnvironment#op}
	Op *string `field:"required" json:"op" yaml:"op"`
	// The list of values associated with the rule clause.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.3/docs/data-sources/feature_flag_environment#values DataLaunchdarklyFeatureFlagEnvironment#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// The context kind associated with this rule clause. If omitted, defaults to user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.3/docs/data-sources/feature_flag_environment#context_kind DataLaunchdarklyFeatureFlagEnvironment#context_kind}
	ContextKind *string `field:"optional" json:"contextKind" yaml:"contextKind"`
	// Whether to negate the rule clause.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.3/docs/data-sources/feature_flag_environment#negate DataLaunchdarklyFeatureFlagEnvironment#negate}
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
	// The type for each of the clause's values.
	//
	// Available types are boolean, string, and number. If omitted, value_type defaults to string
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.3/docs/data-sources/feature_flag_environment#value_type DataLaunchdarklyFeatureFlagEnvironment#value_type}
	ValueType *string `field:"optional" json:"valueType" yaml:"valueType"`
}

