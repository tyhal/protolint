package rules_test

import (
	"reflect"
	"testing"

	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/go-protoparser/v4/parser/meta"

	"github.com/tyhal/protolint/internal/addon/rules"
	"github.com/tyhal/protolint/linter/report"
)

func TestRPCNamesUpperCamelCaseRule_Apply(t *testing.T) {
	tests := []struct {
		name         string
		inputProto   *parser.Proto
		wantFailures []report.Failure
	}{
		{
			name: "no failures for proto without rpc",
			inputProto: &parser.Proto{
				ProtoBody: []parser.Visitee{
					&parser.Service{},
				},
			},
		},
		{
			name: "no failures for proto with valid rpc",
			inputProto: &parser.Proto{
				ProtoBody: []parser.Visitee{
					&parser.Service{
						ServiceBody: []parser.Visitee{
							&parser.RPC{
								RPCName: "RPCName",
							},
						},
					},
				},
			},
		},
		{
			name: "failures for proto with LowerCamelCase",
			inputProto: &parser.Proto{
				ProtoBody: []parser.Visitee{
					&parser.Service{
						ServiceBody: []parser.Visitee{
							&parser.RPC{
								RPCName: "rpcName",
								Meta: meta.Meta{
									Pos: meta.Position{
										Filename: "example.proto",
										Offset:   100,
										Line:     5,
										Column:   10,
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
					"RPC_NAMES_UPPER_CAMEL_CASE",
					`RPC name "rpcName" must be UpperCamelCase`,
				),
			},
		},
		{
			name: "a failure for proto with SnakeCase",
			inputProto: &parser.Proto{
				ProtoBody: []parser.Visitee{
					&parser.Service{
						ServiceBody: []parser.Visitee{
							&parser.RPC{
								RPCName: "RPC_name",
								Meta: meta.Meta{
									Pos: meta.Position{
										Filename: "example.proto",
										Offset:   100,
										Line:     5,
										Column:   10,
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
					"RPC_NAMES_UPPER_CAMEL_CASE",
					`RPC name "RPC_name" must be UpperCamelCase`,
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			rule := rules.NewRPCNamesUpperCamelCaseRule()

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
