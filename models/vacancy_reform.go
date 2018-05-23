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

type vacancyScope struct {
	item *vacancy

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
type VacancyType vacancy
type VacancyF vacancy
type VacancyFilter vacancy

type vacancyLogRow struct {
	vacancy
	LogAuthor  *string
	LogAction  string
	LogDate    time.Time
	LogComment string
}

// Schema returns a schema name in SQL database ("").
type vacancyTableTypeType struct {
	s reform.StructInfo
	z []interface{}
}

func (v vacancyTableTypeType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("vacancies").
func (v vacancyTableTypeType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v vacancyTableTypeType) Columns() []string {
	return []string{"id"}
}

// NewStruct makes a new struct for that view or table.
func (v vacancyTableTypeType) NewStruct() reform.Struct {
	return new(vacancy)
}

// NewRecord makes a new record for that table.
func (v *vacancyTableTypeType) NewRecord() reform.Record {
	return new(vacancy)
}

func (v *vacancyTableTypeType) NewScope() *vacancyScope {
	return &vacancyScope{item: &vacancy{}}
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *vacancyTableTypeType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

func (v vacancyTableTypeType) CreateTableIfNotExists(db *reform.DB) (bool, error) {
	if db == nil {
		db = defaultDB_vacancy
	}
	return db.CreateTableIfNotExists(v.s)
}

func (v vacancyTableTypeType) StructInfo() reform.StructInfo {
	return v.s
}

// vacancyTable represents vacancies view or table in SQL database.
var vacancyTable = &vacancyTableTypeType{
	s: reform.StructInfo{Type: "vacancy", SQLSchema: "", SQLName: "vacancies", Fields: []reform.FieldInfo{{Name: "Id", IsPK: true, IsUnique: false, HasIndex: false, Type: "int", Column: "id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}}, PKFieldIndex: 0, ImitateGorm: false, SkipMethodOrder: false},
	z: new(vacancy).Values(),
}

type vacancyTableTypeType_log struct {
	s reform.StructInfo
	z []interface{}
}

func (v *vacancyTableTypeType_log) Schema() string {
	return v.s.SQLSchema
}

func (v *vacancyTableTypeType_log) Name() string {
	return v.s.SQLName
}

func (v *vacancyTableTypeType_log) Columns() []string {
	return []string{"id", "log_author", "log_action", "log_date", "log_comment"}
}

func (v *vacancyTableTypeType_log) NewStruct() reform.Struct {
	return new(vacancy)
}

func (v *vacancyTableTypeType_log) NewRecord() reform.Record {
	return new(vacancy)
}

func (v *vacancyTableTypeType_log) NewScope() *vacancyScope {
	return &vacancyScope{item: &vacancy{}}
}

func (v *vacancyTableTypeType_log) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

var vacancyTableLogRow = &vacancyTableTypeType_log{
	s: reform.StructInfo{Type: "vacancy", SQLSchema: "", SQLName: "vacancies_log", Fields: []reform.FieldInfo{{Name: "Id", IsPK: true, IsUnique: false, HasIndex: false, Type: "int", Column: "id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogAuthor", IsPK: false, IsUnique: false, HasIndex: false, Type: "*string", Column: "log_author", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogAction", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "log_action", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogDate", IsPK: false, IsUnique: false, HasIndex: false, Type: "time.Time", Column: "log_date", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogComment", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "log_comment", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}}, PKFieldIndex: 0, ImitateGorm: false, SkipMethodOrder: false},
	z: new(vacancyLogRow).Values(),
}

func (s vacancyTableTypeType) ColumnNameByFieldName(fieldName string) string {
	switch fieldName {
	case "Id":
		return "id"
	}
	return ""
}

func (s vacancyTableTypeType_log) ColumnNameByFieldName(fieldName string) string {
	switch fieldName {
	case "Id":
		return "id"
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

func (s *vacancy) FieldPointersByNames(fieldNames []string) (fieldPointers []interface{}) {
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

func (s *vacancyLogRow) FieldPointersByNames(fieldNames []string) (fieldPointers []interface{}) {
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

func (s *vacancy) FieldPointerByName(fieldName string) interface{} {
	switch fieldName {
	case "Id":
		return &s.Id
	}

	return nil
}

func (s *vacancyLogRow) FieldPointerByName(fieldName string) interface{} {
	switch fieldName {
	case "Id":
		return &s.Id
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
func (s vacancy) String() string {
	res := make([]string, 1)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	return strings.Join(res, ", ")
}
func (s vacancyLogRow) String() string {
	res := make([]string, 5)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "LogAuthor: " + reform.Inspect(s.LogAuthor, true)
	res[2] = "LogAction: " + reform.Inspect(s.LogAction, true)
	res[3] = "LogDate: " + reform.Inspect(s.LogDate, true)
	res[4] = "LogComment: " + reform.Inspect(s.LogComment, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *vacancy) Values() []interface{} {
	return []interface{}{
		s.Id,
	}
}
func (s *vacancyLogRow) Values() []interface{} {
	return append(s.vacancy.Values(), []interface{}{
		s.LogAuthor,
		s.LogAction,
		s.LogDate,
		s.LogComment,
	}...)
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *vacancy) Pointers() []interface{} {
	return []interface{}{
		&s.Id,
	}
}
func (s *vacancyLogRow) Pointers() []interface{} {
	return append(s.vacancy.Pointers(), []interface{}{
		&s.LogAuthor,
		&s.LogAction,
		&s.LogDate,
		&s.LogComment,
	}...)
}

// View returns View object for that struct.
func (s vacancy) View() reform.View {
	return vacancyTable
}
func (s vacancyScope) View() reform.View {
	return s.item.View()
}
func (s vacancyLogRow) View() reform.View {
	return vacancyTableLogRow
}

// Generate a scope for object
func (s vacancy) Scope() *vacancyScope {
	return &vacancyScope{item: &s, db: defaultDB_vacancy}
}
func (s *vacancy) PtrScope() *vacancyScope {
	return &vacancyScope{item: s, db: defaultDB_vacancy}
}

// Sets DB to do queries
func (s vacancy) DB(db reform.ReformDBTX) (scope *vacancyScope) { return s.Scope().DB(db) }
func (s *vacancyScope) DB(db reform.ReformDBTX) *vacancyScope {
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
func (s vacancy) GetDB() (db *reform.DB) { return s.Scope().GetDB() }
func (s vacancyScope) GetDB() *reform.DB {
	return s.db.(*reform.DB)
}

func (s vacancy) StartTransaction() (*reform.TX, error) { return s.Scope().StartTransaction() }
func (s vacancyScope) StartTransaction() (*reform.TX, error) {
	return s.db.(*reform.DB).Begin()
}

// Sets default DB (to do not call the scope.DB() method every time)
func (s *vacancy) SetDefaultDB(db *reform.DB) (err error) {
	defaultDB_vacancy = db
	return nil
}

// Compiles SQL tail for defined limit scope
// TODO: should be compiled via dialects
func (s *vacancyScope) getLimitTail() (tail string, args []interface{}, err error) {
	if s.limit <= 0 {
		return
	}

	tail = fmt.Sprintf("%v", s.limit)
	return
}

// Compiles SQL tail for defined group scope
// TODO: should be compiled via dialects
func (s *vacancyScope) getGroupTail() (tail string, args []interface{}, err error) {
	tail = strings.Join(s.groupBy, ", ")

	return
}

// Compiles SQL tail for defined order scope
// TODO: should be compiled via dialects
func (s *vacancyScope) getOrderTail() (tail string, args []interface{}, err error) {
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
func (s *vacancyScope) getWhereTailForFilter(filter VacancyFilter) (tail string, whereTailArgs []interface{}, err error) {
	return s.db.GetWhereTailForFilter(vacancy(filter), nil, "", false)
}

// parseQuerierArgs considers different ways of defning the tail (using scope properties or/and in_args)
func (s vacancyScope) parseWhereTailComponent(in_args []interface{}, placeholderCounter *int) (tail string, args []interface{}, err error) {
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
		case *vacancy:
			in_args[0] = *arg
			return s.parseWhereTailComponent(in_args, placeholderCounter)
		case *VacancyF:
			in_args[0] = *arg
			return s.parseWhereTailComponent(in_args, placeholderCounter)
		case *VacancyFilter:
			in_args[0] = *arg
			return s.parseWhereTailComponent(in_args, placeholderCounter)
		case vacancy:
			if len(in_args) > 1 {
				s = *s.Where(in_args[1], in_args[2:]...)
			}
			tail, args, err = s.getWhereTailForFilter(VacancyFilter(arg))
		case VacancyF:
			if len(in_args) > 1 {
				s = *s.Where(in_args[1], in_args[2:]...)
			}
			tail, args, err = s.getWhereTailForFilter(VacancyFilter(arg))
		case VacancyFilter:
			if len(in_args) > 1 {
				s = *s.Where(in_args[1], in_args[2:]...)
			}
			tail, args, err = s.getWhereTailForFilter(arg)
		default:
			err = fmt.Errorf("Invalid first element of \"in_args\" (%T). It should be a string or VacancyFilter.", arg)
			return
		}
	}

	return
}

// Compiles SQL tail for defined filter
// TODO: should be compiled via dialects
func (s *vacancyScope) getWhereTail() (tail string, whereTailArgs []interface{}, err error) {
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

func (s vacancy) Where(requiredArg interface{}, args ...interface{}) (scope *vacancyScope) {
	return s.Scope().Where(requiredArg, args...)
}
func (s vacancyScope) Where(requiredArg interface{}, in_args ...interface{}) *vacancyScope {
	s.where = append(s.where, append([]interface{}{requiredArg}, in_args...))
	return &s
}
func (s vacancyScope) SetWhere(where [][]interface{}) *vacancyScope {
	s.where = where
	return &s
}
func (s vacancyScope) GetWhere() [][]interface{} {
	return s.where
}

// Sets all scope-related parameters to be equal as in passed scope (as an argument)
func (s vacancyScope) SetScope(anotherScope reform.Scope) *vacancyScope {
	s.where = anotherScope.GetWhere()
	s.order = anotherScope.GetOrder()
	s.groupBy = anotherScope.GetGroup()
	s.limit = anotherScope.GetLimit()
	s.db = anotherScope.GetDB()

	return &s
}
func (s vacancyScope) ISetScope(anotherScope reform.Scope) reform.Scope {
	return s.ISetScope(anotherScope)
}

// Compiles SQL tail for defined db/where/order/limit scope
// TODO: should be compiled via dialects
func (s *vacancyScope) getTail() (tail string, args []interface{}, err error) {
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
func (s vacancy) SelectRows(query string, args ...interface{}) (rows *sql.Rows, err error) {
	return s.Scope().SelectRows(query, args...)
}
func (s *vacancyScope) SelectRows(query string, queryArgs ...interface{}) (rows *sql.Rows, err error) {
	tail, args, err := s.getTail()
	if err != nil {
		return
	}

	return s.db.Query("SELECT "+query+" FROM `"+vacancyTable.Name()+"` "+tail, append(queryArgs, args...)...)
}

func (s *vacancyScope) callStructMethod(str *vacancy, methodName string) error {
	if method := reflect.ValueOf(str).MethodByName(methodName); method.IsValid() {
		switch f := method.Interface().(type) {
		case func():
			f()

		case func(reform.ReformDBTX):
			f(s.db)

		case func(*vacancyScope):
			f(s)

		case func(interface{}): // For compatibility with other ORMs
			f(s.db)

		case func() error:
			return f()

		case func(reform.ReformDBTX) error:
			return f(s.db)

		case func(*vacancyScope) error:
			return f(s)

		case func(interface{}) error: // For compatibility with other ORMS
			return f(s.db)

		default:
			panic("Unknown type of method: \"" + methodName + "\"")
		}
	}
	return nil
}

func (s vacancyScope) checkDb() {
	if s.db == nil {
		panic("s.db == nil")
	}
}

// Select is a handy wrapper for SelectRows() and NextRow(): it makes a query and collects the result into a slice
func (s vacancy) Select(args ...interface{}) (result []vacancy, err error) {
	return s.Scope().Select(args...)
}
func (s vacancyScope) Select(args ...interface{}) (result []vacancy, err error) {
	s.checkDb()

	if len(args) > 0 {
		s = *s.Where(args[0], args[1:]...)
	}
	tail, args, err := s.getTail()
	if err != nil {
		return
	}

	rows, err := s.db.FlexSelectRows(vacancyTable, s.tableQuery, s.fieldsFilter, tail, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		item := vacancy{}
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
func (s vacancy) SelectI(args ...interface{}) (result interface{}, err error) {
	return s.Scope().Select(args...)
}
func (s vacancyScope) SelectI(args ...interface{}) (result interface{}, err error) {
	return s.Select(args...)
}

// "First" a method to select and return only one record.
func (s vacancy) First(args ...interface{}) (result vacancy, err error) {
	return s.Scope().First(args...)
}
func (s vacancyScope) First(args ...interface{}) (result vacancy, err error) {
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
func (s vacancy) FirstI(args ...interface{}) (result interface{}, err error) {
	return s.Scope().First(args...)
}
func (s vacancyScope) FirstI(args ...interface{}) (result interface{}, err error) {
	return s.First(args...)
}

// Sets "GROUP BY".
func (s vacancy) Group(args ...interface{}) (scope *vacancyScope) { return s.Scope().Group(args...) }
func (s vacancyScope) Group(argsI ...interface{}) *vacancyScope {
	for _, argI := range argsI {
		s.groupBy = append(s.groupBy, argI.(string))
	}

	return &s
}
func (s vacancyScope) SetGroup(groupBy []string) *vacancyScope {
	s.groupBy = groupBy
	return &s
}
func (s vacancyScope) GetGroup() []string {
	return s.groupBy
}

// Sets a table query. For example SetTableQuery("table1 JOIN table2 USING(key)")
func (s vacancy) SetTableQuery(query string) (scope *vacancyScope) {
	return s.Scope().SetTableQuery(query)
}
func (s vacancyScope) SetTableQuery(query string) *vacancyScope {
	if query == "" {
		s.tableQuery = nil
	} else {
		s.tableQuery = &query
	}
	return &s
}
func (s vacancyScope) GetTableQuery() string {
	if s.tableQuery != nil {
		return *s.tableQuery
	}
	return s.db.QualifiedView(s.View())
}

// Sets which structure fields should be queried while Select()/First(). For example SetFields("StructField1", "StructIdField", "StructCommentsField"). Could be used just to speed up a query.
// It's not recommended to use this function!
func (s vacancy) SetQueryFieldsByNames(fields ...string) (scope *vacancyScope) {
	return s.Scope().SetQueryFieldsByNames(fields...)
}
func (s vacancyScope) SetQueryFieldsByNames(fields ...string) *vacancyScope {
	s.fieldsFilter = fields
	return &s
}
func (s vacancyScope) GetQueryFields() []string {
	return s.fieldsFilter
}

// Sets order. Arguments should be passed by pairs column-{ASC,DESC}. For example Order("id", "ASC", "value" "DESC")
func (s vacancy) Order(args ...interface{}) (scope *vacancyScope) { return s.Scope().Order(args...) }
func (s vacancyScope) Order(argsI ...interface{}) *vacancyScope {
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
func (s vacancyScope) SetOrder(order []string) *vacancyScope {
	s.order = order
	return &s
}
func (s vacancyScope) GetOrder() []string {
	return s.order
}

func (s vacancy) SetSQLAppend(appendTail string) (scope *vacancyScope) {
	return s.Scope().SetSQLAppend(appendTail)
}
func (s vacancyScope) SetSQLAppend(appendTail string) *vacancyScope {
	s.appendTail = appendTail
	return &s
}

// Sets limit.
func (s vacancy) Limit(limit int) (scope *vacancyScope) { return s.Scope().Limit(limit) }
func (s *vacancyScope) Limit(limit int) *vacancyScope {
	s.limit = limit
	return s
}

// Gets limit
func (s vacancyScope) GetLimit() int {
	return s.limit
}

// "Reload" reloads record using Primary Key
func (s *VacancyFilter) Reload(db *reform.DB) error { return (*vacancy)(s).Reload(db) }
func (s *vacancy) Reload(db *reform.DB) (err error) {
	return db.FindByPrimaryKeyTo(s, s.PKValue())
}

// Create and Insert inserts new record to DB
func (s *vacancy) Create() (err error) { return s.PtrScope().Create() }
func (s *vacancyScope) Create() (err error) {
	return s.Insert()
}
func (s *vacancy) Insert() (err error) { return s.PtrScope().Insert() }
func (s *vacancyScope) Insert() (err error) {
	s.checkDb()
	err = s.db.Insert(s.item)
	if err == nil {
		s.doLog("INSERT")
	}
	return err
}

// Replace "REPLACE INTO" new record to DB
func (s *vacancy) Replace() (err error) { return s.PtrScope().Replace() }
func (s *vacancyScope) Replace() (err error) {
	s.checkDb()
	err = s.db.Replace(s.item)
	if err == nil {
		s.doLog("REPLACE")
	}
	return err
}

// Save inserts new record to DB is PK is zero and updates existing record if PK is not zero
func (s *vacancy) Save() (err error) { return s.PtrScope().Save() }
func (s *vacancyScope) Save() (err error) {
	s.checkDb()
	err = s.db.Save(s.item)
	if err == nil {
		s.doLog("INSERT")
	}
	return err
}

// Update updates existing record in DB
func (s vacancy) Update() (err error) { return s.Scope().Update() }
func (s *vacancyScope) Update() (err error) {
	s.checkDb()
	err = s.db.Update(s.item)
	if err == nil {
		s.doLog("UPDATE")
	}
	return err
}

// Delete deletes existing record in DB
func (s vacancy) Delete() (err error) { return s.Scope().Delete() }
func (s *vacancyScope) Delete() (err error) {
	s.checkDb()
	err = s.db.Delete(s.item)
	if err == nil {
		s.doLog("DELETE")
	}
	return err
}

func (s *vacancyScope) doLog(requestType string) {
	if !s.loggingEnabled {
		return
	}

	var logRow vacancyLogRow
	logRow.vacancy = *s.item
	logRow.LogAuthor = s.loggingAuthor
	logRow.LogAction = requestType
	logRow.LogDate = time.Now()
	logRow.LogComment = s.loggingComment

	s.db.Insert(&logRow)
}

// Enables logging to table "vacancies_log". This table should has the same schema, except:
// - Unique/Primary keys should be removed
// - Should be added next fields: "log_author" (nullable string), "log_date" (timestamp), "log_action" (enum("INSERT", "UPDATE", "DELETE")), "log_comment" (string)
func (s *vacancy) Log(enableLogging bool, author *string, commentFormat string, commentArgs ...interface{}) (scope *vacancyScope) {
	return s.Scope().Log(enableLogging, author, commentFormat, commentArgs...)
}
func (s *vacancyScope) Log(enableLogging bool, author *string, commentFormat string, commentArgs ...interface{}) (scope *vacancyScope) {
	s.loggingEnabled = enableLogging
	s.loggingAuthor = author
	s.loggingComment = fmt.Sprintf(commentFormat, commentArgs...)

	return s
}

// Table returns Table object for that record.
func (s vacancy) Table() reform.Table {
	return vacancyTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s vacancy) PKValue() interface{} {
	return s.Id
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s vacancy) PKPointer() interface{} {
	return &s.Id
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s vacancy) HasPK() bool {
	return s.Id != vacancyTable.z[vacancyTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *VacancyFilter) SetPK(pk interface{}) { (*vacancy)(s).SetPK(pk) }
func (s *vacancy) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.Id = int(i64)
	} else {
		s.Id = pk.(int)
	}
}

var (
	// check interfaces
	_ reform.View   = vacancyTable
	_ reform.Struct = (*vacancy)(nil)
	_ reform.Table  = vacancyTable
	_ reform.Record = (*vacancy)(nil)
	_ fmt.Stringer  = (*vacancy)(nil)

	// querier
	Vacancy           = vacancy{} // Should be read only
	defaultDB_vacancy *reform.DB
)

func init() {
	//parse.AssertUpToDate(&vacancyTable.s, new(vacancy)) // Temporary disabled (doesn't work with arbitary types like "type sliceString []string")
}
