//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LaunchdarklyProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (l *jsiiProxy_LaunchdarklyProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateLaunchdarklyProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLaunchdarklyProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateLaunchdarklyProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func validateNewLaunchdarklyProviderParameters(scope constructs.Construct, id *string, config *LaunchdarklyProviderConfig) error {
	return nil
}

