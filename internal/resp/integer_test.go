package resp

import "testing"

func Test_decodeInteger(t *testing.T) {
	tests := []struct {
		input   string
		want    Value
		wantErr bool
	}{
		{
			input: ":0\r\n",
			want:  NewInt(0),
		},
		{
			input: ":-0\r\n",
			want:  NewInt(0),
		},
		{
			input: ":+0\r\n",
			want:  NewInt(0),
		},
		{
			input: ":123\r\n",
			want:  NewInt(123),
		},
		{
			input: ":+123\r\n",
			want:  NewInt(123),
		},
		{
			input: ":-123\r\n",
			want:  NewInt(-123),
		},
	}
	t.Run("decodeInteger tests", func(t *testing.T) {
		for _, tt := range tests {
			got, gotErr := decodeInteger(tt.input)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("decodeInteger() want %v failed: %v", tt.want, gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("decodeInteger() succeeded unexpectedly")
			}
			if !got.Equals(tt.want) {
				t.Errorf("decodeInteger() = %v, want %v", got, tt.want)
			}
		}
	})
}
