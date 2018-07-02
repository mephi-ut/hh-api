package models

import (
	"github.com/xaionaro-go/extime"
)

//go:generate reform

//reform:applications
type application struct {
	Id         int          `reform:"id,pk"`
	UserId     int          `reform:"user_id"`
	VacancyId  int          `reform:"vacancy_id"`
	Salary     *int         `reform:"salary"`
	Employment float32      `reform:"employment"`
	Answers    string       `reform:"answers" sql_size:"1048576"`
	Email      string       `reform:"email" sql_size:"255"`
	Phone      string       `reform:"phone" sql_size:"255"`
	CreatedAt  *extime.Time `reform:"created_at"`
}

func NewApplication() *application {
	return &application{}
}
