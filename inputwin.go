package libkij

import (
    "fmt"
    "github.com/chrishanli/tui-go"
    "log"
)

type InputWin struct {
    MainTitle       string
    InputWinName    string
    InputBoxNames   []string
    Button          []string
}

func InputWinTest() {
    var tbl InputWin
    tbl.MainTitle = "-- XMU Supermarket --"
    tbl.InputWinName = "-- Search --"
    tbl.InputBoxNames = []string {
        "No.", "Name.", "Date",
    }

    tbl.Button = []string {
        "Search",
        "Read a list from disk",
        "Exit",
    }

    s, texts := NewInputWin(&tbl)

    log.Println("Closed. Selected index:", s, "which indicates", tbl.Button[s])

    log.Println("Got entries:", texts)
}

func NewInputWin(win *InputWin) (selectedButton int, text []string) {
    selectedButton = 0
    inputBoxCount := len(win.InputBoxNames)
    text = make([]string, inputBoxCount)

    // 新建主题
    t := tui.NewTheme()

    // 新建窗口容器
    windowBox := tui.NewVBox()

    // 窗口标题
    windowTitle := &StyledBox {
        Style : "tabletitle",
        Box : tui.NewHBox(tui.NewSpacer(), tui.NewLabel(win.InputWinName), tui.NewSpacer()),
    }
    t.SetStyle("tabletitle", tui.Style{Bg: tui.ColorCyan, Fg: tui.ColorWhite})

    windowBox.Append(tui.NewPadder(1, 1, windowTitle))
    //windowBox.Append(windowTitle)

    // 新建一个vbox用来容纳各个txtBox的标题
    captions := tui.NewVBox()
    for _, ch := range win.InputBoxNames {
        capBox := tui.NewHBox(tui.NewSpacer(), tui.NewLabel(ch))
        captions.Append(tui.NewPadder(1, 1, capBox))
    }
    captions.SetSizePolicy(tui.Maximum, tui.Maximum)

    // 新建一个vbox用来容纳各个txtBox
    txtBoxes := tui.NewVBox()
    widges := make([]tui.Widget, inputBoxCount + len(win.Button))
    txtBoxesCollEntr := make([]*tui.Entry, inputBoxCount)
    for i := 0; i < inputBoxCount; i++ {
        inp := tui.NewEntry()
        inpBox := &StyledBox {
            Style: "txtBox",
            Box:   tui.NewHBox(inp),
        }
        t.SetStyle("txtBox", tui.Style{Bg: tui.ColorWhite, Fg: tui.ColorBlack})

        txtBoxes.Append(tui.NewPadder(1, 1, inpBox))
        widges[i] = inp
        txtBoxesCollEntr[i] = inp
    }

    windowBox.Append(tui.NewHBox(captions, txtBoxes))

    // 新建窗口容纳上述内容，以及有一个蓝色背景
    window := &StyledBox{
        Style: "selectwin",
        //Box:   scrollBox,
        Box:   windowBox,
    }
    t.SetStyle("selectwin", tui.Style{Bg: tui.ColorBlue, Fg: tui.ColorWhite})

    // 新建纵向全局框架
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
        widges[inputBoxCount + i] = but
        but.OnActivated(func(b *tui.Button) {
            selectedButton = choicesMap[b]
            for j, txt := range txtBoxesCollEntr {
                text[j] = txt.Text()
            }
            ui.Quit()
        })
    }
    t.SetStyle("button.focused", tui.Style{Bg: tui.ColorWhite, Fg: tui.ColorBlack})

    buttons := tui.NewHBox(choices...)
    buttons.Append(tui.NewSpacer())

    windowBox.Append(buttons)

    // 设置主题
    ui.SetTheme(t)

    // 设置tab键位
    tui.DefaultFocusChain.Set(widges...)

    if err := ui.Run(); err != nil {
        log.Fatal(err)
    }

    return
}