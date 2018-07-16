package main

import (
	"fmt"
	"github.com/CatmullJason/babybot/learning"
	"github.com/CatmullJason/babybot/detection"
	"github.com/CatmullJason/babybot/player"
	"github.com/CatmullJason/babybot/playlist"
)

var songAmplitudes []float64
var songTotalAmplitude = 0.0
var songMinAmplitude = 20.0
var songMaxAmplitude = 0.0
var songAverageAmplitude = 0.0

func main() {

	fmt.Println("BabyBot has started...")

	learning.InitializeRandomForest()
	learning.ReinforceLearning()
	playlist.InitializeSongsToMap()

	detection.GetBaseline()

	for {
		Run()
	}

}

func Run(){
	value := detection.RunDetection()

	if !player.IsSongPlaying() {
		if len(songAmplitudes) != 0 {
			for _, val := range songAmplitudes {
				if val < songMinAmplitude {
					songMinAmplitude = val
				}
				if val > songMaxAmplitude {
					songMaxAmplitude = val
				}
				songTotalAmplitude += val
			}

			songAverageAmplitude = songTotalAmplitude / float64(len(songAmplitudes))

			songEndingAmplitude := detection.RunDetection()

			sd := learning.SongData{MinAmplitude:songMinAmplitude, MaxAmplitude:songMaxAmplitude, AverageAmplitude: songAverageAmplitude, EndingAmplitude:songEndingAmplitude, Label:player.CurrentSong}
			learning.InsertRow(sd)
			songTotalAmplitude = 0.0
			songMinAmplitude = 20.0
			songMaxAmplitude = 0.0
			songAverageAmplitude = 0.0
			songAmplitudes = []float64{}
		}
		isCry := detection.IsBabyCry(value)
		if isCry {
			player.VolumeZeroOut()
			player.PlaySong()
			player.SetVolume(35)
		}
	} else {
		go player.AdjustVolume()
		songAmplitudes = append(songAmplitudes, value)
	}
}
