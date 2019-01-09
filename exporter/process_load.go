package exporter

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type dataPoint struct {
	pid     string
	ppid    string
	memP    float64
	cpuP    float64
	cpu     int
	command string
	user    string
	start   string
}

func GetProcessStats() []*dataPoint {
	out, err := exec.Command("sh", "-c", "ps -eo pid,ppid,%mem,%cpu,cpu,comm,user,lstart").CombinedOutput()
	if err != nil {
		log.Fatal("CMD Error: ", err)
	}
	trimmed := strings.Trim(string(out), "\n")
	lines := strings.Split(string(trimmed), "\n")
	header := parseRow(lines[0])
	dataPoints := make([]*dataPoint, len(lines)-1, len(lines)-1)
	for i := 1; i < len(lines); i++ {
		dataPoints[i-1] = parseDataPointFromRow(len(header), lines[i])
	}
	return dataPoints
}

func parseDataPointFromRow(headerLength int, row string) *dataPoint {
	fields := parseRow(row)
	timeString := strings.Join(fields[headerLength-1:], " ")
	timestamp, _ := time.Parse(time.ANSIC, timeString)
	memP, _ := strconv.ParseFloat(fields[2], 64)
	cpuP, _ := strconv.ParseFloat(fields[3], 64)
	cpu, _ := strconv.Atoi(fields[4])
	return &dataPoint{
		pid:     fields[0],
		ppid:    fields[1],
		memP:    memP,
		cpuP:    cpuP,
		cpu:     cpu,
		command: fields[5],
		user:    fields[6],
		start:   strconv.FormatInt(timestamp.Unix(), 10),
	}
}

func parseRow(row string) []string {
	return strings.Fields(row)
}

//func main() {
//dataPoints := GetProcessStats()
//fmt.Printf("DataPoints: %s", len(dataPoints))
//}
