package learning

import (
	"fmt"
	"os"
	"encoding/csv"
	"github.com/sjwhitworth/golearn/base"
	"strconv"
	"github.com/sjwhitworth/golearn/ensemble"
	"math/rand"
	"bytes"
	"io"
	"bufio"
	"github.com/CatmullJason/babybot/playlist"
	"time"
)

var tree base.Classifier
var isLearning = false
var rawData *base.DenseInstances
var err error
var count = 40

func ReinforceLearning() {
	rand.Seed(44111342)

	rawData, err = base.ParseCSVToInstances("learning/data/song_dataset.csv", false)
	if err != nil {
		panic(err)
	}

	// Print a pleasant summary of your data.
	fmt.Println(rawData)

	//Initialises a new KNN classifier
	isLearning = true
	tree.Fit(rawData)
	isLearning = false
}

func InitializeRandomForest() {
	tree = ensemble.NewRandomForest(150, 3)
}

func GetBestSong() string {

	tree := tree

	if UsePrediction() {
		rawData, err := base.ParseCSVToTemplatedInstances("learning/data/prediction.csv", false, rawData)
		//Calculates the Euclidean distance and returns the most popular label
		predictions, err := tree.Predict(rawData)
		if err != nil {
			panic(err)
		}
		bestSong := predictions.RowString(0)
		return bestSong
	}

	return playlist.GetRandomSong()
}

type SongData struct {
	MinAmplitude     float64
	MaxAmplitude     float64
	AverageAmplitude float64
	EndingAmplitude  float64
	Label            string
}

type CsvFile struct {
	FilePath string
}

func InsertRow(data SongData) {

	f, err := os.OpenFile("learning/data/song_dataset.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	w := csv.NewWriter(f)
	w.Write(data.toArray())
	w.Flush()

	ReinforceLearning()
}

func lineCounter() (int, error) {
	file, _ := os.Open("learning/data/song_dataset.csv")
	reader := bufio.NewReader(file)

	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := reader.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func UsePrediction() bool {
	lines, _ := lineCounter()

	if lines > count {
		count = 2
	}

	rNum := random(0, count)
	rNum2 := random(0,count)

	if rNum == rNum2 {
		return true
	}

	if count > 2 {
		count--
	}

	return false
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func (sd SongData) toArray() []string {
	var sdArray []string

	sdArray = append(sdArray, strconv.FormatFloat(sd.MinAmplitude, 'f', -1, 64))
	sdArray = append(sdArray, strconv.FormatFloat(sd.MaxAmplitude, 'f', -1, 64))
	sdArray = append(sdArray, strconv.FormatFloat(sd.AverageAmplitude, 'f', -1, 64))
	sdArray = append(sdArray, strconv.FormatFloat(sd.EndingAmplitude, 'f', -1, 64))
	sdArray = append(sdArray, sd.Label)

	return sdArray
}
