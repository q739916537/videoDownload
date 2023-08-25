package http_repo

import "testing"

func TestGetAllInfoIds(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "测试结果",
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllInfoIds()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllInfoIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAllInfoIds() got = %v, want %v", got, tt.want)
			}
		})
	}
}
