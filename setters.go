package simpleGoExif

import (
	"time"

	"github.com/Gasoid/go-dms/dms"
	exif "github.com/dsoprea/go-exif/v2"
)

// SetDescription sets the description of the image.
func (f *Image) SetDescription(description string) error {
	err := f.ifd0Ib.SetStandardWithName(ImageDescriptionTag, description)
	if err != nil {
		return &ExifError{"ifd0Ib.SetStandardWithName failed", err}
	}
	return nil
}

// SetTime sets the time of the image.
func (f *Image) SetTime(date time.Time) error {
	dateTime := exif.ExifFullTimestampString(date)
	err := f.ifd0Ib.SetStandardWithName(DateTimeTag, dateTime)
	if err != nil {
		return &ExifError{"ifd0Ib.SetStandardWithName failed", err}
	}
	return nil
}

// SetGPS sets the GPS coordinates of the image.
func (f *Image) SetGPS(latitude, longitude float64) error {
	if latitude == 0 || longitude == 0 {
		return nil
	}
	childIb, err := exif.GetOrCreateIbFromRootIb(f.rootIb, "IFD/GPSInfo")
	if err != nil {
		return &ExifError{"exif.GetOrCreateIbFromRootIb failed", err}
	}
	lat, lon, err := dms.NewDMS(latitude, longitude)
	if err != nil {
		return &ExifError{"dms.NewDMS", err}
	}
	updatedGiLat := exif.GpsDegrees{
		Degrees: float64(lat.Degrees),
		Minutes: float64(lat.Minutes),
		Seconds: lat.Seconds,
	}

	err = childIb.SetStandardWithName("GPSLatitude", updatedGiLat.Raw())
	if err != nil {
		return &ExifError{"childIb.SetStandardWithName failed", err}
	}
	updatedGiLong := exif.GpsDegrees{
		Degrees: float64(lon.Degrees),
		Minutes: float64(lon.Minutes),
		Seconds: lon.Seconds,
	}

	err = childIb.SetStandardWithName("GPSLongitude", updatedGiLong.Raw())
	if err != nil {
		return &ExifError{"childIb.SetStandardWithName", err}
	}
	return nil
}
