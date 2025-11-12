// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package featureflagenvironment

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FeatureFlagEnvironmentPrerequisitesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FeatureFlagEnvironmentPrerequisitesList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FeatureFlagEnvironmentPrerequisitesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagEnvironmentPrerequisitesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagEnvironmentPrerequisitesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagEnvironmentPrerequisitesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagEnvironmentPrerequisitesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFeatureFlagEnvironmentPrerequisitesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

