package interfaces

import (
    "fmt"
    "github.com/marcusolsson/tui-go"
    "log"
    "strings"
)

type TableWin struct {
    MainTitle       string
    TableName       string
    ColNames        []string
    ColsWidth       []int
    Rows            []string
    Button          []string
}

func TableWinTest() {
    var tbl TableWin
    tbl.MainTitle = "-- XMU Supermarket --"
    tbl.TableName = "-- In-Stock Info --"
    tbl.ColNames = []string {
        "No.", "Name", "Prod.", "Price", "Colour",
    }
    tbl.ColsWidth = []int {
        1, 3, 2, 2, 2,
    }
    tbl.Rows = []string {
        "1\nApple Computer\n3/3/2020\n$885.-\nBlack",
        "2\nMicrosoft Computer\n5/3/2020\n$400.20\n",
        "3\nAir Jordan 1\n12/5/2020\n$199.-\nRed",
        "4\nApple Computer\n3/3/2020\n$885.-\nWhite",
        "5\nMicrosoft Computer\n5/3/2020\n$400.20\n",
        "6\nAir Jordan 1\n12/5/2020\n$199.-\nGreen",
        "7\nApple Computer\n3/3/2020\n$885.-\nSilver",
        "8\nMicrosoft Computer\n5/3/2020\n$400.20\n",
        "9\nAir Jordan 1\n12/5/2020\n$199.-\nBlue",
        "10\nMicrosoft Computer\n5/3/2020\n$400.20\n",
        "11\nAir Jordan 1\n12/5/2020\n$199.-\nGreen",
        "12\nApple Computer\n3/3/2020\n$885.-\nSilver",
        "13\nMicrosoft Computer\n5/3/2020\n$400.20\n",
        "14\nAir Jordan 1\n12/5/2020\n$199.-\nBlue",
        "15\nMicrosoft Computer\n5/3/2020\n$400.20\n",
        "16\nAir Jordan 1\n12/5/2020\n$199.-\nGreen",
        "17\nApple Computer\n3/3/2020\n$885.-\nSilver",
        "16\nAir Jordan 1\n12/5/2020\n$199.-\nBlack",
    }

    tbl.Button = []string {
        "Search",
        "Read a list from disk",
        "Exit",
    }

    s := NewTableWin(&tbl)

    log.Println("Closed. Selected index:", s, "which indicates", tbl.Button[s])

}

func NewTableWin(win *TableWin) (selectedButton int) {
    selectedButton = 0

    // 新建主题
    t := tui.NewTheme()

    // 新建窗口容器
    windowBox := tui.NewVBox()

    // 窗口标题
    windowTitle := &StyledBox {
        Style : "tabletitle",
        Box : tui.NewHBox(tui.NewSpacer(), tui.NewLabel(win.TableName), tui.NewSpacer()),
    }
    t.SetStyle("tabletitle", tui.Style{Bg: tui.ColorCyan, Fg: tui.ColorWhite})

    //windowBox.Append(tui.NewPadder(1, 1, windowTitle))
    windowBox.Append(windowTitle)

    // 表头
    colsCount := len(win.ColNames)
    table := tui.NewGrid(colsCount, 0)
    cols := make([]tui.Widget, colsCount)
    for index, l := range win.ColsWidth {
        // 设定列长度
        table.SetColumnStretch(index, l)
        // 添加列头内容
        cols[index] = tui.NewLabel(win.ColNames[index])
    }
    table.AppendRow(cols...)
    // 表身
    for _, r := range win.Rows {
        // 利用"\n"来分裂该行
        tokens := strings.Split(r, "\n")
        for i, tok := range tokens {
            // 添加该行该列列内容
            cols[i] = tui.NewLabel(tok)
        }
        // 增加一行
        table.AppendRow(cols...)
    }
    table.SetSizePolicy(tui.Expanding, tui.Expanding)

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
    wrapper.Append(tui.NewHBox(window))
    wrapper.Append(tui.NewSpacer())

    // 新建UI及设置它的主题
    ui, err := tui.New(wrapper)
    if err != nil {
        log.Fatal(err)
    }

    // 查看：是否需要滚动
    if len(win.Rows) > 18 {
        // 把表加入窗口
        s := tui.NewScrollArea(table)
        s.SetSizePolicy(tui.Expanding, tui.Expanding)

        scrollBox := tui.NewVBox(s)
        scrollBox.SetBorder(true)
        scrollBox.SetSizePolicy(tui.Expanding, tui.Expanding)

        windowBox.Insert(1, scrollBox)

        // 设置键位
        ui.SetKeybinding("Up", func() {s.Scroll(0, -1)})
        ui.SetKeybinding("Down", func() {s.Scroll(0, 1)})
    } else {
        // 把表加入窗口
        scrollBox := tui.NewVBox(table)
        scrollBox.SetBorder(true)
        scrollBox.SetSizePolicy(tui.Expanding, tui.Expanding)

        windowBox.Insert(1, scrollBox)
    }

    // 按钮
    choicesMap := make(map[*tui.Button]int, len(win.Button))
    choices := make([]tui.Widget, len(win.Button))
    for i, ch := range win.Button {
        but := tui.NewButton(fmt.Sprintf("[%s]", ch))
        // 设置按钮事件
        choicesMap[but] = i
        choices[i] = tui.NewPadder(1, 0, but)
        but.OnActivated(func(b *tui.Button) {
            selectedButton = choicesMap[b]
            ui.Quit()
        })
    }
    choices[0].SetFocused(true)

    t.SetStyle("button.focused", tui.Style{Bg: tui.ColorWhite, Fg: tui.ColorBlack})

    buttons := tui.NewHBox(tui.NewSpacer())
    for _, but := range choices {
        buttons.Append(but)
    }
    buttons.Append(tui.NewSpacer())

    windowBox.Append(buttons)

    // 设置主题
    ui.SetTheme(t)

    // 设置tab键位
    tui.DefaultFocusChain.Set(choices...)

    if err := ui.Run(); err != nil {
        log.Fatal(err)
    }

    return
}