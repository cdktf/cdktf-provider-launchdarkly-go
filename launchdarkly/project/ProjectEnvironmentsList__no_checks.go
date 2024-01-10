// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package project

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProjectEnvironmentsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProjectEnvironmentsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProjectEnvironmentsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProjectEnvironmentsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ProjectEnvironmentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProjectEnvironmentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProjectEnvironmentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProjectEnvironmentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

