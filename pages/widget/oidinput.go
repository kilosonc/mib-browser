package widget

import (
	"regexp"

	//"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	OidInput *tview.InputField
)

func init() {
	oidReg, _ := regexp.Compile(`(?:\.\d{1,3})+\.?`)
	OidInput = tview.NewInputField().
		SetLabel("OID").
		SetFieldWidth(30).
		SetAcceptanceFunc(func(textToCheck string, lastChar rune) bool {
			if textToCheck == "." {
				return true
			}
			if oidReg.FindString(textToCheck) == textToCheck {
				return true
			}
			return false
		})
}
