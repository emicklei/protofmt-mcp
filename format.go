package main

import (
	"bytes"
	"io"
	"os"

	"github.com/emicklei/proto"
	"github.com/emicklei/proto-contrib/pkg/protofmt"
)

func readFormatWrite(filename string) error {
	// open for read
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	// buffer before write
	buf := new(bytes.Buffer)
	if err := format(filename, file, buf); err != nil {
		return err
	}
	// write back to input
	if err := os.WriteFile(filename, buf.Bytes(), os.ModePerm); err != nil {
		return err
	}
	return nil
}

func format(filename string, input io.Reader, output io.Writer) error {
	parser := proto.NewParser(input)
	parser.Filename(filename)
	def, err := parser.Parse()
	if err != nil {
		return err
	}
	protofmt.NewFormatter(output, "  ").Format(def) // 2 spaces
	return nil
}
