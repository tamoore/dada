package config

import (
	"reflect"
	"testing"
)

func Test_newOSXConfigBuilder(t *testing.T) {
	tests := []struct {
		name string
		want *osxConfigBuilder
	}{
		{
			name: "Expect the packageManager property to have 'apt'",
			want: &osxConfigBuilder{packageManager: "brew", dependencies: []string{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newOSXConfigBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newOSXConfigBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
