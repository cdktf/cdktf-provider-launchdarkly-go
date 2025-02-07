// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package flagtrigger


type FlagTriggerInstructions struct {
	// The action to perform when triggering. Currently supported flag actions are `turnFlagOn` and `turnFlagOff`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.22.0/docs/resources/flag_trigger#kind FlagTrigger#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
}

