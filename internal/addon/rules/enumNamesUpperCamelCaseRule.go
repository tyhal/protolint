package rules

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"

	"github.com/tyhal/protolint/linter/report"
	"github.com/tyhal/protolint/linter/strs"
	"github.com/tyhal/protolint/linter/visitor"
)

// EnumNamesUpperCamelCaseRule verifies that all enum names are CamelCase (with an initial capital).
// See https://developers.google.com/protocol-buffers/docs/style#enums.
type EnumNamesUpperCamelCaseRule struct{}

// NewEnumNamesUpperCamelCaseRule creates a new EnumNamesUpperCamelCaseRule.
func NewEnumNamesUpperCamelCaseRule() EnumNamesUpperCamelCaseRule {
	return EnumNamesUpperCamelCaseRule{}
}

// ID returns the ID of this rule.
func (r EnumNamesUpperCamelCaseRule) ID() string {
	return "ENUM_NAMES_UPPER_CAMEL_CASE"
}

// Purpose returns the purpose of this rule.
func (r EnumNamesUpperCamelCaseRule) Purpose() string {
	return "Verifies that all enum names are CamelCase (with an initial capital)."
}

// IsOfficial decides whether or not this rule belongs to the official guide.
func (r EnumNamesUpperCamelCaseRule) IsOfficial() bool {
	return true
}

// Apply applies the rule to the proto.
func (r EnumNamesUpperCamelCaseRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	v := &enumNamesUpperCamelCaseVisitor{
		BaseAddVisitor: visitor.NewBaseAddVisitor(r.ID()),
	}
	return visitor.RunVisitor(v, proto, r.ID())
}

type enumNamesUpperCamelCaseVisitor struct {
	*visitor.BaseAddVisitor
}

// VisitEnum checks the enum.
func (v *enumNamesUpperCamelCaseVisitor) VisitEnum(enum *parser.Enum) bool {
	if !strs.IsUpperCamelCase(enum.EnumName) {
		v.AddFailuref(enum.Meta.Pos, "Enum name %q must be UpperCamelCase", enum.EnumName)
	}
	return false
}
