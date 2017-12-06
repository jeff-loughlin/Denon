# Denon
Serial port control program for Denon AVR receivers.  Command line and web interfaces.

Many Denon AVR receivers have a serial port in the back, and the control protocol is well documented
(see Denon_AVR4311_PROTOCOL_V7.2.0.pdf).  This program allows you to talk to your Denon AVR receiver using that protocol. 

I have mine connected to a Raspberry Pi via a USB-Serial adapter cable.  Using the command-line program, I can talk to my 
receiver using simple commands to power the unit on or off, select inputs and surround modes, control the volume, and lots 
of other useful parameters.  There is also a (very basic) web interface to allow me to talk to my receiver through a web 
browser.

# Building
The command line program is written in Go.  I am new to Go and wanted to learn how to process command line arguments using 
the flags package.  Since this program has lots and lots of command line arguments, I figured this would be a good way to 
learn how to use it.

This assumes you have Go installed already.  If not, do that first.  I'll wait.

To build the command line program, just type

go build denon.go

You'll need to install the serial package first from http://github.com/tarm/serial


# Installation
Copy denon.php, denon.html, and denon.js to your web root directory.
Copy the compiled command line program to someplace where you can find it (e.g. /usr/local/bin)

Type denon -h to see a list of the command line arguments and what they do.
Point your web browser to the denon.html page to use the web interface.

You'll probably need to change some things in denon.php to make the web interface work in your environment.  My network
is set up with a web server on one machine, and the Denon receiver connected to a different machine.  That means the
web server needs to SSH over to the Denon-connected machine to do its magic.  If your web server is on the same machine, 
just change the command in the php file to remove the SSH part.  The php file is only 4 lines long, so it's not hard to 
see what you need to change (even if you don't know PHP).  I should make this configurable, but for now you'll have to do 
it by hand.  Sorry.

# Other Stuff
I have renamed several of the inputs on my AVR receiver to reflect what's actually connected to them, and the web interface 
has them named that way.  You'll probably want to change these to your taste.  I should make that configurable too.  There's 
lots of other room for improvement, and many additional commands available in the serial protocol that I haven't implemented, 
either because I don't need them, or my receiver doesn't support them.  My receiver is an AVR2310CI.  Newer and higher-end 
receivers have some commands that mine doesn't have, but the protocol is otherwise identical on all Denon receivers that 
support it.

There are a few areas where the protocol documentation falls short (that, or the receiver just implements some things badly 
or weirdly).  This is true particularly around the area of Dolby Prologic II and DTS NEO:6 Music and Cinema modes.  There 
are comments in the code describing the weirdness and what I had to do do overcome it.  It seems to work now, at least for 
my setup, but I don't know if this weirdness is peculiar to my model of receiver (AVR-2310CI), or if it's common to all of 
them.  I welcome any reports of similar weirdness on other receivers, or a proper fix if anyone figures out what's really 
going on in this area.

Newer Denon receivers have a network port instead of a serial port.  The protocol is the same, so it should be fairly simple
to adapt this to the newer networked receivers.  I don't have one, so I'm not going there, but if anyone with a networked
receiver wants to give it a shot, please do so, and send me a pull request.  I won't be able to test it, but as long as you 
tell me it works, and as long as it doesn't break anything else, I'll accept it.

The web interface is basic and ugly.  It does what I need it to, but a fancier (and mobile-friendly) interface would be nice.  
That's on the to-do list.

Other suggestions, bug reports, and pull requests are welcome.

