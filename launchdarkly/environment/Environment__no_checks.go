// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package environment

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Environment) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (e *jsiiProxy_Environment) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Environment) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Environment) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Environment) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Environment) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Environment) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Environment) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Environment) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Environment) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Environment) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Environment) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (e *jsiiProxy_Environment) validatePutApprovalSettingsParameters(value interface{}) error {
	return nil
}

func validateEnvironment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEnvironment_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateEnvironment_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Environment) validateSetColorParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Environment) validateSetConfirmChangesParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Environment) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Environment) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Environment) validateSetDefaultTrackEventsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Environment) validateSetDefaultTtlParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Environment) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Environment) validateSetKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Environment) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Environment) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Environment) validateSetProjectKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Environment) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Environment) validateSetRequireCommentsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Environment) validateSetSecureModeParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Environment) validateSetTagsParameters(val *[]*string) error {
	return nil
}

func validateNewEnvironmentParameters(scope constructs.Construct, id *string, config *EnvironmentConfig) error {
	return nil
}

