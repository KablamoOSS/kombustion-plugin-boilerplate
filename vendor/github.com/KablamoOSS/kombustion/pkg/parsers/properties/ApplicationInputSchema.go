package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ApplicationInputSchema Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputschema.html
type ApplicationInputSchema struct {
	RecordEncoding interface{}              `yaml:"RecordEncoding,omitempty"`
	RecordFormat   *ApplicationRecordFormat `yaml:"RecordFormat"`
	RecordColumns  interface{}              `yaml:"RecordColumns"`
}

// ApplicationInputSchema validation
func (resource ApplicationInputSchema) Validate() []error {
	errs := []error{}

	if resource.RecordFormat == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RecordFormat'"))
	} else {
		errs = append(errs, resource.RecordFormat.Validate()...)
	}
	if resource.RecordColumns == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RecordColumns'"))
	}
	return errs
}