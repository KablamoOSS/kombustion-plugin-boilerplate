package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ServiceCatalogCloudFormationProduct Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html
type ServiceCatalogCloudFormationProduct struct {
	Type       string                                        `yaml:"Type"`
	Properties ServiceCatalogCloudFormationProductProperties `yaml:"Properties"`
	Condition  interface{}                                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                   `yaml:"DependsOn,omitempty"`
}

// ServiceCatalogCloudFormationProduct Properties
type ServiceCatalogCloudFormationProductProperties struct {
	AcceptLanguage                 interface{} `yaml:"AcceptLanguage,omitempty"`
	Description                    interface{} `yaml:"Description,omitempty"`
	Distributor                    interface{} `yaml:"Distributor,omitempty"`
	Name                           interface{} `yaml:"Name"`
	Owner                          interface{} `yaml:"Owner"`
	SupportDescription             interface{} `yaml:"SupportDescription,omitempty"`
	SupportEmail                   interface{} `yaml:"SupportEmail,omitempty"`
	SupportUrl                     interface{} `yaml:"SupportUrl,omitempty"`
	ProvisioningArtifactParameters interface{} `yaml:"ProvisioningArtifactParameters"`
	Tags                           interface{} `yaml:"Tags,omitempty"`
}

// NewServiceCatalogCloudFormationProduct constructor creates a new ServiceCatalogCloudFormationProduct
func NewServiceCatalogCloudFormationProduct(properties ServiceCatalogCloudFormationProductProperties, deps ...interface{}) ServiceCatalogCloudFormationProduct {
	return ServiceCatalogCloudFormationProduct{
		Type:       "AWS::ServiceCatalog::CloudFormationProduct",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseServiceCatalogCloudFormationProduct parses ServiceCatalogCloudFormationProduct
func ParseServiceCatalogCloudFormationProduct(
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
	var resource ServiceCatalogCloudFormationProduct
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

// ParseServiceCatalogCloudFormationProduct validator
func (resource ServiceCatalogCloudFormationProduct) Validate() []error {
	return resource.Properties.Validate()
}

// ParseServiceCatalogCloudFormationProductProperties validator
func (resource ServiceCatalogCloudFormationProductProperties) Validate() []error {
	errors := []error{}
	if resource.Name == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Owner == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Owner'"))
	}
	if resource.ProvisioningArtifactParameters == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ProvisioningArtifactParameters'"))
	}
	return errors
}
