package coding

type StackOfPlates struct {
	Plates [][]int
	Cap    int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{make([][]int, 0), cap}
}

func (s *StackOfPlates) Push(val int) {
	if s.Cap == 0 {
		return
	}
	if len(s.Plates) == 0 {
		s.Plates = append(s.Plates, []int{val})
		return
	}
	lastPlate := s.Plates[len(s.Plates)-1]
	if len(lastPlate) == s.Cap {
		s.Plates = append(s.Plates, []int{val})
		return
	}
	lastPlate = append(lastPlate, val)
	s.Plates[len(s.Plates)-1] = lastPlate
}

func (s *StackOfPlates) Pop() int {
	return s.PopAt(len(s.Plates) - 1)
}

func (s *StackOfPlates) PopAt(index int) int {
	if len(s.Plates) == 0 || index >= len(s.Plates) {
		return -1
	}
	plate := s.Plates[index]
	val := plate[len(plate)-1]
	plate = plate[:len(plate)-1]
	s.Plates[index] = plate
	if len(plate) == 0 {
		s.Plates = append(s.Plates[:index], s.Plates[index+1:]...)
	}
	return val
}

/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */
