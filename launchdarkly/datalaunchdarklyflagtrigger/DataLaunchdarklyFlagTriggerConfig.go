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
	// The unique key of the environment the flag trigger will work in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.2/docs/data-sources/flag_trigger#env_key DataLaunchdarklyFlagTrigger#env_key}
	EnvKey *string `field:"required" json:"envKey" yaml:"envKey"`
	// The unique key of the associated flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.2/docs/data-sources/flag_trigger#flag_key DataLaunchdarklyFlagTrigger#flag_key}
	FlagKey *string `field:"required" json:"flagKey" yaml:"flagKey"`
	// The Terraform trigger ID. The unique trigger ID can be found in your saved trigger URL:.
	//
	// ```
	// https://app.launchdarkly.com/webhook/triggers/THIS_IS_YOUR_TRIGGER_ID/aff25a53-17d9-4112-a9b8-12718d1a2e79
	// ```
	//
	// Please note that if you did not save this upon creation of the resource, you will have to reset it to get a new value, which can cause breaking changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.2/docs/data-sources/flag_trigger#id DataLaunchdarklyFlagTrigger#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The unique key of the project encompassing the associated flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.2/docs/data-sources/flag_trigger#project_key DataLaunchdarklyFlagTrigger#project_key}
	ProjectKey *string `field:"required" json:"projectKey" yaml:"projectKey"`
}

