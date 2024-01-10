// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package environment

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EnvironmentApprovalSettingsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EnvironmentApprovalSettingsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EnvironmentApprovalSettingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EnvironmentApprovalSettingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EnvironmentApprovalSettingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EnvironmentApprovalSettingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EnvironmentApprovalSettingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEnvironmentApprovalSettingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

