package simpleGoExif

import (
	"time"

	exif "github.com/dsoprea/go-exif/v2"
)

// GetDescription returns the description of the image.
func (f *Image) GetDescription() string {
	rootIfd, _, err := f.sl.Exif()
	if err != nil {
		return ""
	}
	results, err := rootIfd.FindTagWithName(ImageDescriptionTag)
	if err != nil {
		return ""
	}
	if len(results) == 0 {
		return ""
	}
	valueRaw, err := results[0].Value()
	if err != nil {
		return ""
	}
	return valueRaw.(string)
}

// GetTime returns the time of the image.
func (f *Image) GetTime() time.Time {
	rootIfd, _, err := f.sl.Exif()
	if err != nil {
		return time.Time{}
	}
	results, err := rootIfd.FindTagWithName(DateTimeTag)
	if err != nil {
		return time.Time{}
	}
	if len(results) == 0 {
		return time.Time{}
	}
	valueRaw, err := results[0].Value()
	if err != nil {
		return time.Time{}
	}
	datetime := valueRaw.(string)
	timestamp, err := exif.ParseExifFullTimestamp(datetime)
	if err != nil {
		return time.Time{}
	}
	return timestamp
}
