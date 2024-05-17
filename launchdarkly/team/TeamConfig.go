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
	// The team's unique key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.3/docs/resources/team#key Team#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The team's human-readable name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.3/docs/resources/team#name Team#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of custom role keys for the team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.3/docs/resources/team#custom_role_keys Team#custom_role_keys}
	CustomRoleKeys *[]*string `field:"optional" json:"customRoleKeys" yaml:"customRoleKeys"`
	// The team's description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.3/docs/resources/team#description Team#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.3/docs/resources/team#id Team#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A list of team maintainer IDs as strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.3/docs/resources/team#maintainers Team#maintainers}
	Maintainers *[]*string `field:"optional" json:"maintainers" yaml:"maintainers"`
	// A list of team member IDs as strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.3/docs/resources/team#member_ids Team#member_ids}
	MemberIds *[]*string `field:"optional" json:"memberIds" yaml:"memberIds"`
}

