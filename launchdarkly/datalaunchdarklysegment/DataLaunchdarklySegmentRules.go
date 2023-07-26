package datalaunchdarklysegment


type DataLaunchdarklySegmentRules struct {
	// The attribute by which to group users together.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.3/docs/data-sources/segment#bucket_by DataLaunchdarklySegment#bucket_by}
	BucketBy *string `field:"optional" json:"bucketBy" yaml:"bucketBy"`
	// clauses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.3/docs/data-sources/segment#clauses DataLaunchdarklySegment#clauses}
	Clauses interface{} `field:"optional" json:"clauses" yaml:"clauses"`
	// The context kind associated with this segment rule.
	//
	// This argument is only valid if weight is also specified. If omitted, defaults to 'user'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.3/docs/data-sources/segment#rollout_context_kind DataLaunchdarklySegment#rollout_context_kind}
	RolloutContextKind *string `field:"optional" json:"rolloutContextKind" yaml:"rolloutContextKind"`
	// The integer weight of the rule (between 1 and 100000).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.3/docs/data-sources/segment#weight DataLaunchdarklySegment#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

