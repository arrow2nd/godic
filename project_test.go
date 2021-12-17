package main

import (
	"os"
	"strconv"
	"testing"
)

func TestCodic_GetUserProjectById(t *testing.T) {
	c := NewClient(os.Getenv("CODIC_API_KEY"))

	tests := []struct {
		name    string
		id      string
		want    string
		wantErr bool
	}{
		{
			name:    "取得できるか",
			id:      os.Getenv("PROJECT_ID"),
			want:    os.Getenv("PROJECT_NAME"),
			wantErr: false,
		},
		{
			name:    "存在しないIDの場合にエラーが返るか",
			id:      "1234567890",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetUserProjectById(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Codic.GetUserProjectById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if got.Name != tt.want {
				t.Errorf("got[0].Name = %v, want = %v", got.Name, tt.want)
				return
			}
		})
	}
}

func TestCodic_GetUserProjectList(t *testing.T) {
	c := NewClient(os.Getenv("CODIC_API_KEY"))

	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "取得できるか",
			want:    os.Getenv("PROJECT_ID"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetUserProjectList()
			if (err != nil) != tt.wantErr {
				t.Errorf("Codic.GetUserProjectList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			wantId, _ := strconv.ParseUint(tt.want, 10, 64)
			if got[0].Id != wantId {
				t.Errorf("got[0].Id = %v, want = %v", got[0].Id, wantId)
				return
			}
		})
	}
}
