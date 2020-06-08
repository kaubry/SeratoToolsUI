package main

import (
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"log"
	"os"
)

func main() {
	// Initialize astilectron
	var a, _ = astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName: "Serato Tools",
		//AppIconDefaultPath: "<your .png icon>",  // If path is relative, it must be relative to the data directory
		//AppIconDarwinPath:  "<your .icns icon>", // Same here
		BaseDirectoryPath:  "libs",
		VersionAstilectron: "0.39.0",
		VersionElectron:    "9.0.2",
	})
	defer a.Close()

	// Start astilectron
	a.Start()

	var w, _ = a.NewWindow("./assets/build/index.html", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(600),
		Width:  astikit.IntPtr(600),
	})

	w.Create()

	// Open dev tools
	//w.OpenDevTools()

	// Blocking pattern
	a.Wait()
}
