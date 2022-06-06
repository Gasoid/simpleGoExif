package simpleGoExif

import (
	"bytes"
	"io/ioutil"

	exif "github.com/dsoprea/go-exif/v2"
	exifcommon "github.com/dsoprea/go-exif/v2/common"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure"
)

const (
	ImageDescriptionTag = "ImageDescription"
	DateTimeTag         = "DateTime"
)

// Open creates a new Image struct, initializes the EXIF data, and returns the struct Image
func Open(path string) (*Image, error) {
	image := Image{
		filepath: path,
	}
	err := image.initExif()
	if err != nil {
		return nil, err
	}

	return &image, nil
}

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
