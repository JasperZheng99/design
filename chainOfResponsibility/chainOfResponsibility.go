package chainOfresponsibility

import "fmt"

// 责任链模式

// Patient 病人,病人拿要需要经过多个部门 前台--->医生--->药房
type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
}

// Department 部门接口
type Department interface {
	execute(*Patient)

	setNext(Department)
}

// Reception 前台
type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("挂号好了")
		r.next.execute(p)
		return
	}
	fmt.Println("前台进行挂号")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}

type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("医生诊断好了")
		d.next.execute(p)
		return
	}
	fmt.Println("医生开始诊断")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}

type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("取好药了")
		m.next.execute(p)
		return
	}
	fmt.Println("药房开始拿药")
	p.medicineDone = true

}

func (m *Medical) setNext(next Department) {
	m.next = next
}
