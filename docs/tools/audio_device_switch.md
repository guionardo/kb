---
title: Atalho para trocar dispositivo de áudio no Linux
tags: [linux,audio,automation,bash,golang]
---

Entre reuniões virtuais usando headset, escutar música ou mesmo só assistir uma aula, eu prefiro que o headset seja usado só quando é necessário um isolamento.

Para as outras opções, as caixas de som ambiente são bem mais confortáveis.

No Ubuntu, é meio chatinho ficar chaveando entre os dispositivos de saída de áudio. Aqui na minha máquina, são uma meia dúzia de cliques de mouse.

Então este script aqui em baixo deixou as coisas bem mais interessantes. Fiz algumas alterações, mas a ideia eu peguei desse [link](https://askubuntu.com/questions/156895/how-to-switch-sound-output-with-key-shortcut).


## audio-device-switch.sh

```bash
#!/bin/bash

declare -i sinks_count=`pacmd list-sinks | grep -c index:[[:space:]][[:digit:]]`
declare -i active_sink_index=`pacmd list-sinks | sed -n -e 's/\*[[:space:]]index:[[:space:]]\([[:digit:]]\)/\1/p'`
declare -i major_sink_index=$sinks_count
declare -i next_sink_index=1

if [ $active_sink_index -ne $major_sink_index ] ; then
    next_sink_index=active_sink_index+1
fi

declare -i new_sink_index=$next_sink_index

#change the default sink
pacmd "set-default-sink ${new_sink_index}"


# move all inputs to the new sink: $new_sink_index"

for app in $(pacmd list-sink-inputs | sed -n -e 's/index:[[:space:]]\([[:digit:]]\)/\1/p');
do
    pacmd "move-sink-input $app $new_sink_index"
done

# echo 'display notification'
declare -i ndx=1
pacmd list-sinks | sed -n -e 's/device.description[[:space:]]=[[:space:]]"\(.*\)"/\1/p' | while read line;
do
    if [ $new_sink_index -eq $ndx ] ; then
        notify-send -i audio-volume-high "Sound output switched to" "$line"
        exit
    fi
    ndx=ndx+1
done
```

Salve o arquivo em uma pasta segura, por exemplo /usr/local/bin/audio-device-switch.sh e dê as permissões de execução:

```bash
sudo chmod 755 /usr/local/bin/audio-device-switch.sh
```

## Versão go

```golang
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
```

## Configurações de atalho de teclado

Acesse o menu de setup da sua distribuição, e informe o nome do atalho (Audio Device Switch, no meu caso), o local onde o script está salvo, e o atalho que você deseja usar (Eu escolhi Win+F12).

Prontinho, agora é só aproveitar.