//go:build no_runtime_type_checking

package featureflagenvironment

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FeatureFlagEnvironmentTargetsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FeatureFlagEnvironmentTargetsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagEnvironmentTargetsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagEnvironmentTargetsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagEnvironmentTargetsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagEnvironmentTargetsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFeatureFlagEnvironmentTargetsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

