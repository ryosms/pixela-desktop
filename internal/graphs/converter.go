package graphs

import (
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"image"
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

func convertSvg(svg []byte) (*image.RGBA, error) {
	var pixels graph
	err := xml.Unmarshal(svg, &pixels)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	fmt.Printf("%+v\n", pixels)

	return nil, nil
}
