package rules_test

import (
	"reflect"
	"testing"

	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/go-protoparser/v4/parser/meta"

	"github.com/tyhal/protolint/internal/addon/rules"
	"github.com/tyhal/protolint/linter/report"
)

func TestServiceNamesUpperCamelCaseRule_Apply(t *testing.T) {
	tests := []struct {
		name         string
		inputProto   *parser.Proto
		wantFailures []report.Failure
	}{
		{
			name: "no failures for proto without service",
			inputProto: &parser.Proto{
				ProtoBody: []parser.Visitee{
					&parser.Message{},
				},
			},
		},
		{
			name: "no failures for proto with valid service names",
			inputProto: &parser.Proto{
				ProtoBody: []parser.Visitee{
					&parser.Service{
						ServiceName: "ServiceName",
					},
					&parser.Service{
						ServiceName: "ServiceName2",
					},
				},
			},
		},
		{
			name: "failures for proto with invalid service names",
			inputProto: &parser.Proto{
				ProtoBody: []parser.Visitee{
					&parser.Service{
						ServiceName: "serviceName",
						Meta: meta.Meta{
							Pos: meta.Position{
								Filename: "example.proto",
								Offset:   100,
								Line:     5,
								Column:   10,
							},
						},
					},
					&parser.Service{
						ServiceName: "Service_name",
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
			wantFailures: []report.Failure{
				report.Failuref(
					meta.Position{
						Filename: "example.proto",
						Offset:   100,
						Line:     5,
						Column:   10,
					},
					"SERVICE_NAMES_UPPER_CAMEL_CASE",
					`Service name "serviceName" must be UpperCamelCase`,
				),
				report.Failuref(
					meta.Position{
						Filename: "example.proto",
						Offset:   200,
						Line:     10,
						Column:   20,
					},
					"SERVICE_NAMES_UPPER_CAMEL_CASE",
					`Service name "Service_name" must be UpperCamelCase`,
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			rule := rules.NewServiceNamesUpperCamelCaseRule()

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
