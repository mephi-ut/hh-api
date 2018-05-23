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

type userScope struct {
	item *user

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
type UserType user
type UserF user
type UserFilter user

type userLogRow struct {
	user
	LogAuthor  *string
	LogAction  string
	LogDate    time.Time
	LogComment string
}

// Schema returns a schema name in SQL database ("").
type userTableTypeType struct {
	s reform.StructInfo
	z []interface{}
}

func (v userTableTypeType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("users").
func (v userTableTypeType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v userTableTypeType) Columns() []string {
	return []string{"id", "nickname", "email", "password_hash"}
}

// NewStruct makes a new struct for that view or table.
func (v userTableTypeType) NewStruct() reform.Struct {
	return new(user)
}

// NewRecord makes a new record for that table.
func (v *userTableTypeType) NewRecord() reform.Record {
	return new(user)
}

func (v *userTableTypeType) NewScope() *userScope {
	return &userScope{item: &user{}}
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *userTableTypeType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

func (v userTableTypeType) CreateTableIfNotExists(db *reform.DB) (bool, error) {
	if db == nil {
		db = defaultDB_user
	}
	return db.CreateTableIfNotExists(v.s)
}

func (v userTableTypeType) StructInfo() reform.StructInfo {
	return v.s
}

// userTable represents users view or table in SQL database.
var userTable = &userTableTypeType{
	s: reform.StructInfo{Type: "user", SQLSchema: "", SQLName: "users", Fields: []reform.FieldInfo{{Name: "Id", IsPK: true, IsUnique: false, HasIndex: false, Type: "int", Column: "id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "Nickname", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "nickname", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "Email", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "email", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "PasswordHash", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "password_hash", FieldsPath: []reform.FieldInfo{}, SQLSize: 255, Embedded: "", StructFile: ""}}, PKFieldIndex: 0, ImitateGorm: false, SkipMethodOrder: false},
	z: new(user).Values(),
}

type userTableTypeType_log struct {
	s reform.StructInfo
	z []interface{}
}

func (v *userTableTypeType_log) Schema() string {
	return v.s.SQLSchema
}

func (v *userTableTypeType_log) Name() string {
	return v.s.SQLName
}

func (v *userTableTypeType_log) Columns() []string {
	return []string{"id", "nickname", "email", "password_hash", "log_author", "log_action", "log_date", "log_comment"}
}

func (v *userTableTypeType_log) NewStruct() reform.Struct {
	return new(user)
}

func (v *userTableTypeType_log) NewRecord() reform.Record {
	return new(user)
}

func (v *userTableTypeType_log) NewScope() *userScope {
	return &userScope{item: &user{}}
}

func (v *userTableTypeType_log) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

var userTableLogRow = &userTableTypeType_log{
	s: reform.StructInfo{Type: "user", SQLSchema: "", SQLName: "users_log", Fields: []reform.FieldInfo{{Name: "Id", IsPK: true, IsUnique: false, HasIndex: false, Type: "int", Column: "id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "Nickname", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "nickname", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "Email", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "email", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "PasswordHash", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "password_hash", FieldsPath: []reform.FieldInfo{}, SQLSize: 255, Embedded: "", StructFile: ""}, {Name: "LogAuthor", IsPK: false, IsUnique: false, HasIndex: false, Type: "*string", Column: "log_author", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogAction", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "log_action", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogDate", IsPK: false, IsUnique: false, HasIndex: false, Type: "time.Time", Column: "log_date", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogComment", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "log_comment", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}}, PKFieldIndex: 0, ImitateGorm: false, SkipMethodOrder: false},
	z: new(userLogRow).Values(),
}

func (s userTableTypeType) ColumnNameByFieldName(fieldName string) string {
	switch fieldName {
	case "Id":
		return "id"
	case "Nickname":
		return "nickname"
	case "Email":
		return "email"
	case "PasswordHash":
		return "password_hash"
	}
	return ""
}

func (s userTableTypeType_log) ColumnNameByFieldName(fieldName string) string {
	switch fieldName {
	case "Id":
		return "id"
	case "Nickname":
		return "nickname"
	case "Email":
		return "email"
	case "PasswordHash":
		return "password_hash"
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

func (s *user) FieldPointersByNames(fieldNames []string) (fieldPointers []interface{}) {
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

func (s *userLogRow) FieldPointersByNames(fieldNames []string) (fieldPointers []interface{}) {
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

func (s *user) FieldPointerByName(fieldName string) interface{} {
	switch fieldName {
	case "Id":
		return &s.Id
	case "Nickname":
		return &s.Nickname
	case "Email":
		return &s.Email
	case "PasswordHash":
		return &s.PasswordHash
	}

	return nil
}

func (s *userLogRow) FieldPointerByName(fieldName string) interface{} {
	switch fieldName {
	case "Id":
		return &s.Id
	case "Nickname":
		return &s.Nickname
	case "Email":
		return &s.Email
	case "PasswordHash":
		return &s.PasswordHash
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
func (s user) String() string {
	res := make([]string, 4)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "Nickname: " + reform.Inspect(s.Nickname, true)
	res[2] = "Email: " + reform.Inspect(s.Email, true)
	res[3] = "PasswordHash: " + reform.Inspect(s.PasswordHash, true)
	return strings.Join(res, ", ")
}
func (s userLogRow) String() string {
	res := make([]string, 8)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "Nickname: " + reform.Inspect(s.Nickname, true)
	res[2] = "Email: " + reform.Inspect(s.Email, true)
	res[3] = "PasswordHash: " + reform.Inspect(s.PasswordHash, true)
	res[4] = "LogAuthor: " + reform.Inspect(s.LogAuthor, true)
	res[5] = "LogAction: " + reform.Inspect(s.LogAction, true)
	res[6] = "LogDate: " + reform.Inspect(s.LogDate, true)
	res[7] = "LogComment: " + reform.Inspect(s.LogComment, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *user) Values() []interface{} {
	return []interface{}{
		s.Id,
		s.Nickname,
		s.Email,
		s.PasswordHash,
	}
}
func (s *userLogRow) Values() []interface{} {
	return append(s.user.Values(), []interface{}{
		s.LogAuthor,
		s.LogAction,
		s.LogDate,
		s.LogComment,
	}...)
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *user) Pointers() []interface{} {
	return []interface{}{
		&s.Id,
		&s.Nickname,
		&s.Email,
		&s.PasswordHash,
	}
}
func (s *userLogRow) Pointers() []interface{} {
	return append(s.user.Pointers(), []interface{}{
		&s.LogAuthor,
		&s.LogAction,
		&s.LogDate,
		&s.LogComment,
	}...)
}

// View returns View object for that struct.
func (s user) View() reform.View {
	return userTable
}
func (s userScope) View() reform.View {
	return s.item.View()
}
func (s userLogRow) View() reform.View {
	return userTableLogRow
}

// Generate a scope for object
func (s user) Scope() *userScope {
	return &userScope{item: &s, db: defaultDB_user}
}
func (s *user) PtrScope() *userScope {
	return &userScope{item: s, db: defaultDB_user}
}

// Sets DB to do queries
func (s user) DB(db reform.ReformDBTX) (scope *userScope) { return s.Scope().DB(db) }
func (s *userScope) DB(db reform.ReformDBTX) *userScope {
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
func (s user) GetDB() (db *reform.DB) { return s.Scope().GetDB() }
func (s userScope) GetDB() *reform.DB {
	return s.db.(*reform.DB)
}

func (s user) StartTransaction() (*reform.TX, error) { return s.Scope().StartTransaction() }
func (s userScope) StartTransaction() (*reform.TX, error) {
	return s.db.(*reform.DB).Begin()
}

// Sets default DB (to do not call the scope.DB() method every time)
func (s *user) SetDefaultDB(db *reform.DB) (err error) {
	defaultDB_user = db
	return nil
}

// Compiles SQL tail for defined limit scope
// TODO: should be compiled via dialects
func (s *userScope) getLimitTail() (tail string, args []interface{}, err error) {
	if s.limit <= 0 {
		return
	}

	tail = fmt.Sprintf("%v", s.limit)
	return
}

// Compiles SQL tail for defined group scope
// TODO: should be compiled via dialects
func (s *userScope) getGroupTail() (tail string, args []interface{}, err error) {
	tail = strings.Join(s.groupBy, ", ")

	return
}

// Compiles SQL tail for defined order scope
// TODO: should be compiled via dialects
func (s *userScope) getOrderTail() (tail string, args []interface{}, err error) {
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
func (s *userScope) getWhereTailForFilter(filter UserFilter) (tail string, whereTailArgs []interface{}, err error) {
	return s.db.GetWhereTailForFilter(user(filter), nil, "", false)
}

// parseQuerierArgs considers different ways of defning the tail (using scope properties or/and in_args)
func (s userScope) parseWhereTailComponent(in_args []interface{}, placeholderCounter *int) (tail string, args []interface{}, err error) {
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
		case *user:
			in_args[0] = *arg
			return s.parseWhereTailComponent(in_args, placeholderCounter)
		case *UserF:
			in_args[0] = *arg
			return s.parseWhereTailComponent(in_args, placeholderCounter)
		case *UserFilter:
			in_args[0] = *arg
			return s.parseWhereTailComponent(in_args, placeholderCounter)
		case user:
			if len(in_args) > 1 {
				s = *s.Where(in_args[1], in_args[2:]...)
			}
			tail, args, err = s.getWhereTailForFilter(UserFilter(arg))
		case UserF:
			if len(in_args) > 1 {
				s = *s.Where(in_args[1], in_args[2:]...)
			}
			tail, args, err = s.getWhereTailForFilter(UserFilter(arg))
		case UserFilter:
			if len(in_args) > 1 {
				s = *s.Where(in_args[1], in_args[2:]...)
			}
			tail, args, err = s.getWhereTailForFilter(arg)
		default:
			err = fmt.Errorf("Invalid first element of \"in_args\" (%T). It should be a string or UserFilter.", arg)
			return
		}
	}

	return
}

// Compiles SQL tail for defined filter
// TODO: should be compiled via dialects
func (s *userScope) getWhereTail() (tail string, whereTailArgs []interface{}, err error) {
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

func (s user) Where(requiredArg interface{}, args ...interface{}) (scope *userScope) {
	return s.Scope().Where(requiredArg, args...)
}
func (s userScope) Where(requiredArg interface{}, in_args ...interface{}) *userScope {
	s.where = append(s.where, append([]interface{}{requiredArg}, in_args...))
	return &s
}
func (s userScope) SetWhere(where [][]interface{}) *userScope {
	s.where = where
	return &s
}
func (s userScope) GetWhere() [][]interface{} {
	return s.where
}

// Sets all scope-related parameters to be equal as in passed scope (as an argument)
func (s userScope) SetScope(anotherScope reform.Scope) *userScope {
	s.where = anotherScope.GetWhere()
	s.order = anotherScope.GetOrder()
	s.groupBy = anotherScope.GetGroup()
	s.limit = anotherScope.GetLimit()
	s.db = anotherScope.GetDB()

	return &s
}
func (s userScope) ISetScope(anotherScope reform.Scope) reform.Scope {
	return s.ISetScope(anotherScope)
}

// Compiles SQL tail for defined db/where/order/limit scope
// TODO: should be compiled via dialects
func (s *userScope) getTail() (tail string, args []interface{}, err error) {
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
func (s user) SelectRows(query string, args ...interface{}) (rows *sql.Rows, err error) {
	return s.Scope().SelectRows(query, args...)
}
func (s *userScope) SelectRows(query string, queryArgs ...interface{}) (rows *sql.Rows, err error) {
	tail, args, err := s.getTail()
	if err != nil {
		return
	}

	return s.db.Query("SELECT "+query+" FROM `"+userTable.Name()+"` "+tail, append(queryArgs, args...)...)
}

func (s *userScope) callStructMethod(str *user, methodName string) error {
	if method := reflect.ValueOf(str).MethodByName(methodName); method.IsValid() {
		switch f := method.Interface().(type) {
		case func():
			f()

		case func(reform.ReformDBTX):
			f(s.db)

		case func(*userScope):
			f(s)

		case func(interface{}): // For compatibility with other ORMs
			f(s.db)

		case func() error:
			return f()

		case func(reform.ReformDBTX) error:
			return f(s.db)

		case func(*userScope) error:
			return f(s)

		case func(interface{}) error: // For compatibility with other ORMS
			return f(s.db)

		default:
			panic("Unknown type of method: \"" + methodName + "\"")
		}
	}
	return nil
}

func (s userScope) checkDb() {
	if s.db == nil {
		panic("s.db == nil")
	}
}

// Select is a handy wrapper for SelectRows() and NextRow(): it makes a query and collects the result into a slice
func (s user) Select(args ...interface{}) (result []user, err error) { return s.Scope().Select(args...) }
func (s userScope) Select(args ...interface{}) (result []user, err error) {
	s.checkDb()

	if len(args) > 0 {
		s = *s.Where(args[0], args[1:]...)
	}
	tail, args, err := s.getTail()
	if err != nil {
		return
	}

	rows, err := s.db.FlexSelectRows(userTable, s.tableQuery, s.fieldsFilter, tail, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		item := user{}
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
func (s user) SelectI(args ...interface{}) (result interface{}, err error) {
	return s.Scope().Select(args...)
}
func (s userScope) SelectI(args ...interface{}) (result interface{}, err error) {
	return s.Select(args...)
}

// "First" a method to select and return only one record.
func (s user) First(args ...interface{}) (result user, err error) { return s.Scope().First(args...) }
func (s userScope) First(args ...interface{}) (result user, err error) {
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
func (s user) FirstI(args ...interface{}) (result interface{}, err error) {
	return s.Scope().First(args...)
}
func (s userScope) FirstI(args ...interface{}) (result interface{}, err error) {
	return s.First(args...)
}

// Sets "GROUP BY".
func (s user) Group(args ...interface{}) (scope *userScope) { return s.Scope().Group(args...) }
func (s userScope) Group(argsI ...interface{}) *userScope {
	for _, argI := range argsI {
		s.groupBy = append(s.groupBy, argI.(string))
	}

	return &s
}
func (s userScope) SetGroup(groupBy []string) *userScope {
	s.groupBy = groupBy
	return &s
}
func (s userScope) GetGroup() []string {
	return s.groupBy
}

// Sets a table query. For example SetTableQuery("table1 JOIN table2 USING(key)")
func (s user) SetTableQuery(query string) (scope *userScope) { return s.Scope().SetTableQuery(query) }
func (s userScope) SetTableQuery(query string) *userScope {
	if query == "" {
		s.tableQuery = nil
	} else {
		s.tableQuery = &query
	}
	return &s
}
func (s userScope) GetTableQuery() string {
	if s.tableQuery != nil {
		return *s.tableQuery
	}
	return s.db.QualifiedView(s.View())
}

// Sets which structure fields should be queried while Select()/First(). For example SetFields("StructField1", "StructIdField", "StructCommentsField"). Could be used just to speed up a query.
// It's not recommended to use this function!
func (s user) SetQueryFieldsByNames(fields ...string) (scope *userScope) {
	return s.Scope().SetQueryFieldsByNames(fields...)
}
func (s userScope) SetQueryFieldsByNames(fields ...string) *userScope {
	s.fieldsFilter = fields
	return &s
}
func (s userScope) GetQueryFields() []string {
	return s.fieldsFilter
}

// Sets order. Arguments should be passed by pairs column-{ASC,DESC}. For example Order("id", "ASC", "value" "DESC")
func (s user) Order(args ...interface{}) (scope *userScope) { return s.Scope().Order(args...) }
func (s userScope) Order(argsI ...interface{}) *userScope {
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
func (s userScope) SetOrder(order []string) *userScope {
	s.order = order
	return &s
}
func (s userScope) GetOrder() []string {
	return s.order
}

func (s user) SetSQLAppend(appendTail string) (scope *userScope) {
	return s.Scope().SetSQLAppend(appendTail)
}
func (s userScope) SetSQLAppend(appendTail string) *userScope {
	s.appendTail = appendTail
	return &s
}

// Sets limit.
func (s user) Limit(limit int) (scope *userScope) { return s.Scope().Limit(limit) }
func (s *userScope) Limit(limit int) *userScope {
	s.limit = limit
	return s
}

// Gets limit
func (s userScope) GetLimit() int {
	return s.limit
}

// "Reload" reloads record using Primary Key
func (s *UserFilter) Reload(db *reform.DB) error { return (*user)(s).Reload(db) }
func (s *user) Reload(db *reform.DB) (err error) {
	return db.FindByPrimaryKeyTo(s, s.PKValue())
}

// Create and Insert inserts new record to DB
func (s *user) Create() (err error) { return s.PtrScope().Create() }
func (s *userScope) Create() (err error) {
	return s.Insert()
}
func (s *user) Insert() (err error) { return s.PtrScope().Insert() }
func (s *userScope) Insert() (err error) {
	s.checkDb()
	err = s.db.Insert(s.item)
	if err == nil {
		s.doLog("INSERT")
	}
	return err
}

// Replace "REPLACE INTO" new record to DB
func (s *user) Replace() (err error) { return s.PtrScope().Replace() }
func (s *userScope) Replace() (err error) {
	s.checkDb()
	err = s.db.Replace(s.item)
	if err == nil {
		s.doLog("REPLACE")
	}
	return err
}

// Save inserts new record to DB is PK is zero and updates existing record if PK is not zero
func (s *user) Save() (err error) { return s.PtrScope().Save() }
func (s *userScope) Save() (err error) {
	s.checkDb()
	err = s.db.Save(s.item)
	if err == nil {
		s.doLog("INSERT")
	}
	return err
}

// Update updates existing record in DB
func (s user) Update() (err error) { return s.Scope().Update() }
func (s *userScope) Update() (err error) {
	s.checkDb()
	err = s.db.Update(s.item)
	if err == nil {
		s.doLog("UPDATE")
	}
	return err
}

// Delete deletes existing record in DB
func (s user) Delete() (err error) { return s.Scope().Delete() }
func (s *userScope) Delete() (err error) {
	s.checkDb()
	err = s.db.Delete(s.item)
	if err == nil {
		s.doLog("DELETE")
	}
	return err
}

func (s *userScope) doLog(requestType string) {
	if !s.loggingEnabled {
		return
	}

	var logRow userLogRow
	logRow.user = *s.item
	logRow.LogAuthor = s.loggingAuthor
	logRow.LogAction = requestType
	logRow.LogDate = time.Now()
	logRow.LogComment = s.loggingComment

	s.db.Insert(&logRow)
}

// Enables logging to table "users_log". This table should has the same schema, except:
// - Unique/Primary keys should be removed
// - Should be added next fields: "log_author" (nullable string), "log_date" (timestamp), "log_action" (enum("INSERT", "UPDATE", "DELETE")), "log_comment" (string)
func (s *user) Log(enableLogging bool, author *string, commentFormat string, commentArgs ...interface{}) (scope *userScope) {
	return s.Scope().Log(enableLogging, author, commentFormat, commentArgs...)
}
func (s *userScope) Log(enableLogging bool, author *string, commentFormat string, commentArgs ...interface{}) (scope *userScope) {
	s.loggingEnabled = enableLogging
	s.loggingAuthor = author
	s.loggingComment = fmt.Sprintf(commentFormat, commentArgs...)

	return s
}

// Table returns Table object for that record.
func (s user) Table() reform.Table {
	return userTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s user) PKValue() interface{} {
	return s.Id
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s user) PKPointer() interface{} {
	return &s.Id
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s user) HasPK() bool {
	return s.Id != userTable.z[userTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *UserFilter) SetPK(pk interface{}) { (*user)(s).SetPK(pk) }
func (s *user) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.Id = int(i64)
	} else {
		s.Id = pk.(int)
	}
}

var (
	// check interfaces
	_ reform.View   = userTable
	_ reform.Struct = (*user)(nil)
	_ reform.Table  = userTable
	_ reform.Record = (*user)(nil)
	_ fmt.Stringer  = (*user)(nil)

	// querier
	User           = user{} // Should be read only
	defaultDB_user *reform.DB
)

func init() {
	//parse.AssertUpToDate(&userTable.s, new(user)) // Temporary disabled (doesn't work with arbitary types like "type sliceString []string")
}
