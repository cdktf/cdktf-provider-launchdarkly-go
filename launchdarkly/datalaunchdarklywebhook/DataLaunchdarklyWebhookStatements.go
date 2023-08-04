package datalaunchdarklywebhook


type DataLaunchdarklyWebhookStatements struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/webhook#effect DataLaunchdarklyWebhook#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// An action to perform on a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/webhook#actions DataLaunchdarklyWebhook#actions}
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Targeted actions will be those actions NOT in this list.
	//
	// The 'actions' field must be empty to use this field
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/webhook#not_actions DataLaunchdarklyWebhook#not_actions}
	NotActions *[]*string `field:"optional" json:"notActions" yaml:"notActions"`
	// Targeted resources will be those resources NOT in this list.
	//
	// The 'resources' field must be empty to use this field
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/webhook#not_resources DataLaunchdarklyWebhook#not_resources}
	NotResources *[]*string `field:"optional" json:"notResources" yaml:"notResources"`
	// A list of LaunchDarkly resource specifiers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/webhook#resources DataLaunchdarklyWebhook#resources}
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

