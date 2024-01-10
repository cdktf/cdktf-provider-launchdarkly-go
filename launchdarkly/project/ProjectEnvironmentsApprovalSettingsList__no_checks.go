// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package project

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProjectEnvironmentsApprovalSettingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

