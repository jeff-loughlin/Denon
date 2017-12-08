# Denon
Serial port control program for Denon AVR receivers, with command line and web interfaces.

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

To build the command line program, just type:

go build denon.go

You'll need to install the serial package first from http://github.com/tarm/serial


# Installation
Copy the compiled command line program to someplace where you can find it (e.g. /usr/local/bin).
If you want to use the web interface too, copy denon.php, denon.html, and denon.js to your web root directory (or someplace
where the web server can find them).

Type denon -h to see a list of the command line arguments and what they do.
Point your web browser to the denon.html page on your web server to use the web interface.  

<b>NOTE: There is no security in the web interface.  I don't recommend putting this on a public-facing web server!</b>  If
you do, you are likely to find your receiver blasting Van Halen at full volume at 3:00 in the morning.  Don't blame me if that happens.

If you installed the command line program to someplace other than /usr/local/bin, you'll need to change denon.php to reflect
where it is.  It's only 4 lines long, so you should have no trouble figuring out what to change.


# Other Stuff
I have renamed several of the inputs on my AVR receiver to reflect what's actually connected to them, and the web interface 
has them named that way.  You'll probably want to change these to match your setup.  I should make that configurable.

There's lots of other room for improvement, and many additional commands available in the serial protocol that I haven't 
implemented, either because I don't need them, or my receiver doesn't support them.  My receiver is an AVR-2310CI.  Newer 
and higher-end receivers have some commands that mine doesn't have, but the protocol is otherwise identical on all Denon 
receivers that support it.

There are a few areas where the protocol documentation falls short (either that, or the receiver just implements some things badly 
or weirdly).  This is true particularly around the area of Dolby Prologic II and DTS NEO:6 Music and Cinema modes.  There 
are comments in the code describing the weirdness and what I had to do do overcome it.  It seems to work now, at least for 
my setup, but I don't know if this weirdness is peculiar to my model of receiver, or if it's common to all of them.  I 
welcome any reports of similar weirdness on other receivers, or a proper fix if anyone figures out what's really going on 
in this area.

Newer Denon receivers have a network port instead of a serial port.  The protocol is the same, so it should be fairly simple
to adapt this to the newer networked receivers.  I don't have one, so I'm not going there, but if anyone with a networked
receiver wants to give it a shot, please do so, and send me a pull request.  I won't be able to test it, but as long as you 
tell me it works, and as long as it doesn't break anything else, I'll accept it.

The web interface is basic and ugly.  It does what I need it to, but a fancier (and mobile-friendly) interface would be nice.  
That's on the to-do list.

The serial protocol supports two-way communication, both in the form of a query/response system, and an event-based system.  
The command line program supports the query/response model using the -response flag, which tells it to wait for a response 
after sending a command.  You can use the -cmd and -response flags to send explicit commands to the reciever that the program 
doesn't actively support, and read back responses from the receiver, for example:

denon -cmd="CVC?" -response 

will cause the receiver to respond with the current center channel volume.  The web interface uses this in a very basic way to
set the volume and surround parameter sliders appropriately when the page loads.  Event-based communication causes the receiver 
to send information over the serial port in response to changes at the receiver end (e.g. if someone turns the volume knob 
on the receiver, it will send a series of "MVnn" events over the serial port reflecting the new volume setting).  The program 
doesn't support event based communication, but it could probably be made to work in conjunction with another program that 
opens the serial port in read-only mode and processes events.  I'll leave that project for another day...

Other suggestions, bug reports, and pull requests are welcome.

