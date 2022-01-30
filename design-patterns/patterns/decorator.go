package main

import "fmt"

type MockDB interface {
	Insert()
}

type mockDB struct{}

func (m *mockDB) Insert() {
	fmt.Println("Inserting item...")
}

type Procesor interface {
	Process()
}

type processor struct {
	db MockDB
}

func (p *processor) Process() {
	if p.db != nil {
		fmt.Println("Calling db...")
		p.db.Insert()
		fmt.Println("Processed!")
	}
}

func main() {
	var (
		db = &mockDB{}
		p  = processor{
			db,
		}
	)

	p.Process()
}
