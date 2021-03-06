package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ClusterScalingAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingaction.html
type ClusterScalingAction struct {
	Market                           interface{}                              `yaml:"Market,omitempty"`
	SimpleScalingPolicyConfiguration *ClusterSimpleScalingPolicyConfiguration `yaml:"SimpleScalingPolicyConfiguration"`
}

// ClusterScalingAction validation
func (resource ClusterScalingAction) Validate() []error {
	errors := []error{}

	if resource.SimpleScalingPolicyConfiguration == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'SimpleScalingPolicyConfiguration'"))
	} else {
		errors = append(errors, resource.SimpleScalingPolicyConfiguration.Validate()...)
	}
	return errors
}
