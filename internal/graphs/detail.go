package graphs

import (
	"fmt"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/rect"
	"github.com/ebc-2in2crc/pixela4go"
	"github.com/ryosms/pixela-desktop/internal/types"
	"golang.org/x/mobile/event/key"
	"image"
	"image/color"
	"time"
)

type GraphDetailView struct {
	client      *pixela.Client
	username    string
	graph       *pixela.GraphDefinition
	parent      *nucular.Window
	stats       *pixela.Stats
	statsError  string
	modeShort   bool
	displayDate time.Time
	dateString  string
	dateEditor  nucular.TextEditor
	img         *image.RGBA
	imgError    string
}

var detailView GraphDetailView

func ShowDetail(w *nucular.Window, client *pixela.Client, graphDef *pixela.GraphDefinition) {
	detailView = GraphDetailView{
		client:      client,
		graph:       graphDef,
		parent:      w,
		displayDate: time.Now(),
	}
	detailView.dateEditor.Buffer = []rune(detailView.displayDate.Format("2006-01-02"))

	w.Master().PopupOpen(graphDef.Name, 0, w.Bounds, false, updateDetailView)

	go func() {
		loadSvgImage()
	}()

	go func() {
		stats, err := client.Graph().Stats(&pixela.GraphStatsInput{ID: pixela.String(graphDef.ID)})
		if err != nil {
			detailView.statsError = "failed to get graph stats"
			fmt.Printf("%+v\n", err)
		} else {
			detailView.stats = stats
		}
	}()
}

func updateDetailView(w *nucular.Window) {
	handleKeyEvent(w)
	g := detailView.graph
	w.Row(30).Static(50, 0)
	if w.ButtonText("Back") {
		w.Close()
	}
	w.Label(g.Name, "LC")

	w.Row(20).Static(60, 0)
	w.LabelColored(g.Color, "LT", *(types.DisplayColorByName(g.Color)))
	w.Label(fmt.Sprintf("id: %s", g.ID), "LT")

	if detailView.img != nil {
		showImage(w)
	}
	if len(detailView.imgError) > 0 {
		w.Row(40).Dynamic(1)
		w.LabelColored(detailView.imgError, "LT", color.RGBA{R: 0xFF, A: 0xFF})
	}

	if w.TreePush(nucular.TreeTab, "Stats", false) {
		if len(detailView.statsError) > 0 {
			w.Row(40).Dynamic(1)
			w.LabelColored(detailView.statsError, "LT", color.RGBA{R: 0xFF, A: 0xFF})
		}
		if detailView.stats != nil {
			showStats(w, detailView.stats, g.Unit)
		}
		w.TreePop()
	}

	if w.TreePush(nucular.TreeTab, "Definition", false) {
		w.Row(20).Dynamic(1)
		w.Label(fmt.Sprintf("Unit: %s", g.Unit), "LT")
		w.Label(fmt.Sprintf("Type: %s", g.Type), "LT")
		w.Label(fmt.Sprintf("Timezone: %s", g.TimeZone), "LT")
		w.Label(fmt.Sprintf("SelfSufficient: %v", g.SelfSufficient), "LT")
		w.Label(fmt.Sprintf("Purge Target: %d", len(g.PurgeCacheURLs)), "LT")

		w.TreePop()
	}

	w.Bounds = detailView.parent.Bounds
	w.WidgetBounds()
}

func loadSvgImage() {
	d := &detailView
	var mode *string
	if d.modeShort {
		mode = pixela.String("short")
	}
	input := pixela.GraphGetSVGInput{
		ID:   pixela.String(d.graph.ID),
		Date: pixela.String(d.displayDate.Format("20060102")),
		Mode: mode,
	}
	svg, err := d.client.Graph().GetSVG(&input)
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
	v := &detailView
	w.Row(30).Static(5, 80, 80)
	w.Spacing(1)
	if w.OptionText("Full", !v.modeShort) {
		if v.modeShort {
			v.modeShort = false
			go loadSvgImage()
		}
	}
	if w.OptionText("Short", v.modeShort) {
		if !v.modeShort {
			v.modeShort = true
			go loadSvgImage()
		}
	}

	w.RowScaled(v.img.Rect.Max.Y).StaticScaled(5, v.img.Rect.Max.X)
	w.Spacing(1)
	w.Image(v.img)

	w.Row(30).Static(5, 80, 120, 80, 80)
	w.Spacing(1)
	if w.ButtonText("Prev") {
		if v.modeShort {
			changeSvgDate(v.displayDate.AddDate(0, -3, 0))
		} else {
			changeSvgDate(v.displayDate.AddDate(-1, 0, 0))
		}
	}

	v.dateEditor.Flags = nucular.EditField | nucular.EditIbeamCursor
	v.dateEditor.Maxlen = 10
	v.dateEditor.Edit(w)
	v.dateString = string(v.dateEditor.Buffer)

	if w.ButtonText("Go") {
		moveInputDate(w)
	}

	if w.ButtonText("Next") {
		if v.modeShort {
			changeSvgDate(v.displayDate.AddDate(0, 3, 0))
		} else {
			changeSvgDate(v.displayDate.AddDate(1, 0, 0))
		}
	}
}

func moveInputDate(w *nucular.Window) {
	d, err := time.Parse("2006-01-02", detailView.dateString)
	if err != nil {
		w.Master().PopupOpen("Error", nucular.WindowMovable|nucular.WindowClosable|nucular.WindowTitle|nucular.WindowDynamic|nucular.WindowNoScrollbar, rect.Rect{X: 20, Y: 100, W: 240, H: 150}, true, popUpDateError)
	} else {
		changeSvgDate(d)
	}
}

func handleKeyEvent(w *nucular.Window) {
	if !detailView.dateEditor.Active {
		return
	}
	for _, e := range w.Input().Keyboard.Keys {
		switch e.Code {
		case key.CodeReturnEnter, key.CodeKeypadEnter:
			moveInputDate(w)
		}
	}
}

func showStats(w *nucular.Window, st *pixela.Stats, unit string) {
	w.Row(30).Dynamic(1)
	w.Label(fmt.Sprintf("Today: %v %s", st.TodaysQuantity, unit), "LT")

	w.Row(20).Dynamic(1)
	w.Label(fmt.Sprintf("Total: %v %s", st.TotalQuantity, unit), "LT")
	w.Label(fmt.Sprintf("Max: %v %s", st.MaxQuantity, unit), "LT")
	w.Label(fmt.Sprintf("Min: %v %s", st.MinQuantity, unit), "LT")
	w.Label(fmt.Sprintf("Total Pixels: %v %s", st.TotalPixelsCount, unit), "LT")
}

func changeSvgDate(newDate time.Time) {
	detailView.displayDate = newDate
	detailView.dateEditor.Buffer = []rune(newDate.Format("2006-01-02"))
	go loadSvgImage()
}

func popUpDateError(w *nucular.Window) {
	w.Row(0).Dynamic(1)
	w.LabelWrap("The input value is invalid. Valid input is formatted 'yyyy-MM-dd'.")
}
