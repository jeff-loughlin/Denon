# Denon
Serial port control program for Denon AVR receivers.  Command line and web interfaces.

Many Denon AVR reveivers have a serial port in the back, and the control protocol is well documented
(see Denon_AVR4311_PROTOCOL_V7.2.0.pdf).  This program allows you to talk to your Denon AVR receiver using that protocol. 

I have mine connected to a Raspberry Pi via a USB-Serial adapter cable.  Using the command-line program, I can talk to my receiver using simple commands to power the unit on or off, select inputs and surround modes, control the volume, and lots of other useful parameters.  There is also a (very basic) web interface to allow me to talk to my receiver through a web browser.

# Building
The command line program is written in Go.  I am new to Go and wanted to learn how to process command line arguments using the flags package.  Since this program has lots and lots of command line arguments, I figured this would be a good way to learn how to use it.

This assumes you have Go installed already.  If not, do that first.  I'll wait.

To build the command line program, just type

go build denon.go

You'll need to install the serial package first from http://github.com/tarm/serial


# Installation
Copy denon.php, denon.html, and denon.js to your web root folder.
Copy the compiled command line program to someplace where you can find it (e.g. /usr/local/bin)

Type denon -h to see a list of command line arguments and what they do.
Point your browser to the denon.html page to use the web interface.

You'll probably need to change some things in denon.php to make the web interface work in your environment.  For me, the web server is on a different machine from the one connected to my Denon receiver, so my PHP needs to SSH over to that machine to execute the command line program.  If your web server runs on the same machine, change that to execute the program directly.  That should be all you need to do.

# Other Stuff
I have renamed several of the inputs on my AVR receiver to reflect what's actually connected to them, and the web interface has them named that way (the command interface accepts both the default input names and my renamed ones).  I should make that configurable in the web interface (along with the command interface path, as noted above).  There's lots of other room for improvement, and many additional commands available in the serial protocol that I haven't implemented, either because I don't need them, or my receiver doesn't support them.  My receiver is an AVR2310CI.  Newer and higher-end receivers have some commands that mine doesn't have, but the protocol is otherwise identical on all Denon receivers that support it.


Suggestions and pull requests are welcome.
