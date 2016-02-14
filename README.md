### Changes

#### Feb 14, 2016:

* add ```os.Stat(r.reader.File.Name())``` to ```spoolrecordreader.go```
  * handles disappearing/renamed files
* add sample unified2 files in ```sample_data```
* add examples
  * ```clsreadu2.go``` simple reader with counts
  * ```clsspoolreader.go``` to test **SpoolRecordReader**

***

# go-unified2 [![GoDoc](https://godoc.org/github.com/jasonish/go-unified2?status.png)](https://godoc.org/github.com/jasonish/go-unified2)

A Go(lang) Library for decoding unified2 log files as generated by IDS
applications such as Snort and Suricata.

## Installation

```
go get github.com/jasonish/go-unified2
```

## Documentation

See https://godoc.org/github.com/jasonish/go-unified2

For more information on the unified2 file format see the
[Snort Manual](http://manual.snort.org/node44.html).
