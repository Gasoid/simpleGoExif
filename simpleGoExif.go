package simpleGoExif

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/Gasoid/go-dms/dms"

	exif "github.com/dsoprea/go-exif/v2"
	exifcommon "github.com/dsoprea/go-exif/v2/common"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure"
)

type ExifError struct {
	text string
	Err  error
}

func (e *ExifError) Error() string {
	return e.text
}

func (e *ExifError) Unwrap() error {
	return e.Err
}

type Image struct {
	filepath string
	ifd0Ib   *exif.IfdBuilder
	rootIb   *exif.IfdBuilder
	sl       *jpegstructure.SegmentList
}

// New creates a new Image struct, initializes the EXIF data, and returns the struct Image
func New(path string) (*Image, error) {
	image := Image{
		filepath: path,
	}
	err := image.initExif()
	if err != nil {
		return nil, err
	}

	return &image, nil
}

func (f *Image) initExif() error {
	jmp := jpegstructure.NewJpegMediaParser()
	intfc, err := jmp.ParseFile(f.filepath)
	if err != nil {
		return &ExifError{"Parsing file failed", err}
	}
	f.sl = intfc.(*jpegstructure.SegmentList)
	f.rootIb, err = f.sl.ConstructExifBuilder()
	if err != nil {
		im := exif.NewIfdMappingWithStandard()
		ti := exif.NewTagIndex()
		err := exif.LoadStandardTags(ti)
		if err != nil {
			return &ExifError{"exif.LoadStandardTags failed", err}
		}

		f.rootIb = exif.NewIfdBuilder(im, ti, exifcommon.IfdPathStandard, exifcommon.EncodeDefaultByteOrder)
	}

	f.ifd0Ib, err = exif.GetOrCreateIbFromRootIb(f.rootIb, "IFD0")
	if err != nil {
		return &ExifError{"exif.GetOrCreateIbFromRootIb failed", err}
	}
	return nil
}

// GetRootIb is a getter for the rootIb.
func (f *Image) GetRootIb() *exif.IfdBuilder {
	return f.rootIb
}

// SetDescription sets the description of the image.
func (f *Image) SetDescription(description string) error {
	err := f.ifd0Ib.SetStandardWithName("ImageDescription", description)
	if err != nil {
		return &ExifError{"ifd0Ib.SetStandardWithName failed", err}
	}
	return nil
}

// SetTime sets the time of the image.
func (f *Image) SetTime(date time.Time) error {
	dateTime := exif.ExifFullTimestampString(date)
	err := f.ifd0Ib.SetStandardWithName("DateTime", dateTime)
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

// Close writes the EXIF data to the file.
func (f *Image) Close() error {
	err := f.sl.SetExif(f.rootIb)
	if err != nil {
		return &ExifError{"Couldn't set exif data", err}
	}
	b := bytes.NewBufferString("")
	err = f.sl.Write(b)
	if err != nil {
		return &ExifError{"Couldn't write file", err}
	}
	ioutil.WriteFile(f.filepath, b.Bytes(), 0666)
	return nil
}
