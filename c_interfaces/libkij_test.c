

#include <stdio.h>
#include <stdlib.h>
#include "libkij.h"
#include "libkij_test.h"

void TestInitWin(void) {
    Kij_InitWin* win = (Kij_InitWin *)malloc(sizeof(Kij_InitWin));
    win->logo = KIJ_DEFLOGO;
    win->showPeriod = 1000;
    win->prompt = "My first programme.";
    win->needProgBar = 'T';
    
    Kij_NewInitWin(win);
    free(win);
}

void TestAuthWin(void) {
    Kij_AuthWin* win = (Kij_AuthWin *)malloc(sizeof(Kij_AuthWin));
    win->logo = KIJ_DEFLOGO;
    win->prompt = "Please login.";
    win->statusBarText = "Press Esc to exit.";
    
    Kij_AuthWin_Results* aur = Kij_NewAuthWin(win);
    printf("Username: %s\nPass: %s\n", aur->username, aur->password);
    Kij_Destroy_AuthWin_Results(aur);
    
    free(win);
}

void TestErrWin(void) {
    Kij_ErrorWin* win = (Kij_ErrorWin *)malloc(sizeof(Kij_ErrorWin));
    win->mainTitle = "";
    win->winTitle = "Error!";
    win->errInfo = "Please provide password.";
    win->buttonsCount = 2;
    char* btns[] = {
        "OK",
        "Cancel"
    };
    win->buttons = btns;
    
    int res = Kij_NewErrorWin(win);
    printf("Selected: %s\n\n", win->buttons[res]);
    
    free(win);
}

void TestInfoWin(void) {
    Kij_InfoWin* win = (Kij_InfoWin *)malloc(sizeof(Kij_InfoWin));
    win->mainTitle = "";
    win->winTitle = "Info";
    win->info = "Goodbye!";
    win->buttonsCount = 2;
    char* btns[] = {
        "OK",
        "Cancel"
    };
    win->buttons = btns;
    
    int res = Kij_NewInfoWin(win);
    printf("Selected: %s\n\n", win->buttons[res]);
    
    free(win);
}


void TestInputWin(void) {
    Kij_InputWin* win = (Kij_InputWin *)malloc(sizeof(Kij_InputWin));
    win->mainTitle = "--XMU Market--";
    win->winTitle = "Input";
    win->inputBoxLabelsCount = 2;
    char* lbls[] = {
        "Name",
        "Number"
    };
    win->inputBoxLabels = lbls;
    
    win->buttonsCount = 2;
    char* btns[] = {
        "OK",
        "Cancel"
    };
    win->buttons = btns;
    
    Kij_InputWin_Results* inr = Kij_NewInputWin(win);
    int i = 0;
    for (; i < inr->textsCount; i++) {
        printf("%s: %s\n", win->inputBoxLabels[i], inr->texts[i]);
    }
    Kij_Destroy_InputWin_Results(inr);
    
    free(win);
}


void TestSelectWin(void) {
    Kij_SelectWin* win = (Kij_SelectWin *)malloc(sizeof(Kij_SelectWin));
    win->mainTitle = "-- XMU Market --";
    win->winTitle = "Info";
    win->choiceDesc = "Choose an operation below using arrow keys";
    win->choicesCount = 3;
    char* ch[] = {
        "Check In-stock Info", "Purchase", "Return Products"
    };
    win->choices = ch;
    win->winFootNote = "Press Enter to submit";
    win->statusBarText = "XMU Market 1.0";
    
    int res = Kij_NewSelectWin(win);
    if (res == 255) {
        printf("Selected: Exit\n");
    } else {
        printf("Selected: %s\n", win->choices[res]);
    }
    free(win);
}

void TestTableWin(void) {
    Kij_TableWin* win = (Kij_TableWin *)malloc(sizeof(Kij_TableWin));
    win->mainTitle = "-- XMU Market --";
    win->winTitle = "Info";
    win->columnCount = 5;
    char* cn[] = {
        "No.", "Name", "Prod.", "Price", "Colour"
    };
    win->columnNames = cn;
    int colWid[] = {1, 3, 2, 2, 2};
    win->columnWidths = colWid;
    char* ro[] = {
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
    };
    win->rows = ro;
    win->rowsCount = 16;
    
    win->buttonsCount = 2;
    char* btns[] = {
        "OK",
        "Cancel"
    };
    win->buttons = btns;
    
    Kij_NewTableWin(win);
    free(win);
}
