package main

import (
	"strings"
	"strconv"

	"database/sql"
	"time"
	
	"github.com/mackee/go-sqlla"
)

type exampleSQL struct {
	where sqlla.Where
}

func NewExampleSQL() exampleSQL {
	q := exampleSQL{}
	return q
}

var exampleAllColumns = []string{
	"id","name","created_at",
}

type exampleSelectSQL struct {
	exampleSQL
	Columns     []string
	order       string
	limit       *uint64
	isForUpdate bool
}

func (q exampleSQL) Select() exampleSelectSQL {
	return exampleSelectSQL{
		q,
		exampleAllColumns,
		"",
		nil,
		false,
	}
}

func (q exampleSelectSQL) Limit(l uint64) exampleSelectSQL {
	q.limit = &l
	return q
}

func (q exampleSelectSQL) ForUpdate() exampleSelectSQL {
	q.isForUpdate = true
	return q
}


func (q exampleSelectSQL) ID(v uint64, exprs ...sqlla.Operator) exampleSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint64{Value: v, Op: op, Column: "id"}
	q.where = append(q.where, where)
	return q
}

func (q exampleSelectSQL) IDIn(vs ...uint64) exampleSelectSQL {
	where := sqlla.ExprMultiUint64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "id"}
	q.where = append(q.where, where)
	return q
}



func (q exampleSelectSQL) OrderByID(order sqlla.Order) exampleSelectSQL {
	q.order = " ORDER BY id"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q exampleSelectSQL) Name(v string, exprs ...sqlla.Operator) exampleSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "name"}
	q.where = append(q.where, where)
	return q
}

func (q exampleSelectSQL) NameIn(vs ...string) exampleSelectSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "name"}
	q.where = append(q.where, where)
	return q
}



func (q exampleSelectSQL) OrderByName(order sqlla.Order) exampleSelectSQL {
	q.order = " ORDER BY name"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q exampleSelectSQL) CreatedAt(v time.Time, exprs ...sqlla.Operator) exampleSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprTime{Value: v, Op: op, Column: "created_at"}
	q.where = append(q.where, where)
	return q
}

func (q exampleSelectSQL) CreatedAtIn(vs ...time.Time) exampleSelectSQL {
	where := sqlla.ExprMultiTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "created_at"}
	q.where = append(q.where, where)
	return q
}



func (q exampleSelectSQL) OrderByCreatedAt(order sqlla.Order) exampleSelectSQL {
	q.order = " ORDER BY created_at"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q exampleSelectSQL) ToSql() (string, []interface{}, error) {
	columns := strings.Join(q.Columns, ", ")
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "SELECT " + columns + " FROM example"
	if wheres != "" {
		query += " WHERE" + wheres
	}
	query += q.order
	if q.limit != nil {
		query += " LIMIT " + strconv.FormatUint(*q.limit, 10)
	}

	if q.isForUpdate {
		query += " FOR UPDATE"
	}

	return query + ";", vs, nil
}

func (q exampleSelectSQL) Single(db sqlla.DB) (Example, error) {
	q.Columns = exampleAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return Example{}, err
	}

	row := db.QueryRow(query, args...)
	return q.Scan(row)
}

func (q exampleSelectSQL) All(db sqlla.DB) ([]Example, error) {
	rs := make([]Example, 0, 10)
	q.Columns = exampleAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		r, err := q.Scan(rows)
		if err != nil {
			return nil, err
		}
		rs = append(rs, r)
	}
	return rs, nil
}

func (q exampleSelectSQL) Scan(s sqlla.Scanner) (Example, error) {
	var row Example
	err := s.Scan(
		&row.ID,
		&row.Name,
		&row.CreatedAt,
		
	)
	return row, err
}

type exampleUpdateSQL struct {
	exampleSQL
	setMap	sqlla.SetMap
	Columns []string
}

func (q exampleSQL) Update() exampleUpdateSQL {
	return exampleUpdateSQL{
		exampleSQL: q,
		setMap: sqlla.SetMap{},
	}
}


func (q exampleUpdateSQL) SetID(v uint64) exampleUpdateSQL {
	q.setMap["id"] = v
	return q
}

