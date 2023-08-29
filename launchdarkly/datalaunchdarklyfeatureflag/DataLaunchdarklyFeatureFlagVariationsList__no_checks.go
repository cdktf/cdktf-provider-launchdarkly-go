// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datalaunchdarklyfeatureflag

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagVariationsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagVariationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagVariationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagVariationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagVariationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagVariationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataLaunchdarklyFeatureFlagVariationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

