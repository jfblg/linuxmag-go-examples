package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	exif "github.com/xor-gate/goexif2/exif"
)

func usage(msg string) {
	fmt.Printf("%s\n", msg)
	fmt.Printf("usage: %s image.jpg\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func main() {

	if len(os.Args) != 2 {
		usage("Missing argument")
	}

	img := os.Args[1]

	lat, lon, err := geopos(img)
	if err != nil {
		panic(err)
	}

	latFuzz, lonFuzz := fuzz(lat, lon)
	fmt.Printf("What: %f, %f\n", lat, lon)
	fmt.Printf("Fuzz: %f, %f\n", latFuzz, lonFuzz)
	patch(img, latFuzz, lonFuzz)
}

func patch(path string, lat, lon float64) {
	var out bytes.Buffer

	cmd := exec.Command(
		"exiftool",
		path,
		fmt.Sprintf("-gpslatitude=%f", lat),
		fmt.Sprintf("-gpslongitude=%f", lon),
	)

	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		panic(out.String())
	}
}

func geopos(path string) (float64, float64, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}

	x, err := exif.Decode(f)
	if err != nil {
		return 0, 0, err
	}

	lat, lon, err := x.LatLong()
	if err != nil {
		return 0, 0, err
	}

	return lat, lon, nil
}

func fuzz(lat, lon float64) (float64, float64) {
	r := 1000.0 / 111300 // 1km radius

	// secret center
	lat += 0.066
	lon += 0.024

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	u := r1.Float64()
	v := r1.Float64()

	w := r * math.Sqrt(u)
	t := 2.0 * math.Pi * v
	x := w * math.Cos(t)
	y := w * math.Sin(t)

	x = x / math.Cos(lat*math.Pi/180.0)

	return lat + x, lon + y
}
