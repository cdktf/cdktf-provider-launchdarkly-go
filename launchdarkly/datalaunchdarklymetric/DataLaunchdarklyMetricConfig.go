// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalaunchdarklymetric

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLaunchdarklyMetricConfig struct {
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
	// A unique key that will be used to reference the metric in your code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.16.0/docs/data-sources/metric#key DataLaunchdarklyMetric#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The LaunchDarkly project key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.16.0/docs/data-sources/metric#project_key DataLaunchdarklyMetric#project_key}
	ProjectKey *string `field:"required" json:"projectKey" yaml:"projectKey"`
	// A short description of what the metric will be used for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.16.0/docs/data-sources/metric#description DataLaunchdarklyMetric#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The event key for your metric (if custom metric).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.16.0/docs/data-sources/metric#event_key DataLaunchdarklyMetric#event_key}
	EventKey *string `field:"optional" json:"eventKey" yaml:"eventKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.16.0/docs/data-sources/metric#id DataLaunchdarklyMetric#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the metric is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.16.0/docs/data-sources/metric#is_active DataLaunchdarklyMetric#is_active}
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
	// Whether the metric is numeric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.16.0/docs/data-sources/metric#is_numeric DataLaunchdarklyMetric#is_numeric}
	IsNumeric interface{} `field:"optional" json:"isNumeric" yaml:"isNumeric"`
	// The metric type -available choices are 'pageview', 'click', and 'custom'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.16.0/docs/data-sources/metric#kind DataLaunchdarklyMetric#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// The LaunchDarkly ID of the user who will maintain the metric.
	//
	// If not set, the API will automatically apply the member associated with your Terraform API key or the most recently-set maintainer
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.16.0/docs/data-sources/metric#maintainer_id DataLaunchdarklyMetric#maintainer_id}
	MaintainerId *string `field:"optional" json:"maintainerId" yaml:"maintainerId"`
	// A human-readable name for your metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.16.0/docs/data-sources/metric#name DataLaunchdarklyMetric#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A set of one or more context kinds that this metric can measure events from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.16.0/docs/data-sources/metric#randomization_units DataLaunchdarklyMetric#randomization_units}
	RandomizationUnits *[]*string `field:"optional" json:"randomizationUnits" yaml:"randomizationUnits"`
	// The CSS selector for your metric (if click metric).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.16.0/docs/data-sources/metric#selector DataLaunchdarklyMetric#selector}
	Selector *string `field:"optional" json:"selector" yaml:"selector"`
	// The success criteria for your metric (if numeric metric).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.16.0/docs/data-sources/metric#success_criteria DataLaunchdarklyMetric#success_criteria}
	SuccessCriteria *string `field:"optional" json:"successCriteria" yaml:"successCriteria"`
	// The unit for your metric (if numeric metric).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.16.0/docs/data-sources/metric#unit DataLaunchdarklyMetric#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// urls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.16.0/docs/data-sources/metric#urls DataLaunchdarklyMetric#urls}
	Urls interface{} `field:"optional" json:"urls" yaml:"urls"`
}

