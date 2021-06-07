/*
* @Title:  详细的错误信息（实现errors接口）
* @Author: minibear2333
* @Date:   2020-05-27 23:22
*/
package main

import "fmt"

type FileError struct {
	Op   string
	Name string
	Path string
}

func NewFileError(op string, name string, path string) *FileError {
	return &FileError{Op: op, Name: name, Path: path}
}

func (f *FileError) Error() string {
	return fmt.Sprintf("路径为 %v 的文件 %v，在 %v 操作时出错", f.Path, f.Name, f.Op)
}

func main() {
	f := NewFileError("读", "README.md", "/home/how_to_code/README.md")
	fmt.Println(f.Error())
}
