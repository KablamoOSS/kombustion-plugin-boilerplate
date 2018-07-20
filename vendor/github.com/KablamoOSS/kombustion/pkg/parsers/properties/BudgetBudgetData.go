package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BudgetBudgetData Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html
type BudgetBudgetData struct {
	BudgetName  interface{}       `yaml:"BudgetName,omitempty"`
	BudgetType  interface{}       `yaml:"BudgetType"`
	CostFilters interface{}       `yaml:"CostFilters,omitempty"`
	TimeUnit    interface{}       `yaml:"TimeUnit"`
	TimePeriod  *BudgetTimePeriod `yaml:"TimePeriod,omitempty"`
	BudgetLimit *BudgetSpend      `yaml:"BudgetLimit,omitempty"`
	CostTypes   *BudgetCostTypes  `yaml:"CostTypes,omitempty"`
}

// BudgetBudgetData validation
func (resource BudgetBudgetData) Validate() []error {
	errors := []error{}

	if resource.BudgetType == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'BudgetType'"))
	}
	if resource.TimeUnit == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'TimeUnit'"))
	}
	return errors
}
