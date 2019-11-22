package apis

import (
	"net/url"
)

type TaobaoApiInterface interface {
	GetMethod() string
	GetValues() url.Values
}

type TaobaoApi struct {
	method string
	values url.Values
}

func (tr *TaobaoApi) SetValue(key, value string){
	if tr.values == nil {
		tr.values = url.Values{}
	}
	tr.values.Set(key,value)
}

func (tr *TaobaoApi) GetValue(key string) string{
	return tr.values.Get(key)
}

func (tr *TaobaoApi) SetValues(values url.Values) {
	tr.values = values
}

func (tr *TaobaoApi) GetValues() url.Values{
	return tr.values
}
//
//func (tr *TaobaoApi) SetMethod(method string){
//	tr.method = method
//}

func (tr *TaobaoApi) GetMethod(method string) string{
	return tr.method
}

//类型转换器
type TypeConverter interface {
	StructToJson() string
}
