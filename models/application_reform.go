package models

// Generated with gopkg.in/reform.v1. Do not edit by hand.

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/xaionaro/reform"
)

type applicationScope struct {
	item *application

	db           reform.ReformDBTX
	where        [][]interface{}
	order        []string
	groupBy      []string
	limit        int
	tableQuery   *string
	fieldsFilter []string
	appendTail   string

	loggingEnabled bool
	loggingAuthor  *string
	loggingComment string
}
type ApplicationType application
type ApplicationF application
type ApplicationFilter application

type applicationLogRow struct {
	application
	LogAuthor  *string
	LogAction  string
	LogDate    time.Time
	LogComment string
}

// Schema returns a schema name in SQL database ("").
type applicationTableTypeType struct {
	s reform.StructInfo
	z []interface{}
}

func (v applicationTableTypeType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("applications").
func (v applicationTableTypeType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v applicationTableTypeType) Columns() []string {
	return []string{"id", "user_id", "vacancy_id", "salary", "employment", "answers", "email", "phone", "created_at"}
}

// NewStruct makes a new struct for that view or table.
func (v applicationTableTypeType) NewStruct() reform.Struct {
	return new(application)
}

// NewRecord makes a new record for that table.
func (v *applicationTableTypeType) NewRecord() reform.Record {
	return new(application)
}

func (v *applicationTableTypeType) NewScope() *applicationScope {
	return &applicationScope{item: &application{}}
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *applicationTableTypeType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

func (v applicationTableTypeType) CreateTableIfNotExists(db *reform.DB) (bool, error) {
	if db == nil {
		db = defaultDB_application
	}
	return db.CreateTableIfNotExists(v.s)
}

func (v applicationTableTypeType) StructInfo() reform.StructInfo {
	return v.s
}

// applicationTable represents applications view or table in SQL database.
var applicationTable = &applicationTableTypeType{
	s: reform.StructInfo{Type: "application", SQLSchema: "", SQLName: "applications", Fields: []reform.FieldInfo{{Name: "Id", IsPK: true, IsUnique: false, HasIndex: false, Type: "int", Column: "id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "UserId", IsPK: false, IsUnique: false, HasIndex: false, Type: "int", Column: "user_id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "VacancyId", IsPK: false, IsUnique: false, HasIndex: false, Type: "int", Column: "vacancy_id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "Salary", IsPK: false, IsUnique: false, HasIndex: false, Type: "*int", Column: "salary", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "Employment", IsPK: false, IsUnique: false, HasIndex: false, Type: "float32", Column: "employment", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "Answers", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "answers", FieldsPath: []reform.FieldInfo{}, SQLSize: 1048576, Embedded: "", StructFile: ""}, {Name: "Email", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "email", FieldsPath: []reform.FieldInfo{}, SQLSize: 255, Embedded: "", StructFile: ""}, {Name: "Phone", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "phone", FieldsPath: []reform.FieldInfo{}, SQLSize: 255, Embedded: "", StructFile: ""}, {Name: "CreatedAt", IsPK: false, IsUnique: false, HasIndex: false, Type: "*extime.Time", Column: "created_at", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}}, PKFieldIndex: 0, ImitateGorm: false, SkipMethodOrder: false},
	z: new(application).Values(),
}

type applicationTableTypeType_log struct {
	s reform.StructInfo
	z []interface{}
}

func (v *applicationTableTypeType_log) Schema() string {
	return v.s.SQLSchema
}

func (v *applicationTableTypeType_log) Name() string {
	return v.s.SQLName
}

func (v *applicationTableTypeType_log) Columns() []string {
	return []string{"id", "user_id", "vacancy_id", "salary", "employment", "answers", "email", "phone", "created_at", "log_author", "log_action", "log_date", "log_comment"}
}

func (v *applicationTableTypeType_log) NewStruct() reform.Struct {
	return new(application)
}

func (v *applicationTableTypeType_log) NewRecord() reform.Record {
	return new(application)
}

func (v *applicationTableTypeType_log) NewScope() *applicationScope {
	return &applicationScope{item: &application{}}
}

func (v *applicationTableTypeType_log) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

var applicationTableLogRow = &applicationTableTypeType_log{
	s: reform.StructInfo{Type: "application", SQLSchema: "", SQLName: "applications_log", Fields: []reform.FieldInfo{{Name: "Id", IsPK: true, IsUnique: false, HasIndex: false, Type: "int", Column: "id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "UserId", IsPK: false, IsUnique: false, HasIndex: false, Type: "int", Column: "user_id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "VacancyId", IsPK: false, IsUnique: false, HasIndex: false, Type: "int", Column: "vacancy_id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "Salary", IsPK: false, IsUnique: false, HasIndex: false, Type: "*int", Column: "salary", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "Employment", IsPK: false, IsUnique: false, HasIndex: false, Type: "float32", Column: "employment", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "Answers", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "answers", FieldsPath: []reform.FieldInfo{}, SQLSize: 1048576, Embedded: "", StructFile: ""}, {Name: "Email", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "email", FieldsPath: []reform.FieldInfo{}, SQLSize: 255, Embedded: "", StructFile: ""}, {Name: "Phone", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "phone", FieldsPath: []reform.FieldInfo{}, SQLSize: 255, Embedded: "", StructFile: ""}, {Name: "CreatedAt", IsPK: false, IsUnique: false, HasIndex: false, Type: "*extime.Time", Column: "created_at", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogAuthor", IsPK: false, IsUnique: false, HasIndex: false, Type: "*string", Column: "log_author", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogAction", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "log_action", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogDate", IsPK: false, IsUnique: false, HasIndex: false, Type: "time.Time", Column: "log_date", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogComment", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "log_comment", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}}, PKFieldIndex: 0, ImitateGorm: false, SkipMethodOrder: false},
	z: new(applicationLogRow).Values(),
}

func (s applicationTableTypeType) ColumnNameByFieldName(fieldName string) string {
	switch fieldName {
	case "Id":
		return "id"
	case "UserId":
		return "user_id"
	case "VacancyId":
		return "vacancy_id"
	case "Salary":
		return "salary"
	case "Employment":
		return "employment"
	case "Answers":
		return "answers"
	case "Email":
		return "email"
	case "Phone":
		return "phone"
	case "CreatedAt":
		return "created_at"
	}
	return ""
}

func (s applicationTableTypeType_log) ColumnNameByFieldName(fieldName string) string {
	switch fieldName {
	case "Id":
		return "id"
	case "UserId":
		return "user_id"
	case "VacancyId":
		return "vacancy_id"
	case "Salary":
		return "salary"
	case "Employment":
		return "employment"
	case "Answers":
		return "answers"
	case "Email":
		return "email"
	case "Phone":
		return "phone"
	case "CreatedAt":
		return "created_at"
	case "LogAuthor":
		return "log_author"
	case "LogAction":
		return "log_action"
	case "LogDate":
		return "log_date"
	case "LogComment":
		return "log_comment"
	}
	return ""
}

func (s *application) FieldPointersByNames(fieldNames []string) (fieldPointers []interface{}) {
	if len(fieldNames) == 0 {
		return s.Pointers()
	}

	for _, fieldName := range fieldNames {
		fieldPointer := s.FieldPointerByName(fieldName)
		if fieldPointer == nil {
			panic("Invalid field name:" + fieldName)
		}
		fieldPointers = append(fieldPointers, fieldPointer)
	}

	return
}

func (s *applicationLogRow) FieldPointersByNames(fieldNames []string) (fieldPointers []interface{}) {
	if len(fieldNames) == 0 {
		return s.Pointers()
	}

	for _, fieldName := range fieldNames {
		fieldPointer := s.FieldPointerByName(fieldName)
		if fieldPointer == nil {
			panic("Invalid field name:" + fieldName)
		}
		fieldPointers = append(fieldPointers, fieldPointer)
	}

	return
}

func (s *application) FieldPointerByName(fieldName string) interface{} {
	switch fieldName {
	case "Id":
		return &s.Id
	case "UserId":
		return &s.UserId
	case "VacancyId":
		return &s.VacancyId
	case "Salary":
		return &s.Salary
	case "Employment":
		return &s.Employment
	case "Answers":
		return &s.Answers
	case "Email":
		return &s.Email
	case "Phone":
		return &s.Phone
	case "CreatedAt":
		return &s.CreatedAt
	}

	return nil
}

func (s *applicationLogRow) FieldPointerByName(fieldName string) interface{} {
	switch fieldName {
	case "Id":
		return &s.Id
	case "UserId":
		return &s.UserId
	case "VacancyId":
		return &s.VacancyId
	case "Salary":
		return &s.Salary
	case "Employment":
		return &s.Employment
	case "Answers":
		return &s.Answers
	case "Email":
		return &s.Email
	case "Phone":
		return &s.Phone
	case "CreatedAt":
		return &s.CreatedAt
	case "LogAuthor":
		return &s.LogAuthor
	case "LogAction":
		return &s.LogAction
	case "LogDate":
		return &s.LogDate
	case "LogComment":
		return &s.LogComment
	}

	return nil
}

// String returns a string representation of this struct or record.
func (s application) String() string {
	res := make([]string, 9)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "UserId: " + reform.Inspect(s.UserId, true)
	res[2] = "VacancyId: " + reform.Inspect(s.VacancyId, true)
	res[3] = "Salary: " + reform.Inspect(s.Salary, true)
	res[4] = "Employment: " + reform.Inspect(s.Employment, true)
	res[5] = "Answers: " + reform.Inspect(s.Answers, true)
	res[6] = "Email: " + reform.Inspect(s.Email, true)
	res[7] = "Phone: " + reform.Inspect(s.Phone, true)
	res[8] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	return strings.Join(res, ", ")
}
func (s applicationLogRow) String() string {
	res := make([]string, 13)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "UserId: " + reform.Inspect(s.UserId, true)
	res[2] = "VacancyId: " + reform.Inspect(s.VacancyId, true)
	res[3] = "Salary: " + reform.Inspect(s.Salary, true)
	res[4] = "Employment: " + reform.Inspect(s.Employment, true)
	res[5] = "Answers: " + reform.Inspect(s.Answers, true)
	res[6] = "Email: " + reform.Inspect(s.Email, true)
	res[7] = "Phone: " + reform.Inspect(s.Phone, true)
	res[8] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[9] = "LogAuthor: " + reform.Inspect(s.LogAuthor, true)
	res[10] = "LogAction: " + reform.Inspect(s.LogAction, true)
	res[11] = "LogDate: " + reform.Inspect(s.LogDate, true)
	res[12] = "LogComment: " + reform.Inspect(s.LogComment, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *application) Values() []interface{} {
	return []interface{}{
		s.Id,
		s.UserId,
		s.VacancyId,
		s.Salary,
		s.Employment,
		s.Answers,
		s.Email,
		s.Phone,
		s.CreatedAt,
	}
}
func (s *applicationLogRow) Values() []interface{} {
	return append(s.application.Values(), []interface{}{
		s.LogAuthor,
		s.LogAction,
		s.LogDate,
		s.LogComment,
	}...)
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *application) Pointers() []interface{} {
	return []interface{}{
		&s.Id,
		&s.UserId,
		&s.VacancyId,
		&s.Salary,
		&s.Employment,
		&s.Answers,
		&s.Email,
		&s.Phone,
		&s.CreatedAt,
	}
}
func (s *applicationLogRow) Pointers() []interface{} {
	return append(s.application.Pointers(), []interface{}{
		&s.LogAuthor,
		&s.LogAction,
		&s.LogDate,
		&s.LogComment,
	}...)
}

// View returns View object for that struct.
func (s application) View() reform.View {
	return applicationTable
}
func (s applicationScope) View() reform.View {
	return s.item.View()
}
func (s applicationLogRow) View() reform.View {
	return applicationTableLogRow
}

// Generate a scope for object
func (s application) Scope() *applicationScope {
	return &applicationScope{item: &s, db: defaultDB_application}
}
func (s *application) PtrScope() *applicationScope {
	return &applicationScope{item: s, db: defaultDB_application}
}

// Sets DB to do queries
func (s application) DB(db reform.ReformDBTX) (scope *applicationScope) { return s.Scope().DB(db) }
func (s *applicationScope) DB(db reform.ReformDBTX) *applicationScope {
	if db != nil {
		s.db = db
	}
	afterDBer, ok := interface{}(s).(reform.AfterDBer)
	if ok {
		afterDBer.AfterDB()
	}
	return s
}

// Gets DB
func (s application) GetDB() (db *reform.DB) { return s.Scope().GetDB() }
func (s applicationScope) GetDB() *reform.DB {
	return s.db.(*reform.DB)
}

func (s application) StartTransaction() (*reform.TX, error) { return s.Scope().StartTransaction() }
func (s applicationScope) StartTransaction() (*reform.TX, error) {
	return s.db.(*reform.DB).Begin()
}

// Sets default DB (to do not call the scope.DB() method every time)
func (s *application) SetDefaultDB(db *reform.DB) (err error) {
	defaultDB_application = db
	return nil
}

// Compiles SQL tail for defined limit scope
// TODO: should be compiled via dialects
func (s *applicationScope) getLimitTail() (tail string, args []interface{}, err error) {
	if s.limit <= 0 {
		return
	}

	tail = fmt.Sprintf("%v", s.limit)
	return
}

// Compiles SQL tail for defined group scope
// TODO: should be compiled via dialects
func (s *applicationScope) getGroupTail() (tail string, args []interface{}, err error) {
	tail = strings.Join(s.groupBy, ", ")

	return
}

// Compiles SQL tail for defined order scope
// TODO: should be compiled via dialects
func (s *applicationScope) getOrderTail() (tail string, args []interface{}, err error) {
	var fieldName string
	var orderStringParts []string

	for idx, orderStr := range s.order {
		switch idx % 2 {
		case 0:
			fieldName = orderStr
		case 1:
			orderDirection := orderStr

			orderStringParts = append(orderStringParts, s.db.EscapeTableName(fieldName)+" "+orderDirection)
		}
	}

	tail = strings.Join(orderStringParts, ", ")

	return
}

// Compiles SQL tail for defined filter
// TODO: should be compiled via dialects
func (s *applicationScope) getWhereTailForFilter(filter ApplicationFilter) (tail string, whereTailArgs []interface{}, err error) {
	return s.db.GetWhereTailForFilter(application(filter), nil, "", false)
}

// parseQuerierArgs considers different ways of defning the tail (using scope properties or/and in_args)
func (s applicationScope) parseWhereTailComponent(in_args []interface{}, placeholderCounter *int) (tail string, args []interface{}, err error) {
	if len(in_args) > 0 {
		switch arg := in_args[0].(type) {
		case int:
			tail = "id " + s.db.OperatorAndPlaceholderOfValueForSQL(in_args[0], *placeholderCounter)
			*placeholderCounter++
			args = s.db.ValueForSQL(in_args[0])
		case string:
			tailWords := s.db.SplitConditionByPlaceholders(arg)

			if len(tailWords)-1 != len(in_args[1:]) {
				panic(fmt.Errorf("The pattern doesn't fit for passed arguments (wrong number of question marks?): len(tailWords)-1 != len(in_args[1:]): <%v> <%v>", arg, in_args[1:]))
			}

			for idx, rawNewArgs := range in_args[1:] {
				newArgs := s.db.ValueForSQL(rawNewArgs)
				newTailWords := []string{}
				for range newArgs {
					newTailWords = append(newTailWords, s.db.GetDialect().Placeholder(*placeholderCounter))
					*placeholderCounter++
				}
				tail += tailWords[idx] + strings.Join(newTailWords, ",")
				args = append(args, newArgs...)
			}
			tail += tailWords[len(in_args[1:])]

			return
		case *application:
			in_args[0] = *arg
			return s.parseWhereTailComponent(in_args, placeholderCounter)
		case *ApplicationF:
			in_args[0] = *arg
			return s.parseWhereTailComponent(in_args, placeholderCounter)
		case *ApplicationFilter:
			in_args[0] = *arg
			return s.parseWhereTailComponent(in_args, placeholderCounter)
		case application:
			if len(in_args) > 1 {
				s = *s.Where(in_args[1], in_args[2:]...)
			}
			tail, args, err = s.getWhereTailForFilter(ApplicationFilter(arg))
		case ApplicationF:
			if len(in_args) > 1 {
				s = *s.Where(in_args[1], in_args[2:]...)
			}
			tail, args, err = s.getWhereTailForFilter(ApplicationFilter(arg))
		case ApplicationFilter:
			if len(in_args) > 1 {
				s = *s.Where(in_args[1], in_args[2:]...)
			}
			tail, args, err = s.getWhereTailForFilter(arg)
		default:
			err = fmt.Errorf("Invalid first element of \"in_args\" (%T). It should be a string or ApplicationFilter.", arg)
			return
		}
	}

	return
}

// Compiles SQL tail for defined filter
// TODO: should be compiled via dialects
func (s *applicationScope) getWhereTail() (tail string, whereTailArgs []interface{}, err error) {
	var whereTailStringParts []string

	if len(s.where) == 0 {
		return
	}

	placeholderCounter := 0

	for _, whereComponent := range s.where {
		var whereTailStringPart string
		var whereTailArgsPart []interface{}

		whereTailStringPart, whereTailArgsPart, err = s.parseWhereTailComponent(whereComponent, &placeholderCounter)
		if err != nil {
			return
		}

		if len(whereTailStringPart) > 0 {
			whereTailStringParts = append(whereTailStringParts, whereTailStringPart)
		}
		whereTailArgs = append(whereTailArgs, whereTailArgsPart...)
	}

	if len(whereTailStringParts) == 0 {
		return
	}

	tail = "(" + strings.Join(whereTailStringParts, ") AND (") + ")"

	return
}

func (s application) Where(requiredArg interface{}, args ...interface{}) (scope *applicationScope) {
	return s.Scope().Where(requiredArg, args...)
}
func (s applicationScope) Where(requiredArg interface{}, in_args ...interface{}) *applicationScope {
	s.where = append(s.where, append([]interface{}{requiredArg}, in_args...))
	return &s
}
func (s applicationScope) SetWhere(where [][]interface{}) *applicationScope {
	s.where = where
	return &s
}
func (s applicationScope) GetWhere() [][]interface{} {
	return s.where
}

// Sets all scope-related parameters to be equal as in passed scope (as an argument)
func (s applicationScope) SetScope(anotherScope reform.Scope) *applicationScope {
	s.where = anotherScope.GetWhere()
	s.order = anotherScope.GetOrder()
	s.groupBy = anotherScope.GetGroup()
	s.limit = anotherScope.GetLimit()
	s.db = anotherScope.GetDB()

	return &s
}
func (s applicationScope) ISetScope(anotherScope reform.Scope) reform.Scope {
	return s.ISetScope(anotherScope)
}

// Compiles SQL tail for defined db/where/order/limit scope
// TODO: should be compiled via dialects
func (s *applicationScope) getTail() (tail string, args []interface{}, err error) {
	whereTailString, whereTailArgs, err := s.getWhereTail()

	if err != nil {
		return
	}
	groupTailString, groupTailArgs, err := s.getGroupTail()
	if err != nil {
		return
	}
	orderTailString, orderTailArgs, err := s.getOrderTail()
	if err != nil {
		return
	}
	limitTailString, _, err := s.getLimitTail()
	if err != nil {
		return
	}

	args = append(whereTailArgs, append(groupTailArgs, orderTailArgs...)...)

	if len(whereTailString) > 0 {
		whereTailString = " WHERE " + whereTailString + " "
	}

	if len(groupTailString) > 0 {
		groupTailString = " GROUP BY " + groupTailString + " "
	}

	if len(orderTailString) > 0 {
		orderTailString = " ORDER BY " + orderTailString + " "
	}

	if len(limitTailString) > 0 {
		limitTailString = " LIMIT " + limitTailString + " "
	}

	tail = whereTailString + groupTailString + orderTailString + limitTailString

	if len(s.appendTail) > 0 {
		tail += " " + s.appendTail
	}

	return

}

// SelectRows is a simple wrapper to get raw "sql.Rows"
func (s application) SelectRows(query string, args ...interface{}) (rows *sql.Rows, err error) {
	return s.Scope().SelectRows(query, args...)
}
func (s *applicationScope) SelectRows(query string, queryArgs ...interface{}) (rows *sql.Rows, err error) {
	tail, args, err := s.getTail()
	if err != nil {
		return
	}

	return s.db.Query("SELECT "+query+" FROM `"+applicationTable.Name()+"` "+tail, append(queryArgs, args...)...)
}

func (s *applicationScope) callStructMethod(str *application, methodName string) error {
	if method := reflect.ValueOf(str).MethodByName(methodName); method.IsValid() {
		switch f := method.Interface().(type) {
		case func():
			f()

		case func(reform.ReformDBTX):
			f(s.db)

		case func(*applicationScope):
			f(s)

		case func(interface{}): // For compatibility with other ORMs
			f(s.db)

		case func() error:
			return f()

		case func(reform.ReformDBTX) error:
			return f(s.db)

		case func(*applicationScope) error:
			return f(s)

		case func(interface{}) error: // For compatibility with other ORMS
			return f(s.db)

		default:
			panic("Unknown type of method: \"" + methodName + "\"")
		}
	}
	return nil
}

func (s applicationScope) checkDb() {
	if s.db == nil {
		panic("s.db == nil")
	}
}

// Select is a handy wrapper for SelectRows() and NextRow(): it makes a query and collects the result into a slice
func (s application) Select(args ...interface{}) (result []application, err error) {
	return s.Scope().Select(args...)
}
func (s applicationScope) Select(args ...interface{}) (result []application, err error) {
	s.checkDb()

	if len(args) > 0 {
		s = *s.Where(args[0], args[1:]...)
	}
	tail, args, err := s.getTail()
	if err != nil {
		return
	}

	rows, err := s.db.FlexSelectRows(applicationTable, s.tableQuery, s.fieldsFilter, tail, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		item := application{}
		err = rows.Scan(item.FieldPointersByNames(s.fieldsFilter)...)
		if err != nil {
			return
		}

		s.callStructMethod(&item, "AfterFind")

		result = append(result, item)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return
}
func (s application) SelectI(args ...interface{}) (result interface{}, err error) {
	return s.Scope().Select(args...)
}
func (s applicationScope) SelectI(args ...interface{}) (result interface{}, err error) {
	return s.Select(args...)
}

// "First" a method to select and return only one record.
func (s application) First(args ...interface{}) (result application, err error) {
	return s.Scope().First(args...)
}
func (s applicationScope) First(args ...interface{}) (result application, err error) {
	s.checkDb()

	if len(args) > 0 {
		s = *s.Where(args[0], args[1:]...)
	}
	tail, args, err := s.Limit(1).getTail()
	if err != nil {
		return
	}

	err = s.db.FlexSelectOneTo(&result, s.tableQuery, s.fieldsFilter, tail, args...)

	return
}
func (s application) FirstI(args ...interface{}) (result interface{}, err error) {
	return s.Scope().First(args...)
}
func (s applicationScope) FirstI(args ...interface{}) (result interface{}, err error) {
	return s.First(args...)
}

// Sets "GROUP BY".
func (s application) Group(args ...interface{}) (scope *applicationScope) {
	return s.Scope().Group(args...)
}
func (s applicationScope) Group(argsI ...interface{}) *applicationScope {
	for _, argI := range argsI {
		s.groupBy = append(s.groupBy, argI.(string))
	}

	return &s
}
func (s applicationScope) SetGroup(groupBy []string) *applicationScope {
	s.groupBy = groupBy
	return &s
}
func (s applicationScope) GetGroup() []string {
	return s.groupBy
}

// Sets a table query. For example SetTableQuery("table1 JOIN table2 USING(key)")
func (s application) SetTableQuery(query string) (scope *applicationScope) {
	return s.Scope().SetTableQuery(query)
}
func (s applicationScope) SetTableQuery(query string) *applicationScope {
	if query == "" {
		s.tableQuery = nil
	} else {
		s.tableQuery = &query
	}
	return &s
}
func (s applicationScope) GetTableQuery() string {
	if s.tableQuery != nil {
		return *s.tableQuery
	}
	return s.db.QualifiedView(s.View())
}

// Sets which structure fields should be queried while Select()/First(). For example SetFields("StructField1", "StructIdField", "StructCommentsField"). Could be used just to speed up a query.
// It's not recommended to use this function!
func (s application) SetQueryFieldsByNames(fields ...string) (scope *applicationScope) {
	return s.Scope().SetQueryFieldsByNames(fields...)
}
func (s applicationScope) SetQueryFieldsByNames(fields ...string) *applicationScope {
	s.fieldsFilter = fields
	return &s
}
func (s applicationScope) GetQueryFields() []string {
	return s.fieldsFilter
}

// Sets order. Arguments should be passed by pairs column-{ASC,DESC}. For example Order("id", "ASC", "value" "DESC")
func (s application) Order(args ...interface{}) (scope *applicationScope) {
	return s.Scope().Order(args...)
}
func (s applicationScope) Order(argsI ...interface{}) *applicationScope {
	switch len(argsI) {
	case 0:
	case 1:
		arg := argsI[0].(string)
		args0 := strings.Split(arg, ",")
		var args []string
		for _, arg0 := range args0 {
			args = append(args, strings.Split(arg0, ":")...)
		}
		s.order = args
	default:
		var args []string
		for _, argI := range argsI {
			args = append(args, argI.(string))
		}
		s.order = args
	}

	return &s
}
func (s applicationScope) SetOrder(order []string) *applicationScope {
	s.order = order
	return &s
}
func (s applicationScope) GetOrder() []string {
	return s.order
}

func (s application) SetSQLAppend(appendTail string) (scope *applicationScope) {
	return s.Scope().SetSQLAppend(appendTail)
}
func (s applicationScope) SetSQLAppend(appendTail string) *applicationScope {
	s.appendTail = appendTail
	return &s
}

// Sets limit.
func (s application) Limit(limit int) (scope *applicationScope) { return s.Scope().Limit(limit) }
func (s *applicationScope) Limit(limit int) *applicationScope {
	s.limit = limit
	return s
}

// Gets limit
func (s applicationScope) GetLimit() int {
	return s.limit
}

// "Reload" reloads record using Primary Key
func (s *ApplicationFilter) Reload(db *reform.DB) error { return (*application)(s).Reload(db) }
func (s *application) Reload(db *reform.DB) (err error) {
	return db.FindByPrimaryKeyTo(s, s.PKValue())
}

// Create and Insert inserts new record to DB
func (s *application) Create() (err error) { return s.PtrScope().Create() }
func (s *applicationScope) Create() (err error) {
	return s.Insert()
}
func (s *application) Insert() (err error) { return s.PtrScope().Insert() }
func (s *applicationScope) Insert() (err error) {
	s.checkDb()
	err = s.db.Insert(s.item)
	if err == nil {
		s.doLog("INSERT")
	}
	return err
}

// Replace "REPLACE INTO" new record to DB
func (s *application) Replace() (err error) { return s.PtrScope().Replace() }
func (s *applicationScope) Replace() (err error) {
	s.checkDb()
	err = s.db.Replace(s.item)
	if err == nil {
		s.doLog("REPLACE")
	}
	return err
}

// Save inserts new record to DB is PK is zero and updates existing record if PK is not zero
func (s *application) Save() (err error) { return s.PtrScope().Save() }
func (s *applicationScope) Save() (err error) {
	s.checkDb()
	err = s.db.Save(s.item)
	if err == nil {
		s.doLog("INSERT")
	}
	return err
}

// Update updates existing record in DB
func (s application) Update() (err error) { return s.Scope().Update() }
func (s *applicationScope) Update() (err error) {
	s.checkDb()
	err = s.db.Update(s.item)
	if err == nil {
		s.doLog("UPDATE")
	}
	return err
}

// Delete deletes existing record in DB
func (s application) Delete() (err error) { return s.Scope().Delete() }
func (s *applicationScope) Delete() (err error) {
	s.checkDb()
	err = s.db.Delete(s.item)
	if err == nil {
		s.doLog("DELETE")
	}
	return err
}

func (s *applicationScope) doLog(requestType string) {
	if !s.loggingEnabled {
		return
	}

	var logRow applicationLogRow
	logRow.application = *s.item
	logRow.LogAuthor = s.loggingAuthor
	logRow.LogAction = requestType
	logRow.LogDate = time.Now()
	logRow.LogComment = s.loggingComment

	s.db.Insert(&logRow)
}

// Enables logging to table "applications_log". This table should has the same schema, except:
// - Unique/Primary keys should be removed
// - Should be added next fields: "log_author" (nullable string), "log_date" (timestamp), "log_action" (enum("INSERT", "UPDATE", "DELETE")), "log_comment" (string)
func (s *application) Log(enableLogging bool, author *string, commentFormat string, commentArgs ...interface{}) (scope *applicationScope) {
	return s.Scope().Log(enableLogging, author, commentFormat, commentArgs...)
}
func (s *applicationScope) Log(enableLogging bool, author *string, commentFormat string, commentArgs ...interface{}) (scope *applicationScope) {
	s.loggingEnabled = enableLogging
	s.loggingAuthor = author
	s.loggingComment = fmt.Sprintf(commentFormat, commentArgs...)

	return s
}

// Table returns Table object for that record.
func (s application) Table() reform.Table {
	return applicationTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s application) PKValue() interface{} {
	return s.Id
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s application) PKPointer() interface{} {
	return &s.Id
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s application) HasPK() bool {
	return s.Id != applicationTable.z[applicationTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *ApplicationFilter) SetPK(pk interface{}) { (*application)(s).SetPK(pk) }
func (s *application) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.Id = int(i64)
	} else {
		s.Id = pk.(int)
	}
}

var (
	// check interfaces
	_ reform.View   = applicationTable
	_ reform.Struct = (*application)(nil)
	_ reform.Table  = applicationTable
	_ reform.Record = (*application)(nil)
	_ fmt.Stringer  = (*application)(nil)

	// querier
	Application           = application{} // Should be read only
	defaultDB_application *reform.DB
)

func init() {
	//parse.AssertUpToDate(&applicationTable.s, new(application)) // Temporary disabled (doesn't work with arbitary types like "type sliceString []string")
}
