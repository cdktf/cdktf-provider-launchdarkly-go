package datalaunchdarklymetric

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/datalaunchdarklymetric/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/metric launchdarkly_metric}.
type DataLaunchdarklyMetric interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EventKey() *string
	SetEventKey(val *string)
	EventKeyInput() *string
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
	IsActive() interface{}
	SetIsActive(val interface{})
	IsActiveInput() interface{}
	IsNumeric() interface{}
	SetIsNumeric(val interface{})
	IsNumericInput() interface{}
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	Kind() *string
	SetKind(val *string)
	KindInput() *string
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
	RandomizationUnits() *[]*string
	SetRandomizationUnits(val *[]*string)
	RandomizationUnitsInput() *[]*string
	// Experimental.
	RawOverrides() interface{}
	Selector() *string
	SetSelector(val *string)
	SelectorInput() *string
	SuccessCriteria() *string
	SetSuccessCriteria(val *string)
	SuccessCriteriaInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	Urls() DataLaunchdarklyMetricUrlsList
	UrlsInput() interface{}
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
	PutUrls(value interface{})
	ResetDescription()
	ResetEventKey()
	ResetId()
	ResetIsActive()
	ResetIsNumeric()
	ResetKind()
	ResetMaintainerId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRandomizationUnits()
	ResetSelector()
	ResetSuccessCriteria()
	ResetTags()
	ResetUnit()
	ResetUrls()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataLaunchdarklyMetric
type jsiiProxy_DataLaunchdarklyMetric struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataLaunchdarklyMetric) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) EventKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) EventKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) IsActive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isActive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) IsActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isActiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) IsNumeric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNumeric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) IsNumericInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNumericInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) MaintainerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) MaintainerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) ProjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) ProjectKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) RandomizationUnits() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"randomizationUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) RandomizationUnitsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"randomizationUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) Selector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) SelectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) SuccessCriteria() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) SuccessCriteriaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) Urls() DataLaunchdarklyMetricUrlsList {
	var returns DataLaunchdarklyMetricUrlsList
	_jsii_.Get(
		j,
		"urls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyMetric) UrlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/metric launchdarkly_metric} Data Source.
func NewDataLaunchdarklyMetric(scope constructs.Construct, id *string, config *DataLaunchdarklyMetricConfig) DataLaunchdarklyMetric {
	_init_.Initialize()

	if err := validateNewDataLaunchdarklyMetricParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLaunchdarklyMetric{}

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.dataLaunchdarklyMetric.DataLaunchdarklyMetric",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/metric launchdarkly_metric} Data Source.
func NewDataLaunchdarklyMetric_Override(d DataLaunchdarklyMetric, scope constructs.Construct, id *string, config *DataLaunchdarklyMetricConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.dataLaunchdarklyMetric.DataLaunchdarklyMetric",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetEventKey(val *string) {
	if err := j.validateSetEventKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventKey",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetIsActive(val interface{}) {
	if err := j.validateSetIsActiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isActive",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetIsNumeric(val interface{}) {
	if err := j.validateSetIsNumericParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isNumeric",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetMaintainerId(val *string) {
	if err := j.validateSetMaintainerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintainerId",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetProjectKey(val *string) {
	if err := j.validateSetProjectKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectKey",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetRandomizationUnits(val *[]*string) {
	if err := j.validateSetRandomizationUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"randomizationUnits",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetSelector(val *string) {
	if err := j.validateSetSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selector",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetSuccessCriteria(val *string) {
	if err := j.validateSetSuccessCriteriaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successCriteria",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyMetric)SetUnit(val *string) {
	if err := j.validateSetUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unit",
		val,
	)
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
func DataLaunchdarklyMetric_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataLaunchdarklyMetric_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.dataLaunchdarklyMetric.DataLaunchdarklyMetric",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataLaunchdarklyMetric_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataLaunchdarklyMetric_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.dataLaunchdarklyMetric.DataLaunchdarklyMetric",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataLaunchdarklyMetric_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataLaunchdarklyMetric_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.dataLaunchdarklyMetric.DataLaunchdarklyMetric",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataLaunchdarklyMetric_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-launchdarkly.dataLaunchdarklyMetric.DataLaunchdarklyMetric",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataLaunchdarklyMetric) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLaunchdarklyMetric) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLaunchdarklyMetric) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLaunchdarklyMetric) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLaunchdarklyMetric) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLaunchdarklyMetric) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLaunchdarklyMetric) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLaunchdarklyMetric) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLaunchdarklyMetric) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLaunchdarklyMetric) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLaunchdarklyMetric) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) PutUrls(value interface{}) {
	if err := d.validatePutUrlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUrls",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ResetEventKey() {
	_jsii_.InvokeVoid(
		d,
		"resetEventKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ResetIsActive() {
	_jsii_.InvokeVoid(
		d,
		"resetIsActive",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ResetIsNumeric() {
	_jsii_.InvokeVoid(
		d,
		"resetIsNumeric",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ResetKind() {
	_jsii_.InvokeVoid(
		d,
		"resetKind",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ResetMaintainerId() {
	_jsii_.InvokeVoid(
		d,
		"resetMaintainerId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ResetRandomizationUnits() {
	_jsii_.InvokeVoid(
		d,
		"resetRandomizationUnits",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ResetSelector() {
	_jsii_.InvokeVoid(
		d,
		"resetSelector",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ResetSuccessCriteria() {
	_jsii_.InvokeVoid(
		d,
		"resetSuccessCriteria",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ResetUnit() {
	_jsii_.InvokeVoid(
		d,
		"resetUnit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ResetUrls() {
	_jsii_.InvokeVoid(
		d,
		"resetUrls",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyMetric) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyMetric) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

