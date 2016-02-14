package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/cleesmith/go-unified2"
)

// go run clsspoolreader.go ../sample_data snort.log

func main() {
	folder := os.Args[1]     // sample_data
	filePrefix := os.Args[2] // snort.log
	reader := unified2.NewSpoolRecordReader(folder, filePrefix)
	log.Printf("reader=%T=%v\n", reader, reader)

	closeHookCount := 0
	reader.CloseHook = func(filename string) {
		closeHookCount++
	}

	for {
		record, err := reader.Next()
		log.Printf("reader=%T=%v\n", reader, reader)
		if err != nil {
			if err == io.EOF {
				// EOF is returned when the end of the last spool file
				// is reached and there is nothing else to read.  For
				// the purposes of the example, just sleep for a
				// moment and try again.
				time.Sleep(time.Millisecond)
			} else {
				log.Fatal("Unexpected error: ", err) // Unexpected error.
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
		log.Printf("Current position: filename=%s; offset=%d", filename, offset)
	}

}
