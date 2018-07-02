package models

import (
	"github.com/xaionaro-go/extime"
)

//go:generate reform

//reform:vacancies
type vacancy struct {
	Id        int          `reform:"id,pk"`
	Name      string       `reform:"name"`
	CreatedAt *extime.Time `reform:"created_at"`
	ClosedAt  *extime.Time `reform:"closed_at"`
}
