package pages

import (
	"fmt"
	"strings"
	"time"

	"github.com/gosnmp/gosnmp"
	"github.com/kiloson/mib-browser/pages/widget"
	"github.com/rivo/tview"
)

func GetPage(app *tview.Application, port uint16, community string, ip string) *tview.Grid {
	widget.MibTree.SetSelectedFunc(func(node *tview.TreeNode) {
		children := node.GetChildren()
		if len(children) != 0 {
			node.SetExpanded(!node.IsExpanded())
		} else {
			widget.OidInput.SetText(node.GetReference().(string))
		}
	})

	textView := tview.NewTextView()

	//btnSend := tview.NewButton("Send").SetSelectedFunc(sendAndDisplay(ip,community,port,app,textView))

	//flex := tview.NewFlex().SetDirection(tview.FlexColumn).
	//	AddItem(widget.OidInput, 0, 2, false).
	//	AddItem(widget.DropDown, 0, 2, false).
	//	AddItem(btnSend, 0, 2, false)
	form := tview.NewForm().SetHorizontal(true).AddFormItem(widget.OidInput).AddFormItem(widget.DropDown).AddButton("Send", sendAndDisplayFunc(ip, community, port, app, textView))

	grid := tview.NewGrid().
		SetRows(3, 0).
		SetColumns(30, 0).
		SetBorders(true).
		AddItem(widget.MibTree, 1, 0, 1, 1, 0, 0, true).
		AddItem(form, 0, 0, 1, 2, 0, 0, false).
		AddItem(textView, 1, 1, 1, 1, 0, 0, false)
	return grid
}

func result2String(res *gosnmp.SnmpPacket) string {
	builder := strings.Builder{}
	for _, variable := range res.Variables {
		builder.WriteString(fmt.Sprintf("%v(%v) : ", variable.Name, variable.Type))
		switch variable.Type {
		case gosnmp.OctetString:
			bytes := variable.Value.([]byte)

			builder.Write(bytes)
			builder.Write([]byte{' '})
		default:
			builder.WriteString(gosnmp.ToBigInt(variable.Value).String())
			builder.Write([]byte{' '})
		}
	}
	return builder.String()
}

func sendAndDisplayFunc(ip, community string, port uint16, app *tview.Application, textView *tview.TextView) func() {
	return func() {
		oid := widget.OidInput.GetText()
		method := widget.Method
		go func() {
			g := &gosnmp.GoSNMP{
				Target:             ip,
				Port:               port,
				Transport:          "udp",
				Community:          community,
				Version:            gosnmp.Version2c,
				Timeout:            time.Duration(2) * time.Second,
				Retries:            3,
				ExponentialTimeout: true,
				MaxOids:            gosnmp.MaxOids,
			}
			if err := g.Connect(); err != nil {
				panic(err)
			}
			defer g.Conn.Close()

			switch method {
			case widget.MethodGetRequest:
				oid = fmt.Sprintf("%s.0", oid)
				if res, err := g.Get([]string{oid}); err != nil {
					app.QueueUpdateDraw(func() {
						textView.SetText(err.Error())
					})
				} else {
					app.QueueUpdateDraw(func() {
						textView.SetText(result2String(res))
					})
				}
			case widget.MethodGetNextRequest:
				if res, err := g.GetNext([]string{oid}); err != nil {
					app.QueueUpdateDraw(func() {
						textView.SetText(err.Error())
					})
				} else {
					app.QueueUpdateDraw(func() {
						textView.SetText(result2String(res))
					})
				}
			}
		}()
	}
}
