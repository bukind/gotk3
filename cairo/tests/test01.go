package main

import (
	"fmt"
	"github.com/gotk3/gotk3/cairo"
)

func main() {
	wx := 200.0
	wy := 40.0

	target, err := cairo.NewSurfaceImage(cairo.FORMAT_ARGB32, int(wx), int(wy))
	if err != nil {
	    panic(err)
	}
	cr := cairo.Create(target)

	fmt.Println("format=", target.GetImageFormat())
	fmt.Println("wx=", target.GetImageWidth(),
		", wy=", target.GetImageHeight())

	cr.SetLineWidth(3.0)
	cr.SetSourceRGB(1.0, 0.0, 0.0)
	cr.MoveTo(0., 0.)
	cr.LineTo(wx, wy)
	cr.Stroke()
	cr.MoveTo(wx, 0.)
	cr.LineTo(0., wy)
	cr.Stroke()

	cr.SelectFontFace("serif", cairo.FONT_SLANT_NORMAL,
		cairo.FONT_WEIGHT_NORMAL)
	cr.SetFontSize(20.0)
	cr.SetSourceRGB(0.0, 0.0, 1.0)
	cr.MoveTo(0., wy/2)
	cr.ShowText("Привет!")

	if err := target.WriteToPng("test.png"); err != nil {
	    panic(err)
	}
}
