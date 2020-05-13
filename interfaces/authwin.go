package interfaces

import (
    "github.com/marcusolsson/tui-go"
    "log"
)

var defLogo =
` _ _ _     _    _ _
| (_) |   | |  (_|_)
| |_| |__ | | ___ _
| | | '_ \| |/ / | |
| | | |_) |   <| | |
|_|_|_.__/|_|\_\_| |
                _/ |
               |__/`

type AuthWin struct {
    Logo            string
    Prompt          string
    StatusBarText   string
}

func AuthWinTest() {
    var a AuthWin
    a.Logo = defLogo
    a.Prompt = "Enter your username and password"
    a.StatusBarText = "Press Esc to Exit!"

    u, p := NewAuthWindow(&a)

    log.Println("Exited. Username:", u, ", password:", p)
}

func NewAuthWindow(win *AuthWin) (username, passcode string) {

    t := tui.NewTheme()

    // 用户名
    userEntry := tui.NewEntry()
    userEntry.SetFocused(true)

    // 密码
    pwEntry := tui.NewEntry()
    pwEntry.SetEchoMode(tui.EchoModePassword)

    // 新建一个vbox用来容纳user、pass的Label
    captions := tui.NewVBox()
    captions.Append(tui.NewPadder(1, 1, tui.NewHBox(tui.NewSpacer(), tui.NewLabel("Username"))))
    captions.Append(tui.NewPadder(1, 0, tui.NewHBox(tui.NewSpacer(), tui.NewLabel("Password"))))
    captions.SetSizePolicy(tui.Maximum, tui.Maximum)

    // 该vbox容纳user、pass的文本框
    txtBoxes := tui.NewVBox()
    txtBoxUsername := &StyledBox {
        Style: "txtBox",
        Box:   tui.NewHBox(userEntry),
    }
    t.SetStyle("txtBox", tui.Style{Bg: tui.ColorWhite, Fg: tui.ColorBlack})
    txtBoxes.Append(tui.NewPadder(1, 1, txtBoxUsername))
    txtBoxPass := &StyledBox {
        Style: "txtBox",
        Box:   tui.NewHBox(pwEntry),
    }
    txtBoxes.Append(tui.NewPadder(1, 0, txtBoxPass))

    // 状态栏
    status := tui.NewStatusBar(win.StatusBarText)

    // 登陆及登陆事件
    btnLogin := tui.NewButton("[Login]")

    // 收集按钮
    buttons := tui.NewHBox (
        tui.NewSpacer(),
        tui.NewPadder(1, 0, btnLogin),
    )

    // 新建窗口容纳上述内容，及有一个border
    window := tui.NewVBox(
        tui.NewPadder(10, 1, tui.NewHBox(tui.NewSpacer(), tui.NewLabel(win.Logo), tui.NewSpacer())),
        tui.NewPadder(12, 0, tui.NewHBox(tui.NewSpacer(), tui.NewLabel(win.Prompt), tui.NewSpacer())),
        tui.NewPadder(1, 1, tui.NewHBox(captions, txtBoxes)),
        buttons,
    )
    window.SetBorder(true)

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
        status,
    )

    // 设置默认tab顺序
    tui.DefaultFocusChain.Set(userEntry, pwEntry, btnLogin)

    // 新建UI
    ui, err := tui.New(root)
    if err != nil {
        log.Fatal(err)
    }

    btnLogin.OnActivated(func(b *tui.Button) {
        username = userEntry.Text()
        passcode = pwEntry.Text()
        ui.Quit()

    })
    t.SetStyle("button.focused", tui.Style{Bg: tui.ColorWhite, Fg: tui.ColorBlack})

    // 新建键位绑定
    ui.SetKeybinding("Esc", func() { ui.Quit() })
    ui.SetTheme(t)

    if err := ui.Run(); err != nil {
        log.Fatal(err)
    }

    return
}