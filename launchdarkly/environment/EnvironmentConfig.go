// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package environment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EnvironmentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The color swatch as an RGB hex value with no leading `#`. For example: `000000`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.0/docs/resources/environment#color Environment#color}
	Color *string `field:"required" json:"color" yaml:"color"`
	// The project-unique key for the environment.
	//
	// A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.0/docs/resources/environment#key Environment#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The name of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.0/docs/resources/environment#name Environment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The LaunchDarkly project key.
	//
	// A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.0/docs/resources/environment#project_key Environment#project_key}
	ProjectKey *string `field:"required" json:"projectKey" yaml:"projectKey"`
	// approval_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.0/docs/resources/environment#approval_settings Environment#approval_settings}
	ApprovalSettings interface{} `field:"optional" json:"approvalSettings" yaml:"approvalSettings"`
	// Set to `true` if this environment requires confirmation for flag and segment changes.
	//
	// This field will default to `false` when not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.0/docs/resources/environment#confirm_changes Environment#confirm_changes}
	ConfirmChanges interface{} `field:"optional" json:"confirmChanges" yaml:"confirmChanges"`
	// Denotes whether the environment is critical.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.0/docs/resources/environment#critical Environment#critical}
	Critical interface{} `field:"optional" json:"critical" yaml:"critical"`
	// Set to `true` to enable data export for every flag created in this environment after you configure this argument.
	//
	// This field will default to `false` when not set. To learn more, read [Data Export](https://docs.launchdarkly.com/home/data-export).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.0/docs/resources/environment#default_track_events Environment#default_track_events}
	DefaultTrackEvents interface{} `field:"optional" json:"defaultTrackEvents" yaml:"defaultTrackEvents"`
	// The TTL for the environment.
	//
	// This must be between 0 and 60 minutes. The TTL setting only applies to environments using the PHP SDK. This field will default to `0` when not set. To learn more, read [TTL settings](https://docs.launchdarkly.com/home/organize/environments#ttl-settings).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.0/docs/resources/environment#default_ttl Environment#default_ttl}
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.0/docs/resources/environment#id Environment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set to `true` if this environment requires comments for flag and segment changes.
	//
	// This field will default to `false` when not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.0/docs/resources/environment#require_comments Environment#require_comments}
	RequireComments interface{} `field:"optional" json:"requireComments" yaml:"requireComments"`
	// Set to `true` to ensure a user of the client-side SDK cannot impersonate another user.
	//
	// This field will default to `false` when not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.0/docs/resources/environment#secure_mode Environment#secure_mode}
	SecureMode interface{} `field:"optional" json:"secureMode" yaml:"secureMode"`
	// Tags associated with your resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.0/docs/resources/environment#tags Environment#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

