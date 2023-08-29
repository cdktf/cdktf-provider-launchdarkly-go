// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package metric

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MetricUrlsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MetricUrlsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MetricUrlsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MetricUrlsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MetricUrlsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MetricUrlsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMetricUrlsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

