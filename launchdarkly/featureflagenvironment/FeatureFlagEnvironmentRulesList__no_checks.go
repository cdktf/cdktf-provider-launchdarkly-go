//go:build no_runtime_type_checking

package featureflagenvironment

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FeatureFlagEnvironmentRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FeatureFlagEnvironmentRulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagEnvironmentRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagEnvironmentRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagEnvironmentRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagEnvironmentRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFeatureFlagEnvironmentRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

