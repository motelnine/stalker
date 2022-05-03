package stalker

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Files struct {
	Folder string `json:"fileFolder"`
	Name   string `json:"fileName"`
	Action string `json:"fileAction"`
}

type Folders struct {
	Location string `json:"folderLocation"`
	Action   string `json:"folderAction"`
}

type GitCommands struct {
	Message string `json:"message"`
	Add     string `json:"add"`
	Commit  string `json:"commit"`
	Push    string `json:"push"`
}

type Config struct {
	Interval time.Duration `json:"interval"`
	Git      GitCommands   `json:"git"`
	Folders  []Folders     `json:"folders"`
	Files    []Files       `json:"files"`
	Shell    string        `json:"shell"`
    DryRun   bool          `json:"dryRun"`
}

// ParseConfig attaches the configuration to app.Config
func (app *App) ParseConfig() {

	// Convert file data to string
	configString := app.readConfig()

	// Unmarshal config json
	var config Config
	json.Unmarshal([]byte(configString), &config)
	app.Config = config

	// Parse config tokens
	app.ParseTokens()
}

// readConfig returns file config as string
func (app *App) readConfig() string {
	fileName := "/etc/stalker.json"
	content := ReadFile(fileName)
	return(string(content))
}


// ParseTokens populates tokens in app.Config
func (app *App) ParseTokens() {

	// Append git.message to git.commit
	app.Config.Git.Commit = strings.ReplaceAll(app.Config.Git.Commit, "{git.message}", app.Config.Git.Message)

	// Populate app.Config.Folders tokens
	for key, val := range app.Config.Folders {
		folder :=  val.Location
		val.Action = strings.ReplaceAll(val.Action, "{cd}", "cd "+ folder)
		val.Action = strings.ReplaceAll(val.Action, "{git.add}", app.Config.Git.Add)
		val.Action = strings.ReplaceAll(val.Action, "{git.commit}", app.Config.Git.Commit)
		val.Action = strings.ReplaceAll(val.Action, "{git.push}", app.Config.Git.Push)

		app.Config.Folders[key].Action = val.Action
	}

	// Populate app.Config.Files tokens
	for key, val := range app.Config.Files {
		folder :=  val.Folder
		val.Action = strings.ReplaceAll(val.Action, "{cd}", "cd "+ folder)
		val.Action = strings.ReplaceAll(val.Action, "{cd}", "cd "+ folder)
		val.Action = strings.ReplaceAll(val.Action, "{git.add}", app.Config.Git.Add)
		val.Action = strings.ReplaceAll(val.Action, "{git.commit}", app.Config.Git.Commit)
		val.Action = strings.ReplaceAll(val.Action, "{git.push}", app.Config.Git.Push)

		app.Config.Files[key].Action = val.Action
	}
}

// DumpRules 
func (app *App) DumpRules() {
	data, _ := json.MarshalIndent(app.Config, "", "  ")
	fmt.Println(`"Rules": `+ string(data))
}
