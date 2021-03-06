package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// IPSetIPSetDescriptor Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-ipset-ipsetdescriptor.html
type IPSetIPSetDescriptor struct {
	Type  interface{} `yaml:"Type"`
	Value interface{} `yaml:"Value"`
}

// IPSetIPSetDescriptor validation
func (resource IPSetIPSetDescriptor) Validate() []error {
	errors := []error{}

	if resource.Type == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Type'"))
	}
	if resource.Value == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Value'"))
	}
	return errors
}
