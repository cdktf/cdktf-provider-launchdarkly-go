package project


type ProjectEnvironments struct {
	// A color swatch (as an RGB hex value with no leading '#', e.g. C8C8C8).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/resources/project#color Project#color}
	Color *string `field:"required" json:"color" yaml:"color"`
	// A project-unique key for the new environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/resources/project#key Project#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The name of the new environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/resources/project#name Project#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// approval_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/resources/project#approval_settings Project#approval_settings}
	ApprovalSettings interface{} `field:"optional" json:"approvalSettings" yaml:"approvalSettings"`
	// Whether or not to require confirmation for flag and segment changes in this environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/resources/project#confirm_changes Project#confirm_changes}
	ConfirmChanges interface{} `field:"optional" json:"confirmChanges" yaml:"confirmChanges"`
	// Whether or not to default to sending data export events for flags created in the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/resources/project#default_track_events Project#default_track_events}
	DefaultTrackEvents interface{} `field:"optional" json:"defaultTrackEvents" yaml:"defaultTrackEvents"`
	// The TTL for the environment.
	//
	// This must be between 0 and 60 minutes. The TTL setting only applies to environments using the PHP SDK
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/resources/project#default_ttl Project#default_ttl}
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// Whether or not to require comments for flag and segment changes in this environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/resources/project#require_comments Project#require_comments}
	RequireComments interface{} `field:"optional" json:"requireComments" yaml:"requireComments"`
	// Whether or not to use secure mode.
	//
	// Secure mode ensures a user of the client-side SDK cannot impersonate another user
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/resources/project#secure_mode Project#secure_mode}
	SecureMode interface{} `field:"optional" json:"secureMode" yaml:"secureMode"`
	// Tags associated with your resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/resources/project#tags Project#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

