package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

type errorWriter struct{}

func (ew *errorWriter) Write(_ []byte) (n int, err error) {
	return 0, fmt.Errorf("simulated writer error")
}

func TestUpdateChangelog(t *testing.T) {
	cases := []struct {
		name       string
		version    string
		shortDesc  string
		longDesc   string
		wantOutput string
		wantErr    bool
	}{
		{
			name:       "simple update with no long description",
			version:    "v1.0.2",
			shortDesc:  "Fixed bug in user interface",
			longDesc:   "",
			wantOutput: "## v1.0.2\n- Fixed bug in user interface\n",
			wantErr:    false,
		},
		{
			name:       "update with long description",
			version:    "v1.1.0",
			shortDesc:  "Added new login feature",
			longDesc:   "This update provides a new secure login feature that improves security.",
			wantOutput: "## v1.1.0\n- Added new login feature\n  This update provides a new secure login feature that improves security.\n",
			wantErr:    false,
		},
		{
			name:       "handling error from writer",
			version:    "v1.2.0",
			shortDesc:  "Update that fails",
			longDesc:   "This should fail due to writer error.",
			wantOutput: "",
			wantErr:    true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var buf io.Writer = &bytes.Buffer{} // Use interface type for flexibility
			if tc.wantErr {
				buf = &errorWriter{} // Use the error writer to simulate an error
			}

			err := updateChangelog(buf, tc.version, tc.shortDesc, tc.longDesc)
			if (err != nil) != tc.wantErr {
				t.Errorf("updateChangelog() error = %v, wantErr %v", err, tc.wantErr)
			}
			if !tc.wantErr {
				if gotOutput := buf.(*bytes.Buffer).String(); gotOutput != tc.wantOutput {
					t.Errorf("updateChangelog() got output = %v, want %v", gotOutput, tc.wantOutput)
				}
			}
		})
	}
}
