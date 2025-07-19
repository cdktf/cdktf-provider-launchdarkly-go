// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package metric

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MetricConfig struct {
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
	// The unique key that references the metric.
	//
	// A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#key Metric#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The metric type.
	//
	// Available choices are `click`, `custom`, and `pageview`. A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#kind Metric#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// The human-friendly name for the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#name Metric#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The metrics's project key.
	//
	// A change in this field will force the destruction of the existing resource and the creation of a new one. A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#project_key Metric#project_key}
	ProjectKey *string `field:"required" json:"projectKey" yaml:"projectKey"`
	// The method for analyzing metric events. Available choices are `mean` and `percentile`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#analysis_type Metric#analysis_type}
	AnalysisType *string `field:"optional" json:"analysisType" yaml:"analysisType"`
	// The description of the metric's purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#description Metric#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The event key for your metric (if custom metric).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#event_key Metric#event_key}
	EventKey *string `field:"optional" json:"eventKey" yaml:"eventKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#id Metric#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Include units that did not send any events and set their value to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#include_units_without_events Metric#include_units_without_events}
	IncludeUnitsWithoutEvents interface{} `field:"optional" json:"includeUnitsWithoutEvents" yaml:"includeUnitsWithoutEvents"`
	// Ignored. All metrics are considered active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#is_active Metric#is_active}
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
	// Whether a `custom` metric is a numeric metric or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#is_numeric Metric#is_numeric}
	IsNumeric interface{} `field:"optional" json:"isNumeric" yaml:"isNumeric"`
	// The LaunchDarkly member ID of the member who will maintain the metric.
	//
	// If not set, the API will automatically apply the member associated with your Terraform API key or the most recently-set maintainer
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#maintainer_id Metric#maintainer_id}
	MaintainerId *string `field:"optional" json:"maintainerId" yaml:"maintainerId"`
	// The percentile for the analysis method.
	//
	// An integer denoting the target percentile between 0 and 100. Required when analysis_type is percentile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#percentile_value Metric#percentile_value}
	PercentileValue *float64 `field:"optional" json:"percentileValue" yaml:"percentileValue"`
	// A set of one or more context kinds that this metric can measure events from.
	//
	// Metrics can only use context kinds marked as "Available for experiments." For more information, read [Allocating experiment audiences](https://docs.launchdarkly.com/home/creating-experiments/allocation).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#randomization_units Metric#randomization_units}
	RandomizationUnits *[]*string `field:"optional" json:"randomizationUnits" yaml:"randomizationUnits"`
	// The CSS selector for your metric (if click metric).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#selector Metric#selector}
	Selector *string `field:"optional" json:"selector" yaml:"selector"`
	// The success criteria for your metric (if numeric metric). Available choices are `HigherThanBaseline` and `LowerThanBaseline`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#success_criteria Metric#success_criteria}
	SuccessCriteria *string `field:"optional" json:"successCriteria" yaml:"successCriteria"`
	// Tags associated with your resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#tags Metric#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// (Required for kind `custom`) The unit for numeric `custom` metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#unit Metric#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// The method by which multiple unit event values are aggregated. Available choices are `average` and `sum`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#unit_aggregation_type Metric#unit_aggregation_type}
	UnitAggregationType *string `field:"optional" json:"unitAggregationType" yaml:"unitAggregationType"`
	// urls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.3/docs/resources/metric#urls Metric#urls}
	Urls interface{} `field:"optional" json:"urls" yaml:"urls"`
}

