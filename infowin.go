package libkij

import (
    "fmt"
    "github.com/chrishanli/tui-go"
    "log"
)

type InfoWin struct {
    MainTitle       string
    InfoWinTitle    string
    Info            string
    Button          []string
}

func InfoWinTest() {
    var sel InfoWin
    sel.MainTitle = " -- XMU Supermarket --"
    sel.InfoWinTitle = "Information"
    sel.Info =
`1. Please note that Apple's 
   guarantee has expired!

2. Please note that Pear's 
   guarantee has expired!`
    sel.Button = []string {
        "OK", "Cancel",
    }

    s := NewInfoWin(&sel)

    log.Println("Closed. Selected index:", s, "which indicates", sel.Button[s])

}

func NewInfoWin(win *InfoWin) (selectedButton int) {
    selectedButton = 0

    // 新建主题
    t := tui.NewTheme()

    // 新建窗口容器
    windowBox := tui.NewVBox()

    // 窗口标题
    windowTitle := &StyledBox {
        Style : "tabletitle",
        Box : tui.NewHBox(tui.NewSpacer(), tui.NewLabel(win.InfoWinTitle), tui.NewSpacer()),
    }
    t.SetStyle("tabletitle", tui.Style{Bg: tui.ColorCyan, Fg: tui.ColorWhite})
    windowBox.Append(tui.NewPadder(1, 1, windowTitle))

    // 窗口描述
    if win.Info != "" {
        windowBox.Append(tui.NewPadder(1, 0, tui.NewLabel(win.Info)))
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
    tui.DefaultFocusChain.Set(choices...)

    // 设置主题
    ui.SetTheme(t)

    if err := ui.Run(); err != nil {
        log.Fatal(err)
    }

    return

    return
}