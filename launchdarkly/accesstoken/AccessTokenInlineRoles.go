// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesstoken


type AccessTokenInlineRoles struct {
	// Either `allow` or `deny`.
	//
	// This argument defines whether the statement allows or denies access to the named resources and actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.2/docs/resources/access_token#effect AccessToken#effect}
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// The list of action specifiers defining the actions to which the statement applies.
	//
	// Either `actions` or `not_actions` must be specified. For a list of available actions read [Actions reference](https://docs.launchdarkly.com/home/account-security/custom-roles/actions#actions-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.2/docs/resources/access_token#actions AccessToken#actions}
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// The list of action specifiers defining the actions to which the statement does not apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.2/docs/resources/access_token#not_actions AccessToken#not_actions}
	NotActions *[]*string `field:"optional" json:"notActions" yaml:"notActions"`
	// The list of resource specifiers defining the resources to which the statement does not apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.2/docs/resources/access_token#not_resources AccessToken#not_resources}
	NotResources *[]*string `field:"optional" json:"notResources" yaml:"notResources"`
	// The list of resource specifiers defining the resources to which the statement applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.2/docs/resources/access_token#resources AccessToken#resources}
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

