package graphs

import (
	"fmt"
	"github.com/aarzilli/nucular"
	"github.com/ryosms/pixela-desktop/pkg/pixela"
	"image"
)

type GraphDetailView struct {
	username   string
	graph      pixela.GraphDefinition
	parent     *nucular.Window
	stats      pixela.GraphStats
	statsError string
	img        *image.RGBA
	imgError   string
}

var detailView GraphDetailView

func ShowDetail(w *nucular.Window, username string, graphDef pixela.GraphDefinition) {
	detailView = GraphDetailView{
		username: username,
		graph:    graphDef,
		parent:   w,
	}

	w.Master().PopupOpen(graphDef.Name, 0, w.Bounds, false, updateDetailView)

	go func() {
		loadSvgImage("20190101")
	}()

	go func() {
		stats, err := pixela.GetGraphStats(detailView.username, detailView.graph.Id)
		if err != nil {
			detailView.statsError = "failed to get graph stats"
			fmt.Printf("%+v\n", err)
		} else {
			detailView.stats = *stats
		}
	}()
}

func updateDetailView(w *nucular.Window) {
	w.Row(30).Static(50, 0)
	if w.ButtonText("Back") {
		w.Close()
	}
	w.Label(detailView.graph.Name, "LC")

	if detailView.img != nil {
		w.Row(detailView.img.Rect.Max.Y).Dynamic(1)
		w.Image(detailView.img)
	}
	if len(detailView.imgError) > 0 {
		w.Row(40).Dynamic(1)
		w.Label(detailView.imgError, "LT")
	}

	if len(detailView.statsError) > 0 {
		w.Row(40).Dynamic(1)
		w.Label(detailView.statsError, "LT")
	}

	w.Bounds = detailView.parent.Bounds
	w.WidgetBounds()
}

func loadSvgImage(date string) {
	d := &date
	if len(date) == 0 {
		d = nil
	}
	svg, err := pixela.GetGraphSvg(detailView.username, detailView.graph.Id, d)
	if err != nil {
		detailView.imgError = "failed to get svg data"
		fmt.Printf("%+v\n", err)
	} else {
		img, err := convertSvg(svg)
		if err != nil {
			detailView.imgError = "failed converting svg to image"
			fmt.Printf("%+v\n", err)
		} else {
			detailView.imgError = ""
			detailView.img = img
		}
	}

}
