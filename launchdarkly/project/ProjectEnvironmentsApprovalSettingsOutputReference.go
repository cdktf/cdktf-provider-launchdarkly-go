// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package project

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/v5/jsii"

	"github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/v5/project/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectEnvironmentsApprovalSettingsOutputReference interface {
	cdktf.ComplexObject
	AutoApplyApprovedChanges() interface{}
	SetAutoApplyApprovedChanges(val interface{})
	AutoApplyApprovedChangesInput() interface{}
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
	ResetAutoApplyApprovedChanges()
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

// The jsii proxy struct for ProjectEnvironmentsApprovalSettingsOutputReference
type jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) AutoApplyApprovedChanges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoApplyApprovedChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) AutoApplyApprovedChangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoApplyApprovedChangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) CanApplyDeclinedChanges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canApplyDeclinedChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) CanApplyDeclinedChangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canApplyDeclinedChangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) CanReviewOwnRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canReviewOwnRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) CanReviewOwnRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canReviewOwnRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) MinNumApprovals() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumApprovals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) MinNumApprovalsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumApprovalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) Required() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) RequiredApprovalTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredApprovalTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) RequiredApprovalTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredApprovalTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) RequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) ServiceConfig() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"serviceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) ServiceConfigInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"serviceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) ServiceKind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) ServiceKindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewProjectEnvironmentsApprovalSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ProjectEnvironmentsApprovalSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewProjectEnvironmentsApprovalSettingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.project.ProjectEnvironmentsApprovalSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewProjectEnvironmentsApprovalSettingsOutputReference_Override(p ProjectEnvironmentsApprovalSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.project.ProjectEnvironmentsApprovalSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference)SetAutoApplyApprovedChanges(val interface{}) {
	if err := j.validateSetAutoApplyApprovedChangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoApplyApprovedChanges",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference)SetCanApplyDeclinedChanges(val interface{}) {
	if err := j.validateSetCanApplyDeclinedChangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canApplyDeclinedChanges",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference)SetCanReviewOwnRequest(val interface{}) {
	if err := j.validateSetCanReviewOwnRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canReviewOwnRequest",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference)SetMinNumApprovals(val *float64) {
	if err := j.validateSetMinNumApprovalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNumApprovals",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference)SetRequired(val interface{}) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference)SetRequiredApprovalTags(val *[]*string) {
	if err := j.validateSetRequiredApprovalTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredApprovalTags",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference)SetServiceConfig(val *map[string]*string) {
	if err := j.validateSetServiceConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceConfig",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference)SetServiceKind(val *string) {
	if err := j.validateSetServiceKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceKind",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) ResetAutoApplyApprovedChanges() {
	_jsii_.InvokeVoid(
		p,
		"resetAutoApplyApprovedChanges",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) ResetCanApplyDeclinedChanges() {
	_jsii_.InvokeVoid(
		p,
		"resetCanApplyDeclinedChanges",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) ResetCanReviewOwnRequest() {
	_jsii_.InvokeVoid(
		p,
		"resetCanReviewOwnRequest",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) ResetMinNumApprovals() {
	_jsii_.InvokeVoid(
		p,
		"resetMinNumApprovals",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) ResetRequired() {
	_jsii_.InvokeVoid(
		p,
		"resetRequired",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) ResetRequiredApprovalTags() {
	_jsii_.InvokeVoid(
		p,
		"resetRequiredApprovalTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) ResetServiceConfig() {
	_jsii_.InvokeVoid(
		p,
		"resetServiceConfig",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) ResetServiceKind() {
	_jsii_.InvokeVoid(
		p,
		"resetServiceKind",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectEnvironmentsApprovalSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

