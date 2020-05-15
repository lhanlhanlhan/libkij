//
//  libkij.h
//  Han-Li C-Language
//
//  Created by Han Li on 13/5/2020.
//  Copyright © 2020 Han Li. All rights reserved.
//

#ifndef libkij_h
#define libkij_h

#ifdef __cplusplus
extern "C" {
#endif

#include <stdio.h>

// a_1) 登陆界面参数
typedef struct {
    const char* logo;
    const char* prompt;
    const char* statusBarText;
} Kij_AuthWin;

// a_2) 登陆界面返回
typedef struct {
    const char* username;
    const char* password;
} Kij_AuthWin_Results;

// b) 错误弹框参数
typedef struct {
    const char*   mainTitle;
    const char*   winTitle;
    const char*   errInfo;
    const char**  buttons;
    int     buttonsCount;
} Kij_ErrorWin;

// c) 信息弹框参数
typedef struct {
    const char*   mainTitle;
    const char*   winTitle;
    const char*   info;
    const char**  buttons;
    int     buttonsCount;
} Kij_InfoWin;

// d) 初始化窗口参数
typedef struct {
    const char*   logo;
    const char*   prompt;
    int     showPeriod;
    char    needProgBar;
} Kij_InitWin;

// e_1) 键入窗口参数
typedef struct {
    const char*   mainTitle;
    const char*   winTitle;
    const char**  inputBoxLabels;
    int     inputBoxLabelsCount;
    const char**  buttons;
    int     buttonsCount;
} Kij_InputWin;

// e_2) 键入窗口返回
typedef struct {
    int     selectedButton;
    const char**  texts;
    int     textsCount;
} Kij_InputWin_Results;

// f) 选择窗口参数
typedef struct {
    const char*   mainTitle;
    const char*   winTitle;
    const char*   choiceDesc;
    const char**  choices;
    int     choicesCount;
    const char*   winFootNote;
    const char*   statusBarText;
} Kij_SelectWin;

// g) 表格窗口参数
typedef struct {
    const char*   mainTitle;
    const char*   winTitle;
    const char**  columnNames;
    int*    columnWidths;
    int     columnCount;
    char**  rows;
    int     rowsCount;
    const char**  buttons;
    int     buttonsCount;
} Kij_TableWin;

// logo
#define KIJ_DEFLOGO "` _ _ _     _    _ _\n| (_) |   | |  (_|_)\n| |_| |__ | | ___ _\n| | | '_ \\| |/ / | |\n| | | |_) |   <| | |\n|_|_|_.__/|_|\\_\\_| |\n                _/ |\n               |__/`"


// 1.1 1 新建登陆界面
extern Kij_AuthWin_Results* Kij_NewAuthWin(Kij_AuthWin* win);

// 1.2 2 销毁登陆界面参数
extern void Kij_Destroy_AuthWin(Kij_AuthWin* win);

// 1.3 3 销毁登陆界面返回
extern void Kij_Destroy_AuthWin_Results(Kij_AuthWin_Results* win);


// 2.1 4 新建错误弹框
extern int Kij_NewErrorWin(Kij_ErrorWin* win);

// 2.2 5 销毁错误弹框参数
extern void Kij_Destroy_ErrorWin(Kij_ErrorWin* win);


// 3.1 6 新建信息弹框
extern int Kij_NewInfoWin(Kij_InfoWin* win);

// 3.2 7 销毁信息弹框参数
extern void Kij_Destroy_InfoWin(Kij_InfoWin* win);


// 4.1 8 新建键入界面
extern Kij_InputWin_Results* Kij_NewInputWin(Kij_InputWin* win);

// 4.2 9 销毁键入界面参数
extern void Kij_Destroy_InputWin(Kij_InputWin* win);

// 4.3 10 销毁键入界面返回
extern void Kij_Destroy_InputWin_Results(Kij_InputWin_Results* win);


// 5.1 11 新建选择窗口
extern int Kij_NewSelectWin(Kij_SelectWin* win);

// 5.2 12 销毁选择窗口参数
extern void Kij_Destroy_SelectWin(Kij_SelectWin* win);


// 6.1 13 新建表格窗口
extern int Kij_NewTableWin(Kij_TableWin* win);

// 6.2 14 销毁表格窗口参数
extern void Kij_Destroy_TableWin(Kij_TableWin* win);

// 7.1 15 新建初始化窗口
extern void Kij_NewInitWin(Kij_InitWin* win);

// 7.1 16 销毁初始化窗口参数
extern void Kij_Destroy_InitWin(Kij_InitWin* win);


#ifdef __cplusplus
}
#endif

#endif /* libkij_h */
