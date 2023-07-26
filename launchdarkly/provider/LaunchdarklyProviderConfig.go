package provider


type LaunchdarklyProviderConfig struct {
	// The LaunchDarkly API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs#access_token LaunchdarklyProvider#access_token}
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs#alias LaunchdarklyProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The LaunchDarkly host address, e.g. https://app.launchdarkly.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs#api_host LaunchdarklyProvider#api_host}
	ApiHost *string `field:"optional" json:"apiHost" yaml:"apiHost"`
	// The LaunchDarkly OAuth token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs#oauth_token LaunchdarklyProvider#oauth_token}
	OauthToken *string `field:"optional" json:"oauthToken" yaml:"oauthToken"`
}

