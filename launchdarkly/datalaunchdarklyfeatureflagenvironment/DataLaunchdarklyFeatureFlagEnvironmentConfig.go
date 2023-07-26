package datalaunchdarklyfeatureflagenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLaunchdarklyFeatureFlagEnvironmentConfig struct {
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
	// The LaunchDarkly environment key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/feature_flag_environment#env_key DataLaunchdarklyFeatureFlagEnvironment#env_key}
	EnvKey *string `field:"required" json:"envKey" yaml:"envKey"`
	// The global feature flag's unique id in the format `<project_key>/<flag_key>`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/feature_flag_environment#flag_id DataLaunchdarklyFeatureFlagEnvironment#flag_id}
	FlagId *string `field:"required" json:"flagId" yaml:"flagId"`
	// context_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/feature_flag_environment#context_targets DataLaunchdarklyFeatureFlagEnvironment#context_targets}
	ContextTargets interface{} `field:"optional" json:"contextTargets" yaml:"contextTargets"`
	// fallthrough block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/feature_flag_environment#fallthrough DataLaunchdarklyFeatureFlagEnvironment#fallthrough}
	Fallthrough *DataLaunchdarklyFeatureFlagEnvironmentFallthrough `field:"optional" json:"fallthrough" yaml:"fallthrough"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/feature_flag_environment#id DataLaunchdarklyFeatureFlagEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The index of the variation to serve if targeting is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/feature_flag_environment#off_variation DataLaunchdarklyFeatureFlagEnvironment#off_variation}
	OffVariation *float64 `field:"optional" json:"offVariation" yaml:"offVariation"`
	// Whether targeting is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/feature_flag_environment#on DataLaunchdarklyFeatureFlagEnvironment#on}
	On interface{} `field:"optional" json:"on" yaml:"on"`
	// prerequisites block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/feature_flag_environment#prerequisites DataLaunchdarklyFeatureFlagEnvironment#prerequisites}
	Prerequisites interface{} `field:"optional" json:"prerequisites" yaml:"prerequisites"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/feature_flag_environment#rules DataLaunchdarklyFeatureFlagEnvironment#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/feature_flag_environment#targets DataLaunchdarklyFeatureFlagEnvironment#targets}
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
	// Whether to send event data back to LaunchDarkly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/feature_flag_environment#track_events DataLaunchdarklyFeatureFlagEnvironment#track_events}
	TrackEvents interface{} `field:"optional" json:"trackEvents" yaml:"trackEvents"`
}

