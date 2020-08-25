package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	regLuaClass           = regexp.MustCompile(`class_\<[a-zA-Z]+\>\(\"([a-zA-z]+)\"\)`)
	regLuaProperty        = regexp.MustCompile(`\.property\(\"([a-zA-Z]+)\", [\&a-zA-Z\_:]+\)`)
	regLuaDef             = regexp.MustCompile(`\.def\(\"([a-zA-Z]+)\",`)
	regLuaFunction        = regexp.MustCompile(`([a-zA-Z_]+) ([a-zA-Z_]+)::([a-zA-Z]+)\(([0-9a-zA-Z*,_ ]+)\) `)
	regLuaFunctionNoParam = regexp.MustCompile(`([a-zA-Z_]+) ([a-zA-Z_]+)::([a-zA-Z]+)\(\) `)
)

// Scope retains information about the current scan's scope
type Scope struct {
	class string
}

var scope = Scope{}

func main() {
	err := run()
	if err != nil {
		fmt.Println("failed:", err.Error())
		os.Exit(1)
	}
}

func run() error {
	err := filepath.Walk("../../../server/zone", walk)
	if err != nil {
		return fmt.Errorf("walk: %w", err)
	}
	return nil
}

func walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return fmt.Errorf("walk %q: %w", path, err)
	}

	if info.IsDir() {
		return nil
	}

	// only parse .cpp
	if strings.ToLower(filepath.Ext(path)) != ".cpp" {
		return nil
	}
	if !strings.Contains(path, "lua") {
		return nil
	}

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open txt: %w", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		err = audit(line)
		if err != nil {
			return fmt.Errorf("audit %d: %w", lineNumber, err)
		}
		//fmt.Println(line)
	}
	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("scanner: %w", err)
	}

	return nil
}

func audit(line string) error {
	if err := class(line); err != nil {
		return fmt.Errorf("class: %w", err)
	}
	if err := property(line); err != nil {
		return fmt.Errorf("property: %w", err)
	}
	if err := def(line); err != nil {
		return fmt.Errorf("def: %w", err)
	}
	if err := function(line); err != nil {
		return fmt.Errorf("function: %w", err)
	}
	return nil
}

func class(line string) error {
	matches := regLuaClass.FindAllStringSubmatch(line, -1)
	if len(matches) < 1 {
		return nil
	}
	if len(matches[0]) < 2 {
		return nil
	}
	scope.class = matches[0][1]
	//fmt.Println("class set to", matches[0][1])
	return nil
}

func property(line string) error {
	matches := regLuaProperty.FindAllStringSubmatch(line, -1)
	if len(matches) < 1 {
		return nil
	}
	if len(matches[0]) < 2 {
		return nil
	}

	if scope.class != "Spell" {
		return nil
	}
	fmt.Printf("%s.%s\n", scope.class, matches[0][1])
	return nil
}

func def(line string) error {
	matches := regLuaDef.FindAllStringSubmatch(line, -1)
	if len(matches) == 0 {
		return nil
	}
	if len(matches) < 1 {
		return nil
	}
	if len(matches[0]) < 2 {
		return nil
	}
	if scope.class != "Spell" {
		return nil
	}
	fmt.Printf("%s:%s\n", scope.class, matches[0][1])
	//fmt.Println(matches)
	return nil
}

func function(line string) error {
	matches := regLuaFunction.FindAllStringSubmatch(line, -1)
	if len(matches) == 0 {
		return nil
	}
	if len(matches) < 1 {
		return nil
	}
	if len(matches[0]) < 2 {
		return nil
	}

	returnType := matches[0][1]
	class := matches[0][2]

	if strings.ToLower(class[0:4]) == "lua_" {
		class = class[4:]
	}
	fileClass := strings.ToLower(class)
	if fileClass == "luaparser" {
		return nil
	}
	filenames := map[string]string{
		"statbonuses": "stat_bonuses",
		"hateentry":   "hate_entry",
		"entitylist":  "entity_list",
	}
	for k, v := range filenames {
		if fileClass != k {
			continue
		}
		fileClass = v
		break
	}

	funcName := matches[0][3]
	params := matches[0][4]
	language := "lua"
	path := fmt.Sprintf("../../content/%s/%s/%s.md", language, fileClass, strings.ToLower(funcName))
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s is not documented: %s %s.%s(%s)\n", path, returnType, class, funcName, params)
			return nil
		}
		return fmt.Errorf("stat %s: %w", path, err)
	}

	/*f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer f.Close()*/

	//fmt.Printf("%s:%s\n", scope.class, matches[0][1])
	//fmt.Println(matches)
	return nil
}
