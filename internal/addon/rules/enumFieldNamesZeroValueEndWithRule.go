package rules

import (
	"strings"

	"github.com/yoheimuta/go-protoparser/v4/parser"

	"github.com/tyhal/protolint/linter/report"
	"github.com/tyhal/protolint/linter/visitor"
)

const (
	defaultSuffix = "UNSPECIFIED"
)

// EnumFieldNamesZeroValueEndWithRule verifies that the zero value enum should have the suffix (e.g. "UNSPECIFIED", "INVALID").
// See https://developers.google.com/protocol-buffers/docs/style#enums.
type EnumFieldNamesZeroValueEndWithRule struct {
	suffix string
}

// NewEnumFieldNamesZeroValueEndWithRule creates a new EnumFieldNamesZeroValueEndWithRule.
func NewEnumFieldNamesZeroValueEndWithRule(
	suffix string,
) EnumFieldNamesZeroValueEndWithRule {
	if len(suffix) == 0 {
		suffix = defaultSuffix
	}
	return EnumFieldNamesZeroValueEndWithRule{
		suffix: suffix,
	}
}

// ID returns the ID of this rule.
func (r EnumFieldNamesZeroValueEndWithRule) ID() string {
	return "ENUM_FIELD_NAMES_ZERO_VALUE_END_WITH"
}

// Purpose returns the purpose of this rule.
func (r EnumFieldNamesZeroValueEndWithRule) Purpose() string {
	return `Verifies that the zero value enum should have the suffix (e.g. "UNSPECIFIED", "INVALID").`
}

// IsOfficial decides whether or not this rule belongs to the official guide.
func (r EnumFieldNamesZeroValueEndWithRule) IsOfficial() bool {
	return true
}

// Apply applies the rule to the proto.
func (r EnumFieldNamesZeroValueEndWithRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	v := &enumFieldNamesZeroValueEndWithVisitor{
		BaseAddVisitor: visitor.NewBaseAddVisitor(r.ID()),
		suffix:         r.suffix,
	}
	return visitor.RunVisitor(v, proto, r.ID())
}

type enumFieldNamesZeroValueEndWithVisitor struct {
	*visitor.BaseAddVisitor
	suffix string
}

// VisitEnumField checks the enum field.
func (v *enumFieldNamesZeroValueEndWithVisitor) VisitEnumField(field *parser.EnumField) bool {
	if field.Number == "0" && !strings.HasSuffix(field.Ident, v.suffix) {
		v.AddFailuref(field.Meta.Pos, "EnumField name %q with zero value should have the suffix %q", field.Ident, v.suffix)
	}
	return false
}
