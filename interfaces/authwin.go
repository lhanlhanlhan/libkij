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

func NewAuthWindow() {
    // 用户名
    user := tui.NewEntry()
    user.SetFocused(true)

    // 密码
    password := tui.NewEntry()
    password.SetEchoMode(tui.EchoModePassword)

    // 新建网格
    form := tui.NewGrid(0, 0)
    form.AppendRow(tui.NewLabel("User"), user)
    form.AppendRow(tui.NewLabel("Password"), password)

    // 状态栏
    status := tui.NewStatusBar("Ready.")

    // 登陆及登陆事件
    btnLogin := tui.NewButton("[Login]")
    btnLogin.OnActivated(func(b *tui.Button) {
        status.SetText("Logged in.")
    })

    // 收集按钮
    buttons := tui.NewHBox(
        tui.NewSpacer(),
        tui.NewPadder(1, 0, btnLogin),
    )

    // 新建窗口容纳上述内容，及有一个border
    window := tui.NewVBox(
        tui.NewPadder(10, 1, tui.NewLabel(defLogo)),
        tui.NewPadder(12, 0, tui.NewLabel("Welcome!")),
        tui.NewPadder(1, 1, form),
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
    tui.DefaultFocusChain.Set(user, password, btnLogin)

    // 新建UI
    ui, err := tui.New(root)
    if err != nil {
        log.Fatal(err)
    }

    // 新建键位绑定
    ui.SetKeybinding("Esc", func() { ui.Quit() })

    if err := ui.Run(); err != nil {
        log.Fatal(err)
    }
}