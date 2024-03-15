// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package segment


type SegmentIncludedContexts struct {
	// The context kind associated with this segment target. To target on user contexts, use the included and excluded attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/segment#context_kind Segment#context_kind}
	ContextKind *string `field:"required" json:"contextKind" yaml:"contextKind"`
	// List of target object keys included in or excluded from the segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.1/docs/resources/segment#values Segment#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

