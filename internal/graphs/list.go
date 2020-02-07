package graphs

import (
	"fmt"
	"github.com/aarzilli/nucular"
	"github.com/ryosms/pixela-desktop/pkg/pixela"
	"golang.org/x/mobile/event/key"
	"time"
)

type GraphListView struct {
	graphs        *[]pixela.GraphDefinition
	size          int
	selectedIndex int
	centering     bool

	clickedIndex int
	clickedTime  time.Time

	parent *nucular.Window
}

const windowFlag = nucular.WindowTitle | nucular.WindowClosable

var view GraphListView

func ShowList(w *nucular.Window, graphs *[]pixela.GraphDefinition) {
	view = GraphListView{
		graphs:        graphs,
		size:          len(*graphs),
		selectedIndex: -1,

		clickedIndex: -1,
		parent:       w,
	}
	w.Master().PopupOpen("graph list", windowFlag, w.Bounds, false, updateListView)
}

func updateListView(w *nucular.Window) {
	for _, e := range w.Input().Keyboard.Keys {
		switch e.Code {
		case key.CodeDownArrow:
			view.selectedIndex++
			if view.selectedIndex >= view.size {
				view.selectedIndex = view.size - 1
			}
			view.centering = true
		case key.CodeUpArrow:
			view.selectedIndex--
			if view.selectedIndex < 0 {
				view.selectedIndex = 0
			}
			view.centering = true
		}
	}

	w.Row(0).Dynamic(1)
	if gl, w := nucular.GroupListStart(w, view.size, "graph list", nucular.WindowNoHScrollbar); w != nil {
		w.Row(40).Dynamic(1)
		for gl.Next() {
			i := gl.Index()
			graph := (*view.graphs)[i]
			selected := i == view.selectedIndex
			label := fmt.Sprintf("%s: %s\n    %s",
				graph.Id, graph.Name, graph.Color)
			if w.SelectableLabel(label, "LT", &selected) {
				if doubleClick(i, view.clickedIndex, view.clickedTime) {
					fmt.Println("double clicked!" + graph.Id)
					return
				}
				view.clickedIndex = i
				view.clickedTime = time.Now()
			}
			if selected {
				view.selectedIndex = i
				if view.centering {
					gl.Center()
				}
			}
		}
	}
	w.Bounds = view.parent.Bounds
	w.WidgetBounds()
}

func doubleClick(currentIndex int, previousIndex int, previousTime time.Time) bool {
	if currentIndex != previousIndex {
		return false
	}
	return time.Since(previousTime) < 200*time.Millisecond
}
