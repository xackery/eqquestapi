package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	ac        Autocomplete
	class     string
	classDesc string
)

// Autocomplete struct
type Autocomplete map[string]Class

// Class represents an entry
type Class struct {
	Name       string              `json:"name"`
	Type       string              `json:"type"`
	Inherits   []string            `json:"inherits"`
	Properties map[string]Property `json:"properties"`
	Doc        string              `json:"doc"`
}

// Property struct
type Property struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Mode string `json:"mode"`
	Args []Arg  `json:"args"`
	Doc  string `json:"doc"`
}

// Arg for functions
type Arg struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Doc  string `json:"doc"`
}

func main() {
	err := run()
	if err != nil {
		fmt.Println("failed", err)
		os.Exit(1)
	}
	fmt.Println("done")
}

func run() error {
	ac = Autocomplete{}
	ac = make(map[string]Class)

	/*	prop := Property{
			Name: "spawn2",
			Type: "function",
			Doc:  "spawns an NPC from database",
		}
		prop.Args = []Arg{
			{
				Name: "npc_type",
				Type: "number",
				Doc:  "npc_type :: [npc_type](https://questapi.firebaseapp.com/scripts/npc_type)",
			},
			{
				Name: "grid",
				Type: "number",
				Doc:  "grid number, if any, default to 0",
			},
			{
				Name: "unused",
				Type: "number",
				Doc:  "unusued field, default to 0",
			},
			{
				Name: "x",
				Type: "number",
			},
		}
		c := Class{
			Name: "eq",
			Type: "eq",
			Doc:  "EverQuest Global Namespace",
		}

		c.Properties = make(map[string]Property)
		c.Properties["spawn2"] = prop

		ac["eq"] = c
		enc := json.NewEncoder(os.Stdout)
		enc.Encode(ac)*/
	err := filepath.Walk("../../content/lua", scrape)
	if err != nil {
		return fmt.Errorf("walk: %w", err)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(ac)
	return nil
}

func scrape(path string, info os.FileInfo, err error) error {
	if err != nil {
		return fmt.Errorf("walk %q: %w", path, err)
	}

	if info.IsDir() {
		class = info.Name()
		ac[class] = Class{
			Name:       class,
			Type:       class,
			Properties: make(map[string]Property),
		}
		classDesc = ""
		fmt.Println(class)
		return nil
	}

	// only parse markdown files
	if strings.ToLower(filepath.Ext(path)) != ".md" {
		return nil
	}

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open md: %w", err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("read data: %w", err)
	}
	f.Close()

	isMarkdown := false
	lines := strings.Split(string(data), "\n")

	isDocs := false
	prop := Property{}
	for _, line := range lines {
		if !isMarkdown && strings.TrimSpace(line) == "---" {
			isMarkdown = true
			continue
		}
		if isMarkdown && strings.TrimSpace(line) == "---" {
			break
		}
		if !strings.Contains(line, ":") {
			continue
		}
		key := line[0:strings.Index(line, ":")]
		value := strings.TrimSpace(line[len(key)+1:])

		if strings.Contains(path, "_index") {
			switch key {
			case "description":
				classDesc = value
			}
		}
		switch key {
		case "function":
			prop.Name = value
			prop.Type = "function"
		case "property":
			prop.Name = value
			prop.Type = "property"
		case "description":
			if !isDocs {
				prop.Doc = value
			}
		case "docs":
			isDocs = true
			prop.Doc = value
		}
	}
	if prop.Name != "" {
		c := ac[class]
		c.Properties[prop.Name] = prop
		c.Doc = classDesc
		ac[class] = c
	}

	return nil
}
