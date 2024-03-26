// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package project


type ProjectEnvironments struct {
	// The color swatch as an RGB hex value with no leading `#`. For example: `000000`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.2/docs/resources/project#color Project#color}
	Color *string `field:"required" json:"color" yaml:"color"`
	// The project-unique key for the environment.
	//
	// A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.2/docs/resources/project#key Project#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The name of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.2/docs/resources/project#name Project#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// approval_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.2/docs/resources/project#approval_settings Project#approval_settings}
	ApprovalSettings interface{} `field:"optional" json:"approvalSettings" yaml:"approvalSettings"`
	// Set to `true` if this environment requires confirmation for flag and segment changes.
	//
	// This field will default to `false` when not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.2/docs/resources/project#confirm_changes Project#confirm_changes}
	ConfirmChanges interface{} `field:"optional" json:"confirmChanges" yaml:"confirmChanges"`
	// Set to `true` to enable data export for every flag created in this environment after you configure this argument.
	//
	// This field will default to `false` when not set. To learn more, read [Data Export](https://docs.launchdarkly.com/home/data-export).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.2/docs/resources/project#default_track_events Project#default_track_events}
	DefaultTrackEvents interface{} `field:"optional" json:"defaultTrackEvents" yaml:"defaultTrackEvents"`
	// The TTL for the environment.
	//
	// This must be between 0 and 60 minutes. The TTL setting only applies to environments using the PHP SDK. This field will default to `0` when not set. To learn more, read [TTL settings](https://docs.launchdarkly.com/home/organize/environments#ttl-settings).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.2/docs/resources/project#default_ttl Project#default_ttl}
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// Set to `true` if this environment requires comments for flag and segment changes.
	//
	// This field will default to `false` when not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.2/docs/resources/project#require_comments Project#require_comments}
	RequireComments interface{} `field:"optional" json:"requireComments" yaml:"requireComments"`
	// Set to `true` to ensure a user of the client-side SDK cannot impersonate another user.
	//
	// This field will default to `false` when not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.2/docs/resources/project#secure_mode Project#secure_mode}
	SecureMode interface{} `field:"optional" json:"secureMode" yaml:"secureMode"`
	// Tags associated with your resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.2/docs/resources/project#tags Project#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

