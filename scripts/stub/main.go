package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/xackery/log"
)

func main() {
	log := log.New()
	err := run()
	if err != nil {
		log.Error().Err(err).Msg("failed")
		os.Exit(1)
	}
	log.Info().Msg("done")
}

func run() error {
	log := log.New()

	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}

	input := string(data)

	className := "Appearance"
	classPath := "appearance"
	log.Info().Msgf("generating %s", className)
	f, err := os.Open(fmt.Sprintf("../../content/lua/%s/_index.en.md", classPath))
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(strings.ToLower(line), "- [") {
			continue
		}

		//- [ExploreUnknown](exploreunknown) -- {{% lua_type_number %}}
		fileName := line[strings.Index(line, "(")+1:]
		fileName = fileName[:strings.Index(fileName, ")")]
		funcName := line[strings.Index(line, "[")+1:]
		funcName = funcName[:strings.Index(funcName, "]")]

		fw, err := os.Create(fmt.Sprintf("../../content/lua/%s/%s.md", classPath, fileName))
		if err != nil {
			return fmt.Errorf("create: %w", err)
		}
		prep := strings.ReplaceAll(input, "Ammo", strings.Title(funcName))
		prep = strings.ReplaceAll(prep, "Slot", className)

		fw.WriteString(prep)
		fw.Close()
	}

	return nil
}
