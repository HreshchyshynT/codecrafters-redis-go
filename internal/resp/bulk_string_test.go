package resp

import "testing"

func TestBulkStringDecode(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    string
		want    Value
		wantErr bool
	}{
		{
			name:    "decoding bulk string hello",
			data:    "$5\r\nhello\r\n",
			want:    NewString("hello"),
			wantErr: false,
		},
		{
			name:    "decoding bulk string hello world",
			data:    "$11\r\nhello world\r\n",
			want:    NewString("hello world"),
			wantErr: false,
		},
		{
			name:    "decoding bulk string invalid string",
			data:    "$11\r\nhello worl\r\n",
			want:    EmptyValue(),
			wantErr: true,
		},
		{
			name:    "decoding bulk string NULL",
			data:    "$-1\r\n",
			want:    EmptyValue(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := decodeBulkString(tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Decode() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Decode() succeeded unexpectedly")
			}

			if tt.want.Equals(got) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBulkStringEncode(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    Value
		want    string
		wantErr bool
	}{
		{
			name:    "encoding bulk string hello",
			want:    "$5\r\nhello\r\n",
			data:    NewString("hello"),
			wantErr: false,
		},
		{
			name:    "encoding bulk string hello world",
			want:    "$11\r\nhello world\r\n",
			data:    NewString("hello world"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := encodeBulkString(tt.data)
			if tt.wantErr {
				t.Fatal("Decode() succeeded unexpectedly")
			}

			if tt.want != got {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
