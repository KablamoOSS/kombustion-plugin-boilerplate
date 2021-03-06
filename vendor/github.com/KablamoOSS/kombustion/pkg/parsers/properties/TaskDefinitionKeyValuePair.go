package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TaskDefinitionKeyValuePair Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-environment.html
type TaskDefinitionKeyValuePair struct {
	Name  interface{} `yaml:"Name,omitempty"`
	Value interface{} `yaml:"Value,omitempty"`
}

// TaskDefinitionKeyValuePair validation
func (resource TaskDefinitionKeyValuePair) Validate() []error {
	errors := []error{}

	return errors
}
