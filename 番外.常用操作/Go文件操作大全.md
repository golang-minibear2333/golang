Go 官方库的文件操作分散在多个包中，比如`os`、`ioutil`包，我本来想写一篇总结性的 Go 文件操作的文章，却发现已经有人 2015 年已经写了一篇这样的文章，写的非常好，所以我翻译成了中文，强烈推荐你阅读一下。



原文: [Working with Files in Go](http://www.devdungeon.com/content/working-files-go#write_bytes), 作者: [NanoDano](http://www.devdungeon.com/blogs/nanodano)

## 介绍

### 万物皆文件

UNIX 的一个基础设计就是"万物皆文件"(everything is a file)。我们不必知道一个文件到底映射成什么，操作系统的设备驱动抽象成文件。操作系统为设备提供了文件格式的接口。

Go 语言中的 reader 和 writer 接口也类似。我们只需简单的读写字节，不必知道 reader 的数据来自哪里，也不必知道 writer 将数据发送到哪里。  
你可以在`/dev`下查看可用的设备，有些可能需要较高的权限才能访问。

## 基本操作

### 创建空文件

```go
package main

import (

    "log"

    "os"

)

var (

    newFile *os.File

    err     error

)

func main() {

    newFile, err = os.Create("test.txt")

    if err != nil {

        log.Fatal(err)

    }

    log.Println(newFile)

    newFile.Close()

}
```

### Truncate 文件

```go
package main

import (

    "log"

    "os"

)

func main() {

    // 裁剪一个文件到100个字节。

    // 如果文件本来就少于100个字节，则文件中原始内容得以保留，剩余的字节以null字节填充。

    // 如果文件本来超过100个字节，则超过的字节会被抛弃。

    // 这样我们总是得到精确的100个字节的文件。

    // 传入0则会清空文件。

    err := os.Truncate("test.txt", 100)

    if err != nil {

        log.Fatal(err)

    }

}
```

### 得到文件信息

```go
package main

import (

    "fmt"

    "log"

    "os"

)

var (

    fileInfo os.FileInfo

    err      error

)

func main() {

    // 如果文件不存在，则返回错误

    fileInfo, err = os.Stat("test.txt")

    if err != nil {

        log.Fatal(err)

    }

    fmt.Println("File name:", fileInfo.Name())

    fmt.Println("Size in bytes:", fileInfo.Size())

    fmt.Println("Permissions:", fileInfo.Mode())

    fmt.Println("Last modified:", fileInfo.ModTime())

    fmt.Println("Is Directory: ", fileInfo.IsDir())

    fmt.Printf("System interface type: %T\\n", fileInfo.Sys())

    fmt.Printf("System info: %+v\\n\\n", fileInfo.Sys())

}
```

### 重命名和移动

```go
package main

import (

    "log"

    "os"

)

func main() {

    originalPath := "test.txt"

    newPath := "test2.txt"

    err := os.Rename(originalPath, newPath)

    if err != nil {

        log.Fatal(err)

    }

}
```

> 译者按： rename 和 move 原理一样

### 删除文件

```go
package main

import (

    "log"

    "os"

)

func main() {

    err := os.Remove("test.txt")

    if err != nil {

        log.Fatal(err)

    }

}
```

### 打开和关闭文件

```go
package main

import (

    "log"

    "os"

)

func main() {

    // 简单地以只读的方式打开。下面的例子会介绍读写的例子。

    file, err := os.Open("test.txt")

    if err != nil {

        log.Fatal(err)

    }

    file.Close()

    // OpenFile提供更多的选项。

    // 最后一个参数是权限模式permission mode

    // 第二个是打开时的属性

    file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)

    if err != nil {

        log.Fatal(err)

    }

    file.Close()

    // 下面的属性可以单独使用，也可以组合使用。

    // 组合使用时可以使用 OR 操作设置 OpenFile的第二个参数，例如：

    // os.O\_CREATE|os.O\_APPEND

    // 或者 os.O\_CREATE|os.O\_TRUNC|os.O_WRONLY

    // os.O_RDONLY // 只读

    // os.O_WRONLY // 只写

    // os.O_RDWR // 读写

    // os.O_APPEND // 往文件中添建（Append）

    // os.O_CREATE // 如果文件不存在则先创建

    // os.O_TRUNC // 文件打开时裁剪文件

    // os.O\_EXCL // 和O\_CREATE一起使用，文件不能存在

    // os.O_SYNC // 以同步I/O的方式打开

}
```

> 译者按：熟悉 Linux 的读者应该很熟悉权限模式，通过 Linux 命令`chmod`可以更改文件的权限  
> [https://www.linux.com/learn/understanding-linux-file-permissions](https://www.linux.com/learn/understanding-linux-file-permissions)
>
> 补充了原文未介绍的 flag

### 检查文件是否存在

```go
package main

import (

    "log"

    "os"

)

var (

    fileInfo *os.FileInfo

    err      error

)

func main() {

    // 文件不存在则返回error

    fileInfo, err := os.Stat("test.txt")

    if err != nil {

        if os.IsNotExist(err) {

            log.Fatal("File does not exist.")

        }

    }

    log.Println("File does exist. File information:")

    log.Println(fileInfo)

}
```

### 检查读写权限

```go
package main

import (

    "log"

    "os"

)

func main() {

    // 这个例子测试写权限，如果没有写权限则返回error。

    // 注意文件不存在也会返回error，需要检查error的信息来获取到底是哪个错误导致。

    file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)

    if err != nil {

        if os.IsPermission(err) {

            log.Println("Error: Write permission denied.")

        }

    }

    file.Close()

    // 测试读权限

    file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)

    if err != nil {

        if os.IsPermission(err) {

            log.Println("Error: Read permission denied.")

        }

    }

    file.Close()

}
```

### 改变权限、拥有者、时间戳

```go
package main

import (

    "log"

    "os"

    "time"

)

func main() {

    // 使用Linux风格改变文件权限

    err := os.Chmod("test.txt", 0777)

    if err != nil {

        log.Println(err)

    }

    // 改变文件所有者

    err = os.Chown("test.txt", os.Getuid(), os.Getgid())

    if err != nil {

        log.Println(err)

    }

    // 改变时间戳

    twoDaysFromNow := time.Now().Add(48 \* time.Hour)

    lastAccessTime := twoDaysFromNow

    lastModifyTime := twoDaysFromNow

    err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)

    if err != nil {

        log.Println(err)

    }

}
```

### 硬链接和软链接

一个普通的文件是一个指向硬盘的 inode 的地方。  
硬链接创建一个新的指针指向同一个地方。只有所有的链接被删除后文件才会被删除。硬链接只在相同的文件系统中才工作。你可以认为一个硬链接是一个正常的链接。

symbolic link，又叫软连接，和硬链接有点不一样，它不直接指向硬盘中的相同的地方，而是通过名字引用其它文件。他们可以指向不同的文件系统中的不同文件。并不是所有的操作系统都支持软链接。

```go
package main

import (

    "os"

    "log"

    "fmt"

)

func main() {

    // 创建一个硬链接。

    // 创建后同一个文件内容会有两个文件名，改变一个文件的内容会影响另一个。

    // 删除和重命名不会影响另一个。

    err := os.Link("original.txt", "original_also.txt")

    if err != nil {

        log.Fatal(err)

    }

    fmt.Println("creating sym")

    // Create a symlink

    err = os.Symlink("original.txt", "original_sym.txt")

    if err != nil {

        log.Fatal(err)

    }

    // Lstat返回一个文件的信息，但是当文件是一个软链接时，它返回软链接的信息，而不是引用的文件的信息。

    // Symlink在Windows中不工作。

    fileInfo, err := os.Lstat("original_sym.txt")

    if err != nil {

        log.Fatal(err)

    }

    fmt.Printf("Link info: %+v", fileInfo)

    //改变软链接的拥有者不会影响原始文件。

    err = os.Lchown("original_sym.txt", os.Getuid(), os.Getgid())

    if err != nil {

        log.Fatal(err)

    }

}
```

## 读写

### 复制文件

```go
package main

import (

    "os"

    "log"

    "io"

)

func main() {

    // 打开原始文件

    originalFile, err := os.Open("test.txt")

    if err != nil {

        log.Fatal(err)

    }

    defer originalFile.Close()

    // 创建新的文件作为目标文件

    newFile, err := os.Create("test_copy.txt")

    if err != nil {

        log.Fatal(err)

    }

    defer newFile.Close()

    // 从源中复制字节到目标文件

    bytesWritten, err := io.Copy(newFile, originalFile)

    if err != nil {

        log.Fatal(err)

    }

    log.Printf("Copied %d bytes.", bytesWritten)

    // 将文件内容flush到硬盘中

    err = newFile.Sync()

    if err != nil {

        log.Fatal(err)

    }

}
```

### 跳转到文件指定位置(Seek)

```go
package main

import (

    "os"

    "fmt"

    "log"

)

func main() {

    file, _ := os.Open("test.txt")

    defer file.Close()

    // 偏离位置，可以是正数也可以是负数

    var offset int64 = 5

    // 用来计算offset的初始位置

    // 0 = 文件开始位置

    // 1 = 当前位置

    // 2 = 文件结尾处

    var whence int = 0

    newPosition, err := file.Seek(offset, whence)

    if err != nil {

        log.Fatal(err)

    }

    fmt.Println("Just moved to 5:", newPosition)

    // 从当前位置回退两个字节

    newPosition, err = file.Seek(-2, 1)

    if err != nil {

        log.Fatal(err)

    }

    fmt.Println("Just moved back two:", newPosition)

    // 使用下面的技巧得到当前的位置

    currentPosition, err := file.Seek(0, 1)

    fmt.Println("Current position:", currentPosition)

    // 转到文件开始处

    newPosition, err = file.Seek(0, 0)

    if err != nil {

        log.Fatal(err)

    }

    fmt.Println("Position after seeking 0,0:", newPosition)

}
```

#### 写文件

可以使用`os`包写入一个打开的文件。  
因为 Go 可执行包是静态链接的可执行文件，你 import 的每一个包都会增加你的可执行文件的大小。其它的包如`io`、｀ ioutil ｀、｀ bufio ｀提供了一些方法，但是它们不是必须的。

```go
package main

import (

    "os"

    "log"

)

func main() {

    // 可写方式打开文件

    file, err := os.OpenFile(

        "test.txt",

        os.O\_WRONLY|os.O\_TRUNC|os.O_CREATE,

        0666,

    )

    if err != nil {

        log.Fatal(err)

    }

    defer file.Close()

    // 写字节到文件中

    byteSlice := \[\]byte("Bytes!\\n")

    bytesWritten, err := file.Write(byteSlice)

    if err != nil {

        log.Fatal(err)

    }

    log.Printf("Wrote %d bytes.\\n", bytesWritten)

}
```

#### 快写文件

`ioutil`包有一个非常有用的方法`WriteFile()`可以处理创建／打开文件、写字节 slice 和关闭文件一系列的操作。如果你需要简洁快速地写字节 slice 到文件中，你可以使用它。

```go
package main

import (

    "io/ioutil"

    "log"

)

func main() {

    err := ioutil.WriteFile("test.txt", \[\]byte("Hi\\n"), 0666)

    if err != nil {

        log.Fatal(err)

    }

}
```

#### 使用缓存写

`bufio`包提供了带缓存功能的 writer，所以你可以在写字节到硬盘前使用内存缓存。当你处理很多的数据很有用，因为它可以节省操作硬盘 I/O 的时间。在其它一些情况下它也很有用，比如你每次写一个字节，把它们攒在内存缓存中，然后一次写入到硬盘中，减少硬盘的磨损以及提升性能。

```go
package main

import (

    "log"

    "os"

    "bufio"

)

func main() {

    // 打开文件，只写

    file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)

    if err != nil {

        log.Fatal(err)

    }

    defer file.Close()

    // 为这个文件创建buffered writer

    bufferedWriter := bufio.NewWriter(file)

    // 写字节到buffer

    bytesWritten, err := bufferedWriter.Write(

        \[\]byte{65, 66, 67},

    )

    if err != nil {

        log.Fatal(err)

    }

    log.Printf("Bytes written: %d\\n", bytesWritten)

    // 写字符串到buffer

    // 也可以使用 WriteRune() 和 WriteByte()

    bytesWritten, err = bufferedWriter.WriteString(

        "Buffered string\\n",

    )

    if err != nil {

        log.Fatal(err)

    }

    log.Printf("Bytes written: %d\\n", bytesWritten)

    // 检查缓存中的字节数

    unflushedBufferSize := bufferedWriter.Buffered()

    log.Printf("Bytes buffered: %d\\n", unflushedBufferSize)

    // 还有多少字节可用（未使用的缓存大小）

    bytesAvailable := bufferedWriter.Available()

    if err != nil {

        log.Fatal(err)

    }

    log.Printf("Available buffer: %d\\n", bytesAvailable)

    // 写内存buffer到硬盘

    bufferedWriter.Flush()

    // 丢弃还没有flush的缓存的内容，清除错误并把它的输出传给参数中的writer

    // 当你想将缓存传给另外一个writer时有用

    bufferedWriter.Reset(bufferedWriter)

    bytesAvailable = bufferedWriter.Available()

    if err != nil {

        log.Fatal(err)

    }

    log.Printf("Available buffer: %d\\n", bytesAvailable)

    // 重新设置缓存的大小。

    // 第一个参数是缓存应该输出到哪里，这个例子中我们使用相同的writer。

    // 如果我们设置的新的大小小于第一个参数writer的缓存大小， 比如10，我们不会得到一个10字节大小的缓存，

    // 而是writer的原始大小的缓存，默认是4096。

    // 它的功能主要还是为了扩容。

    bufferedWriter = bufio.NewWriterSize(

        bufferedWriter,

        8000,

    )

    // resize后检查缓存的大小

    bytesAvailable = bufferedWriter.Available()

    if err != nil {

        log.Fatal(err)

    }

    log.Printf("Available buffer: %d\\n", bytesAvailable)

}
```

#### 读取最多 N 个字节

`os.File`提供了文件操作的基本功能， 而`io`、`ioutil`、`bufio`提供了额外的辅助函数。

```go
package main

import (

    "os"

    "log"

)

func main() {

    // 打开文件，只读

    file, err := os.Open("test.txt")

    if err != nil {

        log.Fatal(err)

    }

    defer file.Close()

    // 从文件中读取len(b)字节的文件。

    // 返回0字节意味着读取到文件尾了

    // 读取到文件会返回io.EOF的error

    byteSlice := make(\[\]byte, 16)

    bytesRead, err := file.Read(byteSlice)

    if err != nil {

        log.Fatal(err)

    }

    log.Printf("Number of bytes read: %d\\n", bytesRead)

    log.Printf("Data read: %s\\n", byteSlice)

}
```

#### 读取正好 N 个字节

```go
package main

import (

    "os"

    "log"

    "io"

)

func main() {

    // Open file for reading

    file, err := os.Open("test.txt")

    if err != nil {

        log.Fatal(err)

    }

    // file.Read()可以读取一个小文件到大的byte slice中，

    // 但是io.ReadFull()在文件的字节数小于byte slice字节数的时候会返回错误

    byteSlice := make(\[\]byte, 2)

    numBytesRead, err := io.ReadFull(file, byteSlice)

    if err != nil {

        log.Fatal(err)

    }

    log.Printf("Number of bytes read: %d\\n", numBytesRead)

    log.Printf("Data read: %s\\n", byteSlice)

}
```

#### 读取至少 N 个字节

```go
package main

import (

    "os"

    "log"

    "io"

)

func main() {

    // 打开文件，只读

    file, err := os.Open("test.txt")

    if err != nil {

        log.Fatal(err)

    }

    byteSlice := make(\[\]byte, 512)

    minBytes := 8

    // io.ReadAtLeast()在不能得到最小的字节的时候会返回错误，但会把已读的文件保留

    numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)

    if err != nil {

        log.Fatal(err)

    }

    log.Printf("Number of bytes read: %d\\n", numBytesRead)

    log.Printf("Data read: %s\\n", byteSlice)

}
```

#### 读取全部字节

```go
package main

import (

    "os"

    "log"

    "fmt"

    "io/ioutil"

)

func main() {

    file, err := os.Open("test.txt")

    if err != nil {

        log.Fatal(err)

    }

    // os.File.Read(), io.ReadFull() 和

    // io.ReadAtLeast() 在读取之前都需要一个固定大小的byte slice。

    // 但ioutil.ReadAll()会读取reader(这个例子中是file)的每一个字节，然后把字节slice返回。

    data, err := ioutil.ReadAll(file)

    if err != nil {

        log.Fatal(err)

    }

    fmt.Printf("Data as hex: %x\\n", data)

    fmt.Printf("Data as string: %s\\n", data)

    fmt.Println("Number of bytes read:", len(data))

}
```

#### 快读到内存

```go

package main

import (

    "log"

    "io/ioutil"

)

func main() {

    // 读取文件到byte slice中

    data, err := ioutil.ReadFile("test.txt")

    if err != nil {

        log.Fatal(err)

    }

    log.Printf("Data read: %s\\n", data)

}
```

#### 使用缓存读

有缓存写也有缓存读。  
缓存 reader 会把一些内容缓存在内存中。它会提供比`os.File`和`io.Reader`更多的函数,缺省的缓存大小是 4096，最小缓存是 16。

```go
package main

import (

    "os"

    "log"

    "bufio"

    "fmt"

)

func main() {

    // 打开文件，创建buffered reader

    file, err := os.Open("test.txt")

    if err != nil {

        log.Fatal(err)

    }

    bufferedReader := bufio.NewReader(file)

    // 得到字节，当前指针不变

    byteSlice := make(\[\]byte, 5)

    byteSlice, err = bufferedReader.Peek(5)

    if err != nil {

        log.Fatal(err)

    }

    fmt.Printf("Peeked at 5 bytes: %s\\n", byteSlice)

    // 读取，指针同时移动

    numBytesRead, err := bufferedReader.Read(byteSlice)

    if err != nil {

        log.Fatal(err)

    }

    fmt.Printf("Read %d bytes: %s\\n", numBytesRead, byteSlice)

    // 读取一个字节, 如果读取不成功会返回Error

    myByte, err := bufferedReader.ReadByte()

    if err != nil {

        log.Fatal(err)

    }

    fmt.Printf("Read 1 byte: %c\\n", myByte)

    // 读取到分隔符，包含分隔符，返回byte slice

    dataBytes, err := bufferedReader.ReadBytes('\\n')

    if err != nil {

        log.Fatal(err)

    }

    fmt.Printf("Read bytes: %s\\n", dataBytes)

    // 读取到分隔符，包含分隔符，返回字符串

    dataString, err := bufferedReader.ReadString('\\n')

    if err != nil {

        log.Fatal(err)

    }

    fmt.Printf("Read string: %s\\n", dataString)

    //这个例子读取了很多行，所以test.txt应该包含多行文本才不至于出错

}
```

#### 使用 scanner

`Scanner`是`bufio`包下的类型,在处理文件中以分隔符分隔的文本时很有用。  
通常我们使用换行符作为分隔符将文件内容分成多行。在 CSV 文件中，逗号一般作为分隔符。  
`os.File`文件可以被包装成`bufio.Scanner`，它就像一个缓存 reader。  
我们会调用`Scan()`方法去读取下一个分隔符，使用`Text()`或者`Bytes()`获取读取的数据。

分隔符可以不是一个简单的字节或者字符，有一个特殊的方法可以实现分隔符的功能，以及将指针移动多少，返回什么数据。  
如果没有定制的`SplitFunc`提供，缺省的`ScanLines`会使用`newline`字符作为分隔符，其它的分隔函数还包括`ScanRunes`和`ScanWords`,皆在`bufio`包中。

```go
// To define your own split function, match this fingerprint

type SplitFunc func(data \[\]byte, atEOF bool) (advance int, token \[\]byte, err error)

// Returning (0, nil, nil) will tell the scanner

// to scan again, but with a bigger buffer because

// it wasn't enough data to reach the delimiter
```

下面的例子中，为一个文件创建了`bufio.Scanner`，并按照单词逐个读取：

```go
package main

import (

    "os"

    "log"

    "fmt"

    "bufio"

)

func main() {

    file, err := os.Open("test.txt")

    if err != nil {

        log.Fatal(err)

    }

    scanner := bufio.NewScanner(file)

    // 缺省的分隔函数是bufio.ScanLines,我们这里使用ScanWords。

    // 也可以定制一个SplitFunc类型的分隔函数

    scanner.Split(bufio.ScanWords)

    // scan下一个token.

    success := scanner.Scan()

    if success == false {

        // 出现错误或者EOF是返回Error

        err = scanner.Err()

        if err == nil {

            log.Println("Scan completed and reached EOF")

        } else {

            log.Fatal(err)

        }

    }

    // 得到数据，Bytes() 或者 Text()

    fmt.Println("First word found:", scanner.Text())

    // 再次调用scanner.Scan()发现下一个token

}
```

### 压缩

#### 打包(zip) 文件

```go
// This example uses zip but standard library

// also supports tar archives

package main

import (

    "archive/zip"

    "log"

    "os"

)

func main() {

    // 创建一个打包文件

    outFile, err := os.Create("test.zip")

    if err != nil {

        log.Fatal(err)

    }

    defer outFile.Close()

    // 创建zip writer

    zipWriter := zip.NewWriter(outFile)

    // 往打包文件中写文件。

    // 这里我们使用硬编码的内容，你可以遍历一个文件夹，把文件夹下的文件以及它们的内容写入到这个打包文件中。

    var filesToArchive = \[\]struct {

        Name, Body string

    } {

        {"test.txt", "String contents of file"},

        {"test2.txt", "\\x61\\x62\\x63\\n"},

    }

    // 下面将要打包的内容写入到打包文件中，依次写入。

    for _, file := range filesToArchive {

            fileWriter, err := zipWriter.Create(file.Name)

            if err != nil {

                    log.Fatal(err)

            }

            _, err = fileWriter.Write(\[\]byte(file.Body))

            if err != nil {

                    log.Fatal(err)

            }

    }

    // 清理

    err = zipWriter.Close()

    if err != nil {

            log.Fatal(err)

    }

}
```

#### 抽取(unzip) 文件

```go

// This example uses zip but standard library

// also supports tar archives

package main

import (

    "archive/zip"

    "log"

    "io"

    "os"

    "path/filepath"

)

func main() {

    zipReader, err := zip.OpenReader("test.zip")

    if err != nil {

        log.Fatal(err)

    }

    defer zipReader.Close()

    // 遍历打包文件中的每一文件/文件夹

    for _, file := range zipReader.Reader.File {

        // 打包文件中的文件就像普通的一个文件对象一样

        zippedFile, err := file.Open()

        if err != nil {

            log.Fatal(err)

        }

        defer zippedFile.Close()

        // 指定抽取的文件名。

        // 你可以指定全路径名或者一个前缀，这样可以把它们放在不同的文件夹中。

        // 我们这个例子使用打包文件中相同的文件名。

        targetDir := "./"

        extractedFilePath := filepath.Join(

            targetDir,

            file.Name,

        )

        // 抽取项目或者创建文件夹

        if file.FileInfo().IsDir() {

            // 创建文件夹并设置同样的权限

            log.Println("Creating directory:", extractedFilePath)

            os.MkdirAll(extractedFilePath, file.Mode())

        } else {

            //抽取正常的文件

            log.Println("Extracting file:", file.Name)

            outputFile, err := os.OpenFile(

                extractedFilePath,

                os.O\_WRONLY|os.O\_CREATE|os.O_TRUNC,

                file.Mode(),

            )

            if err != nil {

                log.Fatal(err)

            }

            defer outputFile.Close()

            // 通过io.Copy简洁地复制文件内容

            _, err = io.Copy(outputFile, zippedFile)

            if err != nil {

                log.Fatal(err)

            }

        }

    }

}
```

#### 压缩文件

```go
// 这个例子中使用gzip压缩格式，标准库还支持zlib, bz2, flate, lzw

package main

import (

    "os"

    "compress/gzip"

    "log"

)

func main() {

    outputFile, err := os.Create("test.txt.gz")

    if err != nil {

        log.Fatal(err)

    }

    gzipWriter := gzip.NewWriter(outputFile)

    defer gzipWriter.Close()

    // 当我们写如到gizp writer数据时，它会依次压缩数据并写入到底层的文件中。

    // 我们不必关心它是如何压缩的，还是像普通的writer一样操作即可。

    _, err = gzipWriter.Write(\[\]byte("Gophers rule!\\n"))

    if err != nil {

        log.Fatal(err)

    }

    log.Println("Compressed data written to file.")

}
```

#### 解压缩文件

```go
// 这个例子中使用gzip压缩格式，标准库还支持zlib, bz2, flate, lzw

package main

import (

    "compress/gzip"

    "log"

    "io"

    "os"

)

func main() {

    // 打开一个gzip文件。

    // 文件是一个reader,但是我们可以使用各种数据源，比如web服务器返回的gzipped内容，

    // 它的内容不是一个文件，而是一个内存流

    gzipFile, err := os.Open("test.txt.gz")

    if err != nil {

        log.Fatal(err)

    }

    gzipReader, err := gzip.NewReader(gzipFile)

    if err != nil {

        log.Fatal(err)

    }

    defer gzipReader.Close()

    // 解压缩到一个writer,它是一个file writer

    outfileWriter, err := os.Create("unzipped.txt")

    if err != nil {

        log.Fatal(err)

    }

    defer outfileWriter.Close()

    // 复制内容

    _, err = io.Copy(outfileWriter, gzipReader)

    if err != nil {

        log.Fatal(err)

    }

}
```

### 其它

#### 临时文件和目录

`ioutil`提供了两个函数: `TempDir()` 和 `TempFile()`。  
使用完毕后，调用者负责删除这些临时文件和文件夹。  
有一点好处就是当你传递一个空字符串作为文件夹名的时候，它会在操作系统的临时文件夹中创建这些项目（/tmp on Linux）。  
`os.TempDir()`返回当前操作系统的临时文件夹。

```go
package main

import (

     "os"

     "io/ioutil"

     "log"

     "fmt"

)

func main() {

     // 在系统临时文件夹中创建一个临时文件夹

     tempDirPath, err := ioutil.TempDir("", "myTempDir")

     if err != nil {

          log.Fatal(err)

     }

     fmt.Println("Temp dir created:", tempDirPath)

     // 在临时文件夹中创建临时文件

     tempFile, err := ioutil.TempFile(tempDirPath, "myTempFile.txt")

     if err != nil {

          log.Fatal(err)

     }

     fmt.Println("Temp file created:", tempFile.Name())

     // ... 做一些操作 ...

     // 关闭文件

     err = tempFile.Close()

     if err != nil {

        log.Fatal(err)

    }

    // 删除我们创建的资源

     err = os.Remove(tempFile.Name())

     if err != nil {

        log.Fatal(err)

    }

     err = os.Remove(tempDirPath)

     if err != nil {

        log.Fatal(err)

    }

}
```

#### 通过 HTTP 下载文件

```go
package main

import (

     "os"

     "io"

     "log"

     "net/http"

)

func main() {

     newFile, err := os.Create("devdungeon.html")

     if err != nil {

          log.Fatal(err)

     }

     defer newFile.Close()

     url := "http://www.devdungeon.com/archive"

     response, err := http.Get(url)

     defer response.Body.Close()

     // 将HTTP response Body中的内容写入到文件

     // Body满足reader接口，因此我们可以使用ioutil.Copy

     numBytesWritten, err := io.Copy(newFile, response.Body)

     if err != nil {

          log.Fatal(err)

     }

     log.Printf("Downloaded %d byte file.\\n", numBytesWritten)

}
```

#### 哈希和摘要

```go
package main

import (

    "crypto/md5"

    "crypto/sha1"

    "crypto/sha256"

    "crypto/sha512"

    "log"

    "fmt"

    "io/ioutil"

)

func main() {

    // 得到文件内容

    data, err := ioutil.ReadFile("test.txt")

    if err != nil {

        log.Fatal(err)

    }

    // 计算Hash

    fmt.Printf("Md5: %x\\n\\n", md5.Sum(data))

    fmt.Printf("Sha1: %x\\n\\n", sha1.Sum(data))

    fmt.Printf("Sha256: %x\\n\\n", sha256.Sum256(data))

    fmt.Printf("Sha512: %x\\n\\n", sha512.Sum512(data))

}
```

上面的例子复制整个文件内容到内存中，传递给 hash 函数。  
另一个方式是创建一个 hash writer, 使用`Write`、`WriteString`、`Copy`将数据传给它。  
下面的例子使用 md5 hash,但你可以使用其它的 Writer。

```go
package main

import (

    "crypto/md5"

    "log"

    "fmt"

    "io"

    "os"

)

func main() {

    file, err := os.Open("test.txt")

    if err != nil {

        log.Fatal(err)

    }

    defer file.Close()

    //创建一个新的hasher,满足writer接口

    hasher := md5.New()

    _, err = io.Copy(hasher, file)

    if err != nil {

        log.Fatal(err)

    }

    // 计算hash并打印结果。

    // 传递 nil 作为参数，因为我们不通参数传递数据，而是通过writer接口。

    sum := hasher.Sum(nil)

    fmt.Printf("Md5 checksum: %x\\n", sum)

}
```

### 参考

[Go Standard Library Documentation](https://golang.org/pkg)
