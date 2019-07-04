package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	white := color.RGBA{255, 255, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	// Красим фон в белый.
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)
	// Горизонтальные красные полосы
	for y := 50; y <= 500; y = y + 50 {
		for x := 0; x < 500; x++ {
			img.Set(x, y, red)
		}
	}
	// Вертикальные синии полосы
	for x := 50; x <= 500; x = x + 50 {
		for y := 0; y < 500; y++ {
			img.Set(x, y, blue)
		}
	}

	file, err := os.Create("img.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()

	png.Encode(file, img)

}
