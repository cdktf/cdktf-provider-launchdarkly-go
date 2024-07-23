// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureflagenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FeatureFlagEnvironmentConfig struct {
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
	// The environment key.
	//
	// A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/feature_flag_environment#env_key FeatureFlagEnvironment#env_key}
	EnvKey *string `field:"required" json:"envKey" yaml:"envKey"`
	// fallthrough block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/feature_flag_environment#fallthrough FeatureFlagEnvironment#fallthrough}
	Fallthrough *FeatureFlagEnvironmentFallthrough `field:"required" json:"fallthrough" yaml:"fallthrough"`
	// The feature flag's unique `id` in the format `project_key/flag_key`.
	//
	// A change in this field will force the destruction of the existing resource and the creation of a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/feature_flag_environment#flag_id FeatureFlagEnvironment#flag_id}
	FlagId *string `field:"required" json:"flagId" yaml:"flagId"`
	// The index of the variation to serve if targeting is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/feature_flag_environment#off_variation FeatureFlagEnvironment#off_variation}
	OffVariation *float64 `field:"required" json:"offVariation" yaml:"offVariation"`
	// context_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/feature_flag_environment#context_targets FeatureFlagEnvironment#context_targets}
	ContextTargets interface{} `field:"optional" json:"contextTargets" yaml:"contextTargets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/feature_flag_environment#id FeatureFlagEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether targeting is enabled. Defaults to `false` if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/feature_flag_environment#on FeatureFlagEnvironment#on}
	On interface{} `field:"optional" json:"on" yaml:"on"`
	// prerequisites block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/feature_flag_environment#prerequisites FeatureFlagEnvironment#prerequisites}
	Prerequisites interface{} `field:"optional" json:"prerequisites" yaml:"prerequisites"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/feature_flag_environment#rules FeatureFlagEnvironment#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/feature_flag_environment#targets FeatureFlagEnvironment#targets}
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
	// Whether to send event data back to LaunchDarkly. Defaults to `false` if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.20.0/docs/resources/feature_flag_environment#track_events FeatureFlagEnvironment#track_events}
	TrackEvents interface{} `field:"optional" json:"trackEvents" yaml:"trackEvents"`
}

