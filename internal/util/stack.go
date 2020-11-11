package util

// Stack ...
type Stack interface {
	isEmpty() bool
	Push(item interface{})
	Pop() (interface{}, bool)
}

// ConcreteStack ...
type ConcreteStack []interface{}

// isEmpty ...
func (c *ConcreteStack) isEmpty() bool {
	return len(*c) == 0
}

// Push ...
func (c *ConcreteStack) Push(item interface{}) {
	*c = append(*c, item)
}

// Pop ...
func (c *ConcreteStack) Pop() (interface{}, bool) {
	if c.isEmpty() {
		return nil, false
	}

	index := len(*c)
	item := (*c)[index]
	*c = (*c)[:index]
	return item, true
}
