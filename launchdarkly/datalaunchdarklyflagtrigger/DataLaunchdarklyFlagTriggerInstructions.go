package datalaunchdarklyflagtrigger


type DataLaunchdarklyFlagTriggerInstructions struct {
	// The action to perform when triggering. Currently supported flag actions are "turnFlagOn" and "turnFlagOff".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.2/docs/data-sources/flag_trigger#kind DataLaunchdarklyFlagTrigger#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
}

