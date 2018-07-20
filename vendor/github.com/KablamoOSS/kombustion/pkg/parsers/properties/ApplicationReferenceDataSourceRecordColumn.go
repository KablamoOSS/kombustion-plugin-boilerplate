package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ApplicationReferenceDataSourceRecordColumn Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-recordcolumn.html
type ApplicationReferenceDataSourceRecordColumn struct {
	Mapping interface{} `yaml:"Mapping,omitempty"`
	Name    interface{} `yaml:"Name"`
	SqlType interface{} `yaml:"SqlType"`
}

// ApplicationReferenceDataSourceRecordColumn validation
func (resource ApplicationReferenceDataSourceRecordColumn) Validate() []error {
	errors := []error{}

	if resource.Name == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.SqlType == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'SqlType'"))
	}
	return errors
}
