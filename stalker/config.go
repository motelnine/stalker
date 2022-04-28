package stalker

import (
	"fmt"
	"strings"
)

type Config struct {
	Folders map[string]string
	Files map[string]string
}

func (app *App) ParseConfigLine(line string) {

	// Ommit comments
	if string(line[0]) != "#" {

		args := strings.Split(string(line), ":")

		switch args[0]{

			case "folder":
				app.Config.Folders = map[string]string{
					args[1]: args[2],
				}

		//app.Config.Files = map[string]string{
		//	"Shit": "doubleyeah",
		//}
		}
	}
}

func (app *App) ParseConfig() {
	fileName := "/etc/stalker.conf"
	content := ReadFile(fileName)

	lines := strings.Split(string(content), "\n")

	lines = lines[:len(lines)-1]

	fmt.Println(fileName, "has a total of", len(lines), "lines")
	for _, line := range lines {
		app.ParseConfigLine(line)
	}

}
