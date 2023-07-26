package datalaunchdarklyauditlogsubscription


type DataLaunchdarklyAuditLogSubscriptionStatements struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/audit_log_subscription#effect DataLaunchdarklyAuditLogSubscription#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// An action to perform on a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/audit_log_subscription#actions DataLaunchdarklyAuditLogSubscription#actions}
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Targeted actions will be those actions NOT in this list.
	//
	// The 'actions' field must be empty to use this field
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/audit_log_subscription#not_actions DataLaunchdarklyAuditLogSubscription#not_actions}
	NotActions *[]*string `field:"optional" json:"notActions" yaml:"notActions"`
	// Targeted resources will be those resources NOT in this list.
	//
	// The 'resources' field must be empty to use this field
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/audit_log_subscription#not_resources DataLaunchdarklyAuditLogSubscription#not_resources}
	NotResources *[]*string `field:"optional" json:"notResources" yaml:"notResources"`
	// A list of LaunchDarkly resource specifiers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/audit_log_subscription#resources DataLaunchdarklyAuditLogSubscription#resources}
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

