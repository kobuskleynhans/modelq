// Code generated by ModelQ
// user.go contains model for the database table [blog.user]

package mysql

import (
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/mijia/modelq/gmq"
	"strings"
	"time"
)

type User struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	IsMarried  int       `json:"is_married"`
	Age        int       `json:"age"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// Start of the User APIs.

func (obj User) String() string {
	if data, err := json.Marshal(obj); err != nil {
		return fmt.Sprintf("<User Id=%d>", obj.Id)
	} else {
		return string(data)
	}
}

func (obj User) Insert(dbtx gmq.DbTx) (User, error) {
	if result, err := UserObjs.Insert(obj).Run(dbtx); err != nil {
		return obj, err
	} else {
		if dbtx.DriverName() != "postgres" {
			if id, err := result.LastInsertId(); err != nil {
				return obj, err
			} else {
				obj.Id = id
				return obj, err
			}
		}
		return obj, nil
	}
}

func (obj User) Update(dbtx gmq.DbTx) (int64, error) {
	fields := []string{"Name", "Password", "IsMarried", "Age", "CreateTime"}
	filter := UserObjs.FilterId("=", obj.Id)
	if result, err := UserObjs.Update(obj, fields...).Where(filter).Run(dbtx); err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (obj User) Delete(dbtx gmq.DbTx) (int64, error) {
	filter := UserObjs.FilterId("=", obj.Id)
	if result, err := UserObjs.Delete().Where(filter).Run(dbtx); err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

// Start of the inner Query Api

type _UserQuery struct {
	gmq.Query
}

func (q _UserQuery) Where(f gmq.Filter) _UserQuery {
	q.Query = q.Query.Where(f)
	return q
}

func (q _UserQuery) OrderBy(by ...string) _UserQuery {
	tBy := make([]string, 0, len(by))
	for _, b := range by {
		sortDir := ""
		if b[0] == '-' || b[0] == '+' {
			sortDir = string(b[0])
			b = b[1:]
		}
		if col, ok := UserObjs.fcMap[b]; ok {
			tBy = append(tBy, sortDir+col)
		}
	}
	q.Query = q.Query.OrderBy(tBy...)
	return q
}

func (q _UserQuery) GroupBy(by ...string) _UserQuery {
	tBy := make([]string, 0, len(by))
	for _, b := range by {
		if col, ok := UserObjs.fcMap[b]; ok {
			tBy = append(tBy, col)
		}
	}
	q.Query = q.Query.GroupBy(tBy...)
	return q
}

func (q _UserQuery) Limit(offsets ...int64) _UserQuery {
	q.Query = q.Query.Limit(offsets...)
	return q
}

func (q _UserQuery) Page(number, size int) _UserQuery {
	q.Query = q.Query.Page(number, size)
	return q
}

func (q _UserQuery) Run(dbtx gmq.DbTx) (sql.Result, error) {
	return q.Query.Exec(dbtx)
}

type UserRowVisitor func(obj User) bool

func (q _UserQuery) Iterate(dbtx gmq.DbTx, functor UserRowVisitor) error {
	return q.Query.SelectList(dbtx, func(columns []gmq.Column, rb []sql.RawBytes) bool {
		obj := UserObjs.toUser(columns, rb)
		return functor(obj)
	})
}

func (q _UserQuery) One(dbtx gmq.DbTx) (User, error) {
	var obj User
	err := q.Query.SelectOne(dbtx, func(columns []gmq.Column, rb []sql.RawBytes) bool {
		obj = UserObjs.toUser(columns, rb)
		return true
	})
	return obj, err
}

func (q _UserQuery) List(dbtx gmq.DbTx) ([]User, error) {
	result := make([]User, 0, 10)
	err := q.Query.SelectList(dbtx, func(columns []gmq.Column, rb []sql.RawBytes) bool {
		obj := UserObjs.toUser(columns, rb)
		result = append(result, obj)
		return true
	})
	return result, err
}

// Start of the model facade Apis.

type _UserObjs struct {
	fcMap map[string]string
}

func (o _UserObjs) Names() (schema, tbl, alias string) {
	return "blog", "user", "User"
}

func (o _UserObjs) Select(fields ...string) _UserQuery {
	q := _UserQuery{}
	if len(fields) == 0 {
		fields = []string{"Id", "Name", "Password", "IsMarried", "Age", "CreateTime", "UpdateTime"}
	}
	q.Query = gmq.Select(o, o.columns(fields...))
	return q
}

func (o _UserObjs) Insert(obj User) _UserQuery {
	q := _UserQuery{}
	q.Query = gmq.Insert(o, o.columnsWithData(obj, "Name", "Password", "IsMarried", "Age"))
	return q
}

func (o _UserObjs) Update(obj User, fields ...string) _UserQuery {
	q := _UserQuery{}
	q.Query = gmq.Update(o, o.columnsWithData(obj, fields...))
	return q
}

func (o _UserObjs) Delete() _UserQuery {
	q := _UserQuery{}
	q.Query = gmq.Delete(o)
	return q
}

///// Managed Objects Filters definition

func (o _UserObjs) FilterId(op string, p int64, ps ...int64) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("id", op, params...)
}

func (o _UserObjs) FilterName(op string, p string, ps ...string) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("name", op, params...)
}

func (o _UserObjs) FilterPassword(op string, p string, ps ...string) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("password", op, params...)
}

func (o _UserObjs) FilterIsMarried(op string, p int, ps ...int) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("is_married", op, params...)
}

func (o _UserObjs) FilterAge(op string, p int, ps ...int) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("age", op, params...)
}

func (o _UserObjs) FilterCreateTime(op string, p time.Time, ps ...time.Time) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("create_time", op, params...)
}

func (o _UserObjs) FilterUpdateTime(op string, p time.Time, ps ...time.Time) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("update_time", op, params...)
}

///// Managed Objects Columns definition

func (o _UserObjs) ColumnId(p ...int64) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"id", value}
}

func (o _UserObjs) ColumnName(p ...string) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"name", value}
}

func (o _UserObjs) ColumnPassword(p ...string) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"password", value}
}

func (o _UserObjs) ColumnIsMarried(p ...int) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"is_married", value}
}

func (o _UserObjs) ColumnAge(p ...int) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"age", value}
}

func (o _UserObjs) ColumnCreateTime(p ...time.Time) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"create_time", value}
}

func (o _UserObjs) ColumnUpdateTime(p ...time.Time) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"update_time", value}
}

////// Internal helper funcs

func (o _UserObjs) newFilter(name, op string, params ...interface{}) gmq.Filter {
	if strings.ToUpper(op) == "IN" {
		return gmq.InFilter(name, params)
	}
	return gmq.UnitFilter(name, op, params[0])
}

func (o _UserObjs) toUser(columns []gmq.Column, rb []sql.RawBytes) User {
	obj := User{}
	if len(columns) == len(rb) {
		for i := range columns {
			switch columns[i].Name {
			case "id":
				obj.Id = gmq.AsInt64(rb[i])
			case "name":
				obj.Name = gmq.AsString(rb[i])
			case "password":
				obj.Password = gmq.AsString(rb[i])
			case "is_married":
				obj.IsMarried = gmq.AsInt(rb[i])
			case "age":
				obj.Age = gmq.AsInt(rb[i])
			case "create_time":
				obj.CreateTime = gmq.AsTime(rb[i])
			case "update_time":
				obj.UpdateTime = gmq.AsTime(rb[i])
			}
		}
	}
	return obj
}

func (o _UserObjs) columns(fields ...string) []gmq.Column {
	data := make([]gmq.Column, 0, len(fields))
	for _, f := range fields {
		switch f {
		case "Id":
			data = append(data, o.ColumnId())
		case "Name":
			data = append(data, o.ColumnName())
		case "Password":
			data = append(data, o.ColumnPassword())
		case "IsMarried":
			data = append(data, o.ColumnIsMarried())
		case "Age":
			data = append(data, o.ColumnAge())
		case "CreateTime":
			data = append(data, o.ColumnCreateTime())
		case "UpdateTime":
			data = append(data, o.ColumnUpdateTime())
		}
	}
	return data
}

func (o _UserObjs) columnsWithData(obj User, fields ...string) []gmq.Column {
	data := make([]gmq.Column, 0, len(fields))
	for _, f := range fields {
		switch f {
		case "Id":
			data = append(data, o.ColumnId(obj.Id))
		case "Name":
			data = append(data, o.ColumnName(obj.Name))
		case "Password":
			data = append(data, o.ColumnPassword(obj.Password))
		case "IsMarried":
			data = append(data, o.ColumnIsMarried(obj.IsMarried))
		case "Age":
			data = append(data, o.ColumnAge(obj.Age))
		case "CreateTime":
			data = append(data, o.ColumnCreateTime(obj.CreateTime))
		case "UpdateTime":
			data = append(data, o.ColumnUpdateTime(obj.UpdateTime))
		}
	}
	return data
}

var UserObjs _UserObjs

func init() {
	UserObjs.fcMap = map[string]string{
		"Id":         "id",
		"Name":       "name",
		"Password":   "password",
		"IsMarried":  "is_married",
		"Age":        "age",
		"CreateTime": "create_time",
		"UpdateTime": "update_time",
	}
	gob.Register(User{})
}
