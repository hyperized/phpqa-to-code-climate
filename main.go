package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperized/phpqa-to-code-climate/models/codeclimate"
	"github.com/hyperized/phpqa-to-code-climate/models/psalm"
	"log"
	"os"
)

func main() {
	var (
		stdin     = os.Stdin
		stdout    = os.Stdout
		stat, err = stdin.Stat()
	)
	if err != nil {
		log.Fatal(err)
	}

	if !fileIsNamedPipe(stat) {
		fmt.Println("Please pipe your input")
		fmt.Println("Usage: cat results.jsonBytes | phpqa-to-code-climate >> code-climate.jsonBytes")
	}

	var (
		reader                         = bufio.NewReader(stdin)
		byteArray, streamConversionErr = streamToByteArray(reader)
	)
	if streamConversionErr != nil {
		log.Fatal(streamConversionErr)
	}

	collection := psalm.Collection{}
	unmarshalErr := collection.Unmarshal(byteArray)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	log.Printf("This collection contains: %d item(s)", len(collection))

	codeClimate := psalmMapper(collection)
	jsonBytes, marshalErr := json.MarshalIndent(codeClimate, "", "    ")
	if marshalErr != nil {
		log.Fatal(marshalErr)
	}

	bytesOut, writeErr := stdout.Write(jsonBytes)
	log.Printf("Wrote %d bytes to output", bytesOut)
	if writeErr != nil {
		log.Fatal(writeErr)
	}

}

func psalmMapper(collection psalm.Collection) codeclimate.Collection {
	var output = codeclimate.Collection{}
	for _, v := range collection {
		model := codeclimate.CodeClimate{
			Description: v.Message,
		}
		output = append(output, model)
	}
	return output
}

func fileIsNamedPipe(stat os.FileInfo) bool {
	return stat.Mode()&os.ModeNamedPipe != 0
}

func streamToByteArray(stream *bufio.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	bytesRead, err := buf.ReadFrom(stream)
	log.Printf("Read %d bytes from input", bytesRead)
	return buf.Bytes(), err
}
