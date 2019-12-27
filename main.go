package main

import (
	"fmt"
	"io/ioutil"

	"github.com/jszwec/csvutil"
)

type Data struct {
	CardID   string `csv:"作品ID"`
	CardName string `csv:"作品名"`
	CardYomi string `csv:"作品名読み"`
	// PropName string `csv:"ソート用読み"`
	// PropName string `csv:"副題"`
	// PropName string `csv:"副題読み"`
	// PropName string `csv:"原題"`
	// PropName string `csv:"初出"`
	// PropName string `csv:"分類番号"`
	// PropName string `csv:"文字遣い種別"`
	// PropName string `csv:"作品著作権フラグ"`
	// PropName string `csv:"公開日"`
	// PropName string `csv:"最終更新日"`
	CardURL string `csv:"図書カードURL"`
	// PersonID        string `csv:"人物ID"`
	PersonLastName  string `csv:"姓"`
	PersonFirstName string `csv:"名"`
	// PropName string `csv:"姓読み"`
	// PropName string `csv:"名読み"`
	// PropName string `csv:"姓読みソート用"`
	// PropName string `csv:"名読みソート用"`
	// PropName string `csv:"姓ローマ字"`
	// PropName string `csv:"名ローマ字"`
	// PropName string `csv:"役割フラグ"`
	// PropName string `csv:"生年月日"`
	// PropName string `csv:"没年月日"`
	// PropName string `csv:"人物著作権フラグ"`
	// PropName string `csv:"底本名1"`
	// PropName string `csv:"底本出版社名1"`
	// PropName string `csv:"底本初版発行年1"`
	// PropName string `csv:"入力に使用した版1"`
	// PropName string `csv:"校正に使用した版1"`
	// PropName string `csv:"底本の親本名1"`
	// PropName string `csv:"底本の親本出版社名1"`
	// PropName string `csv:"底本の親本初版発行年1"`
	// PropName string `csv:"底本名2"`
	// PropName string `csv:"底本出版社名2"`
	// PropName string `csv:"底本初版発行年2"`
	// PropName string `csv:"入力に使用した版2"`
	// PropName string `csv:"校正に使用した版2"`
	// PropName string `csv:"底本の親本名2"`
	// PropName string `csv:"底本の親本出版社名2"`
	// PropName string `csv:"底本の親本初版発行年2"`
	// PropName string `csv:"入力者"`
	// PropName string `csv:"校正者"`
	FileURL string `csv:"テキストファイルURL"`
	// PropName string `csv:"テキストファイル最終更新日"`
	// PropName string `csv:"テキストファイル符号化方式"`
	// PropName string `csv:"テキストファイル文字集合"`
	// PropName string `csv:"テキストファイル修正回数"`
	// PropName string `csv:"XHTML/HTMLファイルURL"`
	// PropName string `csv:"XHTML/HTMLファイル最終更新日"`
	// PropName string `csv:"XHTML/HTMLファイル符号化方式"`
	// PropName string `csv:"XHTML/HTMLファイル文字集合"`
	// PropName string `csv:"XHTML/HTMLファイル修正回数"`
}

func main() {
	// https://www.aozora.gr.jp/index_pages/person_all.html
	// 「公開中　作家別作品一覧拡充版：全て(CSV形式、UTF-8、zip圧縮）」をダウンロード
	bs, _ := ioutil.ReadFile("./src/list_person_all_extended_utf8.csv")

	// log.Println(string(bs))

	var data []Data
	if err := csvutil.Unmarshal(bs, &data); err != nil {
		fmt.Println("error:", err)
	}

	for _, d := range data {
		fmt.Println(d)
	}
}
