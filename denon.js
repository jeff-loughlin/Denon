
function sendCommand(cmd)
{
    var req = new XMLHttpRequest();
    req.open("POST","./denon.php", true);
    req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    req.setRequestHeader("Content-Length", cmd.length);
    req.onreadystatechange = function() {
        if (req.readyState == 4)
        {
            if (req.status == 200)
            {
//                alert(req.responseText);
            }
        }
    }
    req.send("cmd=" + cmd);
}

function init()
{
    initVolumeSlider();
}

function initVolumeSlider()
{
    var cmd = '-cmd="MV?" -response';
    var req = new XMLHttpRequest();
    req.open("POST","./denon.php", true);
    req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    req.setRequestHeader("Content-Length", cmd.length);
    req.onreadystatechange = function() {
        if (req.readyState == 4)
        {
            if (req.status == 200)
            {
		setVolumeSliderFromResponseText(req.responseText);
		initPl2WidthSlider();
            }
        }
    }
    req.send("cmd=" + cmd);
}

function initPl2WidthSlider()
{
    var cmd = "-cmd='PSCEN ?' -response";
    var req = new XMLHttpRequest();
    req.open("POST","./denon.php", true);
    req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    req.setRequestHeader("Content-Length", cmd.length);
    req.onreadystatechange = function() {
        if (req.readyState == 4)
        {
            if (req.status == 200)
            {
		setPl2WidthSliderFromResponseText(req.responseText);
		initPl2DimensionSlider();
            }
        }
    }
    req.send("cmd=" + cmd);
}

function initPl2DimensionSlider()
{
    var cmd = "-cmd='PSDIM ?' -response";
    var req = new XMLHttpRequest();
    req.open("POST","./denon.php", true);
    req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    req.setRequestHeader("Content-Length", cmd.length);
    req.onreadystatechange = function() {
        if (req.readyState == 4)
        {
            if (req.status == 200)
            {
		setPl2DimensionSliderFromResponseText(req.responseText);
		initNeo6WidthSlider();
            }
        }
    }
    req.send("cmd=" + cmd);
}

function initNeo6WidthSlider()
{
    var cmd = "-cmd='PSCEI ?' -response";
    var req = new XMLHttpRequest();
    req.open("POST","./denon.php", true);
    req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    req.setRequestHeader("Content-Length", cmd.length);
    req.onreadystatechange = function() {
        if (req.readyState == 4)
        {
            if (req.status == 200)
            {
		setNeo6WidthSliderFromResponseText(req.responseText);
		initMatrixDelaySlider();
            }
        }
    }
    req.send("cmd=" + cmd);
}

function initMatrixDelaySlider()
{
    var cmd = "-cmd='PSDEL ?' -response";
    var req = new XMLHttpRequest();
    req.open("POST","./denon.php", true);
    req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    req.setRequestHeader("Content-Length", cmd.length);
    req.onreadystatechange = function() {
        if (req.readyState == 4)
        {
            if (req.status == 200)
            {
		setMatrixDelaySliderFromResponseText(req.responseText);
		initRoomSizeSlider();
            }
        }
    }
    req.send("cmd=" + cmd);
}

function initRoomSizeSlider()
{
    var cmd = "-cmd='PSRSZ ?' -response";
    var req = new XMLHttpRequest();
    req.open("POST","./denon.php", true);
    req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    req.setRequestHeader("Content-Length", cmd.length);
    req.onreadystatechange = function() {
        if (req.readyState == 4)
        {
            if (req.status == 200)
            {
		setRoomSizeSliderFromResponseText(req.responseText);
		initEffectLevelSlider();
            }
        }
    }
    req.send("cmd=" + cmd);
}

function initEffectLevelSlider()
{
    var cmd = "-cmd='PSEFF ?' -response";
    var req = new XMLHttpRequest();
    req.open("POST","./denon.php", true);
    req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    req.setRequestHeader("Content-Length", cmd.length);
    req.onreadystatechange = function() {
        if (req.readyState == 4)
        {
            if (req.status == 200)
            {
		setEffectLevelSliderFromResponseText(req.responseText);
		initPanCheckbox();
            }
        }
    }
    req.send("cmd=" + cmd);
}

function initPanCheckbox()
{
    var cmd = "-cmd='PSPAN ?' -response";
    var req = new XMLHttpRequest();
    req.open("POST","./denon.php", true);
    req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    req.setRequestHeader("Content-Length", cmd.length);
    req.onreadystatechange = function() {
        if (req.readyState == 4)
        {
            if (req.status == 200)
            {
		setPanCheckboxFromResponseText(req.responseText);
            }
        }
    }
    req.send("cmd=" + cmd);
}

function setVolumeSliderFromResponseText(responseText)
{
    var rawVolume = responseText.substr(2);

    // Receiver returns 99 to mean minimum volume.  I don't know why.  Change that to 0 so it becomes -80 after
    // the adjusted volume calculation below
    if (rawVolume == 99)
    {
	rawVolume = 0;
    }
    // Receiver returns a three digit number for half steps (NN.n).  If we got a three digit number, just 
    // divide it by 10 and drop the decimal.  We don't need to be that precise.
    if (rawVolume >= 100)
    {
	rawVolume = rawVolume / 10;
    }

    // Receiver returns volume as a number from 0-98 (with 0 being the lowest and 98 being the highest, but for
    // some reason sets it to 99 if the volume is at the minimum).  We want to convert it to -80 to +20 for the 
    // slider range (which represents volume relative to 0db reference)
    var adjustedVolume = rawVolume*1 - 80;
    document.getElementById('volumeControl').value = adjustedVolume;
}

function setPl2WidthSliderFromResponseText(responseText)
{
    var val = responseText.substr(6);
    document.getElementById('pl2Width').value = val * 1;
}

function setPl2DimensionSliderFromResponseText(responseText)
{
    var val = responseText.substr(6);
    document.getElementById('pl2Dimension').value = val * 1;
}

function setNeo6WidthSliderFromResponseText(responseText)
{
    var val = responseText.substr(6);
    document.getElementById('neo6Width').value = val * 1;
}

function setMatrixDelaySliderFromResponseText(responseText)
{
    var val = responseText.substr(6);
    document.getElementById('matrixDelay').value = val * 1;
}

function setRoomSizeSliderFromResponseText(responseText)
{
    var val = responseText.substr(6,2);
    switch (val)
    {
	case 'N\n': n = 0; break;
	case 'S\n': n = 1; break;
	case 'MS' : n = 2; break;
	case 'M\n': n = 3; break;
	case 'ML' : n = 4; break;
	case 'L\n': n = 5; break;
    }
    document.getElementById('roomSize').value = n;
}

function setEffectLevelSliderFromResponseText(responseText)
{
    var val = responseText.substr(6);
    document.getElementById('effectLevel').value = val * 1;
}

function setPanCheckboxFromResponseText(responseText)
{
    if (responseText.startsWith("PSPAN ON"))
	document.getElementById('pl2PanMode').checked = true;
    else
	document.getElementById('pl2PanMode').checked = false;	
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
	case '0': s = 'N'; break;
	case '1': s = 'S'; break;
	case '2': s = 'MS'; break;
	case '3': s = 'M'; break;
	case '4': s = 'ML'; break;
	case '5': s = 'L'; break;
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
