// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package metric

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/v4/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/v4/metric/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.0/docs/resources/metric launchdarkly_metric}.
type Metric interface {
	cdktf.TerraformResource
	AnalysisType() *string
	SetAnalysisType(val *string)
	AnalysisTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	IncludeUnitsWithoutEvents() interface{}
	SetIncludeUnitsWithoutEvents(val interface{})
	IncludeUnitsWithoutEventsInput() interface{}
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
	PercentileValue() *float64
	SetPercentileValue(val *float64)
	PercentileValueInput() *float64
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
	UnitAggregationType() *string
	SetUnitAggregationType(val *string)
	UnitAggregationTypeInput() *string
	UnitInput() *string
	Urls() MetricUrlsList
	UrlsInput() interface{}
	Version() *float64
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
	PutUrls(value interface{})
	ResetAnalysisType()
	ResetDescription()
	ResetEventKey()
	ResetId()
	ResetIncludeUnitsWithoutEvents()
	ResetIsActive()
	ResetIsNumeric()
	ResetMaintainerId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPercentileValue()
	ResetRandomizationUnits()
	ResetSelector()
	ResetSuccessCriteria()
	ResetTags()
	ResetUnit()
	ResetUnitAggregationType()
	ResetUrls()
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

// The jsii proxy struct for Metric
type jsiiProxy_Metric struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Metric) AnalysisType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysisType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) AnalysisTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysisTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) EventKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) EventKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) IncludeUnitsWithoutEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeUnitsWithoutEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) IncludeUnitsWithoutEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeUnitsWithoutEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) IsActive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isActive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) IsActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isActiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) IsNumeric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNumeric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) IsNumericInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNumericInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) MaintainerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) MaintainerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) PercentileValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentileValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) PercentileValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentileValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) ProjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) ProjectKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) RandomizationUnits() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"randomizationUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) RandomizationUnitsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"randomizationUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Selector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) SelectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) SuccessCriteria() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) SuccessCriteriaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) UnitAggregationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitAggregationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) UnitAggregationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitAggregationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Urls() MetricUrlsList {
	var returns MetricUrlsList
	_jsii_.Get(
		j,
		"urls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) UrlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metric) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.0/docs/resources/metric launchdarkly_metric} Resource.
func NewMetric(scope constructs.Construct, id *string, config *MetricConfig) Metric {
	_init_.Initialize()

	if err := validateNewMetricParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Metric{}

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.metric.Metric",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.21.0/docs/resources/metric launchdarkly_metric} Resource.
func NewMetric_Override(m Metric, scope constructs.Construct, id *string, config *MetricConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.metric.Metric",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_Metric)SetAnalysisType(val *string) {
	if err := j.validateSetAnalysisTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analysisType",
		val,
	)
}

func (j *jsiiProxy_Metric)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Metric)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Metric)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Metric)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Metric)SetEventKey(val *string) {
	if err := j.validateSetEventKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventKey",
		val,
	)
}

func (j *jsiiProxy_Metric)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Metric)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Metric)SetIncludeUnitsWithoutEvents(val interface{}) {
	if err := j.validateSetIncludeUnitsWithoutEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeUnitsWithoutEvents",
		val,
	)
}

func (j *jsiiProxy_Metric)SetIsActive(val interface{}) {
	if err := j.validateSetIsActiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isActive",
		val,
	)
}

func (j *jsiiProxy_Metric)SetIsNumeric(val interface{}) {
	if err := j.validateSetIsNumericParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isNumeric",
		val,
	)
}

func (j *jsiiProxy_Metric)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_Metric)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_Metric)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Metric)SetMaintainerId(val *string) {
	if err := j.validateSetMaintainerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintainerId",
		val,
	)
}

func (j *jsiiProxy_Metric)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Metric)SetPercentileValue(val *float64) {
	if err := j.validateSetPercentileValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"percentileValue",
		val,
	)
}

func (j *jsiiProxy_Metric)SetProjectKey(val *string) {
	if err := j.validateSetProjectKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectKey",
		val,
	)
}

func (j *jsiiProxy_Metric)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Metric)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Metric)SetRandomizationUnits(val *[]*string) {
	if err := j.validateSetRandomizationUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"randomizationUnits",
		val,
	)
}

func (j *jsiiProxy_Metric)SetSelector(val *string) {
	if err := j.validateSetSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selector",
		val,
	)
}

func (j *jsiiProxy_Metric)SetSuccessCriteria(val *string) {
	if err := j.validateSetSuccessCriteriaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successCriteria",
		val,
	)
}

func (j *jsiiProxy_Metric)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Metric)SetUnit(val *string) {
	if err := j.validateSetUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (j *jsiiProxy_Metric)SetUnitAggregationType(val *string) {
	if err := j.validateSetUnitAggregationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unitAggregationType",
		val,
	)
}

// Generates CDKTF code for importing a Metric resource upon running "cdktf plan <stack-name>".
func Metric_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMetric_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.metric.Metric",
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
func Metric_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMetric_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.metric.Metric",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Metric_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMetric_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.metric.Metric",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Metric_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMetric_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.metric.Metric",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Metric_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-launchdarkly.metric.Metric",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_Metric) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_Metric) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_Metric) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_Metric) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_Metric) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_Metric) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_Metric) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_Metric) PutUrls(value interface{}) {
	if err := m.validatePutUrlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putUrls",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Metric) ResetAnalysisType() {
	_jsii_.InvokeVoid(
		m,
		"resetAnalysisType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metric) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metric) ResetEventKey() {
	_jsii_.InvokeVoid(
		m,
		"resetEventKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metric) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metric) ResetIncludeUnitsWithoutEvents() {
	_jsii_.InvokeVoid(
		m,
		"resetIncludeUnitsWithoutEvents",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metric) ResetIsActive() {
	_jsii_.InvokeVoid(
		m,
		"resetIsActive",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metric) ResetIsNumeric() {
	_jsii_.InvokeVoid(
		m,
		"resetIsNumeric",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metric) ResetMaintainerId() {
	_jsii_.InvokeVoid(
		m,
		"resetMaintainerId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metric) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metric) ResetPercentileValue() {
	_jsii_.InvokeVoid(
		m,
		"resetPercentileValue",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metric) ResetRandomizationUnits() {
	_jsii_.InvokeVoid(
		m,
		"resetRandomizationUnits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metric) ResetSelector() {
	_jsii_.InvokeVoid(
		m,
		"resetSelector",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metric) ResetSuccessCriteria() {
	_jsii_.InvokeVoid(
		m,
		"resetSuccessCriteria",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metric) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metric) ResetUnit() {
	_jsii_.InvokeVoid(
		m,
		"resetUnit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metric) ResetUnitAggregationType() {
	_jsii_.InvokeVoid(
		m,
		"resetUnitAggregationType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metric) ResetUrls() {
	_jsii_.InvokeVoid(
		m,
		"resetUrls",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metric) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metric) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

