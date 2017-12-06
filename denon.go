package main

import (
        "log"
        "github.com/tarm/serial"
	"flag"
	"strconv"
	"strings"
	"fmt"
	"time"
	"bufio"
)

const WAIT_TIME time.Duration = 500

func sendCommand(port string, cmd string) {

//	fmt.Println("Sending command: " + cmd)

        c := &serial.Config{Name: port, Baud: 9600}
        s, err := serial.OpenPort(c)
        if err != nil {
                log.Fatal(err)
        }
        
        n, err := s.Write([]byte(cmd + "\r"))
        if err != nil {
                log.Fatal(err)
        }
	n++
}

func readResponse(port string) string {
        c := &serial.Config{Name: port, Baud: 9600}
	c.ReadTimeout = time.Millisecond * 5000
        s, err := serial.OpenPort(c)
        if err != nil {
                log.Fatal(err)
        }



	r := bufio.NewReader(s)

	// reads until delimiter is reached
	data, err := r.ReadBytes('\x0d')
	if err != nil {
	    // stops execution
	    log.Fatal(err)
	}
//	fmt.Println("Response: " + string(data[:]))
	return string(data[:])
}



func main() {
	powerOn := flag.Bool("poweron", false, "Power On")
	powerOff := flag.Bool("poweroff", false, "Power Off")

	volume := flag.String("volume", "", "Set volume (-80 to 20")
	volPlus := flag.Bool("vol+", false, "Increase Volume")
	volMinus := flag.Bool("vol-", false, "Decrease volume")
	input := flag.String("input", "", "Select Input (Cable (SAT/CBL), Network (CD), DVD, VCR, XBoxOne (HDP), XBox360 (DVR) )")
	surroundMode :=flag.String("surround", "", "Set surround Mode (Direct, PureDirect, Stereo, Standard, DolbyDigital, DTS, Arena, JazzClub, Matrix, MonoMovie, Game, Virtual)")
	restoreMode := flag.String("restore", "", "Set restore mode (1, 2, 3 or OFF)")
	centerVol := flag.String("center", "", "Set Center channel volume (-12 - 12db, or 38 - 62 with 50=0, or use MAX/MIN/MID)")
	swVol := flag.String("sw", "", "Set Subwoofer channel volume (-12 - 12db, or 38-62 with 50=0, or use MAX/MIN/MID)")
	cmd := flag.String("cmd", "", "Send an explicit command to the receiver and (with the -response flag) read back a response")

	fiveChannel := flag.Bool("5ch", false, "Set 5-channel stereo mode")
	direct := flag.Bool("direct", false, "Set Direct mode (output channels same as input channels)")
	pure :=	flag.Bool("pure", false, "Set Pure Direct mode (output channels same as input channels, no EQ or other processing)")
	stereo := flag.Bool("stereo", false, "Set Stereo mode")
	standard := flag.Bool("standard", false, "Set Standard surround mode (toggle between Dolby PL2 and DTS NEO:6)")
	dolby := flag.Bool("dolby", false, "Set Dolby Digital mode")
	dts := flag.Bool("dts", false, "Set DTS Digital mode")
	arena := flag.Bool("arena", false, "Set Rock Arena surround mode")
	jazz := flag.Bool("jazz", false, "Set Jazz Club surround mode")
	mono := flag.Bool("mono", false, "Set Mono Movie surround mode")
	matrix := flag.Bool("matrix", false, "Set Matrix surround mode")
	game := flag.Bool("game", false, "Set Video Game surround mode")
	virtual := flag.Bool("virtual", false, "Set Virtual surround mode")
	pl2 := flag.Bool("pl2", false, "Set Dolby Pro-Logic II mode (requires -music or -movie flag to set sub-mode)")
	neo6 := flag.Bool("neo6", false, "Set DTS NEO:6 mode (requires -music or -movie flag to set sub-mode)")
	movie := flag.Bool("movie", false, "Set PL2/NEO:6 Cinema mode (requires -pl2 or -neo6 flag to set mode)")
	music := flag.Bool("music", false, "Set PL2/NEO:6 Music mode (requires -pl2 or -neo6 flag to set mode)")
	getResponse := flag.Bool("response", false, "Read response from Denon AVR")
	stadium := flag.Bool("stadium", false, "Set Stadium surround parameters (-arena -roomSize=L -effectLevel=15)")
	resetParams := flag.Bool("resetParams", false, "Set default surround parameters (undo all custom surround/PL2/NEO6 parameters)")
	neo6Width := flag.String("neo6width", "", "Set DTS NEO6:Music mode center image width (0-7, default = 3); only valid in NEO6:Music surround mode")
	pl2Pan := flag.String("pl2pan", "", "Set Dolby PL2 PAN mode (ON/OFF, default = OFF); only valid in PL2 Music surround mode")
	pl2Width := flag.String("pl2width", "", "Set Dolby PL2 center channel width (0-7, default = 3); only valid in PL2 Music surround mode")
	pl2Dim := flag.String("pl2dim", "", "Set Dolby PL2 Dimension (0-7, default = 3); only valid in PL2 Music surround Mode")
	matrixDelay := flag.String("matrixdelay", "", "Set MATRIX delay milliseconds (0-300, default = 30); only valid in MATRIX surround mode")
	effectLevel := flag.String("effectlevel", "", "Set surround effect level (0-15); only valid in ARENA or JAZZ CLUB surround modes")
	roomSize := flag.String("roomsize", "", "Set surround room size (S,MS,M,ML,L, default = M); only valid in ARENA or JAZZ CLUB surround modes")
	
	flag.Parse();

	var port string = flag.Arg(0)
	if port == "" {
		port = "/dev/ttyUSB0"
	}

	if (*cmd != "") {
		sendCommand(port, *cmd)
		time.Sleep(500 * time.Millisecond)
	}

	if (*getResponse) {
		response := readResponse(port)
		fmt.Println(response)
		return
	}

	// Rename some inputs
	if (strings.ToUpper(*input) == "CABLE") {
		*input = "SAT/CBL"
	}
	if (strings.ToUpper(*input) == "NETWORK") {
		*input = "CD"
	}
	if (strings.ToUpper(*input) == "XBOXONE") {
		*input = "HDP"
	}
	if (strings.ToUpper(*input) == "XBOX360") {
		*input = "DVR"
	}
	if (strings.ToUpper(*input) == "AUX") {
		*input = "V.AUX"
	}

	// Rename some surround modes
	if (strings.ToUpper(*surroundMode) == "PUREDIRECT") {
		*surroundMode = "PURE DIRECT"
	}
	if (strings.ToUpper(*surroundMode) == "DOLBYDIGITAL") {
		*surroundMode = "DOLBY DIGITAL"
	}
	if (strings.ToUpper(*surroundMode) == "JAZZCLUB") {
		*surroundMode = "JAZZ CLUB"
	}
	if (strings.ToUpper(*surroundMode) == "ARENA") {
		*surroundMode = "ROCK ARENA"
	}
	if (strings.ToUpper(*surroundMode) == "MONOMOVIE") {
		*surroundMode = "MONO MOVIE"
	}
	if (strings.ToUpper(*surroundMode) == "5CH") {
		*surroundMode = "5CH STEREO"
	}


	// Enforce a few rules (pl2 and neo6 flags require -movie or -music flag)
	if (*pl2 && !(*music || *movie)) {
		fmt.Println("-pl2 flag requires -music or -movie flag")
		return
	}

	if (*neo6 && !(*music || *movie)) {
		fmt.Println("-neo6 flag requires -music or -movie flag")
		return
	}

	if (*movie && !(*pl2 || *neo6)) {
		fmt.Println("-movie flag requires -pl2 or -neo6 flag")
		return
	}

	if (*music && !(*pl2 || *neo6)) {
		fmt.Println("-music flag requires -pl2 or -neo6 flag")
		return
	}



	if (*powerOn) {
		sendCommand(port, "PWON")
	}
	
	if (*powerOff) {
		sendCommand(port, "PWSTANDBY")
	}

	if (*volume != "") {
		vol,err := strconv.Atoi(*volume)
		if (err != nil) {
			fmt.Println("Error: Volume must be between -80 and +20")
			return
		}
		if (vol >= -80 && vol <= 20) {
			sendCommand(port, "MV" + strconv.Itoa(vol + 80))
		}
	}

	if (*surroundMode != "") {
		sendCommand(port, "MS" + strings.ToUpper(*surroundMode))
	}

	if (*input != "") {
		sendCommand(port, "SI" + strings.ToUpper(*input))
	}

	if (*restoreMode != "") {
		if (*restoreMode == "OFF") {
			sendCommand(port, "PSRSTR OFF")
		} else {
			sendCommand(port, "PSRSTR MODE" + *restoreMode)
		}
	}

	if (*centerVol != "") {
		if (*centerVol == "MAX") {
			sendCommand(port, "CVC 62")
		} else if (*centerVol == "MIN") {
			sendCommand(port, "CVC 38")
		} else if (*centerVol == "MID") {
			sendCommand(port, "CVC 50")
		} else if n,_ := strconv.Atoi(*centerVol); n >= 38 && n <= 62 {
			sendCommand(port, "CVC " + *centerVol)
		} else if n,_ := strconv.Atoi(*centerVol); n >= -12 && n <= 12 {
			sendCommand(port, "CVC " + strconv.Itoa(50 + n))
		} else {
			fmt.Println("Center Volume out of range")
		}
	}
	if (*swVol != "") {
		if (*swVol == "MAX") {
			sendCommand(port, "CVSW 62")
		} else if (*swVol == "MIN") {
			sendCommand(port, "CVSW 38")
		} else if (*swVol == "MID") {
			sendCommand(port, "CVSW 50")
		} else if n,_ := strconv.Atoi(*swVol); n >= 38 && n <= 62 {
			sendCommand(port, "CVSW " + *swVol)
		} else if n,_ := strconv.Atoi(*swVol); n >= -12 && n <= 12 {
			sendCommand(port, "CVSW " + strconv.Itoa(50 + n))
		} else {
			fmt.Println("Subwoofer Volume out of range")
		}
	}


	
	if (*fiveChannel) {
		sendCommand(port, "MS5CH STEREO")
	}
	if (*direct) {
		sendCommand(port, "MSDIRECT")
	}
	if (*pure) {
		sendCommand(port, "MSPURE DIRECT")
	}
	if (*stereo) {
		sendCommand(port, "MSSTEREO")
	}
	if (*standard) {
		sendCommand(port, "MSSTANDARD")
	}
	if (*dolby) {
		sendCommand(port, "MSDOLBY DIGITAL")
	}
	if (*dts) {
		sendCommand(port, "MSDTS DIGITAL")
	}
	if (*arena) {
		sendCommand(port, "MSROCK ARENA")
	}
	if (*jazz) {
		sendCommand(port, "MSJAZZ CLUB")
	}
	if (*mono) {
		sendCommand(port, "MSMONO MOVIE")
	}
	if (*matrix) {
		sendCommand(port, "MSMATRIX")
	}
	if (*game) {
		sendCommand(port, "MSVIDEO GAME")
	}
	if (*virtual) {
		sendCommand(port, "MSVIRTUAL")
	}


	// Custom surround parameters
	if (*stadium) {
		sendCommand(port, "MSROCK ARENA")
		wait(WAIT_TIME)
		sendCommand(port, "PSEFF 15") // Effect level
		wait(WAIT_TIME)
		sendCommand(port, "PSRSZ L") // Room Size (S,MS,M,ML,L), M=default
	}

	if (*resetParams) {
		sendCommand(port, "PSDEFAULT")
	}

	if (*neo6Width != "") {
		sendCommand(port, "PSCEI " + *neo6Width)
	}

	if (*pl2Width != "") {
		sendCommand(port, "PSCEN " + *pl2Width)
	}

	if (*pl2Dim != "") {
		sendCommand(port, "PSDIM " + *pl2Dim)
	}

	if (*pl2Pan != "") {
		sendCommand(port, "PSPAN " + *pl2Pan)
	}

	if (*matrixDelay != "") {
		sendCommand(port, "PSDEL " + *matrixDelay) // Or possibly MSDELAY - unclear from Serial protocol documentation - both are documented but not differentiated
	}

	if (*effectLevel != "") {
		sendCommand(port, "PSEFF " + *effectLevel)
	}

	if (*roomSize != "") {
		sendCommand(port, "PSRSZ " + *roomSize)
	}

	if (*movie && *pl2) {
		setMovieMode(port)
		setPl2Mode(port)
		setMovieMode(port)
	}
	if (*music && *pl2) {
		setPl2Mode(port)
		setMusicMode(port)
		setPl2Mode(port)
	}
	if (*movie && *neo6) {
		setMovieMode(port)
		setNeo6Mode(port)
		setMovieMode(port)
	}
	if (*music && *neo6) {
		setMusicMode(port)
		setNeo6Mode(port)
		setMusicMode(port)
	}


	if (*volPlus) {
		sendCommand(port, "MVUP")
	}
	if (*volMinus) {
		sendCommand(port, "NVDOWN")
	}

}

