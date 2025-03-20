// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type LaunchdarklyProviderConfig struct {
	// The [personal access token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#personal-tokens) or [service token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#service-tokens) used to authenticate with LaunchDarkly. You can also set this with the `LAUNCHDARKLY_ACCESS_TOKEN` environment variable. You must provide either `access_token` or `oauth_token`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.0/docs#access_token LaunchdarklyProvider#access_token}
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.0/docs#alias LaunchdarklyProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The LaunchDarkly host address. If this argument is not specified, the default host address is `https://app.launchdarkly.com`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.0/docs#api_host LaunchdarklyProvider#api_host}
	ApiHost *string `field:"optional" json:"apiHost" yaml:"apiHost"`
	// The HTTP timeout (in seconds) when making API calls to LaunchDarkly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.0/docs#http_timeout LaunchdarklyProvider#http_timeout}
	HttpTimeout *float64 `field:"optional" json:"httpTimeout" yaml:"httpTimeout"`
	// An OAuth V2 token you use to authenticate with LaunchDarkly.
	//
	// You can also set this with the `LAUNCHDARKLY_OAUTH_TOKEN` environment variable. You must provide either `access_token` or `oauth_token`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.25.0/docs#oauth_token LaunchdarklyProvider#oauth_token}
	OauthToken *string `field:"optional" json:"oauthToken" yaml:"oauthToken"`
}

