package memento

// 备忘录模式

type Memento struct {
	state string
}

func (m *Memento) SetSate(s string) {
	m.state = s
}

func (m *Memento) GetState() string {
	return m.state
}

// 发起人
type Originator struct {
	state string
}

func (o *Originator) SetSate(s string) {
	o.state = s
}

func (m *Originator) GetState() string {
	return m.state
}

func (o *Originator) CreatMemento() *Memento {
	return &Memento{state: o.state}
}

// 负责人
type Caretaker struct {
	memento *Memento
}

func (c *Caretaker) GetMemento() *Memento {
	return c.memento
}

func (c *Caretaker) SetMemento(m *Memento) {
	c.memento = m
}
