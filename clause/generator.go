package clause

import (
	"fmt"
	"strings"

	"github.com/Curricane/corm/log"
)

// 实现各个子句

type generator func(values ...interface{}) (string, []interface{})

var generators map[Type]generator

func init() {
	generators = make(map[Type]generator)
	generators[INSERT] = _insert
	generators[VALUES] = _values
	generators[SELECT] = _select
	generators[LIMIT] = _limit
	generators[WHERE] = _where
	generators[ORDERBY] = _orderBy
	generators[UPDATE] = _update
	generators[DELETE] = _delete
	generators[COUNT] = _count
}

// 产生 "?, ?, ?"
func genBindVars(num int) string {
	var vars []string
	for i := 0; i < num; i++ {
		vars = append(vars, "?")
	}
	return strings.Join(vars, ", ")
}

//
func _insert(values ...interface{}) (string, []interface{}) {
	// INSERT INTO $tableName ($fields)
	log.Infof("values is %#v", values)

	tableName := values[0]
	fields := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("INSERT INTO %s (%v)", tableName, fields), []interface{}{}
}

/*
values 可以这么被调用，因此要处理两层slice
INSERT INTO table_name(col1, col2, col3, ...) VALUES
    (A1, A2, A3, ...),
    (B1, B2, B3, ...),
    ...
*/
func _values(values ...interface{}) (string, []interface{}) {
	// VALUES ($v1), ($v2), ...
	log.Infof("values is %#v", values)
	var bindStr string
	var sql strings.Builder
	var vars []interface{}
	sql.WriteString("VALUES ")
	for i, value := range values {
		v := value.([]interface{})
		if bindStr == "" {
			bindStr = genBindVars(len(v))
		}
		sql.WriteString(fmt.Sprintf("(%v)", bindStr))
		if i+1 != len(values) {
			sql.WriteString(", ")
		}
		vars = append(vars, v...)
	}
	return sql.String(), vars
}

func _select(values ...interface{}) (string, []interface{}) {
	// SELECT $fields FROM $tableName
	log.Infof("values is %#v", values)
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ", ")
	return fmt.Sprintf("SELECT %v FROM %s", fields, tableName), []interface{}{}
}

func _limit(values ...interface{}) (string, []interface{}) {
	// LIMIT $num
	log.Infof("values is %#v", values)
	return "LIMIT ?", values
}

func _where(values ...interface{}) (string, []interface{}) {
	// WHERE $desc
	log.Infof("values is %#v", values)
	desc, vars := values[0], values[1:]
	return fmt.Sprintf("WHERE %s", desc), vars
}

func _orderBy(values ...interface{}) (string, []interface{}) {
	log.Infof("values is %#v", values)
	return fmt.Sprintf("ORDER BY %s", values[0]), []interface{}{}
}

// 参数 tableName map需要更新的键值对
func _update(values ...interface{}) (string, []interface{}) {
	// UPDATE table_name SET column1=value1,column2=value2,...
	log.Infof("values is %#v", values)
	tableName := values[0]
	m := values[1].(map[string]interface{})
	fields := make([]string, 0, len(m))
	vars := make([]interface{}, 0, len(m))
	sql := strings.Builder{}
	sql.WriteString(fmt.Sprintf("UPDATE %s SET ", tableName))
	for k, v := range m {
		fields = append(fields, fmt.Sprintf("%s = ?", k))
		vars = append(vars, v)
	}
	sql1 := strings.Join(fields, ", ")
	log.Infof("sql1 %s", sql1)
	sql.WriteString(sql1)
	return sql.String(), vars
}

func _delete(values ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("DELETE FROM %s", values[0]), []interface{}{}
}

func _count(values ...interface{}) (string, []interface{}) {
	return _select(values[0], []string{"count(*)"})
}
