package main

type PriorityQueue []*Node

func (pQ PriorityQueue) Len() int {
	return len(pQ)
}

func (pQ PriorityQueue) Less(i, j int) bool {
	return pQ[i].Cost < pQ[j].Cost
}

func (pQ PriorityQueue) Swap(i, j int) {
	pQ[i], pQ[j] = pQ[j], pQ[i]
	pQ[i].Index = i
	pQ[j].Index = j
}

func (pQ *PriorityQueue) Push(x any) {
	n := len(*pQ)
	no := x.(*Node)
	no.Index = n
	*pQ = append(*pQ, no)
}

func (pQ *PriorityQueue) Pop() any {
	old := *pQ
	oldLength := len(old)
	last := old[oldLength-1]
	last.Index = -1
	*pQ = old[0 : oldLength-1]
	return last
}
