package interfaces

import (
    "github.com/marcusolsson/tui-go"
    "log"
    "time"
)

type InitWin struct {
    Logo       string
    Prompt     string
    ShowPeriod int // 毫秒
}

// StyledBox is a Box with an overriden Draw method.
// Embedding a Widget within another allows overriding of some behaviors.
type StyledBox struct {
    Style string
    *tui.Box
}

// Draw decorates the Draw call to the widget with a style.
func (s *StyledBox) Draw(p *tui.Painter) {
    p.WithStyle(s.Style, func(p *tui.Painter) {
        s.Box.Draw(p)
    })
}

func NewInitWindow(win InitWin) {
    // 设置logo
    if win.Logo == "" {
        win.Logo = defLogo
    }
    t := tui.NewTheme()

    // 设置进度条
    progress := tui.NewProgress(100)
    progress.SetSizePolicy(tui.Expanding, tui.Maximum)
    prompt := tui.NewLabel(win.Prompt)

    // 新建窗口容纳上述内容，以及有一个蓝色背景
    window := &StyledBox{
        Style: "initwin",
        Box:   tui.NewVBox(
            tui.NewPadder(10, 1, tui.NewLabel(win.Logo)),
            tui.NewPadder(8, 1, prompt),
        ),
    }
    t.SetStyle("initwin", tui.Style{Bg: tui.ColorBlue, Fg: tui.ColorWhite})

    // 新建纵向框架
    wrapper := tui.NewVBox(
        tui.NewSpacer(),
        window,
        tui.NewSpacer(),
    )
    // 新建横向框架
    content := tui.NewHBox(tui.NewSpacer(), wrapper, tui.NewSpacer())

    // 新建屏幕
    root := tui.NewVBox(
        content,
        tui.NewPadder(0, 0, progress),
    )

    // 新建UI及设置它的主题
    ui, err := tui.New(root)
    ui.SetTheme(t)
    if err != nil {
        log.Fatal(err)
    }

    // 计时退出
    go func(ui tui.UI, totalTime int, progress *tui.Progress) {
        i := 0
        for i < totalTime {
            i += 50
            ui.Update(func() {
                progress.SetCurrent(int(float32(i) / float32(totalTime) * 100))
                progress.SetMax(100)
            })
            time.Sleep(50 * time.Millisecond)
        }
        ui.Quit()
    }(ui, win.ShowPeriod, progress)

    if err := ui.Run(); err != nil {
        log.Fatal(err)
    }

}