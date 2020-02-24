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
	WeekGroups []weekGroup `xml:"g"`
}

type weekGroup struct {
	Pixels []pixel `xml:"rect"`
}

type pixel struct {
	TippyContent string `xml:"data-tippy-content,attr"`
	Class        string `xml:"class,attr"`
	Width        string `xml:"width,attr"`
	Height       string `xml:"height,attr"`
	X            string `xml:"x,attr"`
	Y            string `xml:"y,attr"`
	Fill         string `xml:"fill,attr"`
	DataCount    string `xml:"data-count:attr"`
	DataDate     string `xml:"data-date:attr"`
	DataUnit     string `xml:"data-unit:attr"`
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
