package dtm

type TuringMachine struct {
	words          string
	void_character string
}

func TuringMachineSet(w string, v_c byte) TuringMachine {
	var turing_machine TuringMachine
	turing_machine.words = w
	turing_machine.void_character = string(v_c)
	return turing_machine
}
func (tm *TuringMachine) overWrite(b byte, i int) {
	tm.words = OverWrite(tm.words, b, i)
}

type TuringPath struct {
	goto_id      int
	replace_char string
	head_moveto  int
}

func errorTuringPath() TuringPath {
	return TuringPath{0, "", 0}
}

type TuringNode struct {
	id    int
	route map[string]TuringPath
}

func (tn *TuringNode) wordPath(input_word string) TuringPath {
	path, status := tn.route[input_word]
	if status {
		return path
	} else {
		return errorTuringPath()
	}

}

func (tm *TuringMachine) p0Node() TuringNode {
	id := 1
	path1 := TuringPath{2, "o", 1}
	path2 := TuringPath{3, "p", 1}
	path3 := TuringPath{7, "#", -1}
	route := map[string]TuringPath{"a": path1, "b": path2, "#": path3}
	return TuringNode{id, route}
}
func (tm *TuringMachine) paNode() TuringNode {
	id := 2
	path1 := TuringPath{2, "a", 1}
	path2 := TuringPath{2, "b", 1}
	path3 := TuringPath{2, "#", 1}
	path4 := TuringPath{4, "a", 1}
	route := map[string]TuringPath{
		"a":               path1,
		"b":               path2,
		"#":               path3,
		tm.void_character: path4,
	}
	return TuringNode{id, route}
}
func (tm *TuringMachine) pbNode() TuringNode {
	id := 3
	path_a := TuringPath{3, "a", 1}
	path_b := TuringPath{3, "b", 1}
	path_sh := TuringPath{3, "#", 1}
	path_void := TuringPath{5, "b", 1}
	route := map[string]TuringPath{
		"a":               path_a,
		"b":               path_b,
		"#":               path_sh,
		tm.void_character: path_void,
	}
	return TuringNode{id, route}
}
func (tm *TuringMachine) padNode() TuringNode {
	id := 4
	path_void := TuringPath{6, "b", -1}
	route := map[string]TuringPath{
		tm.void_character: path_void,
	}
	return TuringNode{id, route}
}
func (tm *TuringMachine) pbdNode() TuringNode {
	id := 5
	path_void := TuringPath{6, "a", -1}
	route := map[string]TuringPath{
		tm.void_character: path_void,
	}
	return TuringNode{id, route}
}
func (tm *TuringMachine) lNode() TuringNode {
	id := 6
	path_a := TuringPath{6, "a", -1}
	path_b := TuringPath{6, "b", -1}
	path_sh := TuringPath{6, "#", -1}
	path_o := TuringPath{1, "o", 1}
	path_p := TuringPath{1, "p", 1}
	route := map[string]TuringPath{
		"a": path_a,
		"b": path_b,
		"#": path_sh,
		"o": path_o,
		"p": path_p,
	}
	return TuringNode{id, route}
}
func (tm *TuringMachine) cNode() TuringNode {
	id := 7
	path_o := TuringPath{7, "a", -1}
	path_p := TuringPath{7, "b", -1}
	path_void := TuringPath{8, tm.void_character, 1}
	route := map[string]TuringPath{
		"o":               path_o,
		"p":               path_p,
		tm.void_character: path_void,
	}
	return TuringNode{id, route}
}
func (tm *TuringMachine) fNode() TuringNode {
	id := 8
	return TuringNode{id, nil}
}

func OverWrite(w string, b byte, i int) string {
	result := []byte(w)
	result[i] = b
	return string(result)
}
