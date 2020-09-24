package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"gopkg.in/yaml.v2"
)

// Lyrics is a struct
type Lyrics struct {
	Song   string `yaml:"song"`
	Artist string `yaml:"artist"`
	Text   string `yaml:"text"`
}

func songFinder(dir string) (map[string]Lyrics, error) {
	lyrics := map[string]Lyrics{}

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			ext := filepath.Ext(path)
			rx := regexp.MustCompile(".ya?ml")
			if !rx.Match([]byte(ext)) {
				return nil
			}
			song, err := parseSongFile(path)
			if err != nil {
				panic("Invalid song file: " + path)
			}
			key := fmt.Sprintf("%s|%s", song.Artist, song.Song)
			lyrics[key] = song
			return nil
		})
	return lyrics, err
}

func parseSongFile(path string) (Lyrics, error) {
	l := Lyrics{}

	d, err := ioutil.ReadFile(path)
	if err != nil {
		return l, err
	}

	err = yaml.Unmarshal([]byte(d), &l)
	if err != nil {
		return l, err
	}
	return l, nil
}
