package datalaunchdarklysegment


type DataLaunchdarklySegmentIncludedContexts struct {
	// The context kind associated with this segment target. To target on user contexts, use the included and excluded attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/data-sources/segment#context_kind DataLaunchdarklySegment#context_kind}
	ContextKind *string `field:"required" json:"contextKind" yaml:"contextKind"`
	// List of target object keys included in or excluded from the segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/data-sources/segment#values DataLaunchdarklySegment#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

