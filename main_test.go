package serifu

import (
	"encoding/json"
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
		want    string
		wantErr bool
	}{
		{
			"test",
			args{
				strings.NewReader(""),
			},
			`{"Pages":[]}`,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.io)
			b, _ := json.Marshal(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(b), tt.want) {
				t.Errorf("Parse() = %v, want %v", string(b), tt.want)
			}
		})
	}
}
