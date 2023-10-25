package services

import (
	"connectivity-checker-wizard/models"
	r "connectivity-checker-wizard/rules"
)

var ruleMap = map[string]r.Rule{
	"addressValidationRule": buildAddressValidationRule(),
	"sourceNSCNPRule":       buildSourceNSCNPRule(),
}

func RuleHandler(ruleName string) models.ResponseData {
	// find the rule from ruleMap
	rule := getRuleHandlerByName(ruleName)
	// execute  rule
	return rule.Execute(nil)
}

func getRuleHandlerByName(ruleName string) r.Rule {
	for name, r := range ruleMap {
		if name == ruleName {
			return r
		}
	}
	return nil
}

func buildAddressValidationRule() *r.AddressValidationRule {
	rule := &r.AddressValidationRule{}
	rule.Name = "AddressValidationRule"
	rule.SetNextRule(nil)
	return rule
}

func buildSourceNSCNPRule() *r.SourceNSCNPRule {
	rule := &r.SourceNSCNPRule{}
	rule.Name = "SourceNSCNPRule"
	return rule
}
