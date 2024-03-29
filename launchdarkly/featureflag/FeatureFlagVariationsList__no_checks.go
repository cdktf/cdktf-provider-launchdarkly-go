// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package featureflag

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FeatureFlagVariationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FeatureFlagVariationsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FeatureFlagVariationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagVariationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagVariationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagVariationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagVariationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFeatureFlagVariationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

