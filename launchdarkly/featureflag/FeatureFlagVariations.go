// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureflag


type FeatureFlagVariations struct {
	// The variation value.
	//
	// The value's type must correspond to the `variation_type` argument. For example: `variation_type = "boolean"` accepts only `true` or `false`. The `number` variation type accepts both floats and ints, but please note that any trailing zeroes on floats will be trimmed (i.e. `1.1` and `1.100` will both be converted to `1.1`).
	//
	// If you wish to define an empty string variation, you must still define the value field on the variations block like so:
	//
	// ```terraform
	// variations {
	//   value = ""
	// }
	// ```
	//
	// -> **Note:** Terraform manages `variations` as an ordered array and identifies them by index. This means that if you change the order of your `variations` block, you may end up destroying and recreating those variations. Additionally, if you delete variations that have targets that have been attached outside of Terraform, those targets may be incorrectly reassigned to a different variation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.22.0/docs/resources/feature_flag#value FeatureFlag#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// The variation's description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.22.0/docs/resources/feature_flag#description FeatureFlag#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the variation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.22.0/docs/resources/feature_flag#name FeatureFlag#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

