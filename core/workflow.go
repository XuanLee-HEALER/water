package core

type WorkFlow struct {
	root *WorkNode[int]
}

func CreateWorkFlow(work Work) WorkFlow {
	return WorkFlow{
		&(WorkNode[int]{
			work: Pair[int, Work]{
				1, work,
			},
		}),
	}
}

func (w WorkFlow) StartWork() {
	w.root.work.V.Work()
}
