package simpleGoExif

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetDescription(t *testing.T) {
	//var e *ExifError
	tests := []struct {
		name string
		arg  string
		want interface{}
	}{
		{
			name: "no error",
			arg:  "right description",
			want: nil,
		},
	}
	image, _ := Open(img)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := image.SetDescription(tt.arg)
			if err == nil {
				assert.Equal(t, tt.want, err)
			} else {
				assert.ErrorAs(t, err, tt.want)
			}
		})
	}
}

func TestSetTime(t *testing.T) {
	//var e *ExifError
	tests := []struct {
		name string
		arg  time.Time
		want interface{}
	}{
		{
			name: "no error",
			arg:  time.Now(),
			want: nil,
		},
	}
	image, _ := Open(img)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := image.SetTime(tt.arg)
			if err == nil {
				assert.Equal(t, tt.want, err)
			} else {
				assert.ErrorAs(t, err, tt.want)
			}
		})
	}
}
