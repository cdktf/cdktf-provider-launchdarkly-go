// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package segment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SegmentConfig struct {
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
	// A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/segment#env_key Segment#env_key}
	EnvKey *string `field:"required" json:"envKey" yaml:"envKey"`
	// The unique key that references the segment.
	//
	// A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/segment#key Segment#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The human-friendly name for the segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/segment#name Segment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The segment's project key.
	//
	// A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/segment#project_key Segment#project_key}
	ProjectKey *string `field:"required" json:"projectKey" yaml:"projectKey"`
	// The description of the segment's purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/segment#description Segment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of user keys excluded from the segment.
	//
	// To target on other context kinds, use the excluded_contexts block attribute. This attribute is not valid when `unbounded` is set to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/segment#excluded Segment#excluded}
	Excluded *[]*string `field:"optional" json:"excluded" yaml:"excluded"`
	// excluded_contexts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/segment#excluded_contexts Segment#excluded_contexts}
	ExcludedContexts interface{} `field:"optional" json:"excludedContexts" yaml:"excludedContexts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/segment#id Segment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of user keys included in the segment.
	//
	// To target on other context kinds, use the included_contexts block attribute. This attribute is not valid when `unbounded` is set to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/segment#included Segment#included}
	Included *[]*string `field:"optional" json:"included" yaml:"included"`
	// included_contexts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/segment#included_contexts Segment#included_contexts}
	IncludedContexts interface{} `field:"optional" json:"includedContexts" yaml:"includedContexts"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/segment#rules Segment#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// Tags associated with your resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/segment#tags Segment#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Whether to create a standard segment (`false`) or a Big Segment (`true`).
	//
	// Standard segments include rule-based and smaller list-based segments. Big Segments include larger list-based segments and synced segments. Only use a Big Segment if you need to add more than 15,000 individual targets. It is not possible to manage the list of targeted contexts for Big Segments with Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/segment#unbounded Segment#unbounded}
	Unbounded interface{} `field:"optional" json:"unbounded" yaml:"unbounded"`
	// For Big Segments, the targeted context kind. If this attribute is not specified it will default to `user`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/segment#unbounded_context_kind Segment#unbounded_context_kind}
	UnboundedContextKind *string `field:"optional" json:"unboundedContextKind" yaml:"unboundedContextKind"`
}

