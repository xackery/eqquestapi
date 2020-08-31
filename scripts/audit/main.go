package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/xackery/log"
)

var (
	regLuaClass       = regexp.MustCompile(`luabind::class.*\(\"([a-zA-z]+)\"\)`)
	regLuaNamespace   = regexp.MustCompile(`luabind::namespace.*\(\"([a-zA-z]+)\"\)`)
	regLuaProperty    = regexp.MustCompile(`.*property\(\"([\w]+)\"`)
	regLuaDef         = regexp.MustCompile(`.*def\(\"([\w]+)\"`)
	regLuaValue       = regexp.MustCompile(`.*value\(\"([\w]+)\"`)
	regPerlClass      = regexp.MustCompile(`.*XS\(boot_([\w]+)`)
	regPerlProto      = regexp.MustCompile(`.*newXSproto\(strcpy\(buf, "(\w+)"\), XS_(\w+)_`)
	regPerlDefinition = regexp.MustCompile(`.*XS\(XS_([\w]+)\);`)
)

// Scope retains information about the current scan's scope
type Scope struct {
	class     string
	namespace string
	language  string
	defType   string
}

func (s Scope) syntax(funcName string) string {
	if s.language == "lua" {
		if s.class != "" {
			if s.defType == "value" {
				return fmt.Sprintf("%s.%s", s.class, funcName)
			}
			return fmt.Sprintf("%s:%s", s.class, funcName)
		}
		return fmt.Sprintf("%s.%s", s.namespace, funcName)
	}

	return fmt.Sprintf("$%s->%s", strings.ToLower(s.class), funcName)
}

func (s Scope) path(funcName string) string {
	v := ""
	if s.class != "" {
		v = s.class
	} else {
		v = s.namespace
	}

	return fmt.Sprintf("../../content/%s/%s/%s.md", s.language, pathize(v), strings.ToLower(funcName))
}

func (s Scope) pathIndex() string {
	v := ""
	if s.class != "" {
		v = s.class
	} else {
		v = s.namespace
	}

	return fmt.Sprintf("../../content/%s/%s/_index.en.md", s.language, pathize(v))
}

var scope = Scope{}

