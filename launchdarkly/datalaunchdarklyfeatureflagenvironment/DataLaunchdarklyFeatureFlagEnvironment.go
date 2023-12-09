// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalaunchdarklyfeatureflagenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/v3/datalaunchdarklyfeatureflagenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.17.0/docs/data-sources/feature_flag_environment launchdarkly_feature_flag_environment}.
type DataLaunchdarklyFeatureFlagEnvironment interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContextTargets() DataLaunchdarklyFeatureFlagEnvironmentContextTargetsList
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnvKey() *string
	SetEnvKey(val *string)
	EnvKeyInput() *string
	Fallthrough() DataLaunchdarklyFeatureFlagEnvironmentFallthroughList
	FlagId() *string
	SetFlagId(val *string)
	FlagIdInput() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OffVariation() *float64
	On() cdktf.IResolvable
	Prerequisites() DataLaunchdarklyFeatureFlagEnvironmentPrerequisitesList
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Rules() DataLaunchdarklyFeatureFlagEnvironmentRulesList
	Targets() DataLaunchdarklyFeatureFlagEnvironmentTargetsList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TrackEvents() cdktf.IResolvable
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataLaunchdarklyFeatureFlagEnvironment
type jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) ContextTargets() DataLaunchdarklyFeatureFlagEnvironmentContextTargetsList {
	var returns DataLaunchdarklyFeatureFlagEnvironmentContextTargetsList
	_jsii_.Get(
		j,
		"contextTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) EnvKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"envKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) EnvKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"envKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) Fallthrough() DataLaunchdarklyFeatureFlagEnvironmentFallthroughList {
	var returns DataLaunchdarklyFeatureFlagEnvironmentFallthroughList
	_jsii_.Get(
		j,
		"fallthrough",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) FlagId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flagId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) FlagIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flagIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) OffVariation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offVariation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) On() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"on",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) Prerequisites() DataLaunchdarklyFeatureFlagEnvironmentPrerequisitesList {
	var returns DataLaunchdarklyFeatureFlagEnvironmentPrerequisitesList
	_jsii_.Get(
		j,
		"prerequisites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) Rules() DataLaunchdarklyFeatureFlagEnvironmentRulesList {
	var returns DataLaunchdarklyFeatureFlagEnvironmentRulesList
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) Targets() DataLaunchdarklyFeatureFlagEnvironmentTargetsList {
	var returns DataLaunchdarklyFeatureFlagEnvironmentTargetsList
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) TrackEvents() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"trackEvents",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.17.0/docs/data-sources/feature_flag_environment launchdarkly_feature_flag_environment} Data Source.
func NewDataLaunchdarklyFeatureFlagEnvironment(scope constructs.Construct, id *string, config *DataLaunchdarklyFeatureFlagEnvironmentConfig) DataLaunchdarklyFeatureFlagEnvironment {
	_init_.Initialize()

	if err := validateNewDataLaunchdarklyFeatureFlagEnvironmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment{}

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.dataLaunchdarklyFeatureFlagEnvironment.DataLaunchdarklyFeatureFlagEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.17.0/docs/data-sources/feature_flag_environment launchdarkly_feature_flag_environment} Data Source.
func NewDataLaunchdarklyFeatureFlagEnvironment_Override(d DataLaunchdarklyFeatureFlagEnvironment, scope constructs.Construct, id *string, config *DataLaunchdarklyFeatureFlagEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.dataLaunchdarklyFeatureFlagEnvironment.DataLaunchdarklyFeatureFlagEnvironment",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment)SetEnvKey(val *string) {
	if err := j.validateSetEnvKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"envKey",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment)SetFlagId(val *string) {
	if err := j.validateSetFlagIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flagId",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataLaunchdarklyFeatureFlagEnvironment resource upon running "cdktf plan <stack-name>".
func DataLaunchdarklyFeatureFlagEnvironment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataLaunchdarklyFeatureFlagEnvironment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.dataLaunchdarklyFeatureFlagEnvironment.DataLaunchdarklyFeatureFlagEnvironment",
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
func DataLaunchdarklyFeatureFlagEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataLaunchdarklyFeatureFlagEnvironment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.dataLaunchdarklyFeatureFlagEnvironment.DataLaunchdarklyFeatureFlagEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataLaunchdarklyFeatureFlagEnvironment_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataLaunchdarklyFeatureFlagEnvironment_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.dataLaunchdarklyFeatureFlagEnvironment.DataLaunchdarklyFeatureFlagEnvironment",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataLaunchdarklyFeatureFlagEnvironment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataLaunchdarklyFeatureFlagEnvironment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.dataLaunchdarklyFeatureFlagEnvironment.DataLaunchdarklyFeatureFlagEnvironment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataLaunchdarklyFeatureFlagEnvironment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-launchdarkly.dataLaunchdarklyFeatureFlagEnvironment.DataLaunchdarklyFeatureFlagEnvironment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

