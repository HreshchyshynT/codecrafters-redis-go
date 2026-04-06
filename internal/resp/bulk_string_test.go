package resp

import "testing"

func TestBulkStringDecode(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    string
		want    string
		wantErr bool
	}{
		{
			name:    "decoding bulk string hello",
			data:    "$5\r\nhello\r\n",
			want:    "hello",
			wantErr: false,
		},
		{
			name:    "decoding bulk string hello world",
			data:    "$11\r\nhello world\r\n",
			want:    "hello world",
			wantErr: false,
		},
		{
			name:    "decoding bulk string invalid string",
			data:    "$11\r\nhello worl\r\n",
			want:    "",
			wantErr: true,
		},
		{
			name:    "decoding bulk string NULL",
			data:    "$-1\r\n",
			want:    "",
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

			if tt.want != got {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBulkStringEncode(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    string
		want    string
		wantErr bool
	}{
		{
			name:    "encoding bulk string hello",
			want:    "$5\r\nhello\r\n",
			data:    "hello",
			wantErr: false,
		},
		{
			name:    "encoding bulk string hello world",
			want:    "$11\r\nhello world\r\n",
			data:    "hello world",
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
