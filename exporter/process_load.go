package exporter

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// ProcInfo contains information about a process
type ProcInfo struct {
	pid     string
	ppid    string
	memP    float64
	cpuP    float64
	cpu     float64
	command string
	user    string
	start   string
}

// GetProcessStats returns current process information
func GetProcessStats() []*ProcInfo {
	// out, err := exec.Command("sh", "-c", "ps -eo pid,ppid,%mem,%cpu,cpu,comm,user,lstart").CombinedOutput()
	out, err := exec.Command("sh", "-c", "ps -eo pid,ppid,pmem,pcpu,cpu,user,etime,command").CombinedOutput()
	if err != nil {
		log.Fatal("CMD Error: ", err)
	}
	trimmed := strings.Trim(string(out), "\n")
	lines := strings.Split(string(trimmed), "\n")
	header := parseRow(lines[0])
	procInfos := make([]*ProcInfo, len(lines)-1, len(lines)-1)
	for i := 1; i < len(lines); i++ {
		procInfos[i-1] = parseDataPointFromRow(len(header), lines[i])
	}
	return procInfos
}

func parseDataPointFromRow(headerLength int, row string) *ProcInfo {
	fields := parseRow(row)
	memP, _ := strconv.ParseFloat(fields[2], 64)
	cpuP, _ := strconv.ParseFloat(fields[3], 64)
	cpu, _ := strconv.ParseFloat(fields[4], 64)
	return &ProcInfo{
		pid:     fields[0],
		ppid:    fields[1],
		memP:    memP,
		cpuP:    cpuP,
		cpu:     cpu,
		command: fields[7],
		user:    fields[5],
		start:   fields[6],
	}
}

func parseRow(row string) []string {
	return strings.Fields(row)
}
