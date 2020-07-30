/*
* @Title:   缓存工具
* @Author:  minibear2333
* @Date:    2020/7/28 9:37 下午
* @url:     https://github.com/minibear2333/how_to_code
 */
package utils

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
	ep: utils.SetCache("DepotsResult", r, 60*10)
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
	if v := utils.GetCache("DepotsResult"); v != nil {
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