func wait(milliseconds time.Duration) {
	time.Sleep(time.Millisecond * milliseconds)
}

func setMovieMode(port string) {
	// Query the device for the current mode.  If already in Movie mode, do nothing.
	wait(WAIT_TIME)
	sendCommand(port, "PSMODE: ?")
	wait(WAIT_TIME)
	var response = readResponse(port)
	if (! strings.Contains(response, "DTS NEO:6 C") && ! strings.Contains(response, "DOLBY PL2 C") && ! strings.Contains(response, "PSMODE:CINEMA")) {
		sendCommand(port, "PSMODE:CINEMA")
		wait(WAIT_TIME)
	}
}

func setMusicMode(port string) {
	// Query the device for the current mode.  If already in Music mode, do nothing.
	wait(WAIT_TIME)
	sendCommand(port, "PSMODE: ?")
	wait(WAIT_TIME)
	var response = readResponse(port)
	if (! strings.Contains(response, "DTS NEO:6 M") && ! strings.Contains(response, "DOLBY PL2 M") && ! strings.Contains(response, "MODE:MUSIC")) {
		sendCommand(port, "PSMODE:MUSIC")
		wait(WAIT_TIME)
	}
}

func setPl2Mode(port string) {
	// Query device for the current surround mode
	wait(WAIT_TIME)
	sendCommand(port, "MS?")
	wait(WAIT_TIME)
	var response string = readResponse(port)
	wait(WAIT_TIME)
	if (! strings.Contains(response, "PL2")) {
		// If already PL2 mode, do nothing, otherwise toggle standard mode to switch to PL2
		sendCommand(port, "MSSTANDARD")
		wait(WAIT_TIME)
	}		
}

func setNeo6Mode(port string) {
	// Query device for PL2/NEO6 mode
	wait(WAIT_TIME)
	sendCommand(port, "MS?")
	wait(WAIT_TIME)
	var response string = readResponse(port)
	wait(WAIT_TIME)
	if (! strings.Contains(response, "DTS NEO:6")) {
		// If already DTS NEO:6 mode, do nothing, otherwise toggle standard mode to switch to DTS NEO:6
		sendCommand(port, "MSSTANDARD")
		wait(WAIT_TIME)
	}
}
