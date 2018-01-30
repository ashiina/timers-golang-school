package main

import "fmt"

type Accessor interface {
	GetText() string
	SetText(string)
}

type Document struct {
	text string
}

func (d *Document) GetText() string {
	return d.text
}

func (d *Document) SetText(text string) {
	d.text = text
}

type Page struct {
	Document
	Page int
}

func SetAndGet(a Accessor) {
	a.SetText("accessor")
	fmt.Println(a.GetText())
}

func main() {
	var d *Document = &Document{}
	d.SetText("document")
	var a Accessor = &Document{}
	a.SetText("accessor")

	fmt.Println(d.GetText())
	fmt.Println(a.GetText())

	SetAndGet(&Page{})
	SetAndGet(&Document{})
}
