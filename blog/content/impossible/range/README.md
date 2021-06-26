# for range的一个坑

for range是值拷贝出来的副本

在使用for range的时候，要注意的是，不管是slice还是map，循环的值都是被range值拷贝出来的副本值。
举个简单的例了

对于list

```go
    var t []Test
	t = append(t, Test{Index: 1, Num: 1})
	t = append(t, Test{Index: 2, Num: 2})

	// 实际上没有成功修改t.Num，因为是副本复制
	for _, v := range t {
		v.Num += 100
	}

	for _, v := range t {
		// 输出
		// 1 1
		// 2 2
		fmt.Println(v.Index, v.Num)
	}
```

对于 map, 也不能这么搞，实际上都是复制
```go

	m := make(map[int]Test)
	m[0] = Test{Index: 1, Num: 1}
	m[1] = Test{Index: 2, Num: 2}
	for _, v := range m {
		v.Num += 100
	}
	for _, v := range m {
		// 输出(可以乱序)
		// 1 1
		// 2 2
		fmt.Println(v.Index, v.Num)
	}
```
### 怎么做？

两个办法，用下标（map也一样）
```go
	for i := range t {
		t[i].Num += 100
		fmt.Println(t[i].Num)
		// 输出(可以乱序)
		// 101 102
	}
```
	
用指针

```go
	var t2 []*Test
	t2 = append(t2, &Test{Index: 1, Num: 1})
	t2 = append(t2, &Test{Index: 2, Num: 2})

	for k, v := range t2 {
		v.Num += 100
		fmt.Println(t2[k].Num)
		// 输出(可以乱序)
		// 101 102
	}
```

### for range 原理
通过查看 [源代码](https://github.com/golang/gofrontend) ，我们可以发现for range的实现是：

```go
# statements.cc:6419 (441f3f1 on 4 Oct)
// Arrange to do a loop appropriate for the type.  We will produce
//   for INIT ; COND ; POST {
//           ITER_INIT
//           INDEX = INDEX_TEMP
//           VALUE = VALUE_TEMP // If there is a value
//           original statements
//   }
```

并且对于Slice,Map等各有具体不同的编译实现,我们先看看for range slice的具体实现
```go
# statements.cc:6638  (441f3f1 on 4 Oct)
// The loop we generate:
//   for_temp := range
//   len_temp := len(for_temp)
//   for index_temp = 0; index_temp < len_temp; index_temp++ {
//           value_temp = for_temp[index_temp]
//           index = index_temp
//           value = value_temp
//           original body
//   }
```

先是对要遍历的 Slice 做一个拷贝，获取长度大小，然后使用常规for循环进行遍历，并且返回值的拷贝。
再看看for range map的具体实现：

```go
# statements.cc:6891  (441f3f1 on 4 Oct)
// The loop we generate:
//   var hiter map_iteration_struct
//   for mapiterinit(type, range, &hiter); hiter.key != nil; mapiternext(&hiter) {
//           index_temp = *hiter.key
//           value_temp = *hiter.val
//           index = index_temp
//           value = value_temp
//           original body
//   }
```

也是先对map进行了初始化，因为map是hashmap，所以这里其实是一个hashmap指针的拷贝。

引用：[Go 中for range的一个坑](https://studygolang.com/articles/25094?fr=sidebar)
