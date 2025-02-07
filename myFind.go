package main

import (
	"flag"
	"fmt"
)

func main() {
	var d, f, sl bool
	var ext string
	slFlag := flag.Bool("sl", false, "Print symbolic links")
	dFlag := flag.Bool("d", false, "Print directories")
	fFlag := flag.Bool("f", false, "Print files")
	extFlag := flag.String("ext", "", "Print only files with a specific extension.")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Enter one path")
		return
	}
	if *slFlag {
		sl = true
	}
	if *dFlag {
		d = true
	}
	if *fFlag {
		f = true
	}
	if *extFlag != "" && f {
		ext = *extFlag
		if ext[0] != '.' {
			ext = "." + ext
		}
	} else if *extFlag != "" && !f {
		fmt.Println("the -ext flag must be used with the -f flag")
		return
	}
	if !f && !d && !sl {
		f = true
		d = true
		sl = true
	}
	path := flag.Arg(0)

	if err := basis(path, d, f, sl, ext); err != nil {
		fmt.Println(err)
	}
}
