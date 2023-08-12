package datalaunchdarklyauditlogsubscription

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLaunchdarklyAuditLogSubscriptionConfig struct {
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
	// The audit log subscription ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/data-sources/audit_log_subscription#id DataLaunchdarklyAuditLogSubscription#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The integration key.
	//
	// Supported integration keys are `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `logdna`, `msteams`, `new-relic-apm`, `signalfx`, `slack`, and `splunk`. A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/data-sources/audit_log_subscription#integration_key DataLaunchdarklyAuditLogSubscription#integration_key}
	IntegrationKey *string `field:"required" json:"integrationKey" yaml:"integrationKey"`
}

