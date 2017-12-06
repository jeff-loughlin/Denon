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


This program is very basic and there is a lot of room for improvement.  Suggestions and pull requests are welcome.
