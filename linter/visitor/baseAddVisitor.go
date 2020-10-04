package visitor

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/go-protoparser/v4/parser/meta"

	"github.com/tyhal/protolint/linter/report"
)

// BaseAddVisitor represents a base visitor which can accumulate failures.
type BaseAddVisitor struct {
	BaseVisitor
	ruleID   string
	failures []report.Failure
}

// NewBaseAddVisitor creates a BaseAddVisitor.
func NewBaseAddVisitor(ruleID string) *BaseAddVisitor {
	return &BaseAddVisitor{
		ruleID: ruleID,
	}
}

// Failures returns the accumulated failures.
func (v *BaseAddVisitor) Failures() []report.Failure {
	return v.failures
}

// AddFailuref adds to the internal buffer and the formatting works like fmt.Sprintf.
func (v *BaseAddVisitor) AddFailuref(
	pos meta.Position,
	format string,
	a ...interface{},
) {
	v.failures = append(v.failures, report.Failuref(pos, v.ruleID, format, a...))
}

// AddFailurefWithProtoMeta adds to the internal buffer and the formatting works like fmt.Sprintf.
func (v *BaseAddVisitor) AddFailurefWithProtoMeta(
	p *parser.ProtoMeta,
	format string,
	a ...interface{},
) {
	v.AddFailuref(
		meta.Position{
			Filename: p.Filename,
			Offset:   0,
			Line:     1,
			Column:   1,
		},
		format,
		a...,
	)
}
