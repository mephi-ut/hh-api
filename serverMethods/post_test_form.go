package serverMethods

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/mephi-ut/hh-api/jwt"
	"github.com/mephi-ut/hh-api/models"
	"github.com/mephi-ut/hh-api/serverMethods/helpers"
	mw "github.com/mephi-ut/hh-api/serverMethods/middleware"
	"github.com/xaionaro-go/extime"
)

func PostTestForm(c *gin.Context) {
	vacancyName := c.Param("vacancy_name")
	if vacancyName == "" {
		jsonError(c, fmt.Errorf("vacancy name is not set"))
		return
	}

	vacancy, err := models.Vacancy.Where("closed_at IS NULL").First(models.VacancyF{Name: strings.ToLower(vacancyName)})
	if err != nil {
		jsonError(c, err)
		return
	}

	data := map[string]string{}
	if err := c.ShouldBindJSON(&data); err != nil {
		jsonError(c, err)
		return
	}

	expectedSalary, err := strconv.Atoi(data["expected_salary"])
	if err != nil && data["expected_salary"] != ""{
		jsonError(c, err)
		return
	}

	expectedEmployment, err := strconv.ParseFloat(data["expected_employment"], 32)
	if err != nil && data["expected_employment"] != "" {
		jsonError(c, err)
		return
	}

	var email, phone string
	myId := 0
	jwtMiddleware := jwt.GetJwtMiddleware(nil)
	jwtMiddleware.MiddlewareFunc()(c)
	if mw.TryAuthed(c) == nil {
		me := helpers.GetMe(c)
		myId = me.GetId()
		if me.GetEmail() != nil {
			email = *me.GetEmail()
		}
		if me.GetPhone() != nil {
			phone = *me.GetPhone()
		}
	}

	if len(data["email"]) > 0 {
		email = data["email"]
	}
	if len(data["phone"]) > 4 {
		phone = data["phone"]
	}

	if email == "" && phone == "" {
		jsonError(c, fmt.Errorf("email and phone are not set (at least one of them should be set)"))
		return
	}

	body, err := json.Marshal(data)
	if err != nil {
		jsonError(c, fmt.Errorf("internal error"))
		return
	}

	application := models.NewApplication()
	application.CreatedAt  = &[]extime.Time{extime.Now()}[0]
	application.UserId     = myId
	application.VacancyId  = vacancy.Id
	application.Employment = float32(expectedEmployment)
	application.Answers    = string(body)
	application.Email      = email
	application.Phone      = phone
	if expectedSalary > 0 {
		application.Salary = &expectedSalary
	}
	err = application.Create()
	if err != nil {
		jsonError(c, err)
		return
	}

	jsonOK(c)
}
