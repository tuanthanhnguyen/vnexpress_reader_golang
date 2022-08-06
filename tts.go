package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/hajimehoshi/oto/v2"
	htgotts "github.com/hegedustibor/htgo-tts"
	voices "github.com/hegedustibor/htgo-tts/voices"

	"time"

	"github.com/hajimehoshi/go-mp3"
)

func gg_recognise(text, path string) string {
	dir := "./output/"
	path = dir + path
	_ = os.RemoveAll(path)
	speech := htgotts.Speech{Folder: path, Language: voices.Vietnamese}
	speech.Speak(text)
	files, _ := ioutil.ReadDir(path)
	var name string = files[0].Name()
	return path + "/" + name
}
func play(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, ready, err := oto.NewContext(d.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-ready

	p := c.NewPlayer(d)
	defer p.Close()
	p.Play()

	fmt.Printf("Length: %d[bytes]\n", d.Length())
	for {
		time.Sleep(time.Second)
		if !p.IsPlaying() {
			break
		}
	}

	return nil
}
func splitter(text string) []string {
	max := 100
	lines := strings.Split(text, "\n")
	rst := []string{}
	for _, line := range lines {
		if len(line) <= max {
			rst = append(rst, line)
		} else {
			line = strings.ReplaceAll(line, ", ", ",")
			line = strings.ReplaceAll(line, "; ", ";")
			sentences := strings.Split(line, ".")
			for _, sen := range sentences {
				if len(sen) <= max {
					check := strings.ReplaceAll(sen, " ", "")
					check = strings.ReplaceAll(check, ".", "")
					check = strings.ReplaceAll(check, ",", "")
					if check != "" {
						rst = append(rst, sen+". ")
					}
				} else {
					halves := strings.Split(sen, ",")
					for _, hv := range halves {
						if len(hv) <= max {
							rst = append(rst, hv+"")
							//fmt.Println(a)
						} else {
							//fmt.Println(a, "LONG")
							words := strings.Split(hv, " ")
							phrase := ""
							for _, w := range words {
								if len(phrase+w) > max {
									rst = append(rst, phrase)
									phrase = w + " "
								} else {
									phrase += w + " "
								}
							}
							rst = append(rst, phrase+",")
						}
					}
				}
			}
		}
	}
	payload := []string{}
	cell := ""
	for _, chunk := range rst {
		if len(cell+chunk) > max {
			payload = append(payload, cell)
			cell = chunk
		} else {
			cell += chunk

		}
	}
	return payload
}
