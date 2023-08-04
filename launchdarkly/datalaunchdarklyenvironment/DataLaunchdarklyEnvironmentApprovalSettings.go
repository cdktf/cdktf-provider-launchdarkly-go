package datalaunchdarklyenvironment


type DataLaunchdarklyEnvironmentApprovalSettings struct {
	// Whether changes can be applied as long as minNumApprovals is met, regardless of whether any reviewers have declined a request.
	//
	// Defaults to true
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/environment#can_apply_declined_changes DataLaunchdarklyEnvironment#can_apply_declined_changes}
	CanApplyDeclinedChanges interface{} `field:"optional" json:"canApplyDeclinedChanges" yaml:"canApplyDeclinedChanges"`
	// Whether requesters can approve or decline their own request. They may always comment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/environment#can_review_own_request DataLaunchdarklyEnvironment#can_review_own_request}
	CanReviewOwnRequest interface{} `field:"optional" json:"canReviewOwnRequest" yaml:"canReviewOwnRequest"`
	// The number of approvals required before an approval request can be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/environment#min_num_approvals DataLaunchdarklyEnvironment#min_num_approvals}
	MinNumApprovals *float64 `field:"optional" json:"minNumApprovals" yaml:"minNumApprovals"`
	// Whether any changes to flags in this environment will require approval.
	//
	// You may only set required or requiredApprovalTags, not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/environment#required DataLaunchdarklyEnvironment#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// An array of tags used to specify which flags with those tags require approval.
	//
	// You may only set requiredApprovalTags or required, not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/launchdarkly/launchdarkly/2.13.4/docs/data-sources/environment#required_approval_tags DataLaunchdarklyEnvironment#required_approval_tags}
	RequiredApprovalTags *[]*string `field:"optional" json:"requiredApprovalTags" yaml:"requiredApprovalTags"`
}

