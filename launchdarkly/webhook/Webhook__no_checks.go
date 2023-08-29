// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package webhook

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_Webhook) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (w *jsiiProxy_Webhook) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Webhook) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Webhook) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Webhook) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Webhook) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Webhook) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Webhook) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Webhook) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Webhook) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Webhook) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Webhook) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (w *jsiiProxy_Webhook) validatePutStatementsParameters(value interface{}) error {
	return nil
}

func validateWebhook_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWebhook_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateWebhook_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Webhook) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Webhook) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Webhook) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Webhook) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Webhook) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Webhook) validateSetOnParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Webhook) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Webhook) validateSetSecretParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Webhook) validateSetTagsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Webhook) validateSetUrlParameters(val *string) error {
	return nil
}

func validateNewWebhookParameters(scope constructs.Construct, id *string, config *WebhookConfig) error {
	return nil
}

