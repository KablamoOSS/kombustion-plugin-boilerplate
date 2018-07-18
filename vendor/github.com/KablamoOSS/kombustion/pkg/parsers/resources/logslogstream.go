package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// LogsLogStream Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html
type LogsLogStream struct {
	Type       string                  `yaml:"Type"`
	Properties LogsLogStreamProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

// LogsLogStream Properties
type LogsLogStreamProperties struct {
	LogGroupName  interface{} `yaml:"LogGroupName"`
	LogStreamName interface{} `yaml:"LogStreamName,omitempty"`
}

// NewLogsLogStream constructor creates a new LogsLogStream
func NewLogsLogStream(properties LogsLogStreamProperties, deps ...interface{}) LogsLogStream {
	return LogsLogStream{
		Type:       "AWS::Logs::LogStream",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseLogsLogStream parses LogsLogStream
func ParseLogsLogStream(name string, data string) (cf types.TemplateObject, err error) {
	var resource LogsLogStream
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: LogsLogStream - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseLogsLogStream validator
func (resource LogsLogStream) Validate() []error {
	return resource.Properties.Validate()
}

// ParseLogsLogStreamProperties validator
func (resource LogsLogStreamProperties) Validate() []error {
	errs := []error{}
	if resource.LogGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'LogGroupName'"))
	}
	return errs
}