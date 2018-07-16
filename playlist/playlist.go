package playlist

var Playlist = []string {"Loops/1.wav", "Loops/2.wav"}//, "Loops/3.wav", "Loops/4.wav", "Loops/5.wav", "Loops/5.wav"}
var PlayedSongs = make(map[string]int)

func InitializeSongsToMap() {
	for _, val := range Playlist {
		PlayedSongs[val] = 0
	}
}

func GetRandomSong() string {
	return getSongBasedOnPlayedLeast()
}

func AddToPlayedSongs(song string) {
	timesPlayed := PlayedSongs[song]
	PlayedSongs[song] = timesPlayed + 1
}

func GetTimesPlayed(song string) int {
	return PlayedSongs[song]
}

func getSongBasedOnPlayedLeast() string {
	leastPlayed := ""
	lowestNumber := 1000

	for k, v := range PlayedSongs {
		if v < lowestNumber {
			leastPlayed = k
			lowestNumber = v
		}
	}

	return leastPlayed
}