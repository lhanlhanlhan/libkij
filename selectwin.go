package libkij

import (
    "github.com/chrishanli/tui-go"
    "log"
)

type SelectWin struct {
    MainTitle           string
    ChoicePadTitle      string
    ChoicePadDesc       string
    Choices []          string
    ChoicePadFootnote   string
    StatusBarContent    string
}

func SelectWinTest() {
    var sel SelectWin
    sel.MainTitle = " -- XMU Supermarket --"
    sel.ChoicePadTitle = "World!"
    sel.ChoicePadDesc = "Please select!"
    sel.Choices = []string {
        "jujuju", "pangpangpang",
    }
    sel.ChoicePadFootnote = "Select and press Enter to confirm"
    sel.StatusBarContent = "i am status"

    s := NewSelectWin(&sel)

    if s == 255 {
        log.Println("Selected 'Close'")
    } else {
        log.Println("Closed. Selected index:", s, "which indicates", sel.Choices[s])
    }
}

func NewSelectWin(win *SelectWin) (selectedItem int) {
    selectedItem = 255
    // 新建主题
    t := tui.NewTheme()

    // 新建窗口容器
    windowBox := tui.NewVBox()

    // 窗口标题
    windowBox.Append(tui.NewPadder(1, 1, tui.NewHBox(tui.NewSpacer(), tui.NewLabel(win.ChoicePadTitle), tui.NewSpacer())))
    // 窗口描述
    if win.ChoicePadDesc != "" {
        windowBox.Append(tui.NewPadder(1, 0, tui.NewLabel(win.ChoicePadDesc)))
    }

    // 选项
    choices := tui.NewList()
    for _, ch := range win.Choices {
        choices.AddItems(ch)
    }
    choices.AddItems("Quit")
    choices.SetFocused(true)
    choices.SetSelected(0)
    t.SetStyle("list.item", tui.Style{Bg: tui.ColorBlue, Fg: tui.ColorWhite})
    t.SetStyle("list.item.selected", tui.Style{Bg: tui.ColorWhite, Fg: tui.ColorBlack})
    windowBox.Append(tui.NewPadder(5, 1, choices))

    // 窗口脚注
    if win.ChoicePadFootnote != "" {
        windowBox.Append(tui.NewPadder(1, 1, tui.NewLabel(win.ChoicePadFootnote)))
    }

    // 设置状态栏
    status := tui.NewStatusBar(win.StatusBarContent)

    // 新建窗口容纳上述内容，以及有一个蓝色背景
    window := &StyledBox{
        Style: "selectwin",
        Box:   windowBox,
    }
    t.SetStyle("selectwin", tui.Style{Bg: tui.ColorBlue, Fg: tui.ColorWhite})

    // 全局标题
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

    // 新建纵向框架
    wrapper := tui.NewVBox(
        title,
        tui.NewSpacer(),
        // 新建横向框架
        tui.NewHBox(tui.NewSpacer(), window, tui.NewSpacer()),
        tui.NewSpacer(),
    )

    // 新建主显示
    root := tui.NewVBox(
        wrapper,
        tui.NewPadder(0, 0, status),
    )

    // 新建UI及设置它的主题
    ui, err := tui.New(root)
    if err != nil {
        log.Fatal(err)
    }

    // 设置点按事件
    choices.OnItemActivated(func(t *tui.List) {
        // 返回值
        selectedItem = t.Selected()
        if selectedItem == len(win.Choices) {
            // 255 表示点按了退出
            selectedItem = 255
        }
        ui.Quit()
    })

    // 设置主题
    ui.SetTheme(t)

    if err := ui.Run(); err != nil {
        log.Fatal(err)
    }
    return
}