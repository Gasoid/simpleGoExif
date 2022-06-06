package simpleGoExif

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	img = "naruto.jpg"
)

func TestOpen(t *testing.T) {
	var e *ExifError
	tests := []struct {
		name string
		arg  string
		want interface{}
	}{
		{
			name: "no error",
			arg:  img,
			want: nil,
		},
		{
			name: "error",
			arg:  "naruto.jp",
			want: &e,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Open(tt.arg)
			if err == nil {
				assert.Equal(t, tt.want, err)
			} else {
				assert.ErrorAs(t, err, tt.want)
			}
		})
	}
}
