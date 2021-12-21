package main

/*
go run main.go -a WorkDir
go run main.go -a -r WorkDir
*/

import (
	"FindDuplicate/process"
	"flag"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
)

var (
	flagR *bool
	flagA *bool
)

func main() {
	flagR = flag.Bool("r", false, "Удаляет найденные дубликаты в подкаталогах")
	flagA = flag.Bool("a", false, "Перед удалением не спрашивать подтверждения")
	flag.Parse()

	log.SetFormatter(&log.JSONFormatter{})
	fieldsStandard := log.Fields{"duplicate": "service"}
	hLog := log.WithFields(fieldsStandard)

	if len(os.Args) == 1 {
		hLog.Fatal("Число параметров должно быть больше одного")
	}

	hLog.WithFields(log.Fields{"main": "block"}).Info("Start")

	fDir, err := os.Open(flag.Args()[0])
	if err != nil {
		hLog.Fatal("Не верно указана стартовая директория")
	}

	// Запуск обхода указанной директории
	wg := sync.WaitGroup{}
	options := process.OptionsNew(!*flagA, *flagR, 10, hLog)

	err = process.StartWatch(options, fDir, &wg)
	if err != nil {
		hLog.Panicf("Ошибка старта работы службы (%v)", err)
	}

	hLog.WithFields(log.Fields{"main": "block"}).Info("Finish")
}
