package main

import (
	"fmt"
)

/* 鸭子类型 */
type Duck interface {
	Gaga()
	Walk()
	Swimming()
}
type pskDuck struct {
	myTool string
}

func (p *pskDuck) Gaga() {
	fmt.Println("pskDuck Gaga")
}
func (p *pskDuck) Walk() {
	fmt.Println("pskDuck Walk")
}
func (p *pskDuck) Swimming() {
	fmt.Println("pskDuck Swimming")
}

// =====================>
type MyWriter interface {
	Write(string) error
}
type fileCloser struct {
	MyWriter // 继承其他结构体的方法
}
type fileWriter struct {
	filepath string
}
type databaseWriter struct {
	host string
	port string
	db   string
}

func (f *fileWriter) Write(string) error {
	fmt.Println("fileWriter Write")
	return nil
}
func (f *databaseWriter) Write(string) error {
	fmt.Println("databaseWriter Write")
	return nil
}
func (f *fileCloser) Close() error {
	fmt.Println("fileCloser Close")
	return nil
}

func main() {
	// pskDuck 必须拥有Duck接口的全部方法
	var psk1 Duck = &pskDuck{}
	psk1.Gaga()

	// =======>
	var myWriter MyWriter = &fileCloser{
		// &fileWriter{},
		&databaseWriter{}, // 扩展不同的类型方法
	}
	myWriter.Write("hello")
}
