package godic

import (
	"os"
	"testing"
)

func TestCodic_GetTranslate(t *testing.T) {
	c := NewClient(os.Getenv("CODIC_API_KEY"))

	type params struct {
		text   string
		casing string
	}

	tests := []struct {
		name    string
		params  params
		want    string
		wantErr bool
	}{
		{
			name: "変換できるか",
			params: params{
				text:   "こんにちは世界",
				casing: "camel",
			},
			want:    "helloWorld",
			wantErr: false,
		},
		{
			name: "パラメータ不足でのエラーが返るか",
			params: params{
				text:   "",
				casing: "",
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CreateTranslateParam(tt.params.text)
			p.Add("casing", tt.params.casing)

			got, err := c.GetTranslate(p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Codic.GetTranslate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if got[0].TranslatedText != tt.want {
				t.Errorf("got[0].TranslatedText = %v, want = %v", got[0].TranslatedText, tt.want)
				return
			}
		})
	}
}
