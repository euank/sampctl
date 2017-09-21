package main

import (
	"testing"
)

func Test_fromCache(t *testing.T) {
	type args struct {
		cacheDir string
		version  string
		cwd      string
	}
	tests := []struct {
		name    string
		args    args
		wantHit bool
		wantErr bool
	}{
		{"valid", args{"./testcache", "0.3.7-R2-2-1", "./testspace"}, true, false},
		{"valid", args{"./testcache", "0.4a-RC1", "./testspace"}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHit, err := fromCache(tt.args.cacheDir, tt.args.version, tt.args.cwd)
			if (err != nil) != tt.wantErr {
				t.Errorf("fromCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHit != tt.wantHit {
				t.Errorf("fromCache() = %v, want %v", gotHit, tt.wantHit)
			}
		})
	}
}

func Test_fromNet(t *testing.T) {
	type args struct {
		endpoint string
		cacheDir string
		version  string
		cwd      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{"http://files.sa-mp.com", "./testcache", "0.3.7-R2-2-1", "./testspace"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := fromNet(tt.args.endpoint, tt.args.cacheDir, tt.args.version, tt.args.cwd); (err != nil) != tt.wantErr {
				t.Errorf("fromNet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}