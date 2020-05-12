package main

import (
    "libkij/interfaces"
    "log"
)

func main() {
    //var init interfaces.InitWin
    //init.Logo = "   _  __ __  _____  ____  ___           __        __ \n  | |/ //  |/  / / / /  |/  /___ ______/ /_____  / /_\n  |   // /|_/ / / / / /|_/ / __ `/ ___/ //_/ _ \\/ __/\n /   |/ /  / / /_/ / /  / / /_/ / /  / ,< /  __/ /_  \n/_/|_/_/  /_/\\____/_/  /_/\\__,_/_/  /_/|_|\\___/\\__/"
    //init.Prompt = "Initialising, please wait..."
    //init.ShowPeriod = 2000
    //init.NeedProgBar = false

    var sel interfaces.SelectWin
    sel.MainTitle = "Hello"
    sel.ChoicePadTitle = "World!"
    sel.ChoicePadDesc = "Please select!"
    sel.Choices = []string {
        "jujuju", "pangpangpang",
    }
    sel.ChoicePadFootnote = "Press Ctrl+Q to exit this programme"
    sel.StatusBarContent = "i am status"

    interfaces.NewSelectWin(sel)

    log.Println("Closed.")
}
