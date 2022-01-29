package serifu

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		io io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *Script
		wantErr bool
	}{
		{
			"test",
			args{
				strings.NewReader(""),
			},
			&Script{Pages: []*Page{}},
			false,
		},
		{
			"test",
			args{
				strings.NewReader(`- panel`),
			},
			nil,
			true,
		},
		{
			"test",
			args{
				strings.NewReader(`! asd`),
			},
			nil,
			true,
		},
		{
			"test",
			args{
				strings.NewReader(`Test: One`),
			},
			nil,
			true,
		},
		{
			"test",
			args{
				strings.NewReader(`* test`),
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.io)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
