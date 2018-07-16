package detection

import (
	"fmt"
	"os/exec"
	"strings"
	"strconv"
	"time"
)

/*
	Install sox for recording audio
		sudo apt install sox
	Record short audio file to get amplitude
		arecord -qd 1 volt && sox volt -n stat &> volt.d && sed '4q;d' volt.d
*/

var CurrentAmplitude = 0.0
var CurrentBaseline = 0.0
var LastRecordedAmplitude = 0.0

func GetCurrentBaseline() float64 {
	return CurrentBaseline
}

func GetCurrentAmplitude() float64 {
	return CurrentAmplitude
}

func GetBaseline() float64 {
	maxBaseline := 0.0

	for i := 0; i < 5; i++ {
		amp := RunDetection()
		if maxBaseline < amp {
			maxBaseline = amp
		}
		time.Sleep(200 * time.Millisecond)
	}

	CurrentBaseline = maxBaseline

	fmt.Println(fmt.Sprintf("Baseline is %v", maxBaseline))

	return maxBaseline
}

func RunDetection() float64 {
	cmd := "arecord -qd 1 volt && sox volt -n stat &> volt.d && sed '4q;d' volt.d"

	out, err := exec.Command("bash","-c",cmd).Output()
	if err != nil {
		fmt.Println(err)
	}

	output := strings.Split(fmt.Sprintf("%s",out), " ")
	amplitude := output[len(output)-1]

	amplitude = strings.Replace(amplitude, "\n", "", 1)

	finalAmplitude, err := strconv.ParseFloat(amplitude, 64)

	if err != nil {
		fmt.Println(err)
	}

	CurrentAmplitude = finalAmplitude

	return finalAmplitude
}

func IsBabyCry(amp float64) bool {
	fmt.Println(amp)

	if amp > 0.023 {
		LastRecordedAmplitude = amp
		return true
	}
	return false
}