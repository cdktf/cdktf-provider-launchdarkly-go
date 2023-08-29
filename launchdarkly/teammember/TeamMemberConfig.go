// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teammember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TeamMemberConfig struct {
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
	// The team member's email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/team_member#email TeamMember#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// IDs or keys of custom roles. Team members must have either a role or custom role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/team_member#custom_roles TeamMember#custom_roles}
	CustomRoles *[]*string `field:"optional" json:"customRoles" yaml:"customRoles"`
	// The team member's first name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/team_member#first_name TeamMember#first_name}
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/team_member#id TeamMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The team member's last name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/team_member#last_name TeamMember#last_name}
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// The team member's role.
	//
	// This must be reader, writer, admin, or no_access. Team members must have either a role or custom role
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/team_member#role TeamMember#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

