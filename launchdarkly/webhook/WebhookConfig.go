// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webhook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WebhookConfig struct {
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
	// The URL of the remote webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/webhook#url Webhook#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/webhook#id Webhook#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The webhook's human-readable name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/webhook#name Webhook#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies whether the webhook is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/webhook#on Webhook#on}
	On interface{} `field:"optional" json:"on" yaml:"on"`
	// The secret used to sign the webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/webhook#secret Webhook#secret}
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
	// statements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/webhook#statements Webhook#statements}
	Statements interface{} `field:"optional" json:"statements" yaml:"statements"`
	// Tags associated with your resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/webhook#tags Webhook#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