func main() {
	log := log.New()
	err := run()
	if err != nil {
		log.Error().Err(err).Msg("failed")
		os.Exit(1)
	}
	log.Info().Msg("completed")
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
	if !strings.Contains(path, "lua") && !strings.Contains(path, "perl") {
		return nil
	}
	if strings.Contains(path, "lua_general.cpp") {
		scope.namespace = "eq"
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
	if err := auditLua(line); err != nil {
		return err
	}
	return nil
	scope = Scope{}
	if err := auditPerl(line); err != nil {
		return err
	}

	return nil
}

func auditLua(line string) error {
	scope.language = "lua"
	if err := luaClass(line); err != nil {
		return fmt.Errorf("class: %w", err)
	}
	if err := luaNamespace(line); err != nil {
		return fmt.Errorf("luaClass: %w", err)
	}

	if err := luaProperty(line); err != nil {
		return fmt.Errorf("property: %w", err)
	}
	if err := luaDef(line); err != nil {
		return fmt.Errorf("def: %w", err)
	}
	if err := luaValue(line); err != nil {
		return fmt.Errorf("def: %w", err)
	}
	return nil
}

func auditPerl(line string) error {
	scope.language = "perl"
	if err := perlClass(line); err != nil {
		return fmt.Errorf("perlClass: %w", err)
	}
	if err := perlProto(line); err != nil {
		return fmt.Errorf("perlClass: %w", err)
	}
	if err := perlDefinition(line); err != nil {
		return fmt.Errorf("perlClass: %w", err)
	}
	return nil
}

func luaClass(line string) error {
	log := log.New()
	matches := regLuaClass.FindAllStringSubmatch(line, -1)
	if len(matches) < 1 {
		return nil
	}
	if len(matches[0]) < 2 {
		log.Debug().Str("line", line).Msg("lua class not enough submatches")
		return nil
	}
	scope.class = matches[0][1]
	scope.namespace = ""
	scope.language = "lua"

	//log.Debug().Str("class", scope.class).Msg("class changed")
	return nil
}

func luaNamespace(line string) error {
	log := log.New()
	matches := regLuaNamespace.FindAllStringSubmatch(line, -1)
	if len(matches) < 1 {
		return nil
	}
	if len(matches[0]) < 2 {
		log.Debug().Str("line", line).Msg("lua namespace not enough submatches")
		return nil
	}
	scope.namespace = matches[0][1]
	scope.class = ""

	//log.Debug().Str("namespace", scope.namespace).Msg("namespace changed")
	return nil
}

func luaProperty(line string) error {
	log := log.New()

	matches := regLuaProperty.FindAllStringSubmatch(line, -1)
	if len(matches) < 1 {
		return nil
	}
	if len(matches[0]) < 2 {
		log.Debug().Str("line", line).Msg("lua property not enough submatches")
		return nil
	}

	scope.defType = "property"

	if err := checkDocumented(matches[0][1]); err != nil {
		return fmt.Errorf("checkDocumented: %w", err)
	}
	return nil
}

func luaDef(line string) error {
	log := log.New()

	matches := regLuaProperty.FindAllStringSubmatch(line, -1)
	if len(matches) < 1 {
		return nil
	}
	if len(matches[0]) < 2 {
		log.Debug().Str("line", line).Msg("lua property not enough submatches")
		return nil
	}

	scope.defType = "def"
	if err := checkDocumented(matches[0][1]); err != nil {
		return fmt.Errorf("checkDocumented: %w", err)
	}
	return nil
}

func luaValue(line string) error {
	log := log.New()

	matches := regLuaValue.FindAllStringSubmatch(line, -1)
	if len(matches) < 1 {
		return nil
	}
	if len(matches[0]) < 2 {
		log.Debug().Str("line", line).Msg("lua value not enough submatches")
		return nil
	}

	scope.defType = "value"
	if err := checkDocumented(matches[0][1]); err != nil {
		return fmt.Errorf("checkDocumented: %w", err)
	}
	return nil
}

func perlClass(line string) error {
	log := log.New()

	matches := regPerlClass.FindAllStringSubmatch(line, -1)
	if len(matches) < 1 {
		return nil
	}
	if len(matches[0]) < 2 {
		log.Debug().Str("line", line).Msg("perl class not enough submatches")
		return nil
	}
	if len(matches[0][1]) == 0 {
		return nil
	}

	scope.class = matches[0][1]
	return nil
}

func perlProto(line string) error {
	log := log.New()

	matches := regPerlProto.FindAllStringSubmatch(line, -1)
	if len(matches) < 1 {
		return nil
	}
	if len(matches[0]) < 3 {
		log.Debug().Str("line", line).Msg("perl class not enough submatches")
		return nil
	}

	scope.class = matches[0][2]

	if err := checkDocumented(matches[0][1]); err != nil {
		return fmt.Errorf("checkDocumented: %w", err)
	}
	return nil
}

func perlDefinition(line string) error {
	log := log.New()

	matches := regPerlDefinition.FindAllStringSubmatch(line, -1)
	if len(matches) < 1 {
		return nil
	}

	if len(matches[0]) < 2 {
		log.Debug().Str("line", line).Msg("perl class not enough submatches")
		return nil
	}

	splits := strings.SplitN(matches[0][1], "_", 2)
	if len(splits) != 2 {
		log.Warn().Str("line", line).Msg("no class scope?")
		return nil
	}
	funcName := splits[1]
	scope.class = splits[0]

	if scope.class == "" {
		log.Warn().Str("line", line).Msg("no class set")
		return nil
	}

	if err := checkDocumented(funcName); err != nil {
		return fmt.Errorf("checkDocumented: %w", err)
	}
	return nil
}

func checkDocumented(funcName string) error {
	if funcName == "null" || funcName == "valid" {
		return nil
	}
	if err := checkIsIndexed(funcName); err != nil {
		return fmt.Errorf("checkIsIndexed: %w", err)
	}
	if err := checkHasEntry(funcName); err != nil {
		return fmt.Errorf("checkHasEntry: %w", err)
	}

	return nil
}

func checkIsIndexed(funcName string) error {
	log := log.New()

	f, err := os.Open(scope.pathIndex())
	if err != nil {
		if !os.IsNotExist(err) {
			log.Warn().Str("path", scope.pathIndex()).Msg("index does not exist")
			return nil
		}
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(funcName)) {
			return nil
		}
	}
	log.Warn().Str("path", scope.pathIndex()).Msgf("%s %s is not indexed", scope.defType, scope.syntax(funcName))
	return nil
}

func checkHasEntry(funcName string) error {
	log := log.New()
	_, err := os.Stat(scope.path(funcName))
	if err != nil {
		if os.IsNotExist(err) {
			log.Warn().Str("path", scope.path(funcName)).Msgf("%s %s does not exist", scope.defType, scope.syntax(funcName))
			return nil
		}
		return err
	}
	return nil
}

func pathize(name string) string {
	specials := map[string]string{
		"clientlist":     "client_list",
		"clientversion":  "client_version",
		"corpselist":     "corpse_list",
		"doorlist":       "door_list",
		"entitylist":     "entity_list",
		"hateentry":      "hate_entry",
		"hatelist":       "hate_list",
		"inventorywhere": "inventory_where",
		"journalmode":    "journal_mode",
		"moblist":        "mob_list",
		"npclist":        "npc_list",
		"objectlist":     "object_list",
		"ruleb":          "rule_b",
		"rulei":          "rule_i",
		"ruler":          "rule_r",
		"spawnlist":      "spawn_list",
		"speakmode":      "speak_mode",
		"specialability": "special_ability",
		"statbonuses":    "stat_bonuses",
		"doors":          "door",
	}
	name = strings.ToLower(name)
	for k, v := range specials {
		if name == k {
			return v
		}
	}
	return name
}
