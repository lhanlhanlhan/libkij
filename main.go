package main

import "C"
import (
    "libkij/interfaces"
    "unsafe"
)

func main() {
    interfaces.AuthWinTest()
}

//export NewAuthWin
func NewAuthWin(logo *C.char, prompt *C.char, statusBarTxt *C.char) (username, passcode *C.char) {
    var win interfaces.AuthWin
    win.Logo = C.GoString(logo)
    win.Prompt = C.GoString(prompt)
    win.StatusBarText = C.GoString(statusBarTxt)

    u, p := interfaces.NewAuthWindow(&win)
    return C.CString(u), C.CString(p)
}

//export NewErrWin
func NewErrWin(mainTitle *C.char, errorWinTitle *C.char, errorInfo *C.char, buttons **C.char, buttonsCount C.int) (selected C.int) {
    // 先将传入的按钮assume成一个char* 数组
    cButs := (*[1 << 30]*C.char)(unsafe.Pointer(buttons))[:buttonsCount]
    // 再将上面数组映射至一个go的string数组
    goButs := make([]string, buttonsCount)
    for i, ch := range cButs {
        goButs[i] = C.GoString(ch)
    }

    var win interfaces.ErrorWin
    win.MainTitle = C.GoString(mainTitle)
    win.ErrorWinTitle = C.GoString(errorWinTitle)
    win.Error = C.GoString(errorInfo)
    win.Button = goButs

    selected = C.int(interfaces.NewErrorWin(&win))
    return
}

//export NewInfoWin
func NewInfoWin(mainTitle *C.char, infoWinTitle *C.char, infoContents *C.char, buttons **C.char, buttonsCount C.int) (selected C.int) {
    // 先将传入的按钮assume成一个char* 数组
    cButs := (*[1 << 30]*C.char)(unsafe.Pointer(buttons))[:buttonsCount]
    // 再将上面数组映射至一个go的string数组
    goButs := make([]string, buttonsCount)
    for i, ch := range cButs {
        goButs[i] = C.GoString(ch)
    }

    var win interfaces.InfoWin
    win.MainTitle = C.GoString(mainTitle)
    win.InfoWinTitle = C.GoString(infoWinTitle)
    win.Info = C.GoString(infoContents)
    win.Button = goButs

    selected = C.int(interfaces.NewInfoWin(&win))
    return
}

// 只有needProgBar是T或t时，才会显示进度条！
//export NewInitWin
func NewInitWin(logo *C.char, prompt *C.char, showPeriod C.int, needProgBar C.char) {
    var win interfaces.InitWin
    win.Logo = C.GoString(logo)
    win.Prompt = C.GoString(prompt)
    win.ShowPeriod = int(showPeriod)
    win.NeedProgBar = 'T' == byte(needProgBar) || 't' == byte(needProgBar)

    interfaces.NewInitWindow(&win)
}

// 传出的字符串会和文本框个数一样多，而且请注意要free掉！
//export NewInputWin
func NewInputWin(mainTitle *C.char, inputWinTitle *C.char, inputBoxLabels **C.char, inputBoxCount C.int, buttons **C.char, buttonsCount C.int) (sel C.int, text **C.char, textCount C.int) {
    // 先将传入的标签assume成一个char* 数组
    cLabels := (*[1 << 30]*C.char)(unsafe.Pointer(inputBoxLabels))[:inputBoxCount]
    // 再将上面数组映射至一个go的string数组
    goLabels := make([]string, inputBoxCount)
    for i, ch := range cLabels {
        goLabels[i] = C.GoString(ch)
    }

    // 先将传入的按钮assume成一个char* 数组
    cButs := (*[1 << 30]*C.char)(unsafe.Pointer(buttons))[:buttonsCount]
    // 再将上面数组映射至一个go的string数组
    goButs := make([]string, buttonsCount)
    for i, ch := range cButs {
        goButs[i] = C.GoString(ch)
    }

    var tbl interfaces.InputWin
    tbl.MainTitle = C.GoString(mainTitle)
    tbl.InputWinName = C.GoString(inputWinTitle)
    tbl.InputBoxNames = goLabels
    tbl.Button = goButs

    selected, txt := interfaces.NewInputWin(&tbl)

    // 用C函数新建一个数组
    cTxt := C.malloc(C.size_t(len(txt)) * C.size_t(unsafe.Sizeof(uintptr(0))))
    // 将该数组映射至go数组，并强制写入
    goTxt := (*[1 << 30] *C.char)(cTxt)
    for i, s := range txt {
        goTxt[i] = C.CString(s)
    }

    return C.int(selected), (**C.char)(cTxt), C.int(inputBoxCount)
}


