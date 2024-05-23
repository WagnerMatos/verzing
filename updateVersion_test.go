package main

import (
	"fmt"
	"testing"
)

type MockVersioner struct {
	ReadFunc  func() (string, error)
	WriteFunc func(string) error
}

func (m MockVersioner) ReadVersion() (string, error) {
	return m.ReadFunc()
}

func (m MockVersioner) WriteVersion(version string) error {
	return m.WriteFunc(version)
}

func TestUpdateVersion(t *testing.T) {
	tests := []struct {
		name           string
		readVersion    string
		readErr        error
		writeErr       error
		commitType     string
		breakingChange bool
		wantVersion    string
		wantErr        bool
	}{
		{
			name:           "successful minor update",
			readVersion:    "1.2.3",
			commitType:     "Feat",
			breakingChange: false,
			wantVersion:    "1.3.0",
			wantErr:        false,
		},
		{
			name:           "successful patch update",
			readVersion:    "1.2.3",
			commitType:     "Fix",
			breakingChange: false,
			wantVersion:    "1.2.4",
			wantErr:        false,
		},
		{
			name:           "successful major update with breaking changes",
			readVersion:    "1.2.3",
			commitType:     "Feat",
			breakingChange: true,
			wantVersion:    "2.0.0",
			wantErr:        false,
		},
		{
			name:    "failure reading version",
			readErr: fmt.Errorf("read error"),
			wantErr: true,
		},
		{
			name:        "failure writing version",
			readVersion: "1.2.3",
			writeErr:    fmt.Errorf("write error"),
			wantVersion: "1.3.0",
			wantErr:     true,
		},
		{
			name:        "invalid version format",
			readVersion: "1.2",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := MockVersioner{
				ReadFunc: func() (string, error) {
					return tt.readVersion, tt.readErr
				},
				WriteFunc: func(version string) error {
					if version != tt.wantVersion {
						t.Errorf("WriteVersion() = %v, want %v", version, tt.wantVersion)
					}
					return tt.writeErr
				},
			}

			gotVersion, err := updateVersion(mock, tt.commitType, tt.breakingChange)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && gotVersion != tt.wantVersion {
				t.Errorf("updateVersion() = %v, want %v", gotVersion, tt.wantVersion)
			}
		})
	}
}
