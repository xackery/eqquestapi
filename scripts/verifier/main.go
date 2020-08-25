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
	regLuaClass                  = regexp.MustCompile(`([class|namespace]+)_.*\(\"([a-zA-z]+)\"\)`)
	regLuaProperty               = regexp.MustCompile(`[::|.]property\(\"([a-zA-Z]+)\", [&|].*::a-zA-Z\_:]+\)`)
	regLuaDef                    = regexp.MustCompile(`[::|.]def\(\"([a-zA-Z]+)\", [(&|].*::([a-zA-Z0-9_]+)`)
	regLuaFunction               = regexp.MustCompile(`([0-9a-zA-Z_]+) ([*a-zA-Z_]+)::([a-zA-Z]+)\(([:&0-9a-zA-Z*,_ ]+)\)`)
	regLuaFunctionNoParam        = regexp.MustCompile(`([0-9a-zA-Z_]+) ([*a-zA-Z_]+)::([a-zA-Z]+)\(\)`)
	regLuaGeneralFunction        = regexp.MustCompile(`([0-9a-zA-Z_]+) lua_([a-zA-Z_]+)\(([0-9a-zA-Z*,_ ]+)\)`)
	regLuaGeneralFunctionNoParam = regexp.MustCompile(`([0-9a-zA-Z_]+) lua_([a-zA-Z_]+)\(\)`)
)

// Scope retains information about the current scan's scope
type Scope struct {
	class       string
	isNamespace bool
	functions   map[string]FunctionScope
}

// FunctionScope declares
type FunctionScope struct {
	language   string
	class      string
	returnType string
	funcName   string
	params     string
}

var scope = Scope{}

