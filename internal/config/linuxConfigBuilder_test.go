package config

import (
	"reflect"
	"testing"
)

func TestNewLinuxConfigBuilder(t *testing.T) {
	tests := []struct {
		name string
		want *linuxConfigBuilder
	}{
		{
			name: "Expect the packageManager property to have 'apt'",
			want: &linuxConfigBuilder{packageManager: "apt", dependencies: []string{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newLinuxConfigBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinuxConfigBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
