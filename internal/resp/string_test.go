package resp

import (
	"bufio"
	"bytes"
	"testing"
)

func Test_decodeString(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		data    string
		want    Value
		wantErr bool
	}{
		{
			name: "OK",
			data: "+OK\r\n",
			want: NewString("OK"),
		},
		{
			name: "empty",
			data: "+\r\n",
			want: NewString(""),
		},
		{
			name:    "Invalid string",
			data:    "+Has\nnew line\r\n",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := bufio.NewReader(bytes.NewReader([]byte(tt.data)))
			r.ReadByte()
			got, gotErr := decodeString(r)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("decodeString() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("decodeString() succeeded unexpectedly")
			}
			if !got.Equals(tt.want) {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
