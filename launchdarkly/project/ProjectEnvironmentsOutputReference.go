// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package project

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/v4/jsii"

	"github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/v4/project/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectEnvironmentsOutputReference interface {
	cdktf.ComplexObject
	ApiKey() *string
	ApprovalSettings() ProjectEnvironmentsApprovalSettingsList
	ApprovalSettingsInput() interface{}
	ClientSideId() *string
	Color() *string
	SetColor(val *string)
	ColorInput() *string
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
	ConfirmChanges() interface{}
	SetConfirmChanges(val interface{})
	ConfirmChangesInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Critical() interface{}
	SetCritical(val interface{})
	CriticalInput() interface{}
	DefaultTrackEvents() interface{}
	SetDefaultTrackEvents(val interface{})
	DefaultTrackEventsInput() interface{}
	DefaultTtl() *float64
	SetDefaultTtl(val *float64)
	DefaultTtlInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	MobileKey() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	RequireComments() interface{}
	SetRequireComments(val interface{})
	RequireCommentsInput() interface{}
	SecureMode() interface{}
	SetSecureMode(val interface{})
	SecureModeInput() interface{}
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
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
	PutApprovalSettings(value interface{})
	ResetApprovalSettings()
	ResetConfirmChanges()
	ResetCritical()
	ResetDefaultTrackEvents()
	ResetDefaultTtl()
	ResetRequireComments()
	ResetSecureMode()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ProjectEnvironmentsOutputReference
type jsiiProxy_ProjectEnvironmentsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) ApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) ApprovalSettings() ProjectEnvironmentsApprovalSettingsList {
	var returns ProjectEnvironmentsApprovalSettingsList
	_jsii_.Get(
		j,
		"approvalSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) ApprovalSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) ClientSideId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSideId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) Color() *string {
	var returns *string
	_jsii_.Get(
		j,
		"color",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) ColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) ConfirmChanges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confirmChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) ConfirmChangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confirmChangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) Critical() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"critical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) CriticalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"criticalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) DefaultTrackEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultTrackEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) DefaultTrackEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultTrackEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) DefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) MobileKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobileKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) RequireComments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireComments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) RequireCommentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCommentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) SecureMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) SecureModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewProjectEnvironmentsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ProjectEnvironmentsOutputReference {
	_init_.Initialize()

	if err := validateNewProjectEnvironmentsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjectEnvironmentsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.project.ProjectEnvironmentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewProjectEnvironmentsOutputReference_Override(p ProjectEnvironmentsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.project.ProjectEnvironmentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference)SetColor(val *string) {
	if err := j.validateSetColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"color",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference)SetConfirmChanges(val interface{}) {
	if err := j.validateSetConfirmChangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confirmChanges",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference)SetCritical(val interface{}) {
	if err := j.validateSetCriticalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"critical",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference)SetDefaultTrackEvents(val interface{}) {
	if err := j.validateSetDefaultTrackEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTrackEvents",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference)SetDefaultTtl(val *float64) {
	if err := j.validateSetDefaultTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTtl",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference)SetRequireComments(val interface{}) {
	if err := j.validateSetRequireCommentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireComments",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference)SetSecureMode(val interface{}) {
	if err := j.validateSetSecureModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secureMode",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ProjectEnvironmentsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) PutApprovalSettings(value interface{}) {
	if err := p.validatePutApprovalSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApprovalSettings",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) ResetApprovalSettings() {
	_jsii_.InvokeVoid(
		p,
		"resetApprovalSettings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) ResetConfirmChanges() {
	_jsii_.InvokeVoid(
		p,
		"resetConfirmChanges",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) ResetCritical() {
	_jsii_.InvokeVoid(
		p,
		"resetCritical",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) ResetDefaultTrackEvents() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultTrackEvents",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) ResetRequireComments() {
	_jsii_.InvokeVoid(
		p,
		"resetRequireComments",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) ResetSecureMode() {
	_jsii_.InvokeVoid(
		p,
		"resetSecureMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_ProjectEnvironmentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

