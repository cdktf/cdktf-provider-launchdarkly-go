// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package environment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/v3/jsii"

	"github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/v3/environment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EnvironmentApprovalSettingsOutputReference interface {
	cdktf.ComplexObject
	CanApplyDeclinedChanges() interface{}
	SetCanApplyDeclinedChanges(val interface{})
	CanApplyDeclinedChangesInput() interface{}
	CanReviewOwnRequest() interface{}
	SetCanReviewOwnRequest(val interface{})
	CanReviewOwnRequestInput() interface{}
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MinNumApprovals() *float64
	SetMinNumApprovals(val *float64)
	MinNumApprovalsInput() *float64
	Required() interface{}
	SetRequired(val interface{})
	RequiredApprovalTags() *[]*string
	SetRequiredApprovalTags(val *[]*string)
	RequiredApprovalTagsInput() *[]*string
	RequiredInput() interface{}
	ServiceConfig() *map[string]*string
	SetServiceConfig(val *map[string]*string)
	ServiceConfigInput() *map[string]*string
	ServiceKind() *string
	SetServiceKind(val *string)
	ServiceKindInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetCanApplyDeclinedChanges()
	ResetCanReviewOwnRequest()
	ResetMinNumApprovals()
	ResetRequired()
	ResetRequiredApprovalTags()
	ResetServiceConfig()
	ResetServiceKind()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EnvironmentApprovalSettingsOutputReference
type jsiiProxy_EnvironmentApprovalSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) CanApplyDeclinedChanges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canApplyDeclinedChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) CanApplyDeclinedChangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canApplyDeclinedChangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) CanReviewOwnRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canReviewOwnRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) CanReviewOwnRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canReviewOwnRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) MinNumApprovals() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumApprovals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) MinNumApprovalsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumApprovalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) Required() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) RequiredApprovalTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredApprovalTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) RequiredApprovalTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredApprovalTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) RequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) ServiceConfig() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"serviceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) ServiceConfigInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"serviceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) ServiceKind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) ServiceKindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEnvironmentApprovalSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EnvironmentApprovalSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewEnvironmentApprovalSettingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EnvironmentApprovalSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.environment.EnvironmentApprovalSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEnvironmentApprovalSettingsOutputReference_Override(e EnvironmentApprovalSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.environment.EnvironmentApprovalSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference)SetCanApplyDeclinedChanges(val interface{}) {
	if err := j.validateSetCanApplyDeclinedChangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canApplyDeclinedChanges",
		val,
	)
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference)SetCanReviewOwnRequest(val interface{}) {
	if err := j.validateSetCanReviewOwnRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canReviewOwnRequest",
		val,
	)
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference)SetMinNumApprovals(val *float64) {
	if err := j.validateSetMinNumApprovalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNumApprovals",
		val,
	)
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference)SetRequired(val interface{}) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference)SetRequiredApprovalTags(val *[]*string) {
	if err := j.validateSetRequiredApprovalTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredApprovalTags",
		val,
	)
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference)SetServiceConfig(val *map[string]*string) {
	if err := j.validateSetServiceConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceConfig",
		val,
	)
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference)SetServiceKind(val *string) {
	if err := j.validateSetServiceKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceKind",
		val,
	)
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EnvironmentApprovalSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) ResetCanApplyDeclinedChanges() {
	_jsii_.InvokeVoid(
		e,
		"resetCanApplyDeclinedChanges",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) ResetCanReviewOwnRequest() {
	_jsii_.InvokeVoid(
		e,
		"resetCanReviewOwnRequest",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) ResetMinNumApprovals() {
	_jsii_.InvokeVoid(
		e,
		"resetMinNumApprovals",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) ResetRequired() {
	_jsii_.InvokeVoid(
		e,
		"resetRequired",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) ResetRequiredApprovalTags() {
	_jsii_.InvokeVoid(
		e,
		"resetRequiredApprovalTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) ResetServiceConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) ResetServiceKind() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceKind",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentApprovalSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

