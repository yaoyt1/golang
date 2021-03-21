package util

import "fmt"

//getSqlMapToString主要用于models 拼接map参数语句
func GetSqlMapToString(k string, v interface{}) (sqlStr string) {
	switch value := v.(type) {
	case string:
		sqlStr = fmt.Sprintf(" %s='%s' ", k, value)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		sqlStr = fmt.Sprintf(" %s=%d ", k, value)
	case float32, float64:
		sqlStr = fmt.Sprintf(" %s=%f ", k, value)
	default:
		sqlStr = fmt.Sprintf(" %s=%v ", k, value)
	}
	return
}
