package parse

import (
	"fmt"
	"net/url"
	"os"
	"reflect"
	"strconv"
)

type URLParam struct {
	Name  string `noodles:"name"`
	Score int32  `noodles:"score"`
	//... ...
}

func ParamsUnpack(ptr interface{}, querystr string) error {
	values, err := url.ParseQuery(querystr)
	if err != nil {
		return err
	}

	printError := func(tag, param string) {
		fmt.Fprintln(os.Stderr, "Parse param error: ", tag, param)
	}

	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		esTag := fieldInfo.Tag.Get("noodles")
		if esTag == "" {
			continue
		}

		vi := v.Field(i)

		param := values.Get(esTag)
		switch vi.Kind() {
		case reflect.Float64:
			if vv, e := strconv.ParseFloat(param, 64); e != nil {
				printError(esTag, param)
			} else {
				vi.SetFloat(vv)
			}
		case reflect.Float32:
			if vv, e := strconv.ParseFloat(param, 32); e != nil {
				printError(esTag, param)
			} else {
				vi.SetFloat(vv)
			}
		case reflect.String:
			vi.SetString(param)
		case reflect.Int, reflect.Int64: /* warning! */
			if vv, e := strconv.ParseInt(param, 10, 64); e != nil {
				printError(esTag, param)
			} else {
				vi.SetInt(vv)
			}
		case reflect.Int32:
			if vv, e := strconv.ParseInt(param, 10, 32); e != nil {
				printError(esTag, param)
			} else {
				vi.SetInt(vv)
			}
		case reflect.Bool:
			if vv, e := strconv.ParseBool(param); e != nil {
				printError(esTag, param)
			} else {
				vi.SetBool(vv)
			}
		}
	}

	return nil
}

func ParamsUnpack2(p *URLParam, querystr string) error {
	values, err := url.ParseQuery(querystr)
	if err != nil {
		return err
	}

	name := values.Get("name")
	p.Name = name

	score := values.Get("score")
	if vv, e := strconv.ParseInt(score, 10, 32); e != nil {
		return e
	} else {
		p.Score = int32(vv)
	}

	return nil
}
