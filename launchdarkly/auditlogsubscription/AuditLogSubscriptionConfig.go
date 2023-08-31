// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auditlogsubscription

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuditLogSubscriptionConfig struct {
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
	// The set of configuration fields corresponding to the value defined for `integration_key`.
	//
	// Refer to the `formVariables` field in the corresponding `integrations/<integration_key>/manifest.json` file in [this repo](https://github.com/launchdarkly/integration-framework/tree/master/integrations) for a full list of fields for the integration you wish to configure. **IMPORTANT**: Please note that Terraform will only accept these in snake case, regardless of the case shown in the manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.1/docs/resources/audit_log_subscription#config AuditLogSubscription#config}
	Config *map[string]*string `field:"required" json:"config" yaml:"config"`
	// The integration key.
	//
	// Supported integration keys are `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `logdna`, `msteams`, `new-relic-apm`, `signalfx`, `slack`, and `splunk`. A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.1/docs/resources/audit_log_subscription#integration_key AuditLogSubscription#integration_key}
	IntegrationKey *string `field:"required" json:"integrationKey" yaml:"integrationKey"`
	// A human-friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.1/docs/resources/audit_log_subscription#name AuditLogSubscription#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether or not you want your subscription enabled, i.e. to actively send events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.1/docs/resources/audit_log_subscription#on AuditLogSubscription#on}
	On interface{} `field:"required" json:"on" yaml:"on"`
	// statements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.1/docs/resources/audit_log_subscription#statements AuditLogSubscription#statements}
	Statements interface{} `field:"required" json:"statements" yaml:"statements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.1/docs/resources/audit_log_subscription#id AuditLogSubscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Tags associated with your resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.1/docs/resources/audit_log_subscription#tags AuditLogSubscription#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

