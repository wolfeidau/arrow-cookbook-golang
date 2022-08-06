package main

import (
	"io"
	"os"

	"github.com/alecthomas/kong"
	"github.com/apache/arrow/go/v10/arrow"
	"github.com/apache/arrow/go/v10/arrow/csv"
	"github.com/rs/zerolog"
)

var (
	cli struct {
		Filename string
	}
	logger = zerolog.New(zerolog.NewConsoleWriter())
)

func main() {
	kong.Parse(&cli)

	f, err := os.Open(cli.Filename)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to open file")
	}

	schema, err := buildSchema(f)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to build schema")
	}

	rdr := csv.NewReader(f, schema)

	for rdr.Next() {
		rec := rdr.Record()

		logger.Info().Int64("cols", rec.NumCols()).Msg("read record")
	}
}

// takes the first couple of lines of the CSV and builds a schema
func buildSchema(r io.Reader) (*arrow.Schema, error) {

	schema := new(arrow.Schema)

	return schema, nil
}
