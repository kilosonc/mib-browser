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

package main

import (
	"flag"
	"os"

	"github.com/closetool/mib-browser/pages"
	"github.com/rivo/tview"
)

func main() {
	ip := flag.String("i", "127.0.0.1", "ip")
	port := flag.Uint64("p", 161, "port")
	community := flag.String("c", "public", "community")
	help := flag.Bool("h", false, "help")
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	if err := os.Setenv("LC_CTYPE", "en_US.UTF-8"); err != nil {
		panic(err)
	}
	app := tview.NewApplication()
	if err := app.SetRoot(pages.GetPage(app, uint16(*port), *community, *ip), true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
