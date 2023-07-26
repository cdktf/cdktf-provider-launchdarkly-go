//go:build no_runtime_type_checking

package webhook

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebhookStatementsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WebhookStatementsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WebhookStatementsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WebhookStatementsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WebhookStatementsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WebhookStatementsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWebhookStatementsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

