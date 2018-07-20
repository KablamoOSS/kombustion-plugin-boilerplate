package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ServiceCatalogPortfolioShare Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioshare.html
type ServiceCatalogPortfolioShare struct {
	Type       string                                 `yaml:"Type"`
	Properties ServiceCatalogPortfolioShareProperties `yaml:"Properties"`
	Condition  interface{}                            `yaml:"Condition,omitempty"`
	Metadata   interface{}                            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                            `yaml:"DependsOn,omitempty"`
}

// ServiceCatalogPortfolioShare Properties
type ServiceCatalogPortfolioShareProperties struct {
	AcceptLanguage interface{} `yaml:"AcceptLanguage,omitempty"`
	AccountId      interface{} `yaml:"AccountId"`
	PortfolioId    interface{} `yaml:"PortfolioId"`
}

// NewServiceCatalogPortfolioShare constructor creates a new ServiceCatalogPortfolioShare
func NewServiceCatalogPortfolioShare(properties ServiceCatalogPortfolioShareProperties, deps ...interface{}) ServiceCatalogPortfolioShare {
	return ServiceCatalogPortfolioShare{
		Type:       "AWS::ServiceCatalog::PortfolioShare",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseServiceCatalogPortfolioShare parses ServiceCatalogPortfolioShare
func ParseServiceCatalogPortfolioShare(
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
	var resource ServiceCatalogPortfolioShare
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

// ParseServiceCatalogPortfolioShare validator
func (resource ServiceCatalogPortfolioShare) Validate() []error {
	return resource.Properties.Validate()
}

// ParseServiceCatalogPortfolioShareProperties validator
func (resource ServiceCatalogPortfolioShareProperties) Validate() []error {
	errors := []error{}
	if resource.AccountId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'AccountId'"))
	}
	if resource.PortfolioId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'PortfolioId'"))
	}
	return errors
}
