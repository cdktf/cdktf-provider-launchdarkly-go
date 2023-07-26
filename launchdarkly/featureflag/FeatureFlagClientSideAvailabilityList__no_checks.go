//go:build no_runtime_type_checking

package featureflag

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FeatureFlagClientSideAvailabilityList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FeatureFlagClientSideAvailabilityList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagClientSideAvailabilityList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagClientSideAvailabilityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagClientSideAvailabilityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FeatureFlagClientSideAvailabilityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFeatureFlagClientSideAvailabilityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

