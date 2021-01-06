package session

import (
	"database/sql"
	"strings"

	"github.com/Curricane/corm/log"
)

// Session 与数据库交互
type Session struct {
	db      *sql.DB
	sql     strings.Builder // sql语句
	sqlVars []interface{}   // 存放占位符对应的值
}

// New creates a instance of Session
func New(db *sql.DB) *Session {
	return &Session{db: db}
}

// Clear initialize the state of a session
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

// DB returns *sql.DB
func (s *Session) DB() *sql.DB {
	return s.db
}

// Exec raw sql with sqlVals
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow gets a record from db
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// QueryRows gets a list of records from db
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// Raw appends sql and sqlVars
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}
