// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package accesstoken

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessTokenPolicyStatementsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccessTokenPolicyStatementsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessTokenPolicyStatementsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessTokenPolicyStatementsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessTokenPolicyStatementsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessTokenPolicyStatementsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessTokenPolicyStatementsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessTokenPolicyStatementsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

