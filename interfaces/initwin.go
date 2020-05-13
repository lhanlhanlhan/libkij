package interfaces

import (
    "github.com/chrishanli/tui-go"
    "log"
    "time"
)

type InitWin struct {
    Logo        string
    Prompt      string
    ShowPeriod  int // 毫秒
    NeedProgBar bool
}

func InitWinTest() {
    var init InitWin
    init.Logo = "   _  __ __  _____  ____  ___           __        __ \n  | |/ //  |/  / / / /  |/  /___ ______/ /_____  / /_\n  |   // /|_/ / / / / /|_/ / __ `/ ___/ //_/ _ \\/ __/\n /   |/ /  / / /_/ / /  / / /_/ / /  / ,< /  __/ /_  \n/_/|_/_/  /_/\\____/_/  /_/\\__,_/_/  /_/|_|\\___/\\__/"
    init.Prompt = "Initialising, please wait..."
    init.ShowPeriod = 2000
    init.NeedProgBar = false

    NewInitWindow(&init)

    log.Println("Closed.")
}

func NewInitWindow(win* InitWin) {
    // 设置logo
    if win.Logo == "" {
        win.Logo = defLogo
    }

    // 新建主题
    t := tui.NewTheme()

    // 新建logo
    logo := tui.NewHBox(tui.NewSpacer(), tui.NewLabel(win.Logo), tui.NewSpacer())

    // 新建一个提示信息
    prompt := tui.NewHBox(tui.NewSpacer(), tui.NewLabel(win.Prompt), tui.NewSpacer())

    // 新建窗口容纳上述内容，以及有一个蓝色背景
    window := &StyledBox{
        Style: "initwin",
        Box:   tui.NewVBox(
            tui.NewPadder(10, 1, logo),
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
    var root *tui.Box
    var progress *tui.Progress
    if win.NeedProgBar {
        // 设置进度条
        progress = tui.NewProgress(100)
        progress.SetSizePolicy(tui.Expanding, tui.Maximum)
        root = tui.NewVBox(
            content,
            tui.NewPadder(0, 0, progress),
        )
    } else {
        root = tui.NewVBox(
            content,
        )
    }

    // 新建UI及设置它的主题
    ui, err := tui.New(root)
    if err != nil {
        log.Fatal(err)
    }

    ui.SetTheme(t)
    if win.NeedProgBar {
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
    } else {
        go func(ui tui.UI) {
            time.Sleep(time.Duration(win.ShowPeriod) * time.Millisecond)
            ui.Quit()
        }(ui)
    }

    if err := ui.Run(); err != nil {
        log.Fatal(err)
    }

}