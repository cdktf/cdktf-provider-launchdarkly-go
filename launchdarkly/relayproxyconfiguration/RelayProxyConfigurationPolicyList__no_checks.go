//go:build no_runtime_type_checking

package relayproxyconfiguration

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RelayProxyConfigurationPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RelayProxyConfigurationPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RelayProxyConfigurationPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RelayProxyConfigurationPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RelayProxyConfigurationPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RelayProxyConfigurationPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRelayProxyConfigurationPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

