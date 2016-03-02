package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

const Modulus = 151
const Permuter = 64

var Encoding bool

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "Usage: image-cipher encode <input> <output>\n"+
			"       image-cipher decode <input> <output>")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "encode":
		Encoding = true
	case "decode":
	default:
		fmt.Fprintln(os.Stderr, "unknown sub-command: "+os.Args[1])
		os.Exit(1)
	}
	if img, err := readInput(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		res := cipherImage(img)
		if err := writeOutput(res); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func readInput() (image.Image, error) {
	stream, err := os.Open(os.Args[2])
	if err != nil {
		return nil, err
	}
	defer stream.Close()
	decoded, _, err := image.Decode(stream)
	return decoded, err
}

func cipherImage(input image.Image) image.Image {
	res := image.NewRGBA(image.Rect(0, 0, input.Bounds().Dx(), input.Bounds().Dy()))
	for x := 0; x <= input.Bounds().Dx()-Modulus; x += Modulus {
		for y := 0; y <= input.Bounds().Dy()-Modulus; y += Modulus {
			for i := 0; i < Modulus; i++ {
				for j := 0; j < Modulus; j++ {
					sourceX := x + ((i * Permuter) % Modulus)
					sourceY := y + ((j * Permuter) % Modulus)
					if Encoding {
						pixel := input.At(sourceX+input.Bounds().Min.X, sourceY+input.Bounds().Min.Y)
						res.Set(x+i, y+j, pixel)
					} else {
						pixel := input.At(x+i+input.Bounds().Min.X, y+j+input.Bounds().Min.Y)
						res.Set(sourceX, sourceY, pixel)
					}
				}
			}
		}
	}
	return res
}

func writeOutput(i image.Image) error {
	file := os.Args[3]
	output, err := os.Create(file)
	if err != nil {
		return err
	}
	defer output.Close()
	if strings.HasSuffix(file, ".png") {
		return png.Encode(output, i)
	} else if strings.HasSuffix(file, ".jpg") || strings.HasSuffix(file, ".jpeg") {
		return jpeg.Encode(output, i, &jpeg.Options{Quality: 100})
	} else {
		return errors.New("unknown output format")
	}
}
