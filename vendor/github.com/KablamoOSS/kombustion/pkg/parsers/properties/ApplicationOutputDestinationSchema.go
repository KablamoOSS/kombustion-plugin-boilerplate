package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ApplicationOutputDestinationSchema Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-destinationschema.html
type ApplicationOutputDestinationSchema struct {
	RecordFormatType interface{} `yaml:"RecordFormatType,omitempty"`
}

// ApplicationOutputDestinationSchema validation
func (resource ApplicationOutputDestinationSchema) Validate() []error {
	errors := []error{}

	return errors
}
