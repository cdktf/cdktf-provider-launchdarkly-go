// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalaunchdarklyteammember


type DataLaunchdarklyTeamMemberRoleAttributes struct {
	// The key / name of your role attribute. In the example `$${roleAttribute/testAttribute}`, the key is `testAttribute`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.1/docs/data-sources/team_member#key DataLaunchdarklyTeamMember#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// A list of values for your role attribute.
	//
	// For example, if your policy statement defines the resource `"proj/$${roleAttribute/testAttribute}"`, the values would be the keys of the projects you wanted to assign access to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.1/docs/data-sources/team_member#values DataLaunchdarklyTeamMember#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

