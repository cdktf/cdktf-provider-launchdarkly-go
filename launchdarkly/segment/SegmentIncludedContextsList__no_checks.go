// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package segment

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SegmentIncludedContextsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SegmentIncludedContextsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SegmentIncludedContextsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SegmentIncludedContextsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SegmentIncludedContextsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SegmentIncludedContextsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SegmentIncludedContextsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSegmentIncludedContextsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

