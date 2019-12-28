package aozora

import (
	"io/ioutil"

	"github.com/jszwec/csvutil"
	"github.com/rakyll/statik/fs"

	_ "github.com/tgfjt/aozora/statik"
)

//go:generate statik -src=data

/*
WorkExpanded means 「公開中　作家別作品一覧拡充版：全て(CSV形式、UTF-8、zip圧縮）」
  see also: https://www.aozora.gr.jp/index_pages/person_all.html
  remind: remove BOM and change CRLF to LF.
*/
type WorkExpanded struct {
	CardID           string `csv:"作品ID"`
	CardTitle        string `csv:"作品名"`
	CardTitleReading string `csv:"作品名読み"`
	// PropName string `csv:"ソート用読み"`
	// PropName string `csv:"副題"`
	// PropName string `csv:"副題読み"`
	CardOriginalTitle string `csv:"原題"`
	// PropName string `csv:"初出"`
	// PropName string `csv:"分類番号"`
	// PropName string `csv:"文字遣い種別"`
	// PropName string `csv:"作品著作権フラグ"`
	CardPublishedAt string `csv:"公開日"`
	CardUpdatedAt   string `csv:"最終更新日"`
	CardURL         string `csv:"図書カードURL"`
	// PersonID        string `csv:"人物ID"`
	PersonLastName         string `csv:"姓"`
	PersonFirstName        string `csv:"名"`
	PersonLastNameReading  string `csv:"姓読み"`
	PersonFirstNameReading string `csv:"名読み"`
	// PropName string `csv:"姓読みソート用"`
	// PropName string `csv:"名読みソート用"`
	PersonLastNameEn  string `csv:"姓ローマ字"`
	PersonFirstNameEn string `csv:"名ローマ字"`
	// PropName string `csv:"役割フラグ"`
	PersonBornOn                           string `csv:"生年月日"`
	PersonDiedOn                           string `csv:"没年月日"`
	CopyrightFlag                          string `csv:"人物著作権フラグ"`
	OriginBook1Title                       string `csv:"底本名1"`
	OriginBook1PublisherName               string `csv:"底本出版社名1"`
	OriginBook1FirstEdtionIssuedOn         string `csv:"底本初版発行年1"`
	OriginBook1InputEdtion                 string `csv:"入力に使用した版1"`
	OriginBook1ProofEdtion                 string `csv:"校正に使用した版1"`
	ParentOfOriginBook1Title               string `csv:"底本の親本名1"`
	ParentOfOriginBook1PublisherName       string `csv:"底本の親本出版社名1"`
	ParentOfOriginBook1FirstEdtionIssuedOn string `csv:"底本の親本初版発行年1"`
	OriginBook2Title                       string `csv:"底本名2"`
	OriginBook2PublisherName               string `csv:"底本出版社名2"`
	OriginBook2FirstEdtionIssuedOn         string `csv:"底本初版発行年2"`
	OriginBook2InputEdtion                 string `csv:"入力に使用した版2"`
	OriginBook2ProofEdtion                 string `csv:"校正に使用した版2"`
	ParentOfOriginBook2Title               string `csv:"底本の親本名2"`
	ParentOfOriginBook2PublisherName       string `csv:"底本の親本出版社名2"`
	ParentOfOriginBook2FirstEdtionIssuedOn string `csv:"底本の親本初版発行年2"`
	InputEditorName                        string `csv:"入力者"`
	ProofEditorName                        string `csv:"校正者"`
	// FileURL string `csv:"テキストファイルURL"`
	// PropName string `csv:"テキストファイル最終更新日"`
	// PropName string `csv:"テキストファイル符号化方式"`
	// PropName string `csv:"テキストファイル文字集合"`
	// PropName string `csv:"テキストファイル修正回数"`
	XHTMLURL string `csv:"XHTML/HTMLファイルURL"`
	// PropName string `csv:"XHTML/HTMLファイル最終更新日"`
	// PropName string `csv:"XHTML/HTMLファイル符号化方式"`
	// PropName string `csv:"XHTML/HTMLファイル文字集合"`
	// PropName string `csv:"XHTML/HTMLファイル修正回数"`
}

// OriginBook means 底本
type OriginBook struct {
	Title               string
	PublisherName       string
	FirstEdtionIssuedOn string
	InputEdtion         string
	ProofEdtion         string
	ParentBook          ParentBook
}

// ParentBook means 親本
type ParentBook struct {
	Title               string
	PublisherName       string
	FirstEdtionIssuedOn string
}

// OriginBook return origin book with its parent book
func (w WorkExpanded) OriginBook() []OriginBook {
	return []OriginBook{
		OriginBook{
			w.OriginBook1Title,
			w.OriginBook1PublisherName,
			w.OriginBook1FirstEdtionIssuedOn,
			w.OriginBook1InputEdtion,
			w.OriginBook1ProofEdtion,
			ParentBook{
				w.ParentOfOriginBook1Title,
				w.ParentOfOriginBook1PublisherName,
				w.ParentOfOriginBook1FirstEdtionIssuedOn,
			},
		},
		OriginBook{
			w.OriginBook2Title,
			w.OriginBook2PublisherName,
			w.OriginBook2FirstEdtionIssuedOn,
			w.OriginBook2InputEdtion,
			w.OriginBook2ProofEdtion,
			ParentBook{
				w.ParentOfOriginBook2Title,
				w.ParentOfOriginBook2PublisherName,
				w.ParentOfOriginBook2FirstEdtionIssuedOn,
			},
		},
	}
}

// ListOfWorks return a expanded version list of Works.
func ListOfWorks() ([]WorkExpanded, error) {
	var err error
	var listOfWorks []WorkExpanded

	statikFS, err := fs.New()
	if err != nil {
		return nil, err
	}

	// https://www.aozora.gr.jp/index_pages/person_all.html
	// 「公開中　作家別作品一覧拡充版：全て(CSV形式、UTF-8、zip圧縮）」をダウンロード
	f, err := statikFS.Open("/list_person_all_extended_utf8.csv")
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = csvutil.Unmarshal(b, &listOfWorks)
	if err != nil {
		return nil, err
	}

	return listOfWorks, nil
}
