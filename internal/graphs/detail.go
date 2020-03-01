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
	stats      *pixela.GraphStats
	statsError string
	modeShort  bool
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
		loadSvgImage("")
	}()

	go func() {
		stats, err := pixela.GetGraphStats(detailView.username, detailView.graph.Id)
		if err != nil {
			detailView.statsError = "failed to get graph stats"
			fmt.Printf("%+v\n", err)
		} else {
			detailView.stats = stats
		}
	}()
}

func updateDetailView(w *nucular.Window) {
	g := detailView.graph
	w.Row(30).Static(50, 0)
	if w.ButtonText("Back") {
		w.Close()
	}
	w.Label(g.Name, "LC")

	w.Row(20).Static(60, 0)
	w.LabelColored(g.Color, "LT", *(pixela.DisplayColorByName(g.Color)))
	w.Label(fmt.Sprintf("id: %s", g.Id), "LT")

	if detailView.img != nil {
		showImage(w)
	}
	if len(detailView.imgError) > 0 {
		w.Row(40).Dynamic(1)
		w.Label(detailView.imgError, "LT")
	}

	if len(detailView.statsError) > 0 {
		w.Row(40).Dynamic(1)
		w.Label(detailView.statsError, "LT")
	}
	if detailView.stats != nil {
		showStats(w, detailView.stats, g.Unit)
	}

	if w.TreePush(nucular.TreeTab, "Definition", false) {
		w.Row(20).Dynamic(1)
		w.Label(fmt.Sprintf("Unit: %s", g.Unit), "LT")
		w.Label(fmt.Sprintf("Type: %s", g.Type), "LT")
		w.Label(fmt.Sprintf("Timezone: %s", g.Timezone), "LT")
		w.Label(fmt.Sprintf("SelfSufficient: %v", g.SelfSufficient), "LT")
		w.Label(fmt.Sprintf("Purge Target: %d", len(g.PurgeCacheUrls)), "LT")

		w.TreePop()
	}

	w.Bounds = detailView.parent.Bounds
	w.WidgetBounds()
}

func loadSvgImage(date string) {
	d := &date
	if len(date) == 0 {
		d = nil
	}
	svg, err := pixela.GetGraphSvg(detailView.username, detailView.graph.Id, detailView.modeShort, d)
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

func showImage(w *nucular.Window) {
	w.Row(30).Static(5, 80, 80)
	w.Spacing(1)
	if w.OptionText("Full", !detailView.modeShort) {
		detailView.modeShort = false
		loadSvgImage("")
	}
	if w.OptionText("Short", detailView.modeShort) {
		detailView.modeShort = true
		loadSvgImage("")
	}

	w.RowScaled(detailView.img.Rect.Max.Y).StaticScaled(5, detailView.img.Rect.Max.X)
	w.Spacing(1)
	w.Image(detailView.img)

	w.Row(30).Dynamic(4)

}

func showStats(w *nucular.Window, st *pixela.GraphStats, unit string) {
	if w.TreePush(nucular.TreeTab, "Stats", false) {
		w.Row(30).Dynamic(1)
		w.Label(fmt.Sprintf("Today: %v %s", st.TodaysQuantity, unit), "LT")

		w.Row(20).Dynamic(1)
		w.Label(fmt.Sprintf("Total: %v %s", st.TotalQuantity, unit), "LT")
		w.Label(fmt.Sprintf("Max: %v %s", st.MaxQuantity, unit), "LT")
		w.Label(fmt.Sprintf("Min: %v %s", st.MinQuantity, unit), "LT")
		w.Label(fmt.Sprintf("Total Pixels: %v %s", st.TotalPixelsCount, unit), "LT")

		w.TreePop()
	}
}
