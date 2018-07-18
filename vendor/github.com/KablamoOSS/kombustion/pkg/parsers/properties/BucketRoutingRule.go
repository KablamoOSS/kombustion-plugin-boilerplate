package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BucketRoutingRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html
type BucketRoutingRule struct {
	RoutingRuleCondition *BucketRoutingRuleCondition `yaml:"RoutingRuleCondition,omitempty"`
	RedirectRule         *BucketRedirectRule         `yaml:"RedirectRule"`
}

// BucketRoutingRule validation
func (resource BucketRoutingRule) Validate() []error {
	errs := []error{}

	if resource.RedirectRule == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RedirectRule'"))
	} else {
		errs = append(errs, resource.RedirectRule.Validate()...)
	}
	return errs
}