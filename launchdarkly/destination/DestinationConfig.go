// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package destination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DestinationConfig struct {
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
	// The destination-specific configuration. To learn more, read [Destination-Specific Configs](#destination-specific-configs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.26.0/docs/resources/destination#config Destination#config}
	Config *map[string]*string `field:"required" json:"config" yaml:"config"`
	// The environment key.
	//
	// A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.26.0/docs/resources/destination#env_key Destination#env_key}
	EnvKey *string `field:"required" json:"envKey" yaml:"envKey"`
	// The data export destination type.
	//
	// Available choices are `kinesis`, `google-pubsub`, `mparticle`, `azure-event-hubs`, and `segment`. A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.26.0/docs/resources/destination#kind Destination#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// A human-readable name for your data export destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.26.0/docs/resources/destination#name Destination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The LaunchDarkly project key.
	//
	// A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.26.0/docs/resources/destination#project_key Destination#project_key}
	ProjectKey *string `field:"required" json:"projectKey" yaml:"projectKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.26.0/docs/resources/destination#id Destination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the data export destination is on or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.26.0/docs/resources/destination#on Destination#on}
	On interface{} `field:"optional" json:"on" yaml:"on"`
	// Tags associated with your resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.26.0/docs/resources/destination#tags Destination#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

