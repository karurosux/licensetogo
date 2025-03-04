package app

import "testing"

func TestLicenseToGoServer_setupMigrations(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lts, err := NewLicenseToGoServer()
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			gotErr := lts.setupMigrations()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("setupMigrations() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("setupMigrations() succeeded unexpectedly")
			}
		})
	}
}
