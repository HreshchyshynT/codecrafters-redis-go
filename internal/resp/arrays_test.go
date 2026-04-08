package resp

import (
	"bufio"
	"bytes"
	"testing"
)

func Test_decodeArray(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		data    string
		want    Value
		wantErr bool
	}{
		{
			name: "decoding empty array",
			data: "*0\r\n",
			want: NewArray([]Value{}),
		},
		{
			name: "decoding array of two bulk strings",
			data: "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n",
			want: NewArray([]Value{
				NewBulkString("hello"),
				NewBulkString("world"),
			}),
		},
		{
			name: "decoding nested array",
			data: "*2\r\n*3\r\n:1\r\n:2\r\n:3\r\n*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n",
			want: NewArray(
				[]Value{
					NewArray([]Value{NewInt(1), NewInt(2), NewInt(3)}),
					NewArray([]Value{NewBulkString("hello"), NewBulkString("world")}),
				},
			),
		},
		{
			name:    "decoding invalid array",
			data:    "*n\r\n*3\r\n:1\r\n:2\r\n:3\r\n*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n",
			want:    NewArray([]Value{}),
			wantErr: true,
		},
		{
			name:    "decoding array invalid length",
			data:    "*3\r\n*3\r\n:1\r\n:2\r\n:3\r\n*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n",
			want:    NewArray([]Value{}),
			wantErr: true,
		},
		{
			name: "combined types",
			data: "*4\r\n*3\r\n:1\r\n:2\r\n:3\r\n$5\r\nhello\r\n+OK\r\n:4\r\n",
			want: NewArray([]Value{
				NewArray([]Value{
					NewInt(1),
					NewInt(2),
					NewInt(3),
				}),
				NewBulkString("hello"),
				NewString("OK"),
				NewInt(4),
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := bufio.NewReader(bytes.NewReader([]byte(tt.data)))
			r.ReadRune()
			got, gotErr := decodeArray(r)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("decodeArray() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("decodeArray() succeeded unexpectedly")
			}
			if !got.Equals(tt.want) {
				t.Errorf("decodeArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeArray(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := encodeArray()
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("encodeArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
