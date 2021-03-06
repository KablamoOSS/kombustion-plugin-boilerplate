package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ApplicationRecordFormat Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordformat.html
type ApplicationRecordFormat struct {
	RecordFormatType  interface{}                   `yaml:"RecordFormatType"`
	MappingParameters *ApplicationMappingParameters `yaml:"MappingParameters,omitempty"`
}

// ApplicationRecordFormat validation
func (resource ApplicationRecordFormat) Validate() []error {
	errors := []error{}

	if resource.RecordFormatType == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'RecordFormatType'"))
	}
	return errors
}
