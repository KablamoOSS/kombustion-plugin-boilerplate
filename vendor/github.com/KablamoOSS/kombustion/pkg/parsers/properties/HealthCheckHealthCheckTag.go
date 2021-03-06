package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// HealthCheckHealthCheckTag Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktag.html
type HealthCheckHealthCheckTag struct {
	Key   interface{} `yaml:"Key"`
	Value interface{} `yaml:"Value"`
}

// HealthCheckHealthCheckTag validation
func (resource HealthCheckHealthCheckTag) Validate() []error {
	errors := []error{}

	if resource.Key == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.Value == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Value'"))
	}
	return errors
}
