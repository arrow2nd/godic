package godic

import (
	"net/url"
	"reflect"
	"testing"
)

func TestCreateTranslateParam(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want url.Values
	}{
		{
			name: "生成できるか",
			args: args{
				text: "こんにちは世界",
			},
			want: url.Values{
				"text": []string{"こんにちは世界"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateTranslateParam(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTranslateParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateCEDLookupParam(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want url.Values
	}{
		{
			name: "生成できるか",
			args: args{
				query: "hello",
			},
			want: url.Values{
				"query": []string{"hello"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateCEDLookupParam(tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCEDLookupParam() = %v, want %v", got, tt.want)
			}
		})
	}
}
