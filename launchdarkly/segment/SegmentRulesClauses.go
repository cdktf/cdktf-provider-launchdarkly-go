// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package segment


type SegmentRulesClauses struct {
	// The user attribute to operate on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.2/docs/resources/segment#attribute Segment#attribute}
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// The operator associated with the rule clause.
	//
	// Available options are `in`, `endsWith`, `startsWith`, `matches`, `contains`, `lessThan`, `lessThanOrEqual`, `greaterThanOrEqual`, `before`, `after`, `segmentMatch`, `semVerEqual`, `semVerLessThan`, and `semVerGreaterThan`. Read LaunchDarkly's [Operators](https://docs.launchdarkly.com/sdk/concepts/flag-evaluation-rules#operators) documentation for more information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.2/docs/resources/segment#op Segment#op}
	Op *string `field:"required" json:"op" yaml:"op"`
	// The list of values associated with the rule clause.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.2/docs/resources/segment#values Segment#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// The context kind associated with this rule clause. If omitted, defaults to `user`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.2/docs/resources/segment#context_kind Segment#context_kind}
	ContextKind *string `field:"optional" json:"contextKind" yaml:"contextKind"`
	// Whether to negate the rule clause.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.2/docs/resources/segment#negate Segment#negate}
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
	// The type for each of the clause's values.
	//
	// Available types are `boolean`, `string`, and `number`. If omitted, `value_type` defaults to `string`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.2/docs/resources/segment#value_type Segment#value_type}
	ValueType *string `field:"optional" json:"valueType" yaml:"valueType"`
}

