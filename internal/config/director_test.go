package config

import (
	"reflect"
	"testing"
)

func TestNewDirector(t *testing.T) {
	type args struct {
		b Builder
	}
	tests := []struct {
		name string
		args args
		want *Director
	}{
		{
			name: "Test the factory func for creating new (osx) directors",
			args: args{
				b: GetBuilder(OSX),
			},
			want: &Director{builder: GetBuilder(OSX)},
		},
		{
			name: "Test the factory func for creating new (linux) directors",
			args: args{
				b: GetBuilder(Linux),
			},
			want: &Director{builder: GetBuilder(Linux)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDirector(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDirector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirector_MakeConfig(t *testing.T) {
	deps := []string{
		"stow",
		"git",
		"shellcheck",
		"socat",
		"git-extras",
		"libnotify-bin",
	}

	type fields struct {
		builder Builder
	}
	tests := []struct {
		name   string
		fields fields
		want   Product
	}{
		{
			name: "Test the creation of the ConfigProduct (osx)",
			fields: fields{
				builder: GetBuilder(OSX),
			},
			want: Product{
				packageManager: "brew",
				dependencies:   deps,
			},
		},
		{
			name: "Test the creation of the ConfigProduct (linux)",
			fields: fields{
				builder: GetBuilder(Linux),
			},
			want: Product{
				packageManager: "apt",
				dependencies:   deps,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Director{
				builder: tt.fields.builder,
			}
			if got := d.MakeConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Director.MakeConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
