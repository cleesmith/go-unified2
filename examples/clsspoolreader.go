package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"

	"github.com/cleesmith/go-unified2"
)

// cd examples
// go run clsspoolreader.go ../sample_data snort.log

func main() {
	folder := os.Args[1]     // sample_data
	filePrefix := os.Args[2] // snort.log
	reader := unified2.NewSpoolRecordReader(folder, filePrefix)
	reader.Logger(log.New(os.Stdout, "SpoolRecordReader: ", 0))

	closeHookCount := 0
	reader.CloseHook = func(filepath string) {
		closeHookCount++
		log.Printf("*** reader.CloseHook called ... count=%d filepath=%s", closeHookCount, filepath)
		filedir := path.Dir(filepath)
		filename := path.Base(filepath)
		newname := fmt.Sprintf("/indexed_%v.", time.Now().Unix())
		filepathRename := filedir + newname + filename
		err := os.Rename(filepath, filepathRename)
		if err != nil {
			log.Printf("unable to rename file '%v' to '%v' err: %v", filepath, filepathRename, err)
			return
		}
		log.Printf("Indexed file '%v' and renamed to '%v'", filepath, filepathRename)
	}

	for {
		record, err := reader.Next()
		// filename, offset := reader.Offset()
		// if closeHookCount >= 1 {
		// 	log.Printf("closeHookCount=%d filename=%s; offset=%d", closeHookCount, filename, offset)
		// 	os.Exit(999)
		// }
		if err != nil {
			if err == io.EOF {
				// EOF is returned when the end of the last spool file
				// is reached and there is nothing else to read.  For
				// the purposes of the example, just sleep for a
				// moment and try again.
				// log.Printf("err=%v\n", err)
				// filename, offset := reader.Offset()
				// log.Printf("closeHookCount=%d filename=%s; offset=%d", closeHookCount, filename, offset)
				// os.Exit(999)
				log.Println("sleep")
				time.Sleep(time.Second * 5)
			} else {
				log.Fatalf("Unexpected error: '%v'", err)
			}
		}

		if record == nil {
			// The record and err are nil when there are no files at
			// all to be read.  This will happen if the Next() is
			// called before any files exist in the spool
			// directory. For now, sleep.
			// time.Sleep(time.Millisecond)
			log.Println("sleep")
			log.Printf("reader=%T=%v\n", reader, reader)
			time.Sleep(time.Second * 5)
			continue
		}

		// switch record := record.(type) {
		// case *unified2.EventRecord:
		// 	log.Printf("Event: EventId=%d\n", record.EventId)
		// case *unified2.PacketRecord:
		// 	log.Printf("- Packet: EventId=%d\n", record.EventId)
		// case *unified2.ExtraDataRecord:
		// 	log.Printf("- Extra Data: EventId=%d\n", record.EventId)
		// }

		filename, offset := reader.Offset()
		log.Printf("closeHookCount=%d filename=%s; offset=%d", closeHookCount, filename, offset)
	}

}
