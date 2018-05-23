package models

//go:generate reform

//reform:applications
type application struct {
	Id         int     `reform:"id,pk"`
	UserId     int     `reform:"user_id"`
	VacancyId  int     `reform:"vacancy_id"`
	Salary     int     `reform:"salary"`
	Employment float32 `reform:"employment"`
}
