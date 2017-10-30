package reader

import (
	"github.com/tealeg/xlsx"
)

func Reader(inPath string) Row[]{
	xlFile, err := xlsx.OpenFile(readPath)
	if err != nil {
		panic(err)
	}

	var allRow = xlFile.Sheets[0].Rows

	return allRow
}
