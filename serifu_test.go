// Package serifu contains parser for the serfu markup language
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
			"works with no content",
			args{
				strings.NewReader(""),
			},
			&Script{Pages: []*Page{}},
			false,
		},
		{
			"panel out of page",
			args{
				strings.NewReader(`- panel`),
			},
			nil,
			true,
		},
		{
			"side note out of panel",
			args{
				strings.NewReader(`! asd`),
			},
			nil,
			true,
		},
		{
			"text entry out of panel",
			args{
				strings.NewReader(`Test: One`),
			},
			nil,
			true,
		},
		{
			"sound effect out of panel",
			args{
				strings.NewReader(`* test`),
			},
			nil,
			true,
		},
		{
			"test",
			args{
				strings.NewReader(`test`),
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

func TestScript_String(t *testing.T) {
	type fields struct {
		Pages []*Page
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"generates empty result",
			fields{
				[]*Page{},
			},
			"",
		},
		{
			"generates result",
			fields{
				[]*Page{
					{
						Title:    "PAGE 1",
						IsSpread: false,
					},
				},
			},
			`# PAGE 1

`,
		},
		{
			"generates result for spread page",
			fields{
				[]*Page{
					{
						Title:    "PAGE 1",
						IsSpread: true,
					},
				},
			},
			`## PAGE 1

`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Script{
				Pages: tt.fields.Pages,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("Script.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
