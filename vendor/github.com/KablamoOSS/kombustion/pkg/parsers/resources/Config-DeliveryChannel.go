package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ConfigDeliveryChannel Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html
type ConfigDeliveryChannel struct {
	Type       string                          `yaml:"Type"`
	Properties ConfigDeliveryChannelProperties `yaml:"Properties"`
	Condition  interface{}                     `yaml:"Condition,omitempty"`
	Metadata   interface{}                     `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                     `yaml:"DependsOn,omitempty"`
}

// ConfigDeliveryChannel Properties
type ConfigDeliveryChannelProperties struct {
	Name                             interface{}                                                 `yaml:"Name,omitempty"`
	S3BucketName                     interface{}                                                 `yaml:"S3BucketName"`
	S3KeyPrefix                      interface{}                                                 `yaml:"S3KeyPrefix,omitempty"`
	SnsTopicARN                      interface{}                                                 `yaml:"SnsTopicARN,omitempty"`
	ConfigSnapshotDeliveryProperties *properties.DeliveryChannelConfigSnapshotDeliveryProperties `yaml:"ConfigSnapshotDeliveryProperties,omitempty"`
}

// NewConfigDeliveryChannel constructor creates a new ConfigDeliveryChannel
func NewConfigDeliveryChannel(properties ConfigDeliveryChannelProperties, deps ...interface{}) ConfigDeliveryChannel {
	return ConfigDeliveryChannel{
		Type:       "AWS::Config::DeliveryChannel",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseConfigDeliveryChannel parses ConfigDeliveryChannel
func ParseConfigDeliveryChannel(
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
	var resource ConfigDeliveryChannel
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

// ParseConfigDeliveryChannel validator
func (resource ConfigDeliveryChannel) Validate() []error {
	return resource.Properties.Validate()
}

// ParseConfigDeliveryChannelProperties validator
func (resource ConfigDeliveryChannelProperties) Validate() []error {
	errors := []error{}
	if resource.S3BucketName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'S3BucketName'"))
	}
	return errors
}