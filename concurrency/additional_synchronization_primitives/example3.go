package additional_synchronization_primitives

import "sync"

type Person struct {
	name string
}
type PersonPool struct {
	pool sync.Pool
}

func NewPersonPool() *PersonPool {
	return &PersonPool{
		pool: sync.Pool{
			New: func() interface{} { return new(Person) },
		},
	}
}
func (p *PersonPool) Get() *Person {
	return p.pool.Get().(*Person)
}
func (p *PersonPool) Put(person *Person) {
	p.pool.Put(person)
}
