package reader

import (
	"testing"
)

func TestReader(t *testing.T) {
	rows := Reader("C:/work/validatejson/バリデートパターンマトリクス.xlsx")

	if len(rows) == 0 {
		t.Fatalf("行数が0です")
	}

	for _, row := range rows {
		t.Logf("%t\n", row)
	}
}
