//
//  libkij.c
//  Han-Li C-Language
//
//  Created by Han Li on 13/5/2020.
//  Copyright © 2020 Han Li. All rights reserved.
//  注意：内存由谁分配，就由谁释放！

#include "include/libkij_go.h"
#include <stdlib.h>
#include "libkij.h"

// 1.1 1 新建登陆界面
Kij_AuthWin_Results* Kij_NewAuthWin(Kij_AuthWin* win) {
    struct NewAuthWin_return ret = NewAuthWin(win->logo, win->prompt, win->statusBarText);
    
    Kij_AuthWin_Results* res = (Kij_AuthWin_Results *)malloc(sizeof(Kij_AuthWin_Results));
    res->username = ret.r0;
    res->password = ret.r1;
    
    return res;
}

// 1.2 3 销毁登陆界面返回
void Kij_Destroy_AuthWin_Results(Kij_AuthWin_Results* win) {
    free(win->username);
    free(win->password);
    free(win);
}

// 2.1 4 新建错误弹框
int Kij_NewErrorWin(Kij_ErrorWin* win) {
    return NewErrWin(win->mainTitle, win->winTitle, win->errInfo, win->buttons, win->buttonsCount);
}

// 3.1 6 新建信息弹框
int Kij_NewInfoWin(Kij_InfoWin* win) {
    return NewInfoWin(win->mainTitle, win->winTitle, win->info, win->buttons, win->buttonsCount);
}

// 4.1 8 新建键入界面
Kij_InputWin_Results* Kij_NewInputWin(Kij_InputWin* win) {
    struct NewInputWin_return ret = NewInputWin(win->mainTitle, win->winTitle, win->inputBoxLabels, win->inputBoxLabelsCount, win->buttons, win->buttonsCount);
    
    Kij_InputWin_Results* res = (Kij_InputWin_Results *)malloc(sizeof(Kij_InputWin_Results));
    res->selectedButton = ret.r0;
    res->texts = ret.r1;
    res->textsCount = ret.r2;
    
    return res;
}

// 4.3 10 销毁键入界面返回
void Kij_Destroy_InputWin_Results(Kij_InputWin_Results* win) {
    int i;
    for (i = 0; i < win->textsCount; i++) {
        free(win->texts[i]);
    }
    free(win->texts);
    free(win);
}

// 5.1 11 新建选择窗口
int Kij_NewSelectWin(Kij_SelectWin* win) {
    return NewSelectWin(win->mainTitle, win->winTitle, win->choiceDesc, win->choices, win->choicesCount, win->winFootNote, win->statusBarText);
}

// 6.1 13 新建表格窗口
int Kij_NewTableWin(Kij_TableWin* win) {
    return NewTableWin(win->mainTitle, win->winTitle, win->columnNames, win->columnCount, win->columnWidths, win->rows, win->rowsCount, win->buttons, win->buttonsCount);
}

// 7.1 15 新建初始化窗口
void Kij_NewInitWin(Kij_InitWin* win) {
    NewInitWin(win->logo, win->prompt, win->showPeriod, win->needProgBar);
}
