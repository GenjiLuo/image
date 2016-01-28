package cluster

import (
	"github.com/gonum/matrix/mat64"
	kimage "github.com/kingzbauer/image"
	"github.com/kingzbauer/kmeans"
	"image"
	"log"
	"time"
)

// ClusterImage runs the kmeans algorithm on the image data with K = colors
// returning the best kmeans.AppIter
func ClusterImage(imgData mat64.Matrix, colors, runs int) *kmeans.AppIter {
	app := kmeans.NewApp(imgData, runs, colors, 0)

	app.ParallelRun()
	return app.BestIter
}

// RunCluster runs the kmeans algorithm on the image data with K = clusters
// returning the best kmeans.AppIter, the new image and time take to run the algorithm
func RunCluster(img image.Image, clusters, runs int, parallel bool) (imgObj image.Image, iter *kmeans.AppIter, duration time.Duration) {

	// Convert the image to rgba if not already
	RGBAImage, ok := img.(*image.RGBA)
	// if not RGBA, do the convertion
	if !ok {
		log.Println("Explicit convertion to RGBA")
		RGBAImage = kimage.ConvertToRGBA(img)
	}

	// convert the RGBAImage to a matrix
	imageMatrix := kimage.RGBAToMatrix(RGBAImage)

	// create the kmeans app
	app := kmeans.NewApp(imageMatrix, runs, clusters, 0)
	startTime := time.Now()
	if parallel {
		app.ParallelRun()
	} else {
		app.Run()
	}
	duration = time.Now().Sub(startTime) // track time taken

	// convert the matrix back to an image
	iter = app.BestIter
	matrix := iter.ReconstructXMatrix()
	rawMatrixer := matrix.(mat64.RawMatrixer)
	imgObj = kimage.MatrixToRGBA(rawMatrixer, RGBAImage.Bounds().Size())

	return
}
