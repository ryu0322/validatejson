package jsons

import (
	"fmt"
	"log"
	"strings"
	_"fmt"

	"github.com/ryu0322/validatejson/reader"
)

// アクション情報
type actionIndex struct {
	action    string       `json` // アクション名
	itemIdx []itemIndex `json`  // 項目情報配列
}

// 項目情報
type itemIndex struct {
	item     string           `json`	// 項目名
	valIndex validateIndex `json`		// チェック情報`
}

// validateIndexチェック内容を格納した項目です
type validateIndex struct {
	attribute        string `json: "attribute"`			// 属性チェックパターン
	attributePattern string `json: "attribute_format"`	// 属性チェックフォーマット
	classification   string `json: "classification"`	// 区分値文字列
	rangeMin         string    `json: "range_min"`			// 最小範囲
	rangeMax         string    `json: "range_max"`			// 最大範囲
	length           string `json: "length"`			// 許容桁数
	checks           string `json: "checks"`			// チェックパターン文字列
}

func CreateJsonFile(rows []reader.RowInfo) {
	strAction := ""
	var actIdx = make(map[string]actionIndex)

//	fmt.Println("***************all rows ***********************************")
//	fmt.Printf("%+v\n", rows)
	// 全行ループ
	for _, row := range rows {
		var actwk = actionIndex{}
		var itemwk = itemIndex{}
		var valwk = validateIndex{}

/*		if strAction != row.ActionName {
			actwk = actionIndex{
				action: row.ActionName,
			}
		} else {
			actwk = actIdx[row.ActionName]
		}*/

		// チェック内容を確認
		checkstr := row.Required +
					row.Validate + 
					row.Classifi +
					row.Length +
					row.Range +
					row.Month +
					row.StrikePrice +
					row.Price +
					row.HashChk
		
		valwk.checks = checkstr

		// 各チェックのオプション確認
		// 属性チェックをする場合
		if row.Validate == "1" {
			valwk.attribute = row.ValidPtn

			// 属性チェックが３（日付）指定の場合はフォーマットも保存
			if row.ValidPtn == "3" {
				valwk.attributePattern = row.FormatPtn
			}
		}

		// 区分値チェックをする場合
		if row.Classifi == "1" {
			valwk.classification = row.ClassifiStr
		}

		// 範囲チェックをする場合
		if row.Range == "1" {
			valwk.rangeMin = row.RangeMin
			valwk.rangeMax = row.RangeMax
		}

		// 桁数チェックをする場合
		if row.Length == "1" {
			valwk.length = row.LengthVal
		}

		// 項目情報にチェック情報を格納
		itemwk.valIndex = valwk

		// アクションの切り替わりチェック
		if strAction == row.ActionName {
			actwk = actIdx[strAction]
		} else {
			strAction = row.ActionName
			actwk = actionIndex{ action: strAction }
			actwk.itemIdx = []itemIndex{}
		}

		// アクションとアイテム情報の保管
		itemwk.item = row.ItemName
		actwk.itemIdx = append(actwk.itemIdx, itemwk)
		actIdx[strAction] = actwk
	}

	// JSON文字列生成
	jsstrall := ""
	for _, actIdxRow := range actIdx {
		if jsstrall != "" {
			jsstrall += ","
		}
		jsstrall += "\"" + actIdxRow.action + "\": {"
		
		itemstr := ""
//		fmt.Println("itemInfo**********************************")
//		fmt.Printf("[%s]:%+v\n", actIdxRow.action, actIdxRow.itemIdx)
		for _, itemrow := range actIdxRow.itemIdx {

			if itemstr != "" {
				itemstr += ","
			}
			itemstr += "\"" + itemrow.item + "\": {"
			
			checkstr := ""
			if strings.Index(itemrow.valIndex.checks, "1") >= 0 {
				checkstr = "\"checks\": \"" + itemrow.valIndex.checks + "\""
				
				if itemrow.valIndex.attribute != "" {
					checkstr += ", \"attribute\": \"" + itemrow.valIndex.attribute + "\""
				}
				if itemrow.valIndex.attribute == "3" {
					checkstr += ", \"attribute_format\": \"" + itemrow.valIndex.attributePattern + "\""
				}
				if itemrow.valIndex.classification != "" {
					checkstr += ", \"classification\": \"" + itemrow.valIndex.classification + "\""
				}
				if itemrow.valIndex.rangeMin != "" {
					checkstr += ", \"range_min\": \"" + itemrow.valIndex.rangeMin + "\""
				}
				if itemrow.valIndex.rangeMax != "" {
					checkstr += ", \"range_max\": \"" + itemrow.valIndex.rangeMax + "\""
				}
				if itemrow.valIndex.length != "" {
					checkstr += ", \"length\": \"" + itemrow.valIndex.length + "\""
				}
			} else {
				log.Fatal("CheckPattern Error[%s]*************", itemrow.valIndex.checks)
			}
			itemstr += checkstr + "}"
		}
		jsstrall += itemstr + "}"			
	}

	fmt.Println("{" + jsstrall + "}")


}
/*
// CreateActionGroup アクション単位でグルーピングする
func CreateActionGroup(rows []*xlsx.Row) map[string][]*xlsx.Row {
	var actMap = make(map[string][]*xlsx.Row)
	actRow := []*xlsx.Row{}

	// 全行ループ
	for idx, rowData := range rows {

		// 先頭行は無視（ヘッダレコードのため）
		if idx == 0 || idx == 1 {
			continue
		}
		// アクション指定がある場合
		if rowData.Cells[10].Value != "" {
			//			actRowWk, keyFlg := actMap[rowData.Cells[10].Value]

			/*			// 既にデータがある場合は追加
						if keyFlg {
							actRow = actRowWk
						}
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
	ioutil.WriteFile("C:/work/validatepattern.json", byteJsonBody, os.ModePerm)
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
}*/
