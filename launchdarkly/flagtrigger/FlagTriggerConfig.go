// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package flagtrigger

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FlagTriggerConfig struct {
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
	// Whether the trigger is currently active or not. This property defaults to true upon creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.2/docs/resources/flag_trigger#enabled FlagTrigger#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The LaunchDarkly environment key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.2/docs/resources/flag_trigger#env_key FlagTrigger#env_key}
	EnvKey *string `field:"required" json:"envKey" yaml:"envKey"`
	// The key of the feature flag the trigger acts upon.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.2/docs/resources/flag_trigger#flag_key FlagTrigger#flag_key}
	FlagKey *string `field:"required" json:"flagKey" yaml:"flagKey"`
	// instructions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.2/docs/resources/flag_trigger#instructions FlagTrigger#instructions}
	Instructions *FlagTriggerInstructions `field:"required" json:"instructions" yaml:"instructions"`
	// The unique identifier of the integration you intend to set your trigger up with.
	//
	// "generic-trigger" should be used for integrations not explicitly supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.2/docs/resources/flag_trigger#integration_key FlagTrigger#integration_key}
	IntegrationKey *string `field:"required" json:"integrationKey" yaml:"integrationKey"`
	// The LaunchDarkly project key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.2/docs/resources/flag_trigger#project_key FlagTrigger#project_key}
	ProjectKey *string `field:"required" json:"projectKey" yaml:"projectKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.2/docs/resources/flag_trigger#id FlagTrigger#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

