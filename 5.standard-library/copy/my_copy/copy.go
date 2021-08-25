package my_copy

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

type Foo struct {
	List   []int             `json:"list"`
	FooMap map[string]string `json:"foo_map"`
	IntPtr *int              `json:"int_ptr"`
}

func (f *Foo) Duplicate() Foo {
	var tmp = Foo{
		List:   make([]int, 0, len(f.List)),
		FooMap: make(map[string]string),
		IntPtr: new(int),
	}

	copy(tmp.List, f.List)
	for i := range f.FooMap {
		tmp.FooMap[i] = f.FooMap[i]
	}
	if f.IntPtr != nil {
		*tmp.IntPtr = *f.IntPtr
	} else {
		tmp.IntPtr = nil
	}
	return tmp
}

func DeepCopyByJson(dst, src interface{}) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, dst)

	return err
}
func DeepCopyByGob(dst, src interface{}) error {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}

	return gob.NewDecoder(&buffer).Decode(dst)
}

//
//func main() {
//	fmt.Println("手动拷贝方法")
//	var a = 1
//	var t1 = Foo{IntPtr: &a}
//	t2 := t1.Duplicate()
//	a = 2
//	fmt.Println(*t1.IntPtr)
//	fmt.Println(*t2.IntPtr)
//	fmt.Println("json序列化反序列化方法")
//	a = 3
//	t1 = Foo{IntPtr: &a}
//	t2 = Foo{}
//	_ = DeepCopyByJson(&t2, t1)
//	fmt.Println(*t1.IntPtr)
//	fmt.Println(*t2.IntPtr)
//	fmt.Println("gob序列化反序列化方法")
//	a = 4
//	t1 = Foo{IntPtr: &a}
//	t2 = Foo{}
//	_ = DeepCopyByGob(&t2, t1)
//	fmt.Println(*t1.IntPtr)
//	fmt.Println(*t2.IntPtr)
//}
