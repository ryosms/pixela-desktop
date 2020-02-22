package graphs

import (
	"github.com/aarzilli/nucular"
	"github.com/ryosms/pixela-desktop/pkg/pixela"
)

type GraphDetailView struct {
	graph pixela.GraphDefinition

	parent *nucular.Window
}

var detailView GraphDetailView

func ShowDetail(w *nucular.Window, graphDef pixela.GraphDefinition) {
	detailView = GraphDetailView{
		graph:  graphDef,
		parent: w,
	}

	w.Master().PopupOpen(graphDef.Name, 0, w.Bounds, false, updateDetailView)
}

func updateDetailView(w *nucular.Window) {
	width := detailView.parent.Bounds.W - 50
	if width < 300 {
		width = 300
	}
	w.Row(30).Static(50, 0)
	if w.ButtonText("Back") {
		w.Close()
	}
	w.Label(detailView.graph.Name, "LC")

	w.Bounds = detailView.parent.Bounds
	w.WidgetBounds()
}
