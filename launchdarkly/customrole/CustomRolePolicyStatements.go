package customrole


type CustomRolePolicyStatements struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/resources/custom_role#effect CustomRole#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// An action to perform on a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/resources/custom_role#actions CustomRole#actions}
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Targeted actions will be those actions NOT in this list.
	//
	// The 'actions' field must be empty to use this field
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/resources/custom_role#not_actions CustomRole#not_actions}
	NotActions *[]*string `field:"optional" json:"notActions" yaml:"notActions"`
	// Targeted resources will be those resources NOT in this list.
	//
	// The 'resources' field must be empty to use this field
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/resources/custom_role#not_resources CustomRole#not_resources}
	NotResources *[]*string `field:"optional" json:"notResources" yaml:"notResources"`
	// A list of LaunchDarkly resource specifiers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/resources/custom_role#resources CustomRole#resources}
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

