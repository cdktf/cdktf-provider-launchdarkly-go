package datalaunchdarklyfeatureflagenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/jsii"

	"github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/datalaunchdarklyfeatureflagenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference interface {
	cdktf.ComplexObject
	BucketBy() *string
	SetBucketBy(val *string)
	BucketByInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ContextKind() *string
	SetContextKind(val *string)
	ContextKindInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataLaunchdarklyFeatureFlagEnvironmentFallthrough
	SetInternalValue(val *DataLaunchdarklyFeatureFlagEnvironmentFallthrough)
	RolloutWeights() *[]*float64
	SetRolloutWeights(val *[]*float64)
	RolloutWeightsInput() *[]*float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Variation() *float64
	SetVariation(val *float64)
	VariationInput() *float64
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetBucketBy()
	ResetContextKind()
	ResetRolloutWeights()
	ResetVariation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference
type jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) BucketBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) BucketByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) ContextKind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextKind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) ContextKindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextKindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) InternalValue() *DataLaunchdarklyFeatureFlagEnvironmentFallthrough {
	var returns *DataLaunchdarklyFeatureFlagEnvironmentFallthrough
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) RolloutWeights() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"rolloutWeights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) RolloutWeightsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"rolloutWeightsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) Variation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"variation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) VariationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"variationInput",
		&returns,
	)
	return returns
}


func NewDataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference {
	_init_.Initialize()

	if err := validateNewDataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.dataLaunchdarklyFeatureFlagEnvironment.DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference_Override(d DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.dataLaunchdarklyFeatureFlagEnvironment.DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference)SetBucketBy(val *string) {
	if err := j.validateSetBucketByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketBy",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference)SetContextKind(val *string) {
	if err := j.validateSetContextKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contextKind",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference)SetInternalValue(val *DataLaunchdarklyFeatureFlagEnvironmentFallthrough) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference)SetRolloutWeights(val *[]*float64) {
	if err := j.validateSetRolloutWeightsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rolloutWeights",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference)SetVariation(val *float64) {
	if err := j.validateSetVariationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variation",
		val,
	)
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) ResetBucketBy() {
	_jsii_.InvokeVoid(
		d,
		"resetBucketBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) ResetContextKind() {
	_jsii_.InvokeVoid(
		d,
		"resetContextKind",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) ResetRolloutWeights() {
	_jsii_.InvokeVoid(
		d,
		"resetRolloutWeights",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) ResetVariation() {
	_jsii_.InvokeVoid(
		d,
		"resetVariation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLaunchdarklyFeatureFlagEnvironmentFallthroughOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

