package main

import (
	"log"
	"os"
	"path/filepath"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/load"
	"cuelang.org/go/encoding/gocode"
	_ "cuelang.org/go/pkg"
)

func main() {
	dirs, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		dir := filepath.Join(cwd)
		pkg := "." + string(filepath.Separator) + d.Name()
		inst := cue.Build(load.Instances([]string{pkg}, &load.Config{
			Dir:        dir,
			ModuleRoot: dir,
			Module:     "mynewgoproject/internal/core/domain",
		}))[0]
		if err := inst.Err; err != nil {
			log.Fatal(err)
		}

		goPkg := "../domain"
		b, err := gocode.Generate(goPkg, inst, nil)
		if err != nil {
			log.Fatal(err)
		}

		goFile := filepath.Join("cue_gen.go")
		if err := os.WriteFile(goFile, b, 0644); err != nil {
			log.Fatal(err)
		}
	}
}
