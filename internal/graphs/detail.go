package graphs

import (
	"fmt"
	"github.com/aarzilli/nucular"
	"github.com/ryosms/pixela-desktop/pkg/pixela"
)

type GraphDetailView struct {
	username string
	graph    pixela.GraphDefinition
	parent   *nucular.Window
	stats    pixela.GraphStats
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
		stats, err := pixela.GetGraphStats(detailView.username, detailView.graph.Id)
		if err != nil {
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

	w.Bounds = detailView.parent.Bounds
	w.WidgetBounds()
}