func main() {
	scope.functions = make(map[string]FunctionScope)
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
	if strings.Contains(path, "lua_general.cpp") {
		scope.class = "eq"
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
	if err := functionNoParam(line); err != nil {
		return fmt.Errorf("functionNoParam: %w", err)
	}
	if err := generalFunction(line); err != nil {
		return fmt.Errorf("generalFunction: %w", err)
	}
	if err := generalFunctionNoParam(line); err != nil {
		return fmt.Errorf("generalFunctionNoParam: %w", err)
	}

	return nil
}

func class(line string) error {
	matches := regLuaClass.FindAllStringSubmatch(line, -1)
	if len(matches) < 1 {
		return nil
	}
	if len(matches[0]) < 3 {
		fmt.Println(matches)
		return nil
	}
	if matches[0][1] == "namespace" {
		scope.isNamespace = true
	} else {
		scope.isNamespace = false
	}
	scope.class = matches[0][2]
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

	language := "lua"

	syntax := fmt.Sprintf("%s.%s", scope.class, matches[0][1])

	if err := checkDocumented(language, strings.ToLower(scope.class), matches[0][1], syntax); err != nil {
		return fmt.Errorf("checkDocumented: %w", err)
	}
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
	if len(matches[0]) < 3 {
		return nil
	}

	language := "lua"
	delimiter := ":"
	if scope.isNamespace {
		delimiter = "."
	}

	funcName := matches[0][1]

	key := fmt.Sprintf("lua_%s.%s", strings.ToLower(scope.class), strings.ToLower(matches[0][2]))
	fs, ok := scope.functions[key]
	if !ok && funcName != "null" && funcName != "valid" {
		fmt.Println("missing", key, funcName)
	}
	if funcName == "null" || funcName == "valid" {
		fs.returnType = "bool"
	}

	params := ""
	if len(fs.params) > 0 {
		params = fmt.Sprintf("(%s)", fs.params)
	}
	if len(params) == 0 && ok { //if a function exists
		params = "()"
	}

	syntax := fmt.Sprintf("%s%s%s%s; -- %s", scope.class, delimiter, funcName, params, fs.returnType)

	if err := checkDocumented(language, strings.ToLower(scope.class), funcName, syntax); err != nil {
		return fmt.Errorf("checkDocumented: %w", err)
	}
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
	class = strings.ReplaceAll(class, "*", "")
	funcName := matches[0][3]
	params := matches[0][4]
	language := "lua"

	key := fmt.Sprintf("%s.%s", strings.ToLower(class), strings.ToLower(funcName))

	fs, ok := scope.functions[key]
	if !ok {
		fs = FunctionScope{}
	}
	fs.language = language
	fs.returnType = luaReturns(returnType)
	fs.class = class
	fs.funcName = funcName
	fs.params = params

	scope.functions[key] = fs
	return nil
}

func functionNoParam(line string) error {
	matches := regLuaFunctionNoParam.FindAllStringSubmatch(line, -1)
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
	class = strings.ReplaceAll(class, "*", "")

	funcName := matches[0][3]
	language := "lua"
	key := fmt.Sprintf("%s.%s", strings.ToLower(class), strings.ToLower(funcName))
	fs, ok := scope.functions[key]
	if !ok {
		fs = FunctionScope{}
	}
	fs.language = language
	fs.returnType = luaReturns(returnType)
	fs.class = class
	fs.funcName = funcName

	scope.functions[key] = fs
	return nil
}

func generalFunction(line string) error {
	if scope.class != "eq" {
		return nil
	}

	matches := regLuaGeneralFunction.FindAllStringSubmatch(line, -1)
	if len(matches) < 1 {
		return nil
	}
	if len(matches[0]) < 4 {
		return nil
	}

	returnType := matches[0][1]
	fileClass := strings.ToLower(scope.class)
	funcName := matches[0][2]
	params := matches[0][3]
	language := "lua"

	syntax := fmt.Sprintf("%s.%s(%s) -- %s", scope.class, funcName, params, returnType)

	if err := checkDocumented(language, fileClass, funcName, syntax); err != nil {
		return fmt.Errorf("checkDocumented: %w", err)
	}
	return nil
}

func generalFunctionNoParam(line string) error {
	if scope.class != "eq" {
		return nil
	}

	matches := regLuaGeneralFunctionNoParam.FindAllStringSubmatch(line, -1)
	if len(matches) < 1 {
		return nil
	}
	if len(matches[0]) < 3 {
		return nil
	}

	returnType := matches[0][1]
	fileClass := strings.ToLower(scope.class)
	funcName := matches[0][2]
	language := "lua"

	syntax := fmt.Sprintf("%s.%s() -- %s", scope.class, funcName, returnType)

	if err := checkDocumented(language, fileClass, funcName, syntax); err != nil {
		return fmt.Errorf("checkDocumented: %w", err)
	}
	return nil
}

func checkDocumented(language string, fileClass string, funcName string, syntax string) error {
	filenames := map[string]string{
		"statbonuses": "stat_bonuses",
		"hateentry":   "hate_entry",
		"entitylist":  "entity_list",
		"rulei":       "rule",
		"ruleb":       "rule",
		"ruler":       "rule",
	}
	if fileClass == "luaparser" {
		return nil
	}

	for k, v := range filenames {
		if fileClass != k {
			continue
		}
		fileClass = v
		break
	}

	path := fmt.Sprintf("../../content/%s/%s/%s.md", language, fileClass, strings.ToLower(funcName))

	_, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("stat %s: %w", path, err)
		}
		fmt.Printf("%s is not documented: %s\n", path, syntax)
		err = createStub(path, language, funcName, syntax)
		if err != nil {
			return fmt.Errorf("create stub: %w", err)
		}
	}

	//second check if index references it
	path = fmt.Sprintf("../../content/%s/%s/_index.en.md", language, fileClass)
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	isIndexed := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(funcName)) {
			isIndexed = true
			break
		}
	}
	if !isIndexed {
		entry := "- " + syntax[0:strings.Index(syntax, funcName)] + "[" + funcName + "](" + strings.ToLower(funcName) + ")" + syntax[strings.Index(syntax, funcName)+len(funcName):]
		fo, err := os.OpenFile(path,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("write: %w", err)
		}
		defer fo.Close()
		if _, err := fo.WriteString(entry + "\n"); err != nil {
			return fmt.Errorf("writeString: %w", err)
		}
		fmt.Println("added", entry)
	}

	return nil
}

func createStub(path string, language string, funcName string, syntax string) error {
	f, err := os.Create(path)
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
`, strings.Title(language), funcName, funcName, funcName)
	funcPage += fmt.Sprintf("```%s\n%s\n```", strings.ToLower(language), syntax)
	f.WriteString(funcPage)
	f.Close()

	//		indexMethod := "- " + method[0:strings.Index(method, funcName)] + "[" + funcName + "](" + strings.ToLower(funcName) + ")" + method[strings.Index(method, funcName)+len(funcName):]

	return nil
}

func luaReturns(in string) string {
	switch in {
	case "uint32":
		return "number"
	case "string":
		return "string"
	case "char *":
		return "string"
	case "char":
		return "string"
	case "int":
		return "number"
	case "float":
		return "number"
	case "bool":
		return "bool"
	case "void":
		return "void"
	case "double":
		return "number"
	case "uint64":
		return "number"
	}
	return "unknown - " + in
}
