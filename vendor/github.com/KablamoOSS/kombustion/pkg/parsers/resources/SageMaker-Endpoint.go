package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// SageMakerEndpoint Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html
type SageMakerEndpoint struct {
	Type       string                      `yaml:"Type"`
	Properties SageMakerEndpointProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

// SageMakerEndpoint Properties
type SageMakerEndpointProperties struct {
	EndpointConfigName interface{} `yaml:"EndpointConfigName"`
	EndpointName       interface{} `yaml:"EndpointName,omitempty"`
	Tags               interface{} `yaml:"Tags,omitempty"`
}

// NewSageMakerEndpoint constructor creates a new SageMakerEndpoint
func NewSageMakerEndpoint(properties SageMakerEndpointProperties, deps ...interface{}) SageMakerEndpoint {
	return SageMakerEndpoint{
		Type:       "AWS::SageMaker::Endpoint",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSageMakerEndpoint parses SageMakerEndpoint
func ParseSageMakerEndpoint(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"
	var resource SageMakerEndpoint
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	return
}

// ParseSageMakerEndpoint validator
func (resource SageMakerEndpoint) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSageMakerEndpointProperties validator
func (resource SageMakerEndpointProperties) Validate() []error {
	errors := []error{}
	if resource.EndpointConfigName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'EndpointConfigName'"))
	}
	return errors
}
