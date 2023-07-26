package featureflagenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/jsii"

	"github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/featureflagenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FeatureFlagEnvironmentFallthroughOutputReference interface {
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
	InternalValue() *FeatureFlagEnvironmentFallthrough
	SetInternalValue(val *FeatureFlagEnvironmentFallthrough)
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

// The jsii proxy struct for FeatureFlagEnvironmentFallthroughOutputReference
type jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) BucketBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) BucketByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) ContextKind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextKind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) ContextKindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextKindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) InternalValue() *FeatureFlagEnvironmentFallthrough {
	var returns *FeatureFlagEnvironmentFallthrough
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) RolloutWeights() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"rolloutWeights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) RolloutWeightsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"rolloutWeightsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) Variation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"variation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) VariationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"variationInput",
		&returns,
	)
	return returns
}


func NewFeatureFlagEnvironmentFallthroughOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FeatureFlagEnvironmentFallthroughOutputReference {
	_init_.Initialize()

	if err := validateNewFeatureFlagEnvironmentFallthroughOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.featureFlagEnvironment.FeatureFlagEnvironmentFallthroughOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFeatureFlagEnvironmentFallthroughOutputReference_Override(f FeatureFlagEnvironmentFallthroughOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.featureFlagEnvironment.FeatureFlagEnvironmentFallthroughOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference)SetBucketBy(val *string) {
	if err := j.validateSetBucketByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketBy",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference)SetContextKind(val *string) {
	if err := j.validateSetContextKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contextKind",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference)SetInternalValue(val *FeatureFlagEnvironmentFallthrough) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference)SetRolloutWeights(val *[]*float64) {
	if err := j.validateSetRolloutWeightsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rolloutWeights",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference)SetVariation(val *float64) {
	if err := j.validateSetVariationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variation",
		val,
	)
}

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) ResetBucketBy() {
	_jsii_.InvokeVoid(
		f,
		"resetBucketBy",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) ResetContextKind() {
	_jsii_.InvokeVoid(
		f,
		"resetContextKind",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) ResetRolloutWeights() {
	_jsii_.InvokeVoid(
		f,
		"resetRolloutWeights",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) ResetVariation() {
	_jsii_.InvokeVoid(
		f,
		"resetVariation",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlagEnvironmentFallthroughOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

