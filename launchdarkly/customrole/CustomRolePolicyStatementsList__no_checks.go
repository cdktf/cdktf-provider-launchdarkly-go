//go:build no_runtime_type_checking

package customrole

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomRolePolicyStatementsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CustomRolePolicyStatementsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CustomRolePolicyStatementsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomRolePolicyStatementsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomRolePolicyStatementsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CustomRolePolicyStatementsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCustomRolePolicyStatementsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

