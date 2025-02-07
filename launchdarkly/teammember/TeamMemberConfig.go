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
	// The unique email address associated with the team member.
	//
	// A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.22.0/docs/resources/team_member#email TeamMember#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// The list of custom roles keys associated with the team member.
	//
	// Custom roles are only available to customers on an Enterprise plan. To learn more, [read about our pricing](https://launchdarkly.com/pricing/). To upgrade your plan, [contact LaunchDarkly Sales](https://launchdarkly.com/contact-sales/).
	//
	// -> **Note:** each `launchdarkly_team_member` must have either a `role` or `custom_roles` argument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.22.0/docs/resources/team_member#custom_roles TeamMember#custom_roles}
	CustomRoles *[]*string `field:"optional" json:"customRoles" yaml:"customRoles"`
	// The team member's given name. Once created, this cannot be updated except by the team member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.22.0/docs/resources/team_member#first_name TeamMember#first_name}
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// TThe team member's family name. Once created, this cannot be updated except by the team member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.22.0/docs/resources/team_member#last_name TeamMember#last_name}
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// The role associated with team member.
	//
	// Supported roles are `reader`, `writer`, `no_access`, or `admin`. If you don't specify a role, `reader` is assigned by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.22.0/docs/resources/team_member#role TeamMember#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// role_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.22.0/docs/resources/team_member#role_attributes TeamMember#role_attributes}
	RoleAttributes interface{} `field:"optional" json:"roleAttributes" yaml:"roleAttributes"`
}

