// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalaunchdarklyflagtrigger

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLaunchdarklyFlagTriggerConfig struct {
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
	// The LaunchDarkly environment key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.2/docs/data-sources/flag_trigger#env_key DataLaunchdarklyFlagTrigger#env_key}
	EnvKey *string `field:"required" json:"envKey" yaml:"envKey"`
	// The key of the feature flag the trigger acts upon.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.2/docs/data-sources/flag_trigger#flag_key DataLaunchdarklyFlagTrigger#flag_key}
	FlagKey *string `field:"required" json:"flagKey" yaml:"flagKey"`
	// The flag trigger resource ID.
	//
	// This can be found on your trigger URL - please see docs for more info
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.2/docs/data-sources/flag_trigger#id DataLaunchdarklyFlagTrigger#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The LaunchDarkly project key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.2/docs/data-sources/flag_trigger#project_key DataLaunchdarklyFlagTrigger#project_key}
	ProjectKey *string `field:"required" json:"projectKey" yaml:"projectKey"`
	// Whether the trigger is currently active or not. This property defaults to true upon creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.2/docs/data-sources/flag_trigger#enabled DataLaunchdarklyFlagTrigger#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// instructions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.2/docs/data-sources/flag_trigger#instructions DataLaunchdarklyFlagTrigger#instructions}
	Instructions *DataLaunchdarklyFlagTriggerInstructions `field:"optional" json:"instructions" yaml:"instructions"`
	// The unique identifier of the integration you intend to set your trigger up with.
	//
	// "generic-trigger" should be used for integrations not explicitly supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.2/docs/data-sources/flag_trigger#integration_key DataLaunchdarklyFlagTrigger#integration_key}
	IntegrationKey *string `field:"optional" json:"integrationKey" yaml:"integrationKey"`
}

