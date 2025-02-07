// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package team

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TeamRoleAttributesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TeamRoleAttributesList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TeamRoleAttributesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TeamRoleAttributesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TeamRoleAttributesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TeamRoleAttributesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TeamRoleAttributesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTeamRoleAttributesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

