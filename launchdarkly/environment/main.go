package environment

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-launchdarkly.environment.Environment",
		reflect.TypeOf((*Environment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiKey", GoGetter: "ApiKey"},
			_jsii_.MemberProperty{JsiiProperty: "approvalSettings", GoGetter: "ApprovalSettings"},
			_jsii_.MemberProperty{JsiiProperty: "approvalSettingsInput", GoGetter: "ApprovalSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientSideId", GoGetter: "ClientSideId"},
			_jsii_.MemberProperty{JsiiProperty: "color", GoGetter: "Color"},
			_jsii_.MemberProperty{JsiiProperty: "colorInput", GoGetter: "ColorInput"},
			_jsii_.MemberProperty{JsiiProperty: "confirmChanges", GoGetter: "ConfirmChanges"},
			_jsii_.MemberProperty{JsiiProperty: "confirmChangesInput", GoGetter: "ConfirmChangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTrackEvents", GoGetter: "DefaultTrackEvents"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTrackEventsInput", GoGetter: "DefaultTrackEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTtl", GoGetter: "DefaultTtl"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTtlInput", GoGetter: "DefaultTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mobileKey", GoGetter: "MobileKey"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "projectKey", GoGetter: "ProjectKey"},
			_jsii_.MemberProperty{JsiiProperty: "projectKeyInput", GoGetter: "ProjectKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putApprovalSettings", GoMethod: "PutApprovalSettings"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requireComments", GoGetter: "RequireComments"},
			_jsii_.MemberProperty{JsiiProperty: "requireCommentsInput", GoGetter: "RequireCommentsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetApprovalSettings", GoMethod: "ResetApprovalSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfirmChanges", GoMethod: "ResetConfirmChanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultTrackEvents", GoMethod: "ResetDefaultTrackEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultTtl", GoMethod: "ResetDefaultTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireComments", GoMethod: "ResetRequireComments"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecureMode", GoMethod: "ResetSecureMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberProperty{JsiiProperty: "secureMode", GoGetter: "SecureMode"},
			_jsii_.MemberProperty{JsiiProperty: "secureModeInput", GoGetter: "SecureModeInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_Environment{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-launchdarkly.environment.EnvironmentApprovalSettings",
		reflect.TypeOf((*EnvironmentApprovalSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-launchdarkly.environment.EnvironmentApprovalSettingsList",
		reflect.TypeOf((*EnvironmentApprovalSettingsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_EnvironmentApprovalSettingsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-launchdarkly.environment.EnvironmentApprovalSettingsOutputReference",
		reflect.TypeOf((*EnvironmentApprovalSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "canApplyDeclinedChanges", GoGetter: "CanApplyDeclinedChanges"},
			_jsii_.MemberProperty{JsiiProperty: "canApplyDeclinedChangesInput", GoGetter: "CanApplyDeclinedChangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "canReviewOwnRequest", GoGetter: "CanReviewOwnRequest"},
			_jsii_.MemberProperty{JsiiProperty: "canReviewOwnRequestInput", GoGetter: "CanReviewOwnRequestInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "minNumApprovals", GoGetter: "MinNumApprovals"},
			_jsii_.MemberProperty{JsiiProperty: "minNumApprovalsInput", GoGetter: "MinNumApprovalsInput"},
			_jsii_.MemberProperty{JsiiProperty: "required", GoGetter: "Required"},
			_jsii_.MemberProperty{JsiiProperty: "requiredApprovalTags", GoGetter: "RequiredApprovalTags"},
			_jsii_.MemberProperty{JsiiProperty: "requiredApprovalTagsInput", GoGetter: "RequiredApprovalTagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "requiredInput", GoGetter: "RequiredInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCanApplyDeclinedChanges", GoMethod: "ResetCanApplyDeclinedChanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetCanReviewOwnRequest", GoMethod: "ResetCanReviewOwnRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinNumApprovals", GoMethod: "ResetMinNumApprovals"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequired", GoMethod: "ResetRequired"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequiredApprovalTags", GoMethod: "ResetRequiredApprovalTags"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EnvironmentApprovalSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-launchdarkly.environment.EnvironmentConfig",
		reflect.TypeOf((*EnvironmentConfig)(nil)).Elem(),
	)
}
