package util

// Queue ...
type Queue interface {
	isEmpty() bool
	Push(item interface{})
	Pop() (interface{}, bool)
}

// ConcreteQueue ...
type ConcreteQueue []interface{}

// isEmpty ...
func (c *ConcreteQueue) isEmpty() bool {
	return len(*c) == 0
}

// Push ...
func (c *ConcreteQueue) Push(item interface{}) {
	*c = append(*c, item)
}

// Pop ...
func (c *ConcreteQueue) Pop() (interface{}, bool) {
	item := (*c)[0]
	*c = (*c)[1:]

	return item, !c.isEmpty()
}
