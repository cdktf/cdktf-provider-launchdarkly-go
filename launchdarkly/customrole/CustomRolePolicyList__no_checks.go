// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package customrole

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomRolePolicyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CustomRolePolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CustomRolePolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CustomRolePolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomRolePolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomRolePolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CustomRolePolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCustomRolePolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

