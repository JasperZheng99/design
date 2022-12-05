package observer

import "fmt"

// 观察者模式

// Listener 观察者
type Listener interface {
	// OnTeacherComing 老师来了后做的事情
	OnTeacherComing()
}

// Notifier 通知者
type Notifier interface {
	// Notify 向所有观察者发送通知
	Notify()
	AddListener(listener Listener)
	RemoveListener(listener Listener)
}

type Student struct {
	Name     string
	DoThings string //正在干的事情
}

func (s *Student) OnTeacherComing() {
	s.DoThings = "认真学习"
	s.Doing()

}
func (s *Student) Doing() {
	fmt.Println(s.Name, "正在", s.DoThings)
}

type Monitor struct {
	listeners []Listener
}

func (m *Monitor) Notify() {
	for _, listener := range m.listeners {
		listener.OnTeacherComing()
	}
}

func (m *Monitor) AddListener(listener Listener) {
	m.listeners = append(m.listeners, listener)
}

func (m *Monitor) RemoveListener(listener Listener) {
	for index, l := range m.listeners {
		if listener == l {
			m.listeners = append(m.listeners[:index], m.listeners[index:]...)
			break
		}
	}
}
