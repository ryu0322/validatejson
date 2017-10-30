package reader

import (
	"github.com/tealeg/xlsx"
)

func Reader(inPath string) []*xlsx.Row {
	xlFile, err := xlsx.OpenFile(inPath)
	//_, err := xlsx.OpenFile(inPath)
	if err != nil {
		panic(err)
	}

	allRow := xlFile.Sheets[0].Rows

	return allRow
}
