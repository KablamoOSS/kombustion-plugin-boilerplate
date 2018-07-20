package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// DynamoDBTable Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html
type DynamoDBTable struct {
	Type       string                  `yaml:"Type"`
	Properties DynamoDBTableProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

// DynamoDBTable Properties
type DynamoDBTableProperties struct {
	TableName                        interface{}                                       `yaml:"TableName,omitempty"`
	TimeToLiveSpecification          *properties.TableTimeToLiveSpecification          `yaml:"TimeToLiveSpecification,omitempty"`
	StreamSpecification              *properties.TableStreamSpecification              `yaml:"StreamSpecification,omitempty"`
	SSESpecification                 *properties.TableSSESpecification                 `yaml:"SSESpecification,omitempty"`
	ProvisionedThroughput            *properties.TableProvisionedThroughput            `yaml:"ProvisionedThroughput"`
	PointInTimeRecoverySpecification *properties.TablePointInTimeRecoverySpecification `yaml:"PointInTimeRecoverySpecification,omitempty"`
	KeySchema                        interface{}                                       `yaml:"KeySchema"`
	AttributeDefinitions             interface{}                                       `yaml:"AttributeDefinitions,omitempty"`
	GlobalSecondaryIndexes           interface{}                                       `yaml:"GlobalSecondaryIndexes,omitempty"`
	Tags                             interface{}                                       `yaml:"Tags,omitempty"`
	LocalSecondaryIndexes            interface{}                                       `yaml:"LocalSecondaryIndexes,omitempty"`
}

// NewDynamoDBTable constructor creates a new DynamoDBTable
func NewDynamoDBTable(properties DynamoDBTableProperties, deps ...interface{}) DynamoDBTable {
	return DynamoDBTable{
		Type:       "AWS::DynamoDB::Table",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseDynamoDBTable parses DynamoDBTable
func ParseDynamoDBTable(
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
	var resource DynamoDBTable
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

// ParseDynamoDBTable validator
func (resource DynamoDBTable) Validate() []error {
	return resource.Properties.Validate()
}

// ParseDynamoDBTableProperties validator
func (resource DynamoDBTableProperties) Validate() []error {
	errors := []error{}
	if resource.ProvisionedThroughput == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ProvisionedThroughput'"))
	} else {
		errors = append(errors, resource.ProvisionedThroughput.Validate()...)
	}
	if resource.KeySchema == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'KeySchema'"))
	}
	return errors
}
