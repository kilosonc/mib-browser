package main

import (
	"flag"
	"os"

	"github.com/kiloson/mib-browser/pages"
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
