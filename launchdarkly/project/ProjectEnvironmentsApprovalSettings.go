package project


type ProjectEnvironmentsApprovalSettings struct {
	// Set to `true` if changes can be applied as long as the `min_num_approvals` is met, regardless of whether any reviewers have declined a request.
	//
	// Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/project#can_apply_declined_changes Project#can_apply_declined_changes}
	CanApplyDeclinedChanges interface{} `field:"optional" json:"canApplyDeclinedChanges" yaml:"canApplyDeclinedChanges"`
	// Set to `true` if requesters can approve or decline their own request. They may always comment. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/project#can_review_own_request Project#can_review_own_request}
	CanReviewOwnRequest interface{} `field:"optional" json:"canReviewOwnRequest" yaml:"canReviewOwnRequest"`
	// The number of approvals required before an approval request can be applied.
	//
	// This number must be between 1 and 5. Defaults to 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/project#min_num_approvals Project#min_num_approvals}
	MinNumApprovals *float64 `field:"optional" json:"minNumApprovals" yaml:"minNumApprovals"`
	// Set to `true` for changes to flags in this environment to require approval.
	//
	// You may only set `required` to true if `required_approval_tags` is not set and vice versa. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/project#required Project#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// An array of tags used to specify which flags with those tags require approval.
	//
	// You may only set `required_approval_tags` if `required` is not set to `true` and vice versa.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.15.0/docs/resources/project#required_approval_tags Project#required_approval_tags}
	RequiredApprovalTags *[]*string `field:"optional" json:"requiredApprovalTags" yaml:"requiredApprovalTags"`
}

