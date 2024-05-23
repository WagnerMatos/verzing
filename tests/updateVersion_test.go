package main

import (
	"errors"
	"testing"
)

// MockVersioner is a mock implementation of the Versioner interface.
type MockVersioner struct {
	readVersion  func() (string, error)
	writeVersion func(version string) error
}

func (m MockVersioner) ReadVersion() (string, error) {
	return m.readVersion()
}

func (m MockVersioner) WriteVersion(version string) error {
	return m.writeVersion(version)
}

func TestUpdateVersion(t *testing.T) {
	tests := []struct {
		name           string
		readVersion    func() (string, error)
		writeVersion   func(string) error
		commitType     string
		breakingChange bool
		want           string
		wantErr        bool
	}{
		{
			name: "basic non-breaking feature update",
			readVersion: func() (string, error) {
				return "0.1.0", nil
			},
			writeVersion: func(version string) error {
				if version != "0.2.0" {
					t.Errorf("expected version to be 0.2.0, got %s", version)
				}
				return nil
			},
			commitType:     "Feat",
			breakingChange: false,
			want:           "0.2.0",
			wantErr:        false,
		},
		{
			name: "breaking change update",
			readVersion: func() (string, error) {
				return "0.1.0", nil
			},
			writeVersion: func(version string) error {
				if version != "1.0.0" {
					t.Errorf("expected version to be 1.0.0, got %s", version)
				}
				return nil
			},
			commitType:     "Feat",
			breakingChange: true,
			want:           "1.0.0",
			wantErr:        false,
		},
		{
			name: "error reading version",
			readVersion: func() (string, error) {
				return "", errors.New("failed to read file")
			},
			writeVersion: func(version string) error {
				return nil
			},
			commitType:     "Feat",
			breakingChange: false,
			want:           "",
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := MockVersioner{
				readVersion:  tt.readVersion,
				writeVersion: tt.writeVersion,
			}
			got, err := updateVersion(mock, tt.commitType, tt.breakingChange)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("updateVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
