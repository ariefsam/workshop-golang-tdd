package main

func (p *Person) Cetak() (name string) {
	name = p.FirstName + " " + p.LastName
	return
}
