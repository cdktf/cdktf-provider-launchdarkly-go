// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package metric


type MetricUrls struct {
	// The url type - available choices are 'exact', 'canonical', 'substring' and 'regex'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.3/docs/resources/metric#kind Metric#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// The URL-matching regex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.3/docs/resources/metric#pattern Metric#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// The URL substring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.3/docs/resources/metric#substring Metric#substring}
	Substring *string `field:"optional" json:"substring" yaml:"substring"`
	// The exact or canonical URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.3/docs/resources/metric#url Metric#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