func (q exampleUpdateSQL) WhereID(v uint64, exprs ...sqlla.Operator) exampleUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint64{Value: v, Op: op, Column: "id"}
	q.where = append(q.where, where)
	return q
}


func (q exampleUpdateSQL) SetName(v string) exampleUpdateSQL {
	q.setMap["name"] = v
	return q
}

func (q exampleUpdateSQL) WhereName(v string, exprs ...sqlla.Operator) exampleUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "name"}
	q.where = append(q.where, where)
	return q
}


func (q exampleUpdateSQL) SetCreatedAt(v time.Time) exampleUpdateSQL {
	q.setMap["created_at"] = v
	return q
}

func (q exampleUpdateSQL) WhereCreatedAt(v time.Time, exprs ...sqlla.Operator) exampleUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprTime{Value: v, Op: op, Column: "created_at"}
	q.where = append(q.where, where)
	return q
}


func (q exampleUpdateSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = Example{}
	if t, ok := s.(exampleDefaultUpdateHooker); ok {
		q, err = t.DefaultUpdateHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}
	setColumns, svs, err := q.setMap.ToUpdateSql()
	if err != nil {
		return "", []interface{}{}, err
	}
	wheres, wvs, err := q.where.ToSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	query := "UPDATE example SET" + setColumns
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", append(svs, wvs...), nil
}
func (q exampleUpdateSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

type exampleDefaultUpdateHooker interface {
	DefaultUpdateHook(exampleUpdateSQL) (exampleUpdateSQL, error)
}

type exampleInsertSQL struct {
	exampleSQL
	setMap	sqlla.SetMap
	Columns []string
}

func (q exampleSQL) Insert() exampleInsertSQL {
	return exampleInsertSQL{
		exampleSQL: q,
		setMap: sqlla.SetMap{},
	}
}


func (q exampleInsertSQL) ValueID(v uint64) exampleInsertSQL {
	q.setMap["id"] = v
	return q
}


func (q exampleInsertSQL) ValueName(v string) exampleInsertSQL {
	q.setMap["name"] = v
	return q
}


func (q exampleInsertSQL) ValueCreatedAt(v time.Time) exampleInsertSQL {
	q.setMap["created_at"] = v
	return q
}


func (q exampleInsertSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = Example{}
	if t, ok := s.(exampleDefaultInsertHooker); ok {
		q, err = t.DefaultInsertHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}
	qs, vs, err := q.setMap.ToInsertSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	query := "INSERT INTO example " + qs

	return query + ";", vs, nil
}

func (q exampleInsertSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		
		return nil, err
	}
	result, err := db.Exec(query, args...)
	return result, err
}

type exampleDefaultInsertHooker interface {
	DefaultInsertHook(exampleInsertSQL) (exampleInsertSQL, error)
}

type exampleDeleteSQL struct {
	exampleSQL
}

func (q exampleSQL) Delete() exampleDeleteSQL {
	return exampleDeleteSQL{
		q,
	}
}


func (q exampleDeleteSQL) ID(v uint64, exprs ...sqlla.Operator) exampleDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint64{Value: v, Op: op, Column: "id"}
	q.where = append(q.where, where)
	return q
}


func (q exampleDeleteSQL) IDIn(vs ...uint64) exampleDeleteSQL {
	where := sqlla.ExprMultiUint64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "id"}
	q.where = append(q.where, where)
	return q
}

func (q exampleDeleteSQL) Name(v string, exprs ...sqlla.Operator) exampleDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "name"}
	q.where = append(q.where, where)
	return q
}


func (q exampleDeleteSQL) NameIn(vs ...string) exampleDeleteSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "name"}
	q.where = append(q.where, where)
	return q
}

func (q exampleDeleteSQL) CreatedAt(v time.Time, exprs ...sqlla.Operator) exampleDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprTime{Value: v, Op: op, Column: "created_at"}
	q.where = append(q.where, where)
	return q
}


func (q exampleDeleteSQL) CreatedAtIn(vs ...time.Time) exampleDeleteSQL {
	where := sqlla.ExprMultiTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "created_at"}
	q.where = append(q.where, where)
	return q
}

func (q exampleDeleteSQL) ToSql() (string, []interface{}, error) {
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "DELETE FROM example"
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", vs, nil
}

func ( q exampleDeleteSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

