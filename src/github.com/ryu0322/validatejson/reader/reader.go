package reader

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func Reader(inPath string) []*xlsx.Row {
	xlFile, err := xlsx.OpenFile(inPath)
	//_, err := xlsx.OpenFile(inPath)
	if err != nil {
		panic(err)
	}

	allRow := xlFile.Sheets[0].Rows

	for _, row := range allRow {
		fmt.Printf("%x\n", row)
	}
	return allRow
}
