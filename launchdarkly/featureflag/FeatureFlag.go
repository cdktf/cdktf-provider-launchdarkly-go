// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureflag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/v4/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/v4/featureflag/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.17.0/docs/resources/feature_flag launchdarkly_feature_flag}.
type FeatureFlag interface {
	cdktf.TerraformResource
	Archived() interface{}
	SetArchived(val interface{})
	ArchivedInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientSideAvailability() FeatureFlagClientSideAvailabilityList
	ClientSideAvailabilityInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomProperties() FeatureFlagCustomPropertiesList
	CustomPropertiesInput() interface{}
	Defaults() FeatureFlagDefaultsOutputReference
	DefaultsInput() *FeatureFlagDefaults
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IncludeInSnippet() interface{}
	SetIncludeInSnippet(val interface{})
	IncludeInSnippetInput() interface{}
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaintainerId() *string
	SetMaintainerId(val *string)
	MaintainerIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ProjectKey() *string
	SetProjectKey(val *string)
	ProjectKeyInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	Temporary() interface{}
	SetTemporary(val interface{})
	TemporaryInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Variations() FeatureFlagVariationsList
	VariationsInput() interface{}
	VariationType() *string
	SetVariationType(val *string)
	VariationTypeInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutClientSideAvailability(value interface{})
	PutCustomProperties(value interface{})
	PutDefaults(value *FeatureFlagDefaults)
	PutVariations(value interface{})
	ResetArchived()
	ResetClientSideAvailability()
	ResetCustomProperties()
	ResetDefaults()
	ResetDescription()
	ResetId()
	ResetIncludeInSnippet()
	ResetMaintainerId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTemporary()
	ResetVariations()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for FeatureFlag
type jsiiProxy_FeatureFlag struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FeatureFlag) Archived() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archived",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) ArchivedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archivedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) ClientSideAvailability() FeatureFlagClientSideAvailabilityList {
	var returns FeatureFlagClientSideAvailabilityList
	_jsii_.Get(
		j,
		"clientSideAvailability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) ClientSideAvailabilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSideAvailabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) CustomProperties() FeatureFlagCustomPropertiesList {
	var returns FeatureFlagCustomPropertiesList
	_jsii_.Get(
		j,
		"customProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) CustomPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) Defaults() FeatureFlagDefaultsOutputReference {
	var returns FeatureFlagDefaultsOutputReference
	_jsii_.Get(
		j,
		"defaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) DefaultsInput() *FeatureFlagDefaults {
	var returns *FeatureFlagDefaults
	_jsii_.Get(
		j,
		"defaultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) IncludeInSnippet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeInSnippet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) IncludeInSnippetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeInSnippetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) MaintainerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) MaintainerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) ProjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) ProjectKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) Temporary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"temporary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) TemporaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"temporaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) Variations() FeatureFlagVariationsList {
	var returns FeatureFlagVariationsList
	_jsii_.Get(
		j,
		"variations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) VariationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) VariationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlag) VariationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variationTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.17.0/docs/resources/feature_flag launchdarkly_feature_flag} Resource.
func NewFeatureFlag(scope constructs.Construct, id *string, config *FeatureFlagConfig) FeatureFlag {
	_init_.Initialize()

	if err := validateNewFeatureFlagParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FeatureFlag{}

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.featureFlag.FeatureFlag",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.17.0/docs/resources/feature_flag launchdarkly_feature_flag} Resource.
func NewFeatureFlag_Override(f FeatureFlag, scope constructs.Construct, id *string, config *FeatureFlagConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.featureFlag.FeatureFlag",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FeatureFlag)SetArchived(val interface{}) {
	if err := j.validateSetArchivedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archived",
		val,
	)
}

func (j *jsiiProxy_FeatureFlag)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FeatureFlag)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FeatureFlag)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FeatureFlag)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_FeatureFlag)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FeatureFlag)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FeatureFlag)SetIncludeInSnippet(val interface{}) {
	if err := j.validateSetIncludeInSnippetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeInSnippet",
		val,
	)
}

func (j *jsiiProxy_FeatureFlag)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_FeatureFlag)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FeatureFlag)SetMaintainerId(val *string) {
	if err := j.validateSetMaintainerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintainerId",
		val,
	)
}

func (j *jsiiProxy_FeatureFlag)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FeatureFlag)SetProjectKey(val *string) {
	if err := j.validateSetProjectKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectKey",
		val,
	)
}

func (j *jsiiProxy_FeatureFlag)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FeatureFlag)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_FeatureFlag)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FeatureFlag)SetTemporary(val interface{}) {
	if err := j.validateSetTemporaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"temporary",
		val,
	)
}

func (j *jsiiProxy_FeatureFlag)SetVariationType(val *string) {
	if err := j.validateSetVariationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variationType",
		val,
	)
}

// Generates CDKTF code for importing a FeatureFlag resource upon running "cdktf plan <stack-name>".
func FeatureFlag_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateFeatureFlag_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.featureFlag.FeatureFlag",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func FeatureFlag_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFeatureFlag_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.featureFlag.FeatureFlag",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FeatureFlag_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFeatureFlag_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.featureFlag.FeatureFlag",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FeatureFlag_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFeatureFlag_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.featureFlag.FeatureFlag",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FeatureFlag_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-launchdarkly.featureFlag.FeatureFlag",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FeatureFlag) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FeatureFlag) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FeatureFlag) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlag) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlag) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlag) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlag) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlag) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlag) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlag) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlag) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlag) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlag) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FeatureFlag) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlag) MoveFromId(id *string) {
	if err := f.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveFromId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FeatureFlag) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FeatureFlag) MoveToId(id *string) {
	if err := f.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveToId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FeatureFlag) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FeatureFlag) PutClientSideAvailability(value interface{}) {
	if err := f.validatePutClientSideAvailabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putClientSideAvailability",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureFlag) PutCustomProperties(value interface{}) {
	if err := f.validatePutCustomPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putCustomProperties",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureFlag) PutDefaults(value *FeatureFlagDefaults) {
	if err := f.validatePutDefaultsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putDefaults",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureFlag) PutVariations(value interface{}) {
	if err := f.validatePutVariationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putVariations",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureFlag) ResetArchived() {
	_jsii_.InvokeVoid(
		f,
		"resetArchived",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlag) ResetClientSideAvailability() {
	_jsii_.InvokeVoid(
		f,
		"resetClientSideAvailability",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlag) ResetCustomProperties() {
	_jsii_.InvokeVoid(
		f,
		"resetCustomProperties",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlag) ResetDefaults() {
	_jsii_.InvokeVoid(
		f,
		"resetDefaults",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlag) ResetDescription() {
	_jsii_.InvokeVoid(
		f,
		"resetDescription",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlag) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlag) ResetIncludeInSnippet() {
	_jsii_.InvokeVoid(
		f,
		"resetIncludeInSnippet",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlag) ResetMaintainerId() {
	_jsii_.InvokeVoid(
		f,
		"resetMaintainerId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlag) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlag) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlag) ResetTemporary() {
	_jsii_.InvokeVoid(
		f,
		"resetTemporary",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlag) ResetVariations() {
	_jsii_.InvokeVoid(
		f,
		"resetVariations",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlag) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlag) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlag) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlag) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlag) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlag) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

