package core

// 全API共通
//
// https://www.e-stat.go.jp/api/api-info/e-stat-manual3-0#api_3_1
type CommonParams struct {
	// # アプリケーションID
	//
	// 	取得したアプリケーションIDを指定して下さい。
	AppID string `url:"appId" json:"appId"`

	// # 言語
	//
	// 取得するデータの言語を 以下のいずれかを指定して下さい。
	//
	// ・J：日本語 (省略値)
	//
	// ・E：英語
	Lang string `url:"lang,omitempty" json:"lang,omitempty" xml:"LANG,omitempty"`
}

type TabulatCondition struct {
	LevelTab    string `url:"lvTab,omitempty" json:"lvTab,omitempty" xml:"LEVEL_TAB_COND"`
	CodeTab     string `url:"cdTab,omitempty" json:"cdTab,omitempty" xml:"CODE_TAB_SELECT"`
	CodeTabFrom string `url:"cdTabFrom,omitempty" json:"cdTabFrom,omitempty" xml:"CODE_TAB_FROM"`
	CodeTabTo   string `url:"cdTabTo,omitempty" json:"cdTabTo,omitempty" xml:"CODE_TAB_TO"`
}

type TimeCondition struct {
	LevelTime    string `url:"lvTime,omitempty" json:"lvTime,omitempty" xml:"LEVEL_TIME_COND"`
	CodeTime     string `url:"cdTime,omitempty" json:"cdTime,omitempty" xml:"CODE_TIME_SELECT"`
	CodeTimeFrom string `url:"cdTimeFrom,omitempty" json:"cdTimeFrom,omitempty" xml:"CODE_TIME_FROM"`
	CodeTimeTo   string `url:"cdTimeTo,omitempty" json:"cdTimeTo,omitempty" xml:"CODE_TIME_TO"`
}

type AreaCondition struct {
	LevelArea    string `url:"lvArea,omitempty" json:"lvArea,omitempty" xml:"LEVEL_AREA_COND"`
	CodeArea     string `url:"cdArea,omitempty" json:"cdArea,omitempty" xml:"CODE_AREA_SELECT"`
	CodeAreaFrom string `url:"cdAreaFrom,omitempty" json:"cdAreaFrom,omitempty" xml:"CODE_AREA_FROM"`
	CodeAreaTo   string `url:"cdAreaTo,omitempty" json:"cdAreaTo,omitempty" xml:"CODE_AREA_TO"`
}

