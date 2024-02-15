package todo

type item struct {
	name      string
	_type     string
	completed bool
}

type List []item

func (l List) Complete() {}

func (l List) Add() {}

func (l List) Delete(id int) {}

func (l List) Save() {}

func (l List) Get() {}
