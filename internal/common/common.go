package common

import (
	"fmt"
	ent "main/internal/entities"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	SensorData int
)

func ReturnData(status int, id int64, sensor_data string, message string) ent.ReturnData {
	data := ent.ReturnData{
		Status:     status,
		Id:         id,
		SensorData: sensor_data,
		TimeStamp:  time.Now().UnixMilli(),
		Message:    message,
	}
	return data
}

func NumGenerator(values string) int {
	ranges := Explode("-", values)
	min, _ := strconv.Atoi(ranges[0])
	max, _ := strconv.Atoi(ranges[1])

	rand.Seed(time.Now().UnixNano())
	data := rand.Intn(max-min+1) + min
	return data
}

func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}

func BGGenerator(index int) {
	IntervalIndex := fmt.Sprintf("SENSOR%d_INTERVAL", index)
	val, err := strconv.Atoi(os.Getenv(IntervalIndex))
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(time.Duration(val) * time.Second)
	RangeIndex := fmt.Sprintf("SENSOR%d_RANGE", index)
	for _ = range ticker.C {
		SensorData = NumGenerator(os.Getenv(RangeIndex))
		fmt.Printf("Sensor %d Data : %d \n", index, SensorData)
	}
}

func PostToAPI(uri string) int {
	//TODO HTTP CLIENT native
	return 1
}
