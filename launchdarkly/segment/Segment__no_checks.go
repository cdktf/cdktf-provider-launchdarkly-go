// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package segment

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Segment) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_Segment) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Segment) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Segment) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Segment) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Segment) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Segment) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Segment) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Segment) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Segment) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Segment) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Segment) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Segment) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Segment) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_Segment) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_Segment) validatePutExcludedContextsParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_Segment) validatePutIncludedContextsParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_Segment) validatePutRulesParameters(value interface{}) error {
	return nil
}

func validateSegment_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSegment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSegment_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSegment_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Segment) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Segment) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Segment) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Segment) validateSetEnvKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Segment) validateSetExcludedParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Segment) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Segment) validateSetIncludedParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Segment) validateSetKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Segment) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Segment) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Segment) validateSetProjectKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Segment) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Segment) validateSetTagsParameters(val *[]*string) error {
	return nil
}

func validateNewSegmentParameters(scope constructs.Construct, id *string, config *SegmentConfig) error {
	return nil
}

