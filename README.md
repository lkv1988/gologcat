gologcat
===

Print colorful logcat by golang.

###Install


	go get github.com/airk000/gologcat

> bin file, to be continue...

###Configuration

`gologcat` will read `$HOME/.gologcat` every time. You can config your own colors by this file, such as:

	i:red
	v:green
	e:cyan

All available key is among "IVEWD" and all available value(color):

	none(default color of your shell)
	black
	red
	green
	yellow
	blue
	magenta
	cyan
	white

All the key and values are caseless matching. The default color profile is:

	d:green
	i:cyan
	w:yellow
	e:red
	v:none

###Thanks

[go-colortext](https://github.com/daviddengcn/go-colortext)