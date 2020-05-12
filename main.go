package main

import (
    "libkij/interfaces"
    "log"
)

func main() {
    var init interfaces.InitWin
    init.Logo = "__  ____  __ _   _ ____        \n\\ \\/ /  \\/  | | | / ___| _   _ \n \\  /| |\\/| | | | \\___ \\| | | |\n /  \\| |  | | |_| |___) | |_| |\n/_/\\_\\_|  |_|\\___/|____/ \\__,_|"
    init.Prompt = "Initialising, please wait..."
    init.ShowPeriod = 2000
    init.NeedProgBar = false

    interfaces.NewInitWindow(init)

    log.Println("Closed.")
}
