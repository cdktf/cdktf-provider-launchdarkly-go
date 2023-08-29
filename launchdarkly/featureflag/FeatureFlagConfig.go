// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureflag

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FeatureFlagConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A unique key that will be used to reference the flag in your code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#key FeatureFlag#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// A human-friendly name for the feature flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#name FeatureFlag#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The LaunchDarkly project key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#project_key FeatureFlag#project_key}
	ProjectKey *string `field:"required" json:"projectKey" yaml:"projectKey"`
	// The uniform type for all variations. Can be either "boolean", "string", "number", or "json".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#variation_type FeatureFlag#variation_type}
	VariationType *string `field:"required" json:"variationType" yaml:"variationType"`
	// Whether to archive the flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#archived FeatureFlag#archived}
	Archived interface{} `field:"optional" json:"archived" yaml:"archived"`
	// client_side_availability block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#client_side_availability FeatureFlag#client_side_availability}
	ClientSideAvailability interface{} `field:"optional" json:"clientSideAvailability" yaml:"clientSideAvailability"`
	// custom_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#custom_properties FeatureFlag#custom_properties}
	CustomProperties interface{} `field:"optional" json:"customProperties" yaml:"customProperties"`
	// defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#defaults FeatureFlag#defaults}
	Defaults *FeatureFlagDefaults `field:"optional" json:"defaults" yaml:"defaults"`
	// A short description of what the flag will be used for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#description FeatureFlag#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#id FeatureFlag#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether or not this flag should be made available to the client-side JavaScript SDK.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#include_in_snippet FeatureFlag#include_in_snippet}
	IncludeInSnippet interface{} `field:"optional" json:"includeInSnippet" yaml:"includeInSnippet"`
	// The LaunchDarkly id of the user who will maintain the flag.
	//
	// If not set, the API will automatically apply the member associated with your Terraform API key or the most recently set maintainer
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#maintainer_id FeatureFlag#maintainer_id}
	MaintainerId *string `field:"optional" json:"maintainerId" yaml:"maintainerId"`
	// Tags associated with your resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#tags FeatureFlag#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Whether or not the flag is a temporary flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#temporary FeatureFlag#temporary}
	Temporary interface{} `field:"optional" json:"temporary" yaml:"temporary"`
	// variations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/feature_flag#variations FeatureFlag#variations}
	Variations interface{} `field:"optional" json:"variations" yaml:"variations"`
}

