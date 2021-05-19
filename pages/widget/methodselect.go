package widget

import "github.com/rivo/tview"

var (
	DropDown *tview.DropDown
	Method   string
)

const (
	MethodGetRequest     = "GetRequest"
	MethodGetNextRequest = "GetNextRequest"
	MethodWalkRequest    = "Walk"
)

func init() {
	DropDown = tview.NewDropDown().
		SetLabel("Select method: ").
		SetFieldWidth(30).
		SetOptions([]string{MethodGetRequest, MethodGetNextRequest, MethodWalkRequest}, func(text string, _ int) {
			Method = text
		})
}
