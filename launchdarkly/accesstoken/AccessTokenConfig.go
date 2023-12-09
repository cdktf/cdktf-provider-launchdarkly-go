// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesstoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessTokenConfig struct {
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
	// A list of custom role IDs to use as access limits for the access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.17.0/docs/resources/access_token#custom_roles AccessToken#custom_roles}
	CustomRoles *[]*string `field:"optional" json:"customRoles" yaml:"customRoles"`
	// The default API version for this token.
	//
	// Defaults to the latest API version. A change in this field will force the destruction of the existing token in state and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.17.0/docs/resources/access_token#default_api_version AccessToken#default_api_version}
	DefaultApiVersion *float64 `field:"optional" json:"defaultApiVersion" yaml:"defaultApiVersion"`
	// An expiration time for the current token secret, expressed as a Unix epoch time.
	//
	// Replace the computed token secret with a new value. The expired secret will no longer be able to authorize usage of the LaunchDarkly API. This field argument is **deprecated**. Please update your config to remove `expire` to maintain compatibility with future versions
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.17.0/docs/resources/access_token#expire AccessToken#expire}
	Expire *float64 `field:"optional" json:"expire" yaml:"expire"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.17.0/docs/resources/access_token#id AccessToken#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// inline_roles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.17.0/docs/resources/access_token#inline_roles AccessToken#inline_roles}
	InlineRoles interface{} `field:"optional" json:"inlineRoles" yaml:"inlineRoles"`
	// A human-friendly name for the access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.17.0/docs/resources/access_token#name AccessToken#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// policy_statements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.17.0/docs/resources/access_token#policy_statements AccessToken#policy_statements}
	PolicyStatements interface{} `field:"optional" json:"policyStatements" yaml:"policyStatements"`
	// A built-in LaunchDarkly role. Can be `reader`, `writer`, or `admin`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.17.0/docs/resources/access_token#role AccessToken#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Whether the token will be a [service token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#service-tokens). A change in this field will force the destruction of the existing token and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.17.0/docs/resources/access_token#service_token AccessToken#service_token}
	ServiceToken interface{} `field:"optional" json:"serviceToken" yaml:"serviceToken"`
}

