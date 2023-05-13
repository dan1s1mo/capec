package script

type ContextUnit struct {
	BlockName string
	Index     int
	Length    int
}

func (cu *ContextUnit) isFinished() bool {
	cu.Index += 1
	isFinished := cu.Length-1 == cu.Index
	return isFinished
}

type Stack struct {
	stack []ContextUnit
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack) Push(str ContextUnit) {
	s.stack = append(s.stack, str)
}

func (s *Stack) Pop() (*ContextUnit, bool, bool) {
	if s.IsEmpty() {
		return nil, true, true
	} else {
		index := len(s.stack) - 1
		cu := &(s.stack)[index]
		//fmt.Println(cu.Length-1, cu.Index)
		if cu.isFinished() {
			//fmt.Println("MINUS")
			s.stack = s.stack[:index]
			return cu, true, false
		}
		return cu, false, false
	}
}

func (s *Stack) Init(initCu ContextUnit) {
	s.stack = []ContextUnit{initCu}
}
