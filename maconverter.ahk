#NoEnv  ; Recommended for performance and compatibility with future AutoHotkey releases.
SetWorkingDir %A_ScriptDir%  ; Ensures a consistent starting directory.
#SingleInstance force
SetTitleMatchMode, 2

; Cisco, xxxx.xxxx.xxxx
+^1:: ;CTRL+Shift+1
; Run "D:\repositories\Projects\go\maconverter\maconverter.exe" --cisco, ,Hide
Run "maconverter.exe" --cisco, ,Hide
Return

; Nokia, xx:xx:xx:xx:xx:xx
+^2:: ;CTRL+Shift+2
; Run "D:\repositories\Projects\go\maconverter\maconverter.exe" --nokia, ,Hide
Run "maconverter.exe" --nokia, ,Hide
Return

; Huawei, xxxx-xxxx-xxxx
+^3:: ;CTRL+Shift+3
; Run "D:\repositories\Projects\go\maconverter\maconverter.exe" --huawei, ,Hide
Run "maconverter.exe" --huawei, ,Hide
Return
