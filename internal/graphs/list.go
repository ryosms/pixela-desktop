package graphs

import (
	"fmt"
	"github.com/aarzilli/nucular"
	"github.com/ryosms/pixela-desktop/pkg/pixela"
	"golang.org/x/mobile/event/key"
	"time"
)

type GraphListView struct {
	username      string
	graphs        *[]pixela.GraphDefinition
	size          int
	selectedIndex int
	centering     bool

	clickedIndex int
	clickedTime  time.Time

	parent *nucular.Window
}

const windowFlag = nucular.WindowTitle | nucular.WindowClosable

var listView GraphListView

func ShowList(w *nucular.Window, username string, graphs *[]pixela.GraphDefinition) {
	listView = GraphListView{
		username:      username,
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
			listView.selectedIndex++
			if listView.selectedIndex >= listView.size {
				listView.selectedIndex = listView.size - 1
			}
			listView.centering = true
		case key.CodeUpArrow:
			listView.selectedIndex--
			if listView.selectedIndex < 0 {
				listView.selectedIndex = 0
			}
			listView.centering = true
		case key.CodeReturnEnter, key.CodeKeypadEnter:
			if listView.selectedIndex >= 0 && listView.selectedIndex < listView.size {
				ShowDetail(w, listView.username, (*listView.graphs)[listView.selectedIndex])
			}
		}
	}

	w.Row(0).Dynamic(1)
	if gl, w := nucular.GroupListStart(w, listView.size, "graph list", nucular.WindowNoHScrollbar); w != nil {
		w.Row(40).Dynamic(1)
		for gl.Next() {
			i := gl.Index()
			graph := (*listView.graphs)[i]
			selected := i == listView.selectedIndex
			label := fmt.Sprintf("%s: %s\n    %s",
				graph.Id, graph.Name, graph.Color)
			if w.SelectableLabel(label, "LT", &selected) {
				if doubleClick(i, listView.clickedIndex, listView.clickedTime) {
					ShowDetail(listView.parent, listView.username, graph)
					return
				}
				listView.clickedIndex = i
				listView.clickedTime = time.Now()
			}
			if selected {
				listView.selectedIndex = i
				if listView.centering {
					gl.Center()
				}
			}
		}
	}
	w.Bounds = listView.parent.Bounds
	w.WidgetBounds()
}

func doubleClick(currentIndex int, previousIndex int, previousTime time.Time) bool {
	if currentIndex != previousIndex {
		return false
	}
	return time.Since(previousTime) < 200*time.Millisecond
}
