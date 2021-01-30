package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

var serializationKeyWord = "ini"

//ConfigFileSerialization 序列化文件配置
func ConfigFileSerialization(filepath string, result interface{}) (err error) {
	fileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	err = ConfigSerialization(fileData, result)
	return
}

/**
 * ConfigSerialization 序列化
 * @author yaoyt98
 * @description //TODO
 * @date 7:48 下午 2021/1/30
 * @param
 * @return
 **/
func ConfigSerialization(cofigData []byte, result interface{}) (err error) {
	//判断传出参数
	resultType := reflect.TypeOf(result)
	if resultType.Kind() != reflect.Ptr {
		err = errors.New("result参数需要指针类型，请传入指针类型")
		return
	}

	resultStruct := resultType.Elem()
	if resultStruct.Kind() != reflect.Struct {
		err = errors.New("result参数需要结构体类型，请传入结构体类型")
		return
	}

	//分析配置文件内容
	linArr := strings.Split(string(cofigData), "\n")
	var linLen int                    //每一行的行数
	var linFirst byte                 //每一行第一个文字
	var sectionName string            //配置文件节点
	var lastSectionName string        //上一个配置文件节点
	var sectionItemJudge string = "=" //节点内容判断
	var sectionItemJudgeIndex int

	for _, linItem := range linArr {
		linItem = strings.TrimSpace(linItem)
		linLen = len(linItem)
		if linLen == 0 {
			continue
		}

		linFirst = linItem[0]
		if linFirst == '#' || linFirst == ';' {
			continue
		}

		//分析节点
		if linFirst == '[' {
			if linLen <= 2 {
				err = fmt.Errorf("解析节点时错误：%s", linItem)
				return
			}

			if linItem[linLen-1] != ']' {
				err = fmt.Errorf("解析节点时错误：%s", linItem)
				return
			}

			sectionName = strings.TrimSpace(linItem[1 : linLen-1])

			if len(sectionName) == 0 {
				err = fmt.Errorf("解析到空节点：%s", linItem)
				return
			}

			//解析到结构体
			var field reflect.StructField
			for i := 0; i < resultStruct.NumField(); i++ {
				field = resultStruct.Field(i)
				if field.Tag.Get(serializationKeyWord) == sectionName {
					lastSectionName = field.Name
					break
				}
			}
			continue
		}

		//分析内容类型是否正确
		sectionValueOf := reflect.ValueOf(result).Elem().FieldByName(lastSectionName)
		sectionTypeOf := sectionValueOf.Type()
		if sectionTypeOf.Kind() != reflect.Struct {
			err = errors.New("result参数需要结构体类型，请传入结构体类型")
			return
		}

		//分析节点下面的内容
		sectionItemJudgeIndex = strings.Index(linItem, sectionItemJudge)
		if sectionItemJudgeIndex == -1 {
			err = fmt.Errorf("解析节点内容时错误：%s", linItem)
			return
		}

		if sectionValueOf.NumField() == 0 {
			continue
		}

		key := strings.TrimSpace(linItem[0:sectionItemJudgeIndex])
		value := strings.TrimSpace(linItem[sectionItemJudgeIndex+1:])

		if len(key) == 0 {
			continue
		}

		//匹配key
		var keyFieldName string
		for i := 0; i < sectionTypeOf.NumField(); i++ {
			field := sectionTypeOf.Field(i)
			if field.Tag.Get(serializationKeyWord) == key {
				keyFieldName = field.Name
				break
			}
		}

		fieldValue := sectionValueOf.FieldByName(keyFieldName)
		if fieldValue == reflect.ValueOf(nil) {
			return
		}

		switch fieldKind := fieldValue.Type().Kind(); fieldKind {
		case reflect.String:
			fieldValue.SetString(value)
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
			intval, errInt := strconv.ParseInt(value, 10, 64)
			if errInt != nil {
				err = errInt
				return
			}
			fieldValue.SetInt(intval)
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
			intval, errInt := strconv.ParseUint(value, 10, 64)
			if errInt != nil {
				err = errInt
				return
			}
			fieldValue.SetUint(intval)
		case reflect.Float32, reflect.Float64:
			floatVal, errFloat := strconv.ParseFloat(value, 64)
			if errFloat != nil {
				err = errFloat
				return
			}
			fieldValue.SetFloat(floatVal)
		default:
			err = fmt.Errorf("未匹配到类型%v", fieldKind)
		}
	}

	return
}

//UnConfigFileSerialization 反序列化文件配置
func UnConfigFileSerialization() {

}

//UnConfigSerialization 反序列化
func UnConfigSerialization() {

}
