package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//要求to必须已经分配好空间
func Uint32ToBytes(from uint32) (to []byte) {
	to = make([]byte, 4)
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, from)
	copy(to, buffer.Bytes()[0:4])
	return
}

func BytesToUint32(from []byte) (to uint32) {
	buffer := bytes.NewBuffer(from)
	binary.Read(buffer, binary.BigEndian, &to)
	return
}
func StringToInt(v string) (d int, err error) {
	tmp, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return
	}
	return int(tmp), err
}
func StringToUint(v string) (d uint, err error) {
	tmp, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		return
	}
	return uint(tmp), err
}
func StringToUint32(v string) (d uint32, err error) {
	tmp, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		return
	}
	return uint32(tmp), err
}
func StringToUint8(v string) (d uint8, err error) {
	tmp, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		return
	}
	return uint8(tmp), err
}
func StringToUint64(v string) (d uint64, err error) {
	d, err = strconv.ParseUint(v, 10, 64)
	return
}

func Uint64ToString(from uint64) (to string) {
	to = strconv.FormatUint(from, 10)
	return
}

func Float64ToString(from float64) (to string) {
	to = strconv.FormatFloat(from, 'f', -1, 64)
	return
}
func StringToFloat(v string) (d float32, err error) {
	tmp, err := strconv.ParseFloat(v, 10)
	d = float32(tmp)
	return
}
func StringToFloat64(v string) (d float64, err error) {
	d, err = strconv.ParseFloat(v, 10)
	return
}

func StringToInt64(v string) (d int64, err error) {
	d, err = strconv.ParseInt(v, 10, 64)
	return
}

func Int64ToString(from int64) (to string) {
	to = strconv.FormatInt(from, 10)
	return
}
func IntToString(from int) (to string) {
	to = strconv.FormatInt(int64(from), 10)
	return
}
func Uint32ToString(from uint32) (to string) {
	to = strconv.FormatInt(int64(from), 10)
	return
}

//结构->MAP
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		return nil
	}
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := f.Tag.Get("json")
		if name == "" {
			name = f.Name
		}
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func Struct2MapJson(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		return nil
	}
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := f.Tag.Get("json")
		if name == "" {
			name = f.Name
		}
		data[t.Field(i).Tag.Get("json")] = v.Field(i).Interface()
	}
	return data
}

func Bool2Int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func JoinUnderline(keys ...interface{}) string {
	if len(keys) == 0 {
		return ""
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%v", keys[0]))
	for i := 1; i < len(keys); i++ {
		buf.WriteString(fmt.Sprintf("_%v", keys[i]))
	}
	return buf.String()
}

func JoinLine(keys ...interface{}) string {
	if len(keys) == 0 {
		return ""
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%v", keys[0]))
	for i := 1; i < len(keys); i++ {
		buf.WriteString(fmt.Sprintf("|%v", keys[i]))
	}
	return buf.String()
}

func SplitUnderline(key string) []string {
	return strings.Split(key, "_")
}

func SplitLine(key string) []string {
	return strings.Split(key, "|")
}

func XmlMarshal(v interface{}, safeEncoding bool) ([]byte, error) {
	b, err := xml.Marshal(v)
	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("&amp;"), []byte("&"), -1)
	}
	return b, err
}
