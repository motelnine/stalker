package stalker

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

type App struct {
	Config Config
	State  State
}

type State struct {
	Folders map[string]int64
	Files   map[string]int64
}

// Initialize initializes the stalker app
func (app *App) Initialize() {

	// Parse config
	app.ParseConfig()

	// Get initial state
	app.State.Folders = make(map[string]int64)
	app.State.Files = make(map[string]int64)
	app.State = app.GetState()

	// Monitor for state change
	app.Monitor()
}

// GetState populates the file and folder states
func (app *App) GetState() State {

	var state State
	state.Folders = make(map[string]int64)
	state.Files = make(map[string]int64)

	// Get folder states
	for _, val := range app.Config.Folders {
		state.Folders[val.Location] = DirSize(val.Location)

	}

	for _, val := range app.Config.Files {
		fileName := val.Folder +"/"+ val.Name
		state.Files[fileName] = FileSize(fileName)
	}

	return state
}

// DumpState dumps app state to console
func (app *App) DumpState() {
	data, _ := json.MarshalIndent(app.State, "", "  ")
	fmt.Println(`"State": `+ string(data))
}

// Monitor checks system for state changes
func (app *App) Monitor() {
	for {
		newState := app.GetState()

		
		for key, val := range app.State.Folders {
			if(val != newState.Folders[key]) {
				app.executeFolderRule(key)
			}
		}


		for key, val := range app.State.Files {
			if(val != newState.Files[key]) {
				app.executeFileRule(key)
			}
		}

		app.State = newState

		time.Sleep(app.Config.Interval * time.Second)
	}
}

func (app *App) executeFolderRule(folder string) {

	for _, val := range app.Config.Folders {
		if(folder == val.Location) {
			app.executeCommand(val.Action)
		}
	}
}

func (app *App) executeFileRule(file string) {
	for _, val := range app.Config.Files {
		name := val.Folder +"/"+ val.Name

		if (file == name) {
			app.executeCommand(val.Action)
		}
	}
}

func (app *App) executeCommand(cmd string) {

	if app.Config.DryRun == false {
		//cmd := "cat /proc/cpuinfo | egrep '^model name' | uniq | awk '{print substr($0, index($0,$4))}'"
		out, err := exec.Command("bash","-c",cmd).Output()
		if err != nil {
			fmt.Println("Failed to execute command:"+ cmd)
		}

		fmt.Println(string(out))
	}
}
/*
func (app *App) executeCommand(command string) {
	command = `"`+ command +`"`

	if app.Config.DryRun == false {
		cmd := exec.Command(app.Config.Shell, "-c", command)
	    stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
        	return
		}
		fmt.Println(string(stdout))

	} else {
		fmt.Println(app.Config.Shell +" -c "+ command)
	}
}
*/
