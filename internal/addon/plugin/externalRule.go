package plugin

import (
	"path/filepath"

	"github.com/tyhal/protolint/internal/addon/plugin/shared"

	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/go-protoparser/v4/parser/meta"

	"github.com/tyhal/protolint/internal/addon/plugin/proto"
	"github.com/tyhal/protolint/linter/report"
)

// externalRule represents a customized rule that works as a plugin.
type externalRule struct {
	id      string
	purpose string
	client  shared.RuleSet
}

func newExternalRule(
	id string,
	purpose string,
	client shared.RuleSet,
) externalRule {
	return externalRule{
		id:      id,
		purpose: purpose,
		client:  client,
	}
}

// ID returns the ID of this rule.
func (r externalRule) ID() string {
	return r.id
}

// Purpose returns the purpose of this rule.
func (r externalRule) Purpose() string {
	return r.purpose
}

// IsOfficial decides whether or not this rule belongs to the official guide.
func (r externalRule) IsOfficial() bool {
	return true
}

// Apply applies the rule to the proto.
func (r externalRule) Apply(p *parser.Proto) ([]report.Failure, error) {
	relPath := p.Meta.Filename
	absPath, err := filepath.Abs(relPath)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Apply(&proto.ApplyRequest{
		Id:   r.id,
		Path: absPath,
	})
	if err != nil {
		return nil, err
	}

	var fs []report.Failure
	for _, f := range resp.Failures {
		fs = append(fs, report.Failuref(meta.Position{
			Filename: relPath,
			Offset:   int(f.Pos.Offset),
			Line:     int(f.Pos.Line),
			Column:   int(f.Pos.Column),
		}, r.id, f.Message))
	}
	return fs, nil
}
