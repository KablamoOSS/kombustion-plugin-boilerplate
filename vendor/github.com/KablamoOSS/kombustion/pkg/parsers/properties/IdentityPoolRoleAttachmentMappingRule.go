package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// IdentityPoolRoleAttachmentMappingRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-mappingrule.html
type IdentityPoolRoleAttachmentMappingRule struct {
	Claim     interface{} `yaml:"Claim"`
	MatchType interface{} `yaml:"MatchType"`
	RoleARN   interface{} `yaml:"RoleARN"`
	Value     interface{} `yaml:"Value"`
}

// IdentityPoolRoleAttachmentMappingRule validation
func (resource IdentityPoolRoleAttachmentMappingRule) Validate() []error {
	errors := []error{}

	if resource.Claim == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Claim'"))
	}
	if resource.MatchType == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'MatchType'"))
	}
	if resource.RoleARN == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'RoleARN'"))
	}
	if resource.Value == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Value'"))
	}
	return errors
}
