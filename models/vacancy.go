package models

//go:generate reform

//reform:vacancies
type vacancy struct {
	Id int `reform:"id,pk"`
}
