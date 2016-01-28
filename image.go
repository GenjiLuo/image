package image

import (
	"github.com/gonum/matrix/mat64"
	"image"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func ConvertToRGBA(img image.Image) *image.RGBA {
	imgBounds := img.Bounds()

	RGBAImage := image.NewRGBA(imgBounds)
	// copy the original image to the RGBAImage
	draw.Src.Draw(RGBAImage, imgBounds, img, image.ZP)

	return RGBAImage
}

func RGBAToMatrix(img *image.RGBA) mat64.Matrix {
	bounds := img.Bounds()
	imgSize := bounds.Size()
	rows := imgSize.X * imgSize.Y

	imgMatrix := mat64.NewDense(rows, 4, convertPixToFloat64(img.Pix))

	return imgMatrix
}

func convertPixToFloat64(pix []uint8) []float64 {
	conv := make([]float64, len(pix))
	for i := 0; i < len(pix); i++ {
		conv[i] = float64(pix[i])
	}

	return conv
}

func convertFloat64toPix(data []float64) []uint8 {
	conv := make([]uint8, len(data))
	for i := 0; i < len(data); i++ {
		conv[i] = uint8(data[i])
	}

	return conv
}

func MatrixToRGBA(M mat64.RawMatrixer, dimensions image.Point) *image.RGBA {
	bounds := image.Rect(0, 0, dimensions.X, dimensions.Y)
	RGBAImage := image.NewRGBA(bounds)
	RGBAImage.Pix = convertFloat64toPix(M.RawMatrix().Data)
	RGBAImage.Stride = dimensions.X * 4

	return RGBAImage
}
