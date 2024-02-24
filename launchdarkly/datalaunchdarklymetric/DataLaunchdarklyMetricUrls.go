// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalaunchdarklymetric


type DataLaunchdarklyMetricUrls struct {
	// The url type - available choices are 'exact', 'canonical', 'substring' and 'regex'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.0/docs/data-sources/metric#kind DataLaunchdarklyMetric#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// The URL-matching regex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.0/docs/data-sources/metric#pattern DataLaunchdarklyMetric#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// The URL substring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.0/docs/data-sources/metric#substring DataLaunchdarklyMetric#substring}
	Substring *string `field:"optional" json:"substring" yaml:"substring"`
	// The exact or canonical URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.18.0/docs/data-sources/metric#url DataLaunchdarklyMetric#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

