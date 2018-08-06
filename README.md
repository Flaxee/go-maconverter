# Readme

This is a small application/package written in Go, which converts/translates MAC-adresses from one format to another.<br/>
The following 3 formats are currently available:<br/>
xxxx.xxxx.xxxx<br/>
xx:xx:xx:xx:xx:xx<br/>
xxxx-xxxx-xxxx<br/>

Included is also an AutoHotKey script so you can run the program by using keybindings.<br/>
This AHK script will look for the file **maconverter.exe** within the current directory, so make sure to compile the Go package.<br/>
By default, the keybindings are as follows:<br/>
**CTRL+Shift+1** = xxxx.xxxx.xxxx<br/>
**CTRL+Shift+2** = xx:xx:xx:xx:xx:xx<br/>
**CTRL+Shift+3** = xxxx-xxxx-xxxx<br/>


# Run & Compile
Run the package directly: **go run maconverter.go**<br/>
Compile the package: **go install maconverter.go** OR **go build maconverter.go**<br/>
