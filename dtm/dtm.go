package dtm

type TuringMachine struct {
	words          string
	void_character byte
}

func TuringMachineSet(w string, v_c byte) TuringMachine {
	var turing_machine TuringMachine
	turing_machine.words = w
	turing_machine.void_character = v_c
	return turing_machine
}
func (tm *TuringMachine) overWrite(b byte, i int) {
	tm.words = OverWrite(tm.words, b, i)
}
func (tm *TuringMachine) rootingNode(status_id int, position int) int {
	switch status_id {
	case 1:
		return tm.p0Node(position)
	case 2:
		return tm.paNode(position)
	default:
		return -1
	}
}

func (tm *TuringMachine) p0Node(i int) int {
	// id: 1
	if tm.words[i] == 'a' {
		tm.overWrite('o', i)
		return 2 // to pa
	} else if tm.words[i] == 'b' {
		tm.overWrite('p', i)
		return 3 // to pb
	} else {
		return -1
	}
}
func (tm *TuringMachine) paNode(i int) int {
	return -1
}

func OverWrite(w string, b byte, i int) string {
	result := []byte(w)
	result[i] = b
	return string(result)
}
