package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ReceiptFilterIpFilter Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptfilter-ipfilter.html
type ReceiptFilterIpFilter struct {
	Cidr   interface{} `yaml:"Cidr"`
	Policy interface{} `yaml:"Policy"`
}

// ReceiptFilterIpFilter validation
func (resource ReceiptFilterIpFilter) Validate() []error {
	errors := []error{}

	if resource.Cidr == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Cidr'"))
	}
	if resource.Policy == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Policy'"))
	}
	return errors
}
