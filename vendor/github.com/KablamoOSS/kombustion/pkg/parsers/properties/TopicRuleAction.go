package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TopicRuleAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html
type TopicRuleAction struct {
	Sqs              *TopicRuleSqsAction              `yaml:"Sqs,omitempty"`
	Sns              *TopicRuleSnsAction              `yaml:"Sns,omitempty"`
	S3               *TopicRuleS3Action               `yaml:"S3,omitempty"`
	Republish        *TopicRuleRepublishAction        `yaml:"Republish,omitempty"`
	Lambda           *TopicRuleLambdaAction           `yaml:"Lambda,omitempty"`
	Kinesis          *TopicRuleKinesisAction          `yaml:"Kinesis,omitempty"`
	Firehose         *TopicRuleFirehoseAction         `yaml:"Firehose,omitempty"`
	Elasticsearch    *TopicRuleElasticsearchAction    `yaml:"Elasticsearch,omitempty"`
	DynamoDBv2       *TopicRuleDynamoDBv2Action       `yaml:"DynamoDBv2,omitempty"`
	DynamoDB         *TopicRuleDynamoDBAction         `yaml:"DynamoDB,omitempty"`
	CloudwatchMetric *TopicRuleCloudwatchMetricAction `yaml:"CloudwatchMetric,omitempty"`
	CloudwatchAlarm  *TopicRuleCloudwatchAlarmAction  `yaml:"CloudwatchAlarm,omitempty"`
}

// TopicRuleAction validation
func (resource TopicRuleAction) Validate() []error {
	errors := []error{}

	return errors
}
