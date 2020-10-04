package reporters_test

import (
	"bytes"
	"testing"

	"github.com/yoheimuta/go-protoparser/v4/parser/meta"

	"github.com/tyhal/protolint/internal/linter/report/reporters"
	"github.com/tyhal/protolint/linter/report"
)

func TestPlainReporter_Report(t *testing.T) {
	tests := []struct {
		name          string
		inputFailures []report.Failure
		wantOutput    string
	}{
		{
			name: "Prints failures in the plain format",
			inputFailures: []report.Failure{
				report.Failuref(
					meta.Position{
						Filename: "example.proto",
						Offset:   100,
						Line:     5,
						Column:   10,
					},
					"ENUM_NAMES_UPPER_CAMEL_CASE",
					`EnumField name "fIRST_VALUE" must be CAPITALS_WITH_UNDERSCORES`,
				),
				report.Failuref(
					meta.Position{
						Filename: "example.proto",
						Offset:   200,
						Line:     10,
						Column:   20,
					},
					"ENUM_NAMES_UPPER_CAMEL_CASE",
					`EnumField name "SECOND.VALUE" must be CAPITALS_WITH_UNDERSCORES`,
				),
			},
			wantOutput: `[example.proto:5:10] EnumField name "fIRST_VALUE" must be CAPITALS_WITH_UNDERSCORES
[example.proto:10:20] EnumField name "SECOND.VALUE" must be CAPITALS_WITH_UNDERSCORES
`,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			err := reporters.PlainReporter{}.Report(buf, test.inputFailures)
			if err != nil {
				t.Errorf("got err %v, but want nil", err)
				return
			}
			if buf.String() != test.wantOutput {
				t.Errorf("got %s, but want %s", buf.String(), test.wantOutput)
			}
		})
	}
}
