package config

import "github.com/tyhal/protolint/internal/stringsutil"

// Rules represents the enabled rule set.
type Rules struct {
	NoDefault  bool     `yaml:"no_default"`
	AllDefault bool     `yaml:"all_default"`
	Add        []string `yaml:"add"`
	Remove     []string `yaml:"remove"`
}

func (r Rules) shouldSkipRule(
	ruleID string,
	defaultRuleIDs []string,
) bool {
	var ruleIDs []string
	if !r.NoDefault {
		ruleIDs = append(ruleIDs, defaultRuleIDs...)
	}

	for _, add := range r.Add {
		ruleIDs = append(ruleIDs, add)
	}

	var newRuleIDs []string
	for _, id := range ruleIDs {
		if !stringsutil.ContainsStringInSlice(id, r.Remove) {
			newRuleIDs = append(newRuleIDs, id)
		}
	}

	return !stringsutil.ContainsStringInSlice(ruleID, newRuleIDs)
}
