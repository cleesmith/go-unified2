package main

import (
	"io"
	"log"
	"os"

	"github.com/cleesmith/go-unified2"
)

// cd examples
// go run clsreadu2.go ../sample_data/snort.log.1452978988

func main() {
	var etot, ptot, xtot, ttot int64
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// Read records.
	for {
		record, err := unified2.ReadRecord(file)
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				// End of file is reached.  You may want to break here
				// or sleep and try again if you are expecting more
				// data to be written to the input file.
				break
			} else if err == unified2.DecodingError {
				// Error decoding a record, probably corrupt.
				log.Fatal("DecodingError: ", err)
			}
			log.Fatal("wtf? err: ", err) // Some other error.
		}
		ttot++

		switch record := record.(type) {
		case *unified2.EventRecord:
			etot++
			log.Printf("Event: EventId=%d\n", record.EventId)
		case *unified2.PacketRecord:
			ptot++
			log.Printf("- Packet: EventId=%d\n", record.EventId)
		case *unified2.ExtraDataRecord:
			xtot++
			log.Printf("- Extra Data: EventId=%d\n", record.EventId)
		}
	}

	file.Close()

	log.Println()
	log.Println("______ Unified2 records _____")
	log.Printf("event.....: %d\n", etot)
	log.Printf("packet....: %d\n", ptot)
	log.Printf("extradata.: %d\n", xtot)
	log.Printf("total.....: %d\n", ttot)
	log.Println("_____________________________\n")
}
