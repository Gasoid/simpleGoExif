# Simple Go Exif
You can implement this github.com/dsoprea/go-exif/v2 in your code and feel pure pain.
Otherwise you can use my library. I've simplified it.

Please check out following:

```golang
package main

import (
    "time"
    exif "github.com/Gasoid/simpleGoExif"
)

func main() {
    image, err := exif.Open("image.jpg")
    if err != nil {
        fmt.Println("err:", err.Error())
    }
    defer image.Close()
    image.SetDescription("text")
    image.SetTime(time.Now())
    image.SetGPS(float64(52.5219814), float64(13.4111173))
}

```
