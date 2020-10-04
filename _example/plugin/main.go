package main

import (
	"github.com/tyhal/protolint/_example/plugin/customrules"
	"github.com/tyhal/protolint/internal/addon/rules"
	"github.com/tyhal/protolint/linter/rule"
	"github.com/tyhal/protolint/plugin"
)

func main() {
	plugin.RegisterCustomRules(
		// The purpose of this line just illustrates that you can implement the same as internal linter rules.
		rules.NewEnumsHaveCommentRule(true),

		// A common custom rule example. It's simple.
		customrules.NewEnumNamesLowerSnakeCaseRule(),

		// Wrapping with RuleGen allows referring to command-line flags.
		plugin.RuleGen(func(
			verbose bool,
			fixMode bool,
		) rule.Rule {
			return customrules.NewSimpleRule(verbose, fixMode)
		}),
	)
}
