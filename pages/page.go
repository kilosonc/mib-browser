//MIT License
//
//Copyright (c) 2021 kiloson
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package pages

import (
	"fmt"
	"strings"
	"time"

	"github.com/closetool/mib-browser/pages/widget"
	"github.com/gosnmp/gosnmp"
	"github.com/rivo/tview"
)

func GetPage(app *tview.Application, port uint16, community string, ip string) *tview.Grid {
	widget.MibTree.SetSelectedFunc(func(node *tview.TreeNode) {
		children := node.GetChildren()
		if len(children) != 0 {
			node.SetExpanded(!node.IsExpanded())
		}
		widget.OidInput.SetText(node.GetReference().(string))
	})

	textView := tview.NewTextView()

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

func result2String(res []gosnmp.SnmpPDU) string {
	builder := strings.Builder{}
	for _, variable := range res {
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
		builder.Write([]byte{'\n'})
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
						textView.SetText(result2String(res.Variables))
					})
				}
			case widget.MethodGetNextRequest:
				if res, err := g.GetNext([]string{oid}); err != nil {
					app.QueueUpdateDraw(func() {
						textView.SetText(err.Error())
					})
				} else {
					app.QueueUpdateDraw(func() {
						textView.SetText(result2String(res.Variables))
					})
				}
			case widget.MethodWalkRequest:
				if res, err := g.WalkAll(oid); err != nil {
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