//export NewSelectWin
func NewSelectWin(mainTitle *C.char, winTitle *C.char, winDesc *C.char, choices **C.char, choicesCount C.int, winFootNote *C.char, statusBar *C.char) (sel C.int) {

    // 先将传入的按钮assume成一个char* 数组
    cButs := (*[1 << 30]*C.char)(unsafe.Pointer(choices))[:choicesCount]
    // 再将上面数组映射至一个go的string数组
    goButs := make([]string, choicesCount)
    for i, ch := range cButs {
        goButs[i] = C.GoString(ch)
    }

    var tbl interfaces.SelectWin
    tbl.MainTitle = C.GoString(mainTitle)
    tbl.ChoicePadTitle = C.GoString(winTitle)
    tbl.ChoicePadDesc = C.GoString(winDesc)
    tbl.Choices = goButs
    tbl.ChoicePadFootnote = C.GoString(winFootNote)
    tbl.StatusBarContent = C.GoString(statusBar)

    selected := interfaces.NewSelectWin(&tbl)

    return C.int(selected)
}


//export NewTableWin
func NewTableWin(mainTitle *C.char, tableName *C.char, colNames **C.char, colsCount C.int, cols *C.int, rows **C.char, rowsCount C.int, buttons **C.char, buttonsCount C.int) (sel C.int) {

    // 1) 先将传入的列名assume成一个char* 数组
    cColNames := (*[1 << 30]*C.char)(unsafe.Pointer(colNames))[:colsCount]
    // 再将上面数组映射至一个go的string数组
    goColNames := make([]string, colsCount)
    for i, ch := range cColNames {
        goColNames[i] = C.GoString(ch)
    }

    // 2) 先将传入的列宽数组assume成一个int 数组
    cColsWidth := (*[1 << 30]C.int)(unsafe.Pointer(cols))[:colsCount]
    // 再将上面数组映射至一个go的string数组
    goColsWidth := make([]int, colsCount)
    for i, ch := range cColsWidth {
        goColsWidth[i] = int(ch)
    }

    // 3) 先将传入的行数组assume成一个*char 数组
    cRows := (*[1 << 30]*C.char)(unsafe.Pointer(rows))[:rowsCount]
    // 再将上面数组映射至一个go的string数组
    goRows := make([]string, rowsCount)
    for i, ch := range cRows {
        goRows[i] = C.GoString(ch)
    }

    // 4) 先将传入的按钮assume成一个char* 数组
    cButs := (*[1 << 30]*C.char)(unsafe.Pointer(buttons))[:buttonsCount]
    // 再将上面数组映射至一个go的string数组
    goButs := make([]string, buttonsCount)
    for i, ch := range cButs {
        goButs[i] = C.GoString(ch)
    }

    var tbl interfaces.TableWin
    tbl.MainTitle = C.GoString(mainTitle)
    tbl.TableName = C.GoString(tableName)
    tbl.ColNames = goColNames
    tbl.ColsWidth = goColsWidth
    tbl.Rows = goRows
    tbl.Button = goButs

    selected := interfaces.NewTableWin(&tbl)

    return C.int(selected)
}