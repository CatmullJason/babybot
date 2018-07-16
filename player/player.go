package player

import (
	"os/exec"
	"time"
	"github.com/itchyny/volume-go"
	"github.com/CatmullJason/babybot/learning"
	"github.com/CatmullJason/babybot/detection"
	"github.com/CatmullJason/babybot/playlist"
)

var CurrentSong string
var isPlayingSong = false
var currentVolume = 0

func PlaySong() {
	if !isPlayingSong {
		nextSong := learning.GetBestSong()
		playlist.AddToPlayedSongs(nextSong)
		CurrentSong = nextSong
		cmd := "play " + nextSong

		c := exec.Command("bash", "-c", cmd)

		c.Start() // starts the specified command but does not wait for it to complete
		isPlayingSong = true

		// wait for the program to end in a goroutine
		go func() {
			c.Wait()
			isPlayingSong = false
			// logic to run once process finished. Send err in channel if necessary
		}()
	}
}


func IsSongPlaying() bool {
	return isPlayingSong
}

func VolumeUp(){
	currentVol, _ := volume.GetVolume()
	volume.SetVolume(currentVol + 1)
}

func VolumeDown(){
	currentVol, _ := volume.GetVolume()
	volume.SetVolume(currentVol - 1)
}

func VolumeDown5() {
	currentVol, _ := volume.GetVolume()
	volume.SetVolume(currentVol - 5)
}

func VolumeUp5() {
	currentVol, _ := volume.GetVolume()
	volume.SetVolume(currentVol + 5)
}

func VolumeZeroOut() {
	volume.SetVolume(1)
}

func SetVolume(vol int) {
	current, _ := volume.GetVolume()

	if vol > current {
		diff := vol - current

		for i := 0; i < diff; i++ {
			VolumeUp()
			time.Sleep(50 * time.Millisecond)
		}
	} else {
		diff := current - vol

		for i := 0; i < diff; i++ {
			VolumeDown()
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func AdjustVolumeTo(vol int) {
	current, _ := volume.GetVolume()

	if vol > current {
		diff := vol - current

		for i := 0; i < diff; i++ {
			VolumeUp()
			time.Sleep(300 * time.Millisecond)
		}
	} else {
		diff := current - vol

		for i := 0; i < diff; i++ {
			VolumeDown()
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func AdjustVolume(){
	if detection.GetCurrentAmplitude() < 0.085 {
		AdjustVolumeTo(35)
	}
	if (detection.GetCurrentAmplitude() > 0.085) && (detection.GetCurrentAmplitude() < 0.1) {
		AdjustVolumeTo(50)
	}
	if detection.GetCurrentAmplitude() > 0.1 {
		AdjustVolumeTo(65)
	}
}



