package main

type CQueue struct {
	inStack, outStack []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (cque *CQueue) AppendTail(value int) {
	cque.inStack = append(cque.inStack, value)
}
func (cque *CQueue) DeleteHead() int {
	if len(cque.outStack) == 0 {
		if len(cque.inStack) == 0 {
			return -1

		}
		cque.in2out()
	}
	value := cque.outStack[len(cque.outStack)-1]
	cque.outStack = cque.outStack[:len(cque.outStack)-1]
	return value
}

func (cque *CQueue) in2out() {
	for len(cque.inStack) > 0 {
		cque.outStack = append(cque.outStack, cque.inStack[len(cque.inStack)-1])
		cque.inStack = cque.inStack[:len(cque.inStack)-1]
	}

}
