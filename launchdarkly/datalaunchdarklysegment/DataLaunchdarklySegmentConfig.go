package datalaunchdarklysegment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLaunchdarklySegmentConfig struct {
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
	// The segment's environment key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/segment#env_key DataLaunchdarklySegment#env_key}
	EnvKey *string `field:"required" json:"envKey" yaml:"envKey"`
	// The unique key that references the segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/segment#key DataLaunchdarklySegment#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The segment's project key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/segment#project_key DataLaunchdarklySegment#project_key}
	ProjectKey *string `field:"required" json:"projectKey" yaml:"projectKey"`
	// The description of the segment's purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/segment#description DataLaunchdarklySegment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of user keys excluded from the segment. To target on other context kinds, use the excluded_contexts block attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/segment#excluded DataLaunchdarklySegment#excluded}
	Excluded *[]*string `field:"optional" json:"excluded" yaml:"excluded"`
	// excluded_contexts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/segment#excluded_contexts DataLaunchdarklySegment#excluded_contexts}
	ExcludedContexts interface{} `field:"optional" json:"excludedContexts" yaml:"excludedContexts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/segment#id DataLaunchdarklySegment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of user keys included in the segment. To target on other context kinds, use the included_contexts block attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/segment#included DataLaunchdarklySegment#included}
	Included *[]*string `field:"optional" json:"included" yaml:"included"`
	// included_contexts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/segment#included_contexts DataLaunchdarklySegment#included_contexts}
	IncludedContexts interface{} `field:"optional" json:"includedContexts" yaml:"includedContexts"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/segment#rules DataLaunchdarklySegment#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// Tags associated with your resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/segment#tags DataLaunchdarklySegment#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

