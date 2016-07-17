package main

import (
	"github.com/mlctrez/cdnjslib"
)

func main() {

	vendor := &cdnjslib.LibraryCollection{}

	cdnjslib.ReadFile("vendor.json", vendor)

	for _, l := range vendor.Libraries {
		l.SaveLocal("static/vendor")
	}

}
