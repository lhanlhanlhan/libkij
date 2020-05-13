package main

import "C"
import "libkij/interfaces"

func main() {
    interfaces.AuthWinTest()
}

//export NewAuthWin
func NewAuthWin(logo *C.char, prompt *C.char, statusBarTxt *C.char) (username, passcode string) {
    var win interfaces.AuthWin
    win.Logo = C.GoString(logo)
    win.Prompt = C.GoString(prompt)
    win.StatusBarText = C.GoString(statusBarTxt)

    return interfaces.NewAuthWindow(&win)
}