package main

import (
	"fmt"
	"io/ioutil"

	flag "github.com/spf13/pflag"
)

type Args struct {
	//
	Paths []string
}

func main() {

	args := Args{}

	// os.Args[1:] を取得して、フラグを解析しています。
	// これより前にフラグのセットを終える必要がある
	flag.Parse()
	args.Paths = flag.Args()

	for _, path := range args.Paths {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s :\n", path)
		for _, file := range files {
			fmt.Println(file.Name())
		}
		fmt.Println("")
	}
}
