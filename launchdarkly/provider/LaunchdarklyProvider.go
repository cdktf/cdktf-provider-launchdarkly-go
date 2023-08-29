// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-launchdarkly-go/launchdarkly/v2/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs launchdarkly}.
type LaunchdarklyProvider interface {
	cdktf.TerraformProvider
	AccessToken() *string
	SetAccessToken(val *string)
	AccessTokenInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	ApiHost() *string
	SetApiHost(val *string)
	ApiHostInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpTimeout() *float64
	SetHttpTimeout(val *float64)
	HttpTimeoutInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	OauthToken() *string
	SetOauthToken(val *string)
	OauthTokenInput() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAccessToken()
	ResetAlias()
	ResetApiHost()
	ResetHttpTimeout()
	ResetOauthToken()
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

// The jsii proxy struct for LaunchdarklyProvider
type jsiiProxy_LaunchdarklyProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_LaunchdarklyProvider) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) ApiHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) ApiHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) HttpTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) HttpTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) OauthToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) OauthTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchdarklyProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs launchdarkly} Resource.
func NewLaunchdarklyProvider(scope constructs.Construct, id *string, config *LaunchdarklyProviderConfig) LaunchdarklyProvider {
	_init_.Initialize()

	if err := validateNewLaunchdarklyProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LaunchdarklyProvider{}

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.provider.LaunchdarklyProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs launchdarkly} Resource.
func NewLaunchdarklyProvider_Override(l LaunchdarklyProvider, scope constructs.Construct, id *string, config *LaunchdarklyProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-launchdarkly.provider.LaunchdarklyProvider",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LaunchdarklyProvider)SetAccessToken(val *string) {
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_LaunchdarklyProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_LaunchdarklyProvider)SetApiHost(val *string) {
	_jsii_.Set(
		j,
		"apiHost",
		val,
	)
}

func (j *jsiiProxy_LaunchdarklyProvider)SetHttpTimeout(val *float64) {
	_jsii_.Set(
		j,
		"httpTimeout",
		val,
	)
}

func (j *jsiiProxy_LaunchdarklyProvider)SetOauthToken(val *string) {
	_jsii_.Set(
		j,
		"oauthToken",
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
func LaunchdarklyProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLaunchdarklyProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.provider.LaunchdarklyProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LaunchdarklyProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLaunchdarklyProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.provider.LaunchdarklyProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LaunchdarklyProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLaunchdarklyProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-launchdarkly.provider.LaunchdarklyProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LaunchdarklyProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-launchdarkly.provider.LaunchdarklyProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LaunchdarklyProvider) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LaunchdarklyProvider) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LaunchdarklyProvider) ResetAccessToken() {
	_jsii_.InvokeVoid(
		l,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchdarklyProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		l,
		"resetAlias",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchdarklyProvider) ResetApiHost() {
	_jsii_.InvokeVoid(
		l,
		"resetApiHost",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchdarklyProvider) ResetHttpTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchdarklyProvider) ResetOauthToken() {
	_jsii_.InvokeVoid(
		l,
		"resetOauthToken",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchdarklyProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchdarklyProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchdarklyProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchdarklyProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchdarklyProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

