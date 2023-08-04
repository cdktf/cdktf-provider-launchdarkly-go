package datalaunchdarklyenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLaunchdarklyEnvironmentConfig struct {
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
	// A project-unique key for the new environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/environment#key DataLaunchdarklyEnvironment#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/environment#project_key DataLaunchdarklyEnvironment#project_key}.
	ProjectKey *string `field:"required" json:"projectKey" yaml:"projectKey"`
	// approval_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/environment#approval_settings DataLaunchdarklyEnvironment#approval_settings}
	ApprovalSettings interface{} `field:"optional" json:"approvalSettings" yaml:"approvalSettings"`
	// Whether or not to require confirmation for flag and segment changes in this environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/environment#confirm_changes DataLaunchdarklyEnvironment#confirm_changes}
	ConfirmChanges interface{} `field:"optional" json:"confirmChanges" yaml:"confirmChanges"`
	// Whether or not to default to sending data export events for flags created in the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/environment#default_track_events DataLaunchdarklyEnvironment#default_track_events}
	DefaultTrackEvents interface{} `field:"optional" json:"defaultTrackEvents" yaml:"defaultTrackEvents"`
	// The TTL for the environment.
	//
	// This must be between 0 and 60 minutes. The TTL setting only applies to environments using the PHP SDK
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/environment#default_ttl DataLaunchdarklyEnvironment#default_ttl}
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/environment#id DataLaunchdarklyEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether or not to require comments for flag and segment changes in this environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/environment#require_comments DataLaunchdarklyEnvironment#require_comments}
	RequireComments interface{} `field:"optional" json:"requireComments" yaml:"requireComments"`
	// Whether or not to use secure mode.
	//
	// Secure mode ensures a user of the client-side SDK cannot impersonate another user
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/environment#secure_mode DataLaunchdarklyEnvironment#secure_mode}
	SecureMode interface{} `field:"optional" json:"secureMode" yaml:"secureMode"`
	// Tags associated with your resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/environment#tags DataLaunchdarklyEnvironment#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

