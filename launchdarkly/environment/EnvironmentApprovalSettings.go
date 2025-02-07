// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package environment


type EnvironmentApprovalSettings struct {
	// Set to `true` if changes can be applied as long as the `min_num_approvals` is met, regardless of whether any reviewers have declined a request.
	//
	// Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.22.0/docs/resources/environment#can_apply_declined_changes Environment#can_apply_declined_changes}
	CanApplyDeclinedChanges interface{} `field:"optional" json:"canApplyDeclinedChanges" yaml:"canApplyDeclinedChanges"`
	// Set to `true` if requesters can approve or decline their own request. They may always comment. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.22.0/docs/resources/environment#can_review_own_request Environment#can_review_own_request}
	CanReviewOwnRequest interface{} `field:"optional" json:"canReviewOwnRequest" yaml:"canReviewOwnRequest"`
	// The number of approvals required before an approval request can be applied.
	//
	// This number must be between 1 and 5. Defaults to 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.22.0/docs/resources/environment#min_num_approvals Environment#min_num_approvals}
	MinNumApprovals *float64 `field:"optional" json:"minNumApprovals" yaml:"minNumApprovals"`
	// Set to `true` for changes to flags in this environment to require approval.
	//
	// You may only set `required` to true if `required_approval_tags` is not set and vice versa. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.22.0/docs/resources/environment#required Environment#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// An array of tags used to specify which flags with those tags require approval.
	//
	// You may only set `required_approval_tags` if `required` is not set to `true` and vice versa.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.22.0/docs/resources/environment#required_approval_tags Environment#required_approval_tags}
	RequiredApprovalTags *[]*string `field:"optional" json:"requiredApprovalTags" yaml:"requiredApprovalTags"`
	// The configuration for the service associated with this approval.
	//
	// This is specific to each approval service. For a `service_kind` of `servicenow`, the following fields apply:
	//
	// 	 - `template` (String) The sys_id of the Standard Change Request Template in ServiceNow that LaunchDarkly will use when creating the change request.
	// 	 - `detail_column` (String) The name of the ServiceNow Change Request column LaunchDarkly uses to populate detailed approval request information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.22.0/docs/resources/environment#service_config Environment#service_config}
	ServiceConfig *map[string]*string `field:"optional" json:"serviceConfig" yaml:"serviceConfig"`
	// The kind of service associated with this approval.
	//
	// This determines which platform is used for requesting approval. Valid values are `servicenow`, `launchdarkly`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.22.0/docs/resources/environment#service_kind Environment#service_kind}
	ServiceKind *string `field:"optional" json:"serviceKind" yaml:"serviceKind"`
}

