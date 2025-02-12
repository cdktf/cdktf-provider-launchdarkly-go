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
	// The unique feature flag key that references the flag in your application code.
	//
	// A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag#key FeatureFlag#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The human-readable name of the feature flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag#name FeatureFlag#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The feature flag's project key.
	//
	// A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag#project_key FeatureFlag#project_key}
	ProjectKey *string `field:"required" json:"projectKey" yaml:"projectKey"`
	// The feature flag's variation type: `boolean`, `string`, `number` or `json`.
	//
	// A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag#variation_type FeatureFlag#variation_type}
	VariationType *string `field:"required" json:"variationType" yaml:"variationType"`
	// Specifies whether the flag is archived or not.
	//
	// Note that you cannot create a new flag that is archived, but can update a flag to be archived.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag#archived FeatureFlag#archived}
	Archived interface{} `field:"optional" json:"archived" yaml:"archived"`
	// client_side_availability block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag#client_side_availability FeatureFlag#client_side_availability}
	ClientSideAvailability interface{} `field:"optional" json:"clientSideAvailability" yaml:"clientSideAvailability"`
	// custom_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag#custom_properties FeatureFlag#custom_properties}
	CustomProperties interface{} `field:"optional" json:"customProperties" yaml:"customProperties"`
	// defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag#defaults FeatureFlag#defaults}
	Defaults *FeatureFlagDefaults `field:"optional" json:"defaults" yaml:"defaults"`
	// The feature flag's description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag#description FeatureFlag#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag#id FeatureFlag#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies whether this flag should be made available to the client-side JavaScript SDK using the client-side Id.
	//
	// This value gets its default from your project configuration if not set. `include_in_snippet` is now deprecated. Please migrate to `client_side_availability.using_environment_id` to maintain future compatibility.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag#include_in_snippet FeatureFlag#include_in_snippet}
	IncludeInSnippet interface{} `field:"optional" json:"includeInSnippet" yaml:"includeInSnippet"`
	// The feature flag maintainer's 24 character alphanumeric team member ID.
	//
	// `maintainer_team_key` cannot be set if `maintainer_id` is set. If neither is set, it will automatically be or stay set to the member ID associated with the API key used by your LaunchDarkly Terraform provider or the most recently-set maintainer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag#maintainer_id FeatureFlag#maintainer_id}
	MaintainerId *string `field:"optional" json:"maintainerId" yaml:"maintainerId"`
	// The key of the associated team that maintains this feature flag. `maintainer_id` cannot be set if `maintainer_team_key` is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag#maintainer_team_key FeatureFlag#maintainer_team_key}
	MaintainerTeamKey *string `field:"optional" json:"maintainerTeamKey" yaml:"maintainerTeamKey"`
	// Tags associated with your resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag#tags FeatureFlag#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Specifies whether the flag is a temporary flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag#temporary FeatureFlag#temporary}
	Temporary interface{} `field:"optional" json:"temporary" yaml:"temporary"`
	// variations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.23.0/docs/resources/feature_flag#variations FeatureFlag#variations}
	Variations interface{} `field:"optional" json:"variations" yaml:"variations"`
}

