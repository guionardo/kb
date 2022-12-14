package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

type Sink struct {
	Index   int
	Name    string
	Current bool
}

var (
	sinks               []Sink
	maxIndex            int
	currentSinkIndex    int
	getSinkError        error
	inputIndexes        []int
	getInputIndexeError error
)

func get_pacmd_cmd(args ...string) (lines []string, err error) {
	log.Printf("pacmd %v", args)
	out, err := exec.Command("pacmd", args...).Output()
	if err != nil {
		return nil, fmt.Errorf("Get pacmd %v : %v", args, err)
	}
	raw_lines := strings.Split(string(out), "\n")
	lines = make([]string, 0, len(raw_lines))
	for _, line := range raw_lines {
		line = strings.TrimSpace(strings.ReplaceAll(line, "\t", ""))
		if line != "" {
			lines = append(lines, line)
		}
	}
	return
}

func get_pacmd_sinks(wg *sync.WaitGroup) {
	defer wg.Done()
	var lines []string
	if lines, getSinkError = get_pacmd_cmd("list-sinks"); getSinkError != nil {
		return
	}

	lastIndex := 0
	currentSinkIndex = -1
	currentIndex := false
	lastName := ""
	maxIndex = 0

	for _, line := range lines {
		line = strings.ReplaceAll(strings.TrimSpace(line), "\t", "")
		if strings.Contains(line, "index:") {
			if lastIndex, getSinkError = strconv.Atoi(strings.TrimSpace(strings.Split(line, ":")[1])); getSinkError != nil {
				return
			}
			currentIndex = strings.HasPrefix(line, "*")

			if lastIndex > maxIndex {
				maxIndex = lastIndex
			}
			continue
		}
		if strings.HasPrefix(line, "alsa.name =") {
			lastName = strings.ReplaceAll(strings.TrimSpace(strings.Split(line, "=")[1]), "\"", "")
			sinks = append(sinks, Sink{lastIndex, lastName, currentIndex})
			if currentIndex {
				currentSinkIndex = len(sinks) - 1
			}
			lastIndex = 0
			currentIndex = false

			lastName = ""
			continue
		}
	}
	log.Printf("Sinks: %v", sinks)
	return
}

func get_pacmd_sink_inputs(wg *sync.WaitGroup) {
	defer wg.Done()
	var lines []string
	if lines, getInputIndexeError = get_pacmd_cmd("list-sink-inputs"); getInputIndexeError != nil {
		return
	}

	inputIndexes = make([]int, 0)
	index := 0
	for _, line := range lines {
		line = strings.ReplaceAll(strings.TrimSpace(line), "\t", "")
		if strings.Contains(line, "index:") {
			if index, getInputIndexeError = strconv.Atoi(strings.TrimSpace(strings.Split(line, ":")[1])); getInputIndexeError != nil {
				return
			}
			inputIndexes = append(inputIndexes, index)
		}
	}
	return
}

func set_pacmd_sink(sinkIndex int) (err error) {
	for _, inputIndex := range inputIndexes {
		_, err := get_pacmd_cmd("move-sink-input", strconv.Itoa(inputIndex), strconv.Itoa(sinkIndex))
		if err != nil {
			return fmt.Errorf("Move sink input %v to %v : %v", inputIndex, sinkIndex, err)
		}
	}

	return
}

func notify_send(args ...string) error {
	out, err := exec.Command("notify-send", args...).Output()
	if err != nil {
		log.Printf("Notify-send %v : %v\n%s", args, err, string(out))
		return fmt.Errorf("notify-send %v : %v", args, err)
	}
	return nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go get_pacmd_sinks(&wg)
	go get_pacmd_sink_inputs(&wg)
	wg.Wait()
	if getSinkError != nil {
		log.Fatalf("get_pacmd_sinks %v", getSinkError)
	}
	if getInputIndexeError != nil {
		log.Fatalf("get_pacmd_sink_inputs %v", getInputIndexeError)
	}

	oldSink := sinks[currentSinkIndex]
	if currentSinkIndex == -1 {
		currentSinkIndex = 0
	}
	if currentSinkIndex < len(sinks)-1 {
		currentSinkIndex++
	} else {
		currentSinkIndex = 0
	}
	newSink := sinks[currentSinkIndex]
	_, err := get_pacmd_cmd("set-default-sink", strconv.Itoa(newSink.Index))
	if err != nil {
		log.Fatalf("set-default-sink %v", err)
	}
	log.Printf("Switching from %v to %v", oldSink, newSink)

	if err := set_pacmd_sink(newSink.Index); err != nil {
		log.Fatalf("set_pacmd_sink %v", err)
	}
	if err := notify_send("-i", "audio-volume-high", "Sound output switched to "+newSink.Name); err != nil {
		log.Fatalf("%v", err)
	}

}