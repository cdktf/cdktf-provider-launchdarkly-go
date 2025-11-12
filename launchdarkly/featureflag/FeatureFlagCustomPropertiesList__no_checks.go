// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package featureflag

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FeatureFlagCustomPropertiesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FeatureFlagCustomPropertiesList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FeatureFlagCustomPropertiesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagCustomPropertiesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagCustomPropertiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagCustomPropertiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagCustomPropertiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFeatureFlagCustomPropertiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

