package model


type Coal struct{
	Id int `gorm:"primary_key"`
	Name string `gorm:"type:varchar(128);not null;"`
	Visibility int `gorm:"not null;"`
}

type Coal_analysis_report struct {
	Id int `gorm:"primary_key"`
	Coal_id int `gorm:"not null;"`
	Analysis_category_id int `gorm:"not null;"`
	Date string `gorm:"varchar(128)"`
	Active bool `gorm:"not null;"`
}

type Analysis_category struct {
	Id int `gorm:"primary_key"`
	Name string `gorm:"varchar(128);not null;"`
	Visibility int `gorm:"not null;"`
}
type Report struct{
	Coal_analysis_report_id int `json:"coal_analysis_report_id"`
	Coal_type_attribute_list_id int `json:"coal_type_attribute_list_id"`
	Detail string `json:"detail"`
	Visibility int `json:"visibility"`
}

type Receiver struct {
	Categoryid int `json:"categoryid"`
	Req [] Report `json:"req"`
}

type Report_content struct {
	Id int `json:"id"`
	Val string `json:"val"`
}