package common

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/AsakuraMori/bangumi_go_api/pkg/errno"
)

func Json(i interface{}) string {
	b, err := json.Marshal(i)
	if err != nil {
		//logs.Errorf("json.Marshal error: %v", err)
		return ""
	} else {
		return string(b)
	}
}

func JsonFormat(i interface{}) string {
	b, err := json.MarshalIndent(i, "", "    ")
	if err != nil {
		//logs.Errorf("json.Marshal error: %v", err)
		return ""
	} else {
		return string(b)
	}
}

//ToMap 把object转化为 map 形式， 便于在生成json时添加一些额外的字段
func ToMap(o interface{}) (map[string]interface{}, error) {
	jsonByte, err := json.Marshal(o)
	if err != nil {
		err = errno.Errorf(errno.ErrInternalServer, "json.Marshal error")
		//logs.Errorf("%v", err)
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(jsonByte, &m); err != nil {
		err = errno.Errorf(errno.ErrInternalServer, "json.Unmarshal error, jsonStr: %s", string(jsonByte))
		//logs.Errorf("%v", err)
		return nil, err
	}
	return m, nil
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func StrPtrEmptyNil(s string) *string {
	if s == "" {
		return nil
	}
	return StrPtr(s)
}

func StrPtr(s string) *string {
	res := new(string)
	*res = s
	return res
}

func Int64PtrEmptyNil(s int64) *int64 {
	if s == 0 {
		return nil
	}
	return Int64Ptr(s)
}

func Int64Ptr(s int64) *int64 {
	res := new(int64)
	*res = s
	return res
}

func IfElse(condition bool, o1 string, o2 string) string {
	if condition {
		return o1
	}
	return o2
}
