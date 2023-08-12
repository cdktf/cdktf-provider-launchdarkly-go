package featureflagenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/featureflagenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag_environment launchdarkly_feature_flag_environment}.
type FeatureFlagEnvironment interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContextTargets() FeatureFlagEnvironmentContextTargetsList
	ContextTargetsInput() interface{}
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
	Fallthrough() FeatureFlagEnvironmentFallthroughOutputReference
	FallthroughInput() *FeatureFlagEnvironmentFallthrough
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
	SetOffVariation(val *float64)
	OffVariationInput() *float64
	On() interface{}
	SetOn(val interface{})
	OnInput() interface{}
	Prerequisites() FeatureFlagEnvironmentPrerequisitesList
	PrerequisitesInput() interface{}
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
	Rules() FeatureFlagEnvironmentRulesList
	RulesInput() interface{}
	Targets() FeatureFlagEnvironmentTargetsList
	TargetsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TrackEvents() interface{}
	SetTrackEvents(val interface{})
	TrackEventsInput() interface{}
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
	PutContextTargets(value interface{})
	PutFallthrough(value *FeatureFlagEnvironmentFallthrough)
	PutPrerequisites(value interface{})
	PutRules(value interface{})
	PutTargets(value interface{})
	ResetContextTargets()
	ResetId()
	ResetOn()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrerequisites()
	ResetRules()
	ResetTargets()
	ResetTrackEvents()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for FeatureFlagEnvironment
type jsiiProxy_FeatureFlagEnvironment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FeatureFlagEnvironment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) ContextTargets() FeatureFlagEnvironmentContextTargetsList {
	var returns FeatureFlagEnvironmentContextTargetsList
	_jsii_.Get(
		j,
		"contextTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) ContextTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contextTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) EnvKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"envKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) EnvKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"envKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) Fallthrough() FeatureFlagEnvironmentFallthroughOutputReference {
	var returns FeatureFlagEnvironmentFallthroughOutputReference
	_jsii_.Get(
		j,
		"fallthrough",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) FallthroughInput() *FeatureFlagEnvironmentFallthrough {
	var returns *FeatureFlagEnvironmentFallthrough
	_jsii_.Get(
		j,
		"fallthroughInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) FlagId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flagId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) FlagIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flagIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) OffVariation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offVariation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) OffVariationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offVariationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) On() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"on",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) OnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) Prerequisites() FeatureFlagEnvironmentPrerequisitesList {
	var returns FeatureFlagEnvironmentPrerequisitesList
	_jsii_.Get(
		j,
		"prerequisites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) PrerequisitesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prerequisitesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) Rules() FeatureFlagEnvironmentRulesList {
	var returns FeatureFlagEnvironmentRulesList
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) RulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) Targets() FeatureFlagEnvironmentTargetsList {
	var returns FeatureFlagEnvironmentTargetsList
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) TargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) TrackEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trackEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironment) TrackEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trackEventsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag_environment launchdarkly_feature_flag_environment} Resource.
func NewFeatureFlagEnvironment(scope constructs.Construct, id *string, config *FeatureFlagEnvironmentConfig) FeatureFlagEnvironment {
	_init_.Initialize()

	if err := validateNewFeatureFlagEnvironmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FeatureFlagEnvironment{}

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.featureFlagEnvironment.FeatureFlagEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag_environment launchdarkly_feature_flag_environment} Resource.
func NewFeatureFlagEnvironment_Override(f FeatureFlagEnvironment, scope constructs.Construct, id *string, config *FeatureFlagEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.featureFlagEnvironment.FeatureFlagEnvironment",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironment)SetEnvKey(val *string) {
	if err := j.validateSetEnvKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"envKey",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironment)SetFlagId(val *string) {
	if err := j.validateSetFlagIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flagId",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironment)SetOffVariation(val *float64) {
	if err := j.validateSetOffVariationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"offVariation",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironment)SetOn(val interface{}) {
	if err := j.validateSetOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"on",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironment)SetTrackEvents(val interface{}) {
	if err := j.validateSetTrackEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackEvents",
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
func FeatureFlagEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFeatureFlagEnvironment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.featureFlagEnvironment.FeatureFlagEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FeatureFlagEnvironment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFeatureFlagEnvironment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.featureFlagEnvironment.FeatureFlagEnvironment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FeatureFlagEnvironment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFeatureFlagEnvironment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.featureFlagEnvironment.FeatureFlagEnvironment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FeatureFlagEnvironment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-launchdarkly.featureFlagEnvironment.FeatureFlagEnvironment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FeatureFlagEnvironment) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FeatureFlagEnvironment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FeatureFlagEnvironment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FeatureFlagEnvironment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FeatureFlagEnvironment) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FeatureFlagEnvironment) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FeatureFlagEnvironment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FeatureFlagEnvironment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FeatureFlagEnvironment) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FeatureFlagEnvironment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FeatureFlagEnvironment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FeatureFlagEnvironment) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FeatureFlagEnvironment) PutContextTargets(value interface{}) {
	if err := f.validatePutContextTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putContextTargets",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureFlagEnvironment) PutFallthrough(value *FeatureFlagEnvironmentFallthrough) {
	if err := f.validatePutFallthroughParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putFallthrough",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureFlagEnvironment) PutPrerequisites(value interface{}) {
	if err := f.validatePutPrerequisitesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putPrerequisites",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureFlagEnvironment) PutRules(value interface{}) {
	if err := f.validatePutRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putRules",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureFlagEnvironment) PutTargets(value interface{}) {
	if err := f.validatePutTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTargets",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureFlagEnvironment) ResetContextTargets() {
	_jsii_.InvokeVoid(
		f,
		"resetContextTargets",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlagEnvironment) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlagEnvironment) ResetOn() {
	_jsii_.InvokeVoid(
		f,
		"resetOn",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlagEnvironment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlagEnvironment) ResetPrerequisites() {
	_jsii_.InvokeVoid(
		f,
		"resetPrerequisites",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlagEnvironment) ResetRules() {
	_jsii_.InvokeVoid(
		f,
		"resetRules",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlagEnvironment) ResetTargets() {
	_jsii_.InvokeVoid(
		f,
		"resetTargets",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlagEnvironment) ResetTrackEvents() {
	_jsii_.InvokeVoid(
		f,
		"resetTrackEvents",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlagEnvironment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlagEnvironment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlagEnvironment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlagEnvironment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

