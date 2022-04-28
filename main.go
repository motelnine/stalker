package main

import (
	"fmt"
	"ambrota.com/stalker"
)

func main() {
	app := stalker.App{}
	//configFile := "/etc/stalker.conf"
	//content := app.stalker.ReadFile(configFile)
	app.ParseConfig()
	fmt.Println(app)

	//size, err := stalker.DirSize("/home/matthew/.keys/keepassdb")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(size)
}
