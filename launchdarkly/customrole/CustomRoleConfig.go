// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customrole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomRoleConfig struct {
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
	// A unique key that will be used to reference the custom role in your code.
	//
	// A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.0/docs/resources/custom_role#key CustomRole#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// A name for the custom role. This must be unique within your organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.0/docs/resources/custom_role#name CustomRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The base permission level - either reader or no_access. Defaults to reader.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.0/docs/resources/custom_role#base_permissions CustomRole#base_permissions}
	BasePermissions *string `field:"optional" json:"basePermissions" yaml:"basePermissions"`
	// Description of the custom role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.0/docs/resources/custom_role#description CustomRole#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.0/docs/resources/custom_role#id CustomRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.0/docs/resources/custom_role#policy CustomRole#policy}
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// policy_statements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.0/docs/resources/custom_role#policy_statements CustomRole#policy_statements}
	PolicyStatements interface{} `field:"optional" json:"policyStatements" yaml:"policyStatements"`
}

