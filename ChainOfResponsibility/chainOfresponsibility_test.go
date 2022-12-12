package chainOfresponsibility

import "testing"

func TestChainOfresponsibility(t *testing.T) {
	reception := new(Reception)
	doctor := new(Doctor)
	medical := new(Medical)

	reception.setNext(doctor)
	doctor.setNext(medical)

	patient := &Patient{
		name: "张三",
	}

	reception.execute(patient)
}
