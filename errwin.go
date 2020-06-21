package libkij

import (
    "fmt"
    "github.com/chrishanli/tui-go"
    "log"
)

type ErrorWin struct {
    MainTitle       string
    ErrorWinTitle   string
    Error           string
    Button          []string
}

func ErrorWinTest() {
    var sel ErrorWin
    sel.ErrorWinTitle = "Error"
    sel.Error =
`1. Please note that Apple's 
   guarantee has expired!

2. Please note that Pear's 
   guarantee has expired!`
    sel.Button = []string {
        "OK",
    }

    s := NewErrorWin(&sel)

    log.Println("Closed. Selected index:", s, "which indicates", sel.Button[s])

}

func NewErrorWin(win *ErrorWin) (selectedButton int) {
    selectedButton = 0

    // 新建主题
    t := tui.NewTheme()

    // 新建窗口容器
    windowBox := tui.NewVBox()

    // 窗口标题
    windowTitle := &StyledBox {
        Style : "errortitle",
        Box : tui.NewHBox(tui.NewSpacer(), tui.NewLabel(win.ErrorWinTitle), tui.NewSpacer()),
    }
    t.SetStyle("errortitle", tui.Style{Bg: tui.ColorRed, Fg: tui.ColorWhite})

    windowBox.Append(tui.NewPadder(1, 1, windowTitle))
    // 窗口描述
    if win.Error != "" {
        windowBox.Append(tui.NewPadder(1, 0, tui.NewLabel(win.Error)))
    }

    // 新建窗口容纳上述内容，以及有一个蓝色背景
    window := &StyledBox{
        Style: "selectwin",
        Box:   windowBox,
    }
    t.SetStyle("selectwin", tui.Style{Bg: tui.ColorBlue, Fg: tui.ColorWhite})

    // 新建纵向框架
    wrapper := tui.NewVBox()

    // 如果有全局标题
    if win.MainTitle != "" {
        var titleLabel *tui.Label
        if win.MainTitle != "" {
            titleLabel = tui.NewLabel(win.MainTitle)
        } else {
            titleLabel = tui.NewLabel("Main")
        }
        title := &StyledBox{
            Style: "mainTitle",
            Box:   tui.NewHBox(tui.NewSpacer(), titleLabel, tui.NewSpacer()),
        }
        titleLabel.SetStyleName("mainTitle")
        t.SetStyle("mainTitle", tui.Style{Bg: tui.ColorWhite, Fg: tui.ColorBlack})

        wrapper.Append(title)
    }

    wrapper.Append(tui.NewSpacer())
    wrapper.Append(tui.NewHBox(tui.NewSpacer(), window, tui.NewSpacer()))
    wrapper.Append(tui.NewSpacer())

    // 新建UI及设置它的主题
    ui, err := tui.New(wrapper)
    if err != nil {
        log.Fatal(err)
    }

    // 按钮
    choicesMap := make(map[*tui.Button]int, len(win.Button))
    choices := make([]tui.Widget, len(win.Button))
    for i, ch := range win.Button {
        but := tui.NewButton(fmt.Sprintf("[%s]", ch))
        // 设置按钮事件
        choicesMap[but] = i
        choices[i] = tui.NewPadder(1, 1, but)
        but.OnActivated(func(b *tui.Button) {
            selectedButton = choicesMap[b]
            ui.Quit()
        })
    }
    t.SetStyle("button.focused", tui.Style{Bg: tui.ColorWhite, Fg: tui.ColorBlack})
    choices[0].SetFocused(true)

    buttons := tui.NewHBox(tui.NewSpacer())
    for _, but := range choices {
        buttons.Append(but)
    }
    buttons.Append(tui.NewSpacer())
    windowBox.Append(buttons)

    // 设置主题
    ui.SetTheme(t)
    tui.DefaultFocusChain.Set(choices...)

    if err := ui.Run(); err != nil {
        log.Fatal(err)
    }
    return
}