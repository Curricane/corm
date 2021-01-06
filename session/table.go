package session

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Curricane/corm/log"
	"github.com/Curricane/corm/schema"
)

// Model assigns refTable 更新refTable
func (s *Session) Model(value interface{}) *Session {
	// nil or different model, update refTable
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

// RefTable returns a Schema instance that contains all parsed fields
func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Errorf("Model is not set")
	}
	return s.refTable
}

// CreateTable create a table in database with a model
func (s *Session) CreateTable() error {
	table := s.RefTable()
	var columns []string
	for _, f := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", f.Name, f.Type, f.Tag))
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s);", table.Name, desc)).Exec()
	return err
}

// DropTable drops a table with the name of model
func (s *Session) DropTable() error {
	table := s.RefTable()
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", table.Name)).Exec()
	return err
}

// HasTable returns true of the table exists
func (s *Session) HasTable() bool {
	// 构建查询表名的sql 和 表名参数
	sql, args := s.dialect.TableExistSQL(s.RefTable().Name)
	row := s.Raw(sql, args...).QueryRow() // 查询
	var tmp string
	_ = row.Scan(&tmp) // 获取结果
	return tmp == s.RefTable().Name
}
