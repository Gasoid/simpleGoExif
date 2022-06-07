package simpleGoExif

import (
	"time"

	exif "github.com/dsoprea/go-exif/v2"
)

// GetTagValueString returns the value of tag.
// the method works only for string value
func (f *Image) GetTagValueString(tag string) string {
	rootIfd, _, err := f.sl.Exif()
	if err != nil {
		return ""
	}
	results, err := rootIfd.FindTagWithName(tag)
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

// GetDescription returns the description of the image.
func (f *Image) GetDescription() string {
	return f.GetTagValueString(ImageDescriptionTag)
}

// GetTime returns the time of the image.
func (f *Image) GetTime() time.Time {
	datetime := f.GetTagValueString(DateTimeTag)
	if datetime == "" {
		return time.Time{}
	}
	timestamp, err := exif.ParseExifFullTimestamp(datetime)
	if err != nil {
		return time.Time{}
	}
	return timestamp
}
