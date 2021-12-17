package godic

import (
	"net/url"
	"os"
	"testing"
)

func TestCodic_GetCEDLookup(t *testing.T) {
	c := NewClient(os.Getenv("CODIC_API_KEY"))

	type params struct {
		query string
		count string
	}

	tests := []struct {
		name    string
		params  params
		want    string
		wantErr bool
	}{
		{
			name: "取得できるか",
			params: params{
				query: "terminal",
				count: "2",
			},
			want:    "terminal",
			wantErr: false,
		},
		{
			name: "パラメータ不足でのエラーが返るか",
			params: params{
				query: "",
				count: "",
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := url.Values{}
			v.Add("query", tt.params.query)
			v.Add("count", tt.params.count)

			got, err := c.GetCEDLookup(v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Codic.GetCEDLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if got[0].Title != tt.want {
				t.Errorf("got[0].Title = %v, want = %v", got[0].Title, tt.want)
				return
			}
		})
	}
}

func TestCodic_GetCEDEntries(t *testing.T) {
	c := NewClient(os.Getenv("CODIC_API_KEY"))

	tests := []struct {
		name    string
		id      string
		want    string
		wantErr bool
	}{
		{
			name:    "取得できるか",
			id:      "41941",
			want:    "term",
			wantErr: false,
		},
		{
			name:    "パラメータ不足でのエラーが返るか",
			id:      "",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetCEDEntries(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Codic.GetCEDEntries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if got.Title != tt.want {
				t.Errorf("got.Title = %v, want = %v", got.Title, tt.want)
				return
			}
		})
	}
}
