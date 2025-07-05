package ieservice

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"log"
	"os"
	"time"
)

type IEService interface {
	Import(f *os.File) error
	Export(data interface{}) (*os.File, error)
}

type IEDefault struct {
}

func New() *IEDefault {
	return &IEDefault{}
}

func (e *IEDefault) Import(f *os.File) error {
	return nil
}

func (e *IEDefault) Export(data interface{}) (*os.File, error) {
	f, err := os.Create(
		fmt.Sprintf("export-%s.csv", time.Now().Format("02-01-2006-15:04")),
	)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	if err != nil {
		return nil, err
	}

	if err := gocsv.MarshalFile(data, f); err != nil {
		log.Fatalf("ошибка при записи CSV: %v", err)
	}

	return f, nil
}
