package 工具

import "time"

var Cache map[string]struct {
	t     time.Time
	d     int64 // 时间间隔，秒
	value interface{}
}

func init() {
	Cache = make(map[string]struct {
		t     time.Time
		d     int64
		value interface{}
	})
}

/*
	ep: 工具.SetCache("DepotsResult", r, 60*10)
*/
func SetCache(key string, value interface{}, d int64) {
	Cache[key] = struct {
		t     time.Time
		d     int64
		value interface{}
	}{t: time.Now(), d: d, value: value}
}

/*
	ep:
	if v := 工具.GetCache("DepotsResult"); v != nil {
		if r, ok := v.(DepotsResult); ok {
			return r
		}
	}
*/
func GetCache(key string) interface{} {
	if value, ok := Cache[key]; ok {
		if time.Now().Unix()-value.t.Unix() > value.d {
			return nil
		}
		return value.value
	} else {
		return nil
	}
}
