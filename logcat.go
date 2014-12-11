package main

import (
	"github.com/daviddengcn/go-colortext"
	"fmt"
	"os/exec"
	"bufio"
	"log"
	"os"
	"strings"
	"io"
)

var colorMap map[rune]ct.Color
var LOG_SLOTS = "VDIWE"
var COLOR_VALUES = map[string]ct.Color {
	"NONE"   :ct.None,
	"BLACK"  :ct.Black,
	"RED"    :ct.Red,
	"GREEN"  :ct.Green,
	"YELLOW" :ct.Yellow,
	"BLUE"   :ct.Blue,
	"MAGENTA":ct.Magenta,
	"CYAN"   :ct.Cyan,
	"WHITE"  :ct.White }
var COLOR_KEYS []string

func tryParseConfig() {
	defer func() {
		if pan := recover(); pan != nil {
			fmt.Println("Configure file format error, please check it.")
			return
		}
	}()
	home := os.Getenv("HOME")
	if f, err := os.Open(home + "/.gologcat"); err == nil {
		r := bufio.NewReader(f)
		for {
			s, err := r.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err == nil {
				envs := strings.Split(s, ":")
				key := strings.ToUpper(envs[0])
				value := strings.ToUpper(envs[1])
				value = strings.Split(value, "\n")[0]
				if strings.Contains(LOG_SLOTS, key) && sliceIndexOf(COLOR_KEYS, value) != -1 {
					k := []rune(key)
					colorMap[k[0]] = COLOR_VALUES[value]
				}
			}
		}
	}
}

func sliceIndexOf(sli []string, s string) (position int) {
	position = -1
	length := len(sli)
	for i := 0; i < length; i++ {
		if strings.EqualFold(s, sli[i]) {
			return i
		}
	}
	return
}

func main() {
	for k := range COLOR_VALUES {
		COLOR_KEYS = append(COLOR_KEYS, k)
	}
	colorMap = make(map[rune]ct.Color)
	//default
	for _, i := range LOG_SLOTS {
		switch i {
		case 'D':
			colorMap[i] = COLOR_VALUES["GREEN"]
		case 'I':
			colorMap[i] = COLOR_VALUES["CYAN"]
		case 'W':
			colorMap[i] = COLOR_VALUES["YELLOW"]
		case 'E':
			colorMap[i] = COLOR_VALUES["RED"]
		default:
			colorMap[i] = COLOR_VALUES["NONE"]
		}
	}
	// load config in $HOME/.gologcat
	tryParseConfig()
	cmd := exec.Command("adb", "logcat")
	stdout, err := cmd.StdoutPipe();
	defer stdout.Close()
	if err != nil {
		ct.ChangeColor(ct.Red, false, ct.None, false)
		log.Fatal(err)
		return
	}

	if err := cmd.Start(); err != nil {
		ct.ChangeColor(ct.Blue, false, ct.None, false)
		log.Fatal(err)
		return
	}
	reader := bufio.NewReader(stdout)
	for {
		s, err := reader.ReadString('\n')
        if err == io.EOF {
            break
        }
		if err == nil {
			array := []rune(s)
            color := colorMap[array[0]]
            if color == ct.None {
                ct.ResetColor()
            } else {
			    ct.ChangeColor(colorMap[array[0]], false, ct.None, false)
            }
            fmt.Print(s)
		}
	}

}
