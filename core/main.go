package main

import (
	"github.com/miketheprogrammer/go-thrust/thrust"
	"fmt"
)

func main() {
	thrust.InitLogger()
	thrust.Start()

  var config Config
	config = get_config()
	var baseFile = "index.html"
	fmt.Println(config.RootDir + baseFile)
	fmt.Println("abc")

	thrustWindow := thrust.NewWindow(thrust.WindowOptions{
		RootUrl: config.RootDir + baseFile,
		HasFrame: true,
	})

	thrustWindow.Show()
	thrustWindow.Focus()

	thrust.LockThread()
}
