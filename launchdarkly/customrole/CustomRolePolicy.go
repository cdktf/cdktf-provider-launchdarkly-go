// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customrole


type CustomRolePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/custom_role#actions CustomRole#actions}.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/custom_role#effect CustomRole#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/custom_role#resources CustomRole#resources}.
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
}

