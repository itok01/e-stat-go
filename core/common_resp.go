package core

import "time"

type ResponseResult struct {
	Status   int       `xml:"STATUS"`
	ErrorMsg string    `xml:"ERROR_MSG"`
	Date     time.Time `xml:"DATE"`
}

type ResultInf struct {
	FromNumber int `xml:"FROM_NUMBER"`
	ToNumber   int `xml:"TO_NUMBER"`
	NextKey    int `xml:"NEXT_KEY,omitempty"`
}

type StatName struct {
	Code string `xml:"code,attr"`
	Name string `xml:",innerxml"`
}

type GovOrg struct {
	Code string `xml:"code,attr"`
	Name string `xml:",innerxml"`
}

type MainCategory struct {
	Code string `xml:"code,attr"`
	Name string `xml:",innerxml"`
}

type SubCategory struct {
	Code string `xml:"code,attr"`
	Name string `xml:",innerxml"`
}

type Organization struct {
	Code string `xml:"code,attr"`
	Name string `xml:",innerxml"`
}

type Title struct {
	Number string `xml:"no,attr"`
	Name   string `xml:",innerxml"`
}

type StatisticsNameSpec struct {
	TabulationCategory     string `xml:"TABULATION_CATEGORY"`
	TabulationSubCategory1 string `xml:"TABULATION_SUB_CATEGORY1,omitempty"`
	TabulationSubCategory2 string `xml:"TABULATION_SUB_CATEGORY2,omitempty"`
	TabulationSubCategory3 string `xml:"TABULATION_SUB_CATEGORY3,omitempty"`
	TabulationSubCategory4 string `xml:"TABULATION_SUB_CATEGORY4,omitempty"`
	TabulationSubCategory5 string `xml:"TABULATION_SUB_CATEGORY5,omitempty"`
}

type Description struct {
	TabulationCategoryExplanation     string `xml:"TABULATION_CATEGORY_EXPLANATION"`
	TabulationSubCategoryExplanation1 string `xml:"TABULATION_SUB_CATEGORY_EXPLANATION1,omitempty"`
	TabulationSubCategoryExplanation2 string `xml:"TABULATION_SUB_CATEGORY_EXPLANATION2,omitempty"`
	TabulationSubCategoryExplanation3 string `xml:"TABULATION_SUB_CATEGORY_EXPLANATION3,omitempty"`
	TabulationSubCategoryExplanation4 string `xml:"TABULATION_SUB_CATEGORY_EXPLANATION4,omitempty"`
	TabulationSubCategoryExplanation5 string `xml:"TABULATION_SUB_CATEGORY_EXPLANATION5,omitempty"`
}

type TitleSpec struct {
	TableCategory     string `xml:"TABLE_CATEGORY"`
	TableName         string `xml:"TABLE_NAME"`
	TableExplanation  string `xml:"TABLE_EXPLANATION,omitempty"`
	TableSubCategory1 string `xml:"TABLE_SUB_CATEGORY1,omitempty"`
	TableSubCategory2 string `xml:"TABLE_SUB_CATEGORY2,omitempty"`
	TableSubCategory3 string `xml:"TABLE_SUB_CATEGORY3,omitempty"`
}

type TableInf struct {
	ID                 string             `xml:"id,attr"`
	StatName           StatName           `xml:"STAT_NAME"`
	GovOrg             GovOrg             `xml:"GOV_ORG"`
	StatisticsName     string             `xml:"STATISTICS_NAME"`
	Title              Title              `xml:"TITLE"`
	Cycle              string             `xml:"CYCLE"`
	SurveyDate         string             `xml:"SURVEY_DATE"`
	OpenDate           string             `xml:"OPEN_DATE"`
	SmallArea          int                `xml:"SMALL_AREA"`
	CollectArea        string             `xml:"COLLECT_AREA"`
	MainCategory       MainCategory       `xml:"MAIN_CATEGORY"`
	SubCategory        SubCategory        `xml:"SUB_CATEGORY"`
	OverallTotalNumber int                `xml:"OVERALL_TOTAL_NUMBER"`
	UpdatedDate        string             `xml:"UPDATED_DATE"`
	StatisticsNameSpec StatisticsNameSpec `xml:"STATISTICS_NAME_SPEC"`
	Description        Description        `xml:"DESCRIPTION"`
	TitleSpec          TitleSpec          `xml:"TITLE_SPEC"`
}

type ClassObjExplanation struct {
	ID          string `xml:"id,attr"`
	Explanation string `xml:",innerxml"`
}

type ClassObjClass struct {
	Code       string `xml:"code,attr"`
	Name       string `xml:"name,attr"`
	Level      string `xml:"level,attr"`
	Unit       string `xml:"unit,attr,omitempty"`
	ParentCode string `xml:"parentCode,attr,omitempty"`
	AddInf     string `xml:"addInf,attr,omitempty"`
}

type ClassObj struct {
	ID          string                `xml:"id,attr"`
	Name        string                `xml:"name,attr"`
	Description string                `xml:"description,attr"`
	Class       []ClassObjClass       `xml:"CLASS"`
	Explanation []ClassObjExplanation `xml:"EXPLANATION"`
}

type ClassInf struct {
	ClassObj []ClassObj `xml:"CLASS_OBJ"`
}

type DataInfNote struct {
	Char string `xml:"char,attr"`
	Note string `xml:",innerxml"`
}

type DataInfAnnotation struct {
	Target     string `xml:"annotation,attr"`
	Annotation string `xml:",innerxml"`
}

type DataInfValue struct {
	Tab        string `xml:"tab,attr"`
	Cat01      string `xml:"cat01,attr"`
	Cat02      string `xml:"cat02,attr"`
	Cat03      string `xml:"cat03,attr"`
	Cat04      string `xml:"cat04,attr"`
	Cat05      string `xml:"cat05,attr"`
	Cat06      string `xml:"cat06,attr"`
	Cat07      string `xml:"cat07,attr"`
	Cat08      string `xml:"cat08,attr"`
	Cat09      string `xml:"cat09,attr"`
	Cat10      string `xml:"cat10,attr"`
	Cat11      string `xml:"cat11,attr"`
	Cat12      string `xml:"cat12,attr"`
	Cat13      string `xml:"cat13,attr"`
	Cat14      string `xml:"cat14,attr"`
	Cat15      string `xml:"cat15,attr"`
	Area       string `xml:"area,attr,omitempty"`
	Time       string `xml:"time,attr"`
	Unit       string `xml:"unit,attr"`
	Annotation string `xml:"annotation,attr,omitempty"`
	Value      string `xml:",innerxml"`
}

type DataInf struct {
	Note       []DataInfNote       `xml:"NOTE,omitempty"`
	Annotation []DataInfAnnotation `xml:"ANNOTATION,omitempty"`
	Value      []DataInfValue      `xml:"VALUE,omitempty"`
}
