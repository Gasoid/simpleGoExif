package simpleGoExif

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDescription(t *testing.T) {
	//var e *ExifError
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "no description",
			arg:  "naruto.jpg",
			want: "",
		},
		{
			name: "description exists",
			arg:  "narutoExif.jpg",
			want: "text1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			image, _ := Open(tt.arg)
			descr := image.GetDescription()
			assert.Equal(t, tt.want, descr)
		})
	}
}

func TestGetTime(t *testing.T) {
	//var e *ExifError
	tests := []struct {
		name string
		arg  string
		want time.Time
	}{
		{
			name: "no time",
			arg:  "naruto.jpg",
			want: time.Time{},
		},
		{
			name: "description exists",
			arg:  "narutoExif.jpg",
			want: time.Date(2022, time.June, 5, 14, 34, 34, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			image, _ := Open(tt.arg)
			date := image.GetTime()
			assert.Equal(t, tt.want, date)
		})
	}
}
