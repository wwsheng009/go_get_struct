package main

import (
	"flag"
	"github.com/cartmanis/go_get_struct/generator"
	"github.com/cartmanis/go_get_struct/node"
	"github.com/prometheus/common/log"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	for _, pathFile := range flag.Args() {
		absPath, err := filepath.Abs(pathFile)
		if err != nil {
			log.Errorf("Не удалось  получить абсолютный путь до файла %v. Ошибка: %v", pathFile, err)
			return
		}
		file, err := os.Open(absPath)
		if err != nil {
			log.Errorf("Не удалось открыть файл %v. Ошибка: %v", pathFile, err)
			return
		}
		defer file.Close()
		n, err := node.Parse(file)
		if err != nil {
			log.Errorf("Не удалось распарсить файл %v. Ошибка: %v", pathFile, err)
			return
		}
		if err := generator.CreateStruct(n, absPath); err != nil {
			log.Errorf("Не удалось получить стуктуру go для файла %v. Ошибка: %v", pathFile, err)
			return
		}
	}

}
