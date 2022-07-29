package _io

import (
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	//writer, err := os.OpenFile(p.Title, os.O_CREATE | os.O_WRONLY, 0666)
	//if err != nil {
	//	return err
	//}
	//
	//_, err = writer.Write(p.Body)
	//if err != nil {
	//	return err
	//}

	err := ioutil.WriteFile(p.Title, p.Body, 0666)

	return err
}

func (p *Page) load(title string) (err error) {
	p.Title = title
	p.Body, err = ioutil.ReadFile(p.Title)
	return err
}
