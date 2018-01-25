package reader

import (
	"strings"
	"fmt"
	"log"

	"github.com/tealeg/xlsx"
	"github.com/ahmetb/go-linq"

	_ "github.com/ryu0322/validatejson/utility"
)


// 列定義
const (
	_ = iota
	_
	_
	_
	_
	_
	_
	_
	IMPLEMENTATION	// 実装
	_
	ACTIONNAME		// アクション名
	ITEMNAME		// 項目名
	REQUIRED		// 必須チェック
	VALIDATE		// 属性チェック
	CLASSIFI		// 区分値チェック
	LENGTH			// 桁数チェック
	RANGE			// 範囲チェック
	MONTH			// 限月チェック
	STRIKEPRICE		// 権利行使価格チェック
	PRICE			// 価格入力チェック
	HASHCHK			// 調整用ポジションのハッシュ値チェック
	VALIDPTN		// 属性値チェックパターン
	FORMATPTN		// 属性値フォーマット
	CLASSIFISTR		// 区分値項目名
	RANGEMIN		// 範囲最小値
	RANGEMAX		// 範囲最大値
	LENGTHVAL		// 許容桁数
)

type RowInfo struct {
	ActionName string	// アクション名
	ItemName string		// 項目名
	Required string		// 必須チェック
	Validate string		// 属性チェック
	Classifi string		// 区分値チェック
	Length string		// 桁数チェック
	Range string		// 範囲チェック
	Month string		// 限月チェック
	StrikePrice string	// 権利行使価格チェック
	Price string		// 価格入力チェック
	HashChk string		// 調整用ポジションのハッシュチェック
	ValidPtn string		// 属性値チェックパターン
	FormatPtn string	// 属性値フォーマット
	ClassifiStr string	// 区分値項目名
	RangeMin string		// 範囲最小値
	RangeMax string		// 範囲最大値
	LengthVal string	// 許容桁数
}
func Reader(inPath string) []RowInfo {
	xlFile, err := xlsx.OpenFile(inPath)
	//_, err := xlsx.OpenFile(inPath)
	if err != nil {
		log.Fatal(err)
	}

	allRow := xlFile.Sheets[0].Rows

	fmt.Println("all row gets!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

	var retRows = []RowInfo{}

	// 行ループ
	for idx, rowData := range allRow {
		// 1、2行目は飛ばす
		if idx == 0 || idx == 1 {
			continue
		}


		// 実装列が×、もしくはアクションが未設定の物は無視
		if rowData.Cells[IMPLEMENTATION].Value == "×" ||
			rowData.Cells[ACTIONNAME].Value == "" ||
			rowData.Cells[ACTIONNAME].Value == "-" {
			continue	
		}
		
		// 行情報に値をセット
		var rowinfo = RowInfo{
			ActionName: strings.Replace(rowData.Cells[ACTIONNAME].Value, "-", "", -1),
			ItemName: strings.Replace(rowData.Cells[ITEMNAME].Value, "-", "", -1),
			Required: strings.Replace(rowData.Cells[REQUIRED].Value, "-", "", -1),
			Validate: strings.Replace(rowData.Cells[VALIDATE].Value, "-", "", -1),
			Classifi: strings.Replace(rowData.Cells[CLASSIFI].Value, "-", "", -1),
			Length: strings.Replace(rowData.Cells[LENGTH].Value, "-", "", -1),
			Range: strings.Replace(rowData.Cells[RANGE].Value, "-", "", -1),
			Month: strings.Replace(rowData.Cells[MONTH].Value, "-", "", -1),
			StrikePrice: strings.Replace(rowData.Cells[STRIKEPRICE].Value, "-", "", -1),
			Price: strings.Replace(rowData.Cells[PRICE].Value, "-", "", -1),
			HashChk: strings.Replace(rowData.Cells[HASHCHK].Value, "-", "", -1),
			ValidPtn: strings.Replace(rowData.Cells[VALIDPTN].Value, "-", "", -1),
			FormatPtn: strings.Replace(rowData.Cells[FORMATPTN].Value, "-", "", -1),
			ClassifiStr: strings.Replace(rowData.Cells[CLASSIFISTR].Value, "-", "", -1),
			RangeMin: strings.Replace(rowData.Cells[RANGEMIN].Value, "-", "", -1),
			RangeMax: strings.Replace(rowData.Cells[RANGEMAX].Value, "-", "", -1),
			LengthVal: strings.Replace(rowData.Cells[LENGTHVAL].Value, "-", "", -1),
		}

		retRows = append(retRows, rowinfo)
	}

	// LINQで並べ替えしておく
	linq.From(retRows).
			OrderBy(
				func (item interface{}) interface{} {return item.(RowInfo).ActionName},
			).
			ToSlice(&retRows)
	return retRows
}
