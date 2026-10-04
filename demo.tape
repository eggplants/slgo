Output demo.gif

Require go

Set Shell "bash"
Set FontSize 14
Set Width 1200
Set Height 500
Set Padding 10
Set TypingSpeed 75ms
Set Framerate 20

Hide
Type "go build -o sl ./cmd/sl && export PATH=$PWD:$PATH PS1='> ' && clear"
Enter
Wait@60s
Show

Type "sl"
Sleep 300ms
Enter
Wait@30s
Sleep 500ms

Type "sl -a -F"
Sleep 300ms
Enter
Wait@30s
Sleep 500ms

Type "sl -l"
Sleep 300ms
Enter
Wait@30s
Sleep 1s

Type "sl -p"
Sleep 300ms
Enter
Sleep 15s
