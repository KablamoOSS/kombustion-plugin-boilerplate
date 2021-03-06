package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ApplicationOutputKinesisStreamsOutput Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-kinesisstreamsoutput.html
type ApplicationOutputKinesisStreamsOutput struct {
	ResourceARN interface{} `yaml:"ResourceARN"`
	RoleARN     interface{} `yaml:"RoleARN"`
}

// ApplicationOutputKinesisStreamsOutput validation
func (resource ApplicationOutputKinesisStreamsOutput) Validate() []error {
	errors := []error{}

	if resource.ResourceARN == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ResourceARN'"))
	}
	if resource.RoleARN == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'RoleARN'"))
	}
	return errors
}
