package graphs

import (
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"image"
	"image/color"
	"strings"
)

type graph struct {
	Frame      frame      `xml:"rect"`
	OuterGroup outerGroup `xml:"g"`
}

type frame struct {
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
	X      int    `xml:"x,attr"`
	Y      int    `xml:"y,attr"`
	Fill   string `xml:"fill,attr"`
}

type outerGroup struct {
	Transform  string      `xml:"transform,attr"`
	WeekGroups []weekGroup `xml:"g"`
}

type weekGroup struct {
	Transform string  `xml:"transform,attr"`
	Pixels    []pixel `xml:"rect"`
}

type pixel struct {
	Class  string `xml:"class,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
	X      int    `xml:"x,attr"`
	Y      int    `xml:"y,attr"`
	Fill   string `xml:"fill,attr"`
}

var white = color.RGBA{R: 255, G: 255, B: 255, A: 255}

func convertSvg(svg []byte) (*image.RGBA, error) {
	var pixels graph
	err := xml.Unmarshal(svg, &pixels)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	frame := pixels.Frame
	img := image.NewRGBA(image.Rect(frame.X, frame.Y, frame.X+frame.Width, frame.Y+frame.Height))

	fillRect(img, img.Rect, white)
	err = drawPixels(img, &pixels.OuterGroup)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return img, nil
}

func drawPixels(img *image.RGBA, svg *outerGroup) error {
	topX, topY, err := transformXY(svg.Transform)
	if err != nil {
		return errors.WithStack(err)
	}
	for _, g := range svg.WeekGroups {
		if len(g.Transform) == 0 {
			continue
		}
		startX, startY, err := transformXY(g.Transform)
		if err != nil {
			return errors.WithStack(err)
		}
		startX += topX
		startY += topY
		for _, p := range g.Pixels {
			err := drawPixel(img, startX, startY, p)
			if err != nil {
				return errors.WithStack(err)
			}
		}
	}
	return nil
}

func drawPixel(img *image.RGBA, startX int, startY int, p pixel) error {
	x := startX + p.X
	y := startY + p.Y
	c, err := fillToColor(p.Fill)
	if err != nil {
		return errors.WithStack(err)
	}
	rect := image.Rect(x, y, x+p.Width, y+p.Height)
	fillRect(img, rect, c)
	return nil
}

func fillRect(img *image.RGBA, target image.Rectangle, color color.RGBA) {
	for x := target.Min.X; x <= target.Max.X; x++ {
		for y := target.Min.Y; y <= target.Max.Y; y++ {
			img.Set(x, y, color)
		}
	}
}

func transformXY(transform string) (x int, y int, err error) {
	if !strings.Contains(transform, "translate(") {
		return 0, 0, fmt.Errorf("transform is not valid(actual: %s)", transform)
	}
	_, err = fmt.Sscanf(strings.ReplaceAll(transform, " ", ""), "translate(%d,%d)", &x, &y)
	return
}

func fillToColor(fill string) (c color.RGBA, err error) {
	c.A = 255
	switch len(fill) {
	case 7: // #FFFFFF
		_, err = fmt.Sscanf(fill, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4: // #FFF
		_, err = fmt.Sscanf(fill, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		c.R *= 0x11 // F → FF
		c.G *= 0x11
		c.B *= 0x11
	default:
		err = fmt.Errorf("invalid parameter(fill: %+v)", fill)
	}

	return
}
