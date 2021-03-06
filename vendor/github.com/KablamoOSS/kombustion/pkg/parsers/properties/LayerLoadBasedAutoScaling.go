package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// LayerLoadBasedAutoScaling Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html
type LayerLoadBasedAutoScaling struct {
	Enable      interface{}                 `yaml:"Enable,omitempty"`
	DownScaling *LayerAutoScalingThresholds `yaml:"DownScaling,omitempty"`
	UpScaling   *LayerAutoScalingThresholds `yaml:"UpScaling,omitempty"`
}

// LayerLoadBasedAutoScaling validation
func (resource LayerLoadBasedAutoScaling) Validate() []error {
	errors := []error{}

	return errors
}
