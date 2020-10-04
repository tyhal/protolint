package rules_test

import (
	"reflect"
	"testing"

	"github.com/yoheimuta/go-protoparser/v4/parser/meta"

	"github.com/yoheimuta/go-protoparser/v4/parser"

	"github.com/tyhal/protolint/internal/addon/rules"
	"github.com/tyhal/protolint/linter/report"
)

func TestFieldNamesExcludePrepositionsRule_Apply(t *testing.T) {
	tests := []struct {
		name              string
		inputProto        *parser.Proto
		inputPrepositions []string
		inputExcludes     []string
		wantFailures      []report.Failure
	}{
		{
			name: "no failures for proto without fields",
			inputProto: &parser.Proto{
				ProtoBody: []parser.Visitee{
					&parser.Enum{},
				},
			},
		},
		{
			name: "no failures for proto with valid field names",
			inputProto: &parser.Proto{
				ProtoBody: []parser.Visitee{
					&parser.Service{},
					&parser.Message{
						MessageBody: []parser.Visitee{
							&parser.Field{
								FieldName: "error_reason",
							},
							&parser.Field{
								FieldName: "failure_time_cpu_usage",
							},
							&parser.MapField{
								MapName: "song_name2",
							},
							&parser.Oneof{
								OneofFields: []*parser.OneofField{
									{
										FieldName: "song_name3",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "failures for proto with invalid field names",
			inputProto: &parser.Proto{
				ProtoBody: []parser.Visitee{
					&parser.Message{
						MessageBody: []parser.Visitee{
							&parser.Field{
								FieldName: "reason_for_error",
								Meta: meta.Meta{
									Pos: meta.Position{
										Filename: "example.proto",
										Offset:   100,
										Line:     5,
										Column:   10,
									},
								},
							},
							&parser.Field{
								FieldName: "cpu_usage_at_time_of_failure",
								Meta: meta.Meta{
									Pos: meta.Position{
										Filename: "example.proto",
										Offset:   200,
										Line:     10,
										Column:   20,
									},
								},
							},
							&parser.MapField{
								MapName: "name_of_song",
								Meta: meta.Meta{
									Pos: meta.Position{
										Filename: "example.proto",
										Offset:   210,
										Line:     14,
										Column:   30,
									},
								},
							},
							&parser.Oneof{
								OneofFields: []*parser.OneofField{
									{
										FieldName: "name_of_song2",
										Meta: meta.Meta{
											Pos: meta.Position{
												Filename: "example.proto",
												Offset:   300,
												Line:     21,
												Column:   45,
											},
										},
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
					"FIELD_NAMES_EXCLUDE_PREPOSITIONS",
					`Field name "reason_for_error" should not include a preposition "for"`,
				),
				report.Failuref(
					meta.Position{
						Filename: "example.proto",
						Offset:   200,
						Line:     10,
						Column:   20,
					},
					"FIELD_NAMES_EXCLUDE_PREPOSITIONS",
					`Field name "cpu_usage_at_time_of_failure" should not include a preposition "at"`,
				),
				report.Failuref(
					meta.Position{
						Filename: "example.proto",
						Offset:   200,
						Line:     10,
						Column:   20,
					},
					"FIELD_NAMES_EXCLUDE_PREPOSITIONS",
					`Field name "cpu_usage_at_time_of_failure" should not include a preposition "of"`,
				),
				report.Failuref(
					meta.Position{
						Filename: "example.proto",
						Offset:   210,
						Line:     14,
						Column:   30,
					},
					"FIELD_NAMES_EXCLUDE_PREPOSITIONS",
					`Field name "name_of_song" should not include a preposition "of"`,
				),
				report.Failuref(
					meta.Position{
						Filename: "example.proto",
						Offset:   300,
						Line:     21,
						Column:   45,
					},
					"FIELD_NAMES_EXCLUDE_PREPOSITIONS",
					`Field name "name_of_song2" should not include a preposition "of"`,
				),
			},
		},
		{
			name: "failures for proto with invalid field names, but field name including the excluded keyword is no problem",
			inputProto: &parser.Proto{
				ProtoBody: []parser.Visitee{
					&parser.Message{
						MessageBody: []parser.Visitee{
							&parser.Field{
								FieldName: "end_of_support_version",
							},
							&parser.MapField{
								MapName: "end_of_support_version",
							},
							&parser.Oneof{
								OneofFields: []*parser.OneofField{
									{
										FieldName: "end_of_support_version",
									},
								},
							},
							&parser.Field{
								FieldName: "version_of_support_end",
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
				"end_of_support",
			},
			wantFailures: []report.Failure{
				report.Failuref(
					meta.Position{
						Filename: "example.proto",
						Offset:   200,
						Line:     10,
						Column:   20,
					},
					"FIELD_NAMES_EXCLUDE_PREPOSITIONS",
					`Field name "version_of_support_end" should not include a preposition "of"`,
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			rule := rules.NewFieldNamesExcludePrepositionsRule(test.inputPrepositions, test.inputExcludes)

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
