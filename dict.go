package dict

import (
	"fmt"
	"strings"

	"github.com/valyala/fastjson"
)

type DictMap map[string][]string

func (dm DictMap) Query(term string) (res string) {
	term = strings.ToLower(term)
	if entry, ok := dm[term]; ok {
		res = entry[0]
		for _, v := range entry[1:] {
			res += "\n" + v
		}
	} else {
		res = fmt.Sprintf("Not Found: %v", term)
	}
	return
}

type FmtFunc func(word string, pos string, def string) string

func ParseDict(b []byte, fn FmtFunc) (dm DictMap, err error) {
	if fn == nil {
		fn = func(word string, pos string, def string) string {
			return fmt.Sprintf("%v (%v) — %v", word, pos, def)
		}
	}

	dm = make(DictMap)
	js := fastjson.MustParseBytes(b)
	arr, err := js.Array()
	if err != nil {
		return nil, err
	}

	for _, v := range arr {
		word := string(v.GetStringBytes("word"))
		pos := string(v.GetStringBytes("pos"))
		meaning := string(v.GetStringBytes("meaning"))
		res := fn(word, pos, meaning)
		key := strings.ToLower(string(word))
		entry := dm[key]
		dm[key] = append(entry, res)
	}
	return
}

func QueryRaw(b []byte, term string, fn FmtFunc) (res string) {
	if fn == nil {
		fn = func(word string, pos string, def string) string {
			return fmt.Sprintf("%v (%v) — %v", word, pos, def)
		}
	}

	js := fastjson.MustParseBytes(b)
	arr, err := js.Array()
	if err != nil {
		return
	}

	for _, v := range arr {
		word := string(v.GetStringBytes("word"))
		if strings.ToLower(string(word)) == term {
			res += fn(word, string(v.GetStringBytes("pos")), string(v.GetStringBytes("meaning"))) + "\n"
		}
	}
	return
}
