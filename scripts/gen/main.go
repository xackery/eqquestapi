package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("failed", err)
		os.Exit(1)
	}
	fmt.Println("done")
}

func run() error {
	err := filepath.Walk(".", scrape)
	if err != nil {
		return fmt.Errorf("walk: %w", err)
	}
	return nil
}

func scrape(path string, info os.FileInfo, err error) error {
	if err != nil {
		return fmt.Errorf("walk %q: %w", path, err)
	}

	if info.IsDir() {
		return nil
	}

	// only parse .pl and .lua
	if strings.ToLower(filepath.Ext(path)) != ".txt" {
		return nil
	}
	language := ""
	if strings.Contains(path, "lua") {
		language = "Lua"
	}
	if strings.Contains(path, "perl") {
		language = "Perl"
	}
	if len(language) == 0 {
		return fmt.Errorf("unknown language")
	}

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open txt: %w", err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("read data: %w", err)
	}
	f.Close()

	methods := strings.Split(fmt.Sprintf("%s", data), "\n")
	category := path[strings.Index(path, ".")+1:]
	category = category[0:strings.Index(category, ".")]
	category = strings.Title(category)

	indexFuncs := []string{}
	outPath := fmt.Sprintf("../../content/%s/%s/", strings.ToLower(language), strings.ToLower(category))

	fmt.Println(outPath)
	for _, method := range methods {
		if len(method) == 0 {
			continue
		}

		funcName := method
		if language == "Perl" {
			funcName = funcName[strings.Index(funcName, "->")+2:]
			funcName = funcName[0:strings.Index(funcName, "(")]
		}

		if language == "Lua" {
			funcName = funcName[strings.Index(funcName, ":")+1:]
			funcName = funcName[0:strings.Index(funcName, "(")]
		}

		filename := outPath + strings.ToLower(funcName) + ".md"
		f, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("create: %w", err)
		}
		funcPage := fmt.Sprintf(`---
title: %s %s
weight: 1
hidden: true
menuTitle: %s
---
## %s
`, language, funcName, funcName, funcName)
		funcPage += fmt.Sprintf("```%s\n%s\n```", strings.ToLower(language), method)
		f.WriteString(funcPage)
		f.Close()

		indexMethod := "- " + method[0:strings.Index(method, funcName)] + "[" + funcName + "](" + strings.ToLower(funcName) + ")" + method[strings.Index(method, funcName)+len(funcName):]

		indexFuncs = append(indexFuncs, indexMethod)
	}

	indexPage := fmt.Sprintf(`---
date: 2020-08-24T16:50:16+02:00
title: %s
menuTitle: %s
weight: 25
---

## %s Methods (%s)
`, strings.Title(strings.ReplaceAll(category, "_", " ")), strings.Title(strings.ReplaceAll(category, "_", " ")), strings.Title(strings.ReplaceAll(category, "_", " ")), language)

	indexPage += strings.Join(indexFuncs, "\n")
	f, err = os.Create(outPath + "_index.en.md")
	if err != nil {
		return fmt.Errorf("create index: %w", err)
	}
	f.WriteString(indexPage)
	f.Close()

	return nil
}
