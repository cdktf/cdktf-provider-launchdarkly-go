//go:build no_runtime_type_checking

package segment

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SegmentExcludedContextsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SegmentExcludedContextsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SegmentExcludedContextsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SegmentExcludedContextsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SegmentExcludedContextsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SegmentExcludedContextsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSegmentExcludedContextsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

