package main

import (
	"encoding/csv"
	"io"
	"time"
)

// unmarshalTime unmarshal data in CSV to time
func unmarshalTime(data []byte, t *time.Time) error {
	var err error
	*t, err = time.Parse("2006-01-02", string(data))
	return err
}

// parseData parses data from r and returns a table with columns filled
func parseData(r io.Reader) (Table, error) {
	dec, err := csvutil.NewDecoder(csv.NewReader(r))
	if err != nil {
		return Table{}, err
	}
	dec.Register(unmarshalTime)

	var table Table
	for {
		var row Row
		err := dec.Decode(&row)

		if err == io.EOF {
			break
		}

		if err != nil {
			return Table{}, err
		}

		table.Date = append(table.Date, row.Date)
		table.Price = append(table.Price, row.Close)
		table.Volume = append(table.Volume, row.Volume)
	}

	return table, nil
}
