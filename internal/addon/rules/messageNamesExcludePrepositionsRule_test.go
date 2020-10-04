package rules_test

import (
	"reflect"
	"testing"

	"github.com/yoheimuta/go-protoparser/v4/parser/meta"

	"github.com/yoheimuta/go-protoparser/v4/parser"

	"github.com/tyhal/protolint/internal/addon/rules"
	"github.com/tyhal/protolint/linter/report"
)

func TestMessageNamesExcludePrepositionsRule_Apply(t *testing.T) {
	tests := []struct {
		name              string
		inputProto        *parser.Proto
		inputPrepositions []string
		inputExcludes     []string
		wantFailures      []report.Failure
	}{
		{
			name: "no failures for proto without messages",
			inputProto: &parser.Proto{
				ProtoBody: []parser.Visitee{
					&parser.Enum{},
				},
			},
		},
		{
			name: "no failures for proto with valid message names",
			inputProto: &parser.Proto{
				ProtoBody: []parser.Visitee{
					&parser.Service{},
					&parser.Message{
						MessageName: "AccountStatus",
					},
				},
			},
		},
		{
			name: "failures for proto with invalid message names",
			inputProto: &parser.Proto{
				ProtoBody: []parser.Visitee{
					&parser.Message{
						MessageName: "AccountStatus",
						MessageBody: []parser.Visitee{
							&parser.Message{
								MessageName: "StatusOfAccount",
								Meta: meta.Meta{
									Pos: meta.Position{
										Filename: "example.proto",
										Offset:   100,
										Line:     5,
										Column:   10,
									},
								},
							},
							&parser.Message{
								MessageName: "WithAccountForActive",
								Meta: meta.Meta{
									Pos: meta.Position{
										Filename: "example.proto",
										Offset:   200,
										Line:     10,
										Column:   20,
									},
								},
							},
						},
					},
				},
			},
			wantFailures: []report.Failure{
				report.Failuref(
					meta.Position{
						Filename: "example.proto",
						Offset:   100,
						Line:     5,
						Column:   10,
					},
					"MESSAGE_NAMES_EXCLUDE_PREPOSITIONS",
					`Message name "StatusOfAccount" should not include a preposition "Of"`,
				),
				report.Failuref(
					meta.Position{
						Filename: "example.proto",
						Offset:   200,
						Line:     10,
						Column:   20,
					},
					"MESSAGE_NAMES_EXCLUDE_PREPOSITIONS",
					`Message name "WithAccountForActive" should not include a preposition "With"`,
				),
				report.Failuref(
					meta.Position{
						Filename: "example.proto",
						Offset:   200,
						Line:     10,
						Column:   20,
					},
					"MESSAGE_NAMES_EXCLUDE_PREPOSITIONS",
					`Message name "WithAccountForActive" should not include a preposition "For"`,
				),
			},
		},
		{
			name: "failures for proto with invalid message names, but message name including the excluded keyword is no problem",
			inputProto: &parser.Proto{
				ProtoBody: []parser.Visitee{
					&parser.Message{
						MessageName: "AccountStatus",
						MessageBody: []parser.Visitee{
							&parser.Message{
								MessageName: "SpecialEndOfSupport",
								Meta: meta.Meta{
									Pos: meta.Position{
										Filename: "example.proto",
										Offset:   100,
										Line:     5,
										Column:   10,
									},
								},
							},
							&parser.Message{
								MessageName: "EndOfSales",
								Meta: meta.Meta{
									Pos: meta.Position{
										Filename: "example.proto",
										Offset:   200,
										Line:     10,
										Column:   20,
									},
								},
							},
						},
					},
				},
			},
			inputExcludes: []string{
				"EndOfSupport",
			},
			wantFailures: []report.Failure{
				report.Failuref(
					meta.Position{
						Filename: "example.proto",
						Offset:   200,
						Line:     10,
						Column:   20,
					},
					"MESSAGE_NAMES_EXCLUDE_PREPOSITIONS",
					`Message name "EndOfSales" should not include a preposition "Of"`,
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			rule := rules.NewMessageNamesExcludePrepositionsRule(test.inputPrepositions, test.inputExcludes)

			got, err := rule.Apply(test.inputProto)
			if err != nil {
				t.Errorf("got err %v, but want nil", err)
				return
			}
			if !reflect.DeepEqual(got, test.wantFailures) {
				t.Errorf("got %v, but want %v", got, test.wantFailures)
			}
		})
	}
}
