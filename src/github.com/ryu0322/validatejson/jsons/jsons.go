package jsons

import (
	"io/ioutil"
	"os"

	"github.com/tealeg/xlsx"
)

type actionIndex struct {
	action    string       `json` // アクション名
	itemIndex []*itemIndex `json`
}

type itemIndex struct {
	item     string           `json`
	valIndex []*validateIndex `json`
}

// validateIndexチェック内容を格納した項目です
type validateIndex struct {
	attribute        string `json: "attribute"`
	attributePattern string `json: "attribute"`
	classification   string `json: "classification"`
	rangeMin         int    `json: "attribute"`
	rangeMax         int    `json: "attribute"`
	length           string `json: "attribute"`
	checks           string `json: "checks"`
}

// CreateActionGroup アクション単位でグルーピングする
func CreateActionGroup(rows []*xlsx.Row) map[string][]*xlsx.Row {
	var actMap = make(map[string][]*xlsx.Row)
	actRow := []*xlsx.Row{}

	// 全行ループ
	for _, rowData := range rows {

		// アクション指定がある場合
		if rowData.Cells[10].Value != "" {
			//			actRowWk, keyFlg := actMap[rowData.Cells[10].Value]

			/*			// 既にデータがある場合は追加
						if keyFlg {
							actRow = actRowWk
						}*/
			actRow = append(actRow, rowData)
			actMap[rowData.Cells[10].Value] = actRow
		}
	}
	return actMap
}

// NewJson Jsonファイルを出力します
func NewJson(actList map[string][]*xlsx.Row) {

	strJsonBody := ""

	icnt := 0

	for key, value := range actList {

		if icnt > 0 {
			strJsonBody += ","
		}

		strValBody := ""
		strItemBody := ""
		boolComma := false
		for icnt2 := 0; icnt2 < len(value); icnt2++ {
			if icnt2 > 0 {
				strItemBody += ","
			}
			// 先にチェック内容から詰めていく
			strValBody += createValidateString(*value[icnt2])

			// Item内容を詰めて
			if !boolComma {
				strItemBody += createItemSting(key, *value[icnt2], strValBody)

			} else {
				strItemBody = strItemBody + ",\n" + createItemSting(key, *value[icnt2], strValBody)
			}
		}

		// アクション単位でJson文字列出力
		strJsonBody = key + ": {\n" + strItemBody + "}"
	}

	// 開始と終了のかっこで挟んで出来上がり
	strJsonBody = "{" + strJsonBody + "}"

	// ファイル出力
	byteJsonBody := []byte(strJsonBody)
	ioutil.WriteFile("C:/Myfolder/validatepattern.json", byteJsonBody, os.ModePerm)
}

// createValidateString アイテム名を生成します
func createValidateString(row xlsx.Row) string {
	strChecks := ""
	strVal := ""
	for icnt := 12; icnt < 21; icnt++ {
		if strVal != "" {
			strVal += "\n,"
		}
		if row.Cells[icnt].Value == "1" {
			strChecks += "1"

			switch icnt {
			case 13: // 属性
				strVal += "\"attribute\": \"" + row.Cells[21].Value + "\""
				if row.Cells[21].Value == "3" {
					strVal += ",\n\"attribute_format\": \"" + row.Cells[22].Value + "\""
				}
			case 14: // 区分値
				strVal += "\"classification\": \"" + row.Cells[23].Value + "\""
			case 15: // 桁数
				strVal += "\"length\": \"" + row.Cells[25].Value + "\""
			case 16: // 範囲
				strVal += "\"range_min\": \"" + row.Cells[24].Value + "\""
				strVal += ",\n\"range_max\": \"" + row.Cells[24].Value + "\""
			}
		} else {
			strChecks += "0"
		}
	}
	strVal += strChecks

	return strVal
}

// createItemString アイテム単位でJSON文字列を作成する
func createItemSting(key string, row xlsx.Row, valBody string) string {
	strItemBody := "\"" + row.Cells[11].Value + "\": {" + valBody + "}"
	return strItemBody
}
