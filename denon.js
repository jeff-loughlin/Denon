
function sendCommand(cmd)
{
    var req = new XMLHttpRequest();
    req.open("POST","./denon.php", true);
    req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    req.setRequestHeader("Content-Length", cmd.length); // POST request MUST have a Content-Length header (as per HTTP/1.1)
    req.onreadystatechange = function() {
        if (req.readyState == 4)
        {
            if (req.status == 200)
            {
//              alert(req.responseText);
            }
        }
    }
    req.send("cmd=" + cmd);
}

function powerOn()
{
    sendCommand('-poweron');
}

function powerOff()
{
    sendCommand('-poweroff');
}

function inputCable()
{
    sendCommand('-input=cable');
}

function inputNetwork()
{
    sendCommand('-input=network');
}

function inputDVD()
{
    sendCommand('-input=DVD');
}

function inputVCR()
{
    sendCommand('-input=VCR');
}

function inputXBoxOne()
{
    sendCommand('-input=XBoxOne');
}

function inputXBox360()
{
    sendCommand('-input=XBox360');
}

function surroundDirect()
{
    sendCommand('-direct');
}

function surroundPureDirect()
{
    sendCommand('-pureDirect');
}

function surroundStereo()
{
    sendCommand('-stereo');
}

function surround5ChStereo()
{
    sendCommand('-5ch');
}

function surroundStandard()
{
    sendCommand('-standard');
}

function surroundDolbyDigital()
{
    sendCommand('-dolby');
}

function surroundDTS()
{
    sendCommand('-dts');
}

function surroundArena()
{
    sendCommand('-arena');
}

function surroundJazzClub()
{
    sendCommand('-jazz');
}

function surroundMonoMovie()
{
    sendCommand('-mono');
}

function surroundMatrix()
{
    sendCommand('-matrix');
}

function surroundVideoGame()
{
    sendCommand('-game');
}

function surroundVirtual()
{
    sendCommand('-virtual');
}

function surroundPl2Music()
{
    sendCommand('-music -pl2');
}

function surroundPl2Movie()
{
    sendCommand('-movie -pl2');
}

function surroundNeo6Music()
{
    sendCommand('-music -neo6');
}

function surroundNeo6Movie()
{
    sendCommand('-movie -neo6');
}

function pl2WidthChanged()
{
    width = document.getElementById('pl2Width');
    sendCommand('-pl2width=' + width.value);
}

function pl2DimensionChanged()
{
    dimension = document.getElementById('pl2Dimension')
    sendCommand('-pl2dim=' + dimension.value);
}

function pl2PanModeChanged()
{
    mode = document.getElementById('pl2PanMode')
    if (mode.checked)
	sendCommand('-pl2pan=ON');
    else
	sendCommand('-pl2pan=OFF');
}

function neo6WidthChanged()
{
    width = document.getElementById('neo6Width');
    sendCommand('-neo6width=' + width.value);
}

function matrixDelayChanged()
{
    delay = document.getElementById('matrixDelay');
    sendCommand('-matrixdelay=' + delay.value);
}

function roomSizeChanged()
{
    size = document.getElementById('roomSize');
    var s = '';
    switch (size.value) {
	case '0': s = 'S'; break;
	case '1': s = 'MS'; break;
	case '2': s = 'M'; break;
	case '3': s = 'ML'; break;
	case '4': s = 'L'; break;
    }
    sendCommand('-roomsize=' + s);
}

function effectLevelChanged()
{
    level = document.getElementById('effectLevel');
    sendCommand('-effectlevel=' + level.value);
}

function surroundResetParams()
{
	sendCommand('-resetParams');
}

function volumeChanged()
{
	volume = document.getElementById('volumeControl');
	sendCommand('-volume=' + volume.value);
}
