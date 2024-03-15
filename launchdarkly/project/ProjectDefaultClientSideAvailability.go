// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package project


type ProjectDefaultClientSideAvailability struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/project#using_environment_id Project#using_environment_id}.
	UsingEnvironmentId interface{} `field:"required" json:"usingEnvironmentId" yaml:"usingEnvironmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/project#using_mobile_key Project#using_mobile_key}.
	UsingMobileKey interface{} `field:"required" json:"usingMobileKey" yaml:"usingMobileKey"`
}