type CategoryCondition struct {
	LevelCat01    string `url:"lvCat01,omitempty" json:"lvCat01,omitempty" xml:"LEVEL_CAT01_COND"`
	CodeCat01     string `url:"cdCat01,omitempty" json:"cdCat01,omitempty" xml:"CODE_CAT01_SELECT"`
	CodeCat01From string `url:"cdCat01From,omitempty" json:"cdCat01From,omitempty" xml:"CODE_CAT01_FROM"`
	CodeCat01To   string `url:"cdCat01To,omitempty" json:"cdCat01To,omitempty" xml:"CODE_CAT01_TO"`

	LevelCat02    string `url:"lvCat02,omitempty" json:"lvCat02,omitempty" xml:"LEVEL_CAT02_COND"`
	CodeCat02     string `url:"cdCat02,omitempty" json:"cdCat02,omitempty" xml:"CODE_CAT02_SELECT"`
	CodeCat02From string `url:"cdCat02From,omitempty" json:"cdCat02From,omitempty" xml:"CODE_CAT02_FROM"`
	CodeCat02To   string `url:"cdCat02To,omitempty" json:"cdCat02To,omitempty" xml:"CODE_CAT02_TO"`

	LevelCat03    string `url:"lvCat03,omitempty" json:"lvCat03,omitempty" xml:"LEVEL_CAT03_COND"`
	CodeCat03     string `url:"cdCat03,omitempty" json:"cdCat03,omitempty" xml:"CODE_CAT03_SELECT"`
	CodeCat03From string `url:"cdCat03From,omitempty" json:"cdCat03From,omitempty" xml:"CODE_CAT03_FROM"`
	CodeCat03To   string `url:"cdCat03To,omitempty" json:"cdCat03To,omitempty" xml:"CODE_CAT03_TO"`

	LevelCat04    string `url:"lvCat04,omitempty" json:"lvCat04,omitempty" xml:"LEVEL_CAT04_COND"`
	CodeCat04     string `url:"cdCat04,omitempty" json:"cdCat04,omitempty" xml:"CODE_CAT04_SELECT"`
	CodeCat04From string `url:"cdCat04From,omitempty" json:"cdCat04From,omitempty" xml:"CODE_CAT04_FROM"`
	CodeCat04To   string `url:"cdCat04To,omitempty" json:"cdCat04To,omitempty" xml:"CODE_CAT04_TO"`

	LevelCat05    string `url:"lvCat05,omitempty" json:"lvCat05,omitempty" xml:"LEVEL_CAT05_COND"`
	CodeCat05     string `url:"cdCat05,omitempty" json:"cdCat05,omitempty" xml:"CODE_CAT05_SELECT"`
	CodeCat05From string `url:"cdCat05From,omitempty" json:"cdCat05From,omitempty" xml:"CODE_CAT05_FROM"`
	CodeCat05To   string `url:"cdCat05To,omitempty" json:"cdCat05To,omitempty" xml:"CODE_CAT05_TO"`

	LevelCat06    string `url:"lvCat06,omitempty" json:"lvCat06,omitempty" xml:"LEVEL_CAT06_COND"`
	CodeCat06     string `url:"cdCat06,omitempty" json:"cdCat06,omitempty" xml:"CODE_CAT06_SELECT"`
	CodeCat06From string `url:"cdCat06From,omitempty" json:"cdCat06From,omitempty" xml:"CODE_CAT06_FROM"`
	CodeCat06To   string `url:"cdCat06To,omitempty" json:"cdCat06To,omitempty" xml:"CODE_CAT06_TO"`

	LevelCat07    string `url:"lvCat07,omitempty" json:"lvCat07,omitempty" xml:"LEVEL_CAT07_COND"`
	CodeCat07     string `url:"cdCat07,omitempty" json:"cdCat07,omitempty" xml:"CODE_CAT07_SELECT"`
	CodeCat07From string `url:"cdCat07From,omitempty" json:"cdCat07From,omitempty" xml:"CODE_CAT07_FROM"`
	CodeCat07To   string `url:"cdCat07To,omitempty" json:"cdCat07To,omitempty" xml:"CODE_CAT07_TO"`

	LevelCat08    string `url:"lvCat08,omitempty" json:"lvCat08,omitempty" xml:"LEVEL_CAT08_COND"`
	CodeCat08     string `url:"cdCat08,omitempty" json:"cdCat08,omitempty" xml:"CODE_CAT08_SELECT"`
	CodeCat08From string `url:"cdCat08From,omitempty" json:"cdCat08From,omitempty" xml:"CODE_CAT08_FROM"`
	CodeCat08To   string `url:"cdCat08To,omitempty" json:"cdCat08To,omitempty" xml:"CODE_CAT08_TO"`

	LevelCat09    string `url:"lvCat09,omitempty" json:"lvCat09,omitempty" xml:"LEVEL_CAT09_COND"`
	CodeCat09     string `url:"cdCat09,omitempty" json:"cdCat09,omitempty" xml:"CODE_CAT09_SELECT"`
	CodeCat09From string `url:"cdCat09From,omitempty" json:"cdCat09From,omitempty" xml:"CODE_CAT09_FROM"`
	CodeCat09To   string `url:"cdCat09To,omitempty" json:"cdCat09To,omitempty" xml:"CODE_CAT09_TO"`

	LevelCat10    string `url:"lvCat10,omitempty" json:"lvCat10,omitempty" xml:"LEVEL_CAT10_COND"`
	CodeCat10     string `url:"cdCat10,omitempty" json:"cdCat10,omitempty" xml:"CODE_CAT10_SELECT"`
	CodeCat10From string `url:"cdCat10From,omitempty" json:"cdCat10From,omitempty" xml:"CODE_CAT10_FROM"`
	CodeCat10To   string `url:"cdCat10To,omitempty" json:"cdCat10To,omitempty" xml:"CODE_CAT10_TO"`

	LevelCat11    string `url:"lvCat11,omitempty" json:"lvCat11,omitempty" xml:"LEVEL_CAT11_COND"`
	CodeCat11     string `url:"cdCat11,omitempty" json:"cdCat11,omitempty" xml:"CODE_CAT11_SELECT"`
	CodeCat11From string `url:"cdCat11From,omitempty" json:"cdCat11From,omitempty" xml:"CODE_CAT11_FROM"`
	CodeCat11To   string `url:"cdCat11To,omitempty" json:"cdCat11To,omitempty" xml:"CODE_CAT11_TO"`

	LevelCat12    string `url:"lvCat12,omitempty" json:"lvCat12,omitempty" xml:"LEVEL_CAT12_COND"`
	CodeCat12     string `url:"cdCat12,omitempty" json:"cdCat12,omitempty" xml:"CODE_CAT12_SELECT"`
	CodeCat12From string `url:"cdCat12From,omitempty" json:"cdCat12From,omitempty" xml:"CODE_CAT12_FROM"`
	CodeCat12To   string `url:"cdCat12To,omitempty" json:"cdCat12To,omitempty" xml:"CODE_CAT12_TO"`

	LevelCat13    string `url:"lvCat13,omitempty" json:"lvCat13,omitempty" xml:"LEVEL_CAT13_COND"`
	CodeCat13     string `url:"cdCat13,omitempty" json:"cdCat13,omitempty" xml:"CODE_CAT13_SELECT"`
	CodeCat13From string `url:"cdCat13From,omitempty" json:"cdCat13From,omitempty" xml:"CODE_CAT13_FROM"`
	CodeCat13To   string `url:"cdCat13To,omitempty" json:"cdCat13To,omitempty" xml:"CODE_CAT13_TO"`

	LevelCat14    string `url:"lvCat14,omitempty" json:"lvCat14,omitempty" xml:"LEVEL_CAT14_COND"`
	CodeCat14     string `url:"cdCat14,omitempty" json:"cdCat14,omitempty" xml:"CODE_CAT14_SELECT"`
	CodeCat14From string `url:"cdCat14From,omitempty" json:"cdCat14From,omitempty" xml:"CODE_CAT14_FROM"`
	CodeCat14To   string `url:"cdCat14To,omitempty" json:"cdCat14To,omitempty" xml:"CODE_CAT14_TO"`

	LevelCat15    string `url:"lvCat15,omitempty" json:"lvCat15,omitempty" xml:"LEVEL_CAT15_COND"`
	CodeCat15     string `url:"cdCat15,omitempty" json:"cdCat15,omitempty" xml:"CODE_CAT15_SELECT"`
	CodeCat15From string `url:"cdCat15From,omitempty" json:"cdCat15From,omitempty" xml:"CODE_CAT15_FROM"`
	CodeCat15To   string `url:"cdCat15To,omitempty" json:"cdCat15To,omitempty" xml:"CODE_CAT15_TO"`
}

type NarrowingConditon struct {
	TabulatCondition
	TimeCondition
	AreaCondition
	CategoryCondition
}
