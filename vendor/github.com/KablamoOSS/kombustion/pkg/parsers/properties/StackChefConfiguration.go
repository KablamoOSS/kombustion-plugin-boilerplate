package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// StackChefConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-chefconfiguration.html
type StackChefConfiguration struct {
	BerkshelfVersion interface{} `yaml:"BerkshelfVersion,omitempty"`
	ManageBerkshelf  interface{} `yaml:"ManageBerkshelf,omitempty"`
}

// StackChefConfiguration validation
func (resource StackChefConfiguration) Validate() []error {
	errors := []error{}

	return errors
}
