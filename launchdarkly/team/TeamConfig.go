// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package team

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TeamConfig struct {
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
	// The team key.
	//
	// A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.2/docs/resources/team#key Team#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// A human-friendly name for the team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.2/docs/resources/team#name Team#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of custom role keys the team will access.
	//
	// The referenced custom roles must already exist in LaunchDarkly. If they don't, the provider may behave unexpectedly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.2/docs/resources/team#custom_role_keys Team#custom_role_keys}
	CustomRoleKeys *[]*string `field:"optional" json:"customRoleKeys" yaml:"customRoleKeys"`
	// The team description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.2/docs/resources/team#description Team#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.2/docs/resources/team#id Team#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of member IDs for users who maintain the team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.2/docs/resources/team#maintainers Team#maintainers}
	Maintainers *[]*string `field:"optional" json:"maintainers" yaml:"maintainers"`
	// List of member IDs who belong to the team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.2/docs/resources/team#member_ids Team#member_ids}
	MemberIds *[]*string `field:"optional" json:"memberIds" yaml:"memberIds"`
	// role_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.2/docs/resources/team#role_attributes Team#role_attributes}
	RoleAttributes interface{} `field:"optional" json:"roleAttributes" yaml:"roleAttributes"`
}

