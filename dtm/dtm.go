package dtm

import (
	"fmt"
	"strconv"
)

type TuringMachine struct {
	word           string
	void_character byte
	nodes          []TuringNode
	end_id         []int
}

func TuringMachineSet(void_character byte) TuringMachine {
	// constructor for TuringMachine
	var turing_machine TuringMachine
	turing_machine.void_character = void_character
	turing_machine.constructNodes()

	return turing_machine
}
func (tm *TuringMachine) overWrite(b byte, i int) {
	result := []byte(tm.word)
	if b == tm.void_character {
		return
	} else if i < 0 {
		result = append([]byte{b}, result...)
	} else if i >= len(result) {
		result = append(result, b)
	} else {
		result[i] = b
	}

	tm.word = string(result)
}
func (tm *TuringMachine) constructNodes() {
	// This function set nodes of TuringMachine for TuringMachine instance.
	tm.nodes = []TuringNode{
		tm.p0Node(),
		tm.paNode(),
		tm.pbNode(),
		tm.padNode(),
		tm.pbdNode(),
		tm.lNode(),
		tm.cNode(),
		tm.fNode(),
	}
	tm.end_id = append(tm.end_id, 7)
}

// Execute Turing Machine for words with debug print, and output processed words.
func (tm *TuringMachine) RunWithDebug(word string, limit int) string {
	// This function is inputed words, and this function returns words proccessed by TuringMachine.
	word_position := 0
	node_id := 0
	var path TuringPath
	tm.word = word
	fmt.Println("-----")
	fmt.Println("input word:", tm.word)

	for count := 0; tm.nodes[node_id].route != nil && count < limit; count++ {
		if word_position < 0 || len(tm.word) <= word_position {
			fmt.Println("input:", string(tm.void_character))
			path = tm.nodes[node_id].wordPath(tm.void_character)
		} else {
			fmt.Println("input:", string(tm.word[word_position]))
			path = tm.nodes[node_id].wordPath(tm.word[word_position])
		}

		fmt.Print("\t")
		path.print()

		node_id = path.goto_id
		tm.overWrite(path.replace_char, word_position)
		word_position += path.head_moveto

		fmt.Println("\tword:", tm.word, "position:", word_position, "node_id:", node_id)
	}

	fmt.Println("output:", tm.word)

	return tm.word
}

// Execute Turing Machine for words, and output processed words.
func (tm *TuringMachine) Run(word string) string {
	word_position := 0
	node_status := 0
	var path TuringPath
	tm.word = word

	for tm.nodes[node_status].route != nil {
		if word_position < 0 || len(tm.word) <= word_position {
			// input void_character
			path = tm.nodes[node_status].wordPath(tm.void_character)
		} else {
			// input a character in the word
			path = tm.nodes[node_status].wordPath(tm.word[word_position])
		}
		node_status = path.goto_id
		tm.overWrite(path.replace_char, word_position)
		word_position += path.head_moveto
	}

	return tm.word
}

// Print a status of Turing Machine
func (tm *TuringMachine) Print() {
	fmt.Println("word:", tm.word)
	fmt.Println("void_character:", tm.void_character)
	fmt.Println("end_ids", tm.end_id)
	fmt.Println("nodes:")
	for _, node := range tm.nodes {
		fmt.Println("{" + node.str() + "}")
	}
}

type TuringPath struct {
	goto_id      int
	replace_char byte
	head_moveto  int
}

func errorTuringPath() TuringPath {
	return TuringPath{0, 0, 0}
}
func (tp *TuringPath) print() {
	fmt.Println("path:", tp.goto_id, string(tp.replace_char), tp.head_moveto)
}

// Return string for TuringPath status
func (tp *TuringPath) str() string {
	return "{" + strconv.Itoa(tp.goto_id) +
		" " + string(tp.replace_char) +
		" " + strconv.Itoa(tp.head_moveto) + "}"
}

type TuringNode struct {
	id    int
	route map[string]TuringPath
}

// Return string of TuringNode status
func (tn *TuringNode) str() string {
	turing_paths_str := "["
	i := 0
	for key, path := range tn.route {
		turing_paths_str = turing_paths_str + key + ":" + path.str()
		if i != len(tn.route)-1 {
			turing_paths_str = turing_paths_str + " "
		}
		i++
	}
	turing_paths_str = turing_paths_str + "]"
	return strconv.Itoa(tn.id) + " " + turing_paths_str
}

// Search TuringPath for input_byte
func (tn *TuringNode) wordPath(input_word byte) TuringPath {
	path, status := tn.route[string(input_word)]
	if status {
		return path
	} else {
		return errorTuringPath()
	}

}

func (tm *TuringMachine) p0Node() TuringNode {
	id := 0
	path1 := TuringPath{1, 'o', 1}
	path2 := TuringPath{2, 'p', 1}
	path3 := TuringPath{6, '#', -1}
	route := map[string]TuringPath{"a": path1, "b": path2, "#": path3}
	return TuringNode{id, route}
}
func (tm *TuringMachine) paNode() TuringNode {
	id := 1
	path1 := TuringPath{1, 'a', 1}
	path2 := TuringPath{1, 'b', 1}
	path3 := TuringPath{1, '#', 1}
	path4 := TuringPath{3, 'a', 1}
	route := map[string]TuringPath{
		"a":                       path1,
		"b":                       path2,
		"#":                       path3,
		string(tm.void_character): path4,
	}
	return TuringNode{id, route}
}
func (tm *TuringMachine) pbNode() TuringNode {
	id := 2
	path_a := TuringPath{2, 'a', 1}
	path_b := TuringPath{2, 'b', 1}
	path_sh := TuringPath{2, '#', 1}
	path_void := TuringPath{4, 'b', 1}
	route := map[string]TuringPath{
		"a":                       path_a,
		"b":                       path_b,
		"#":                       path_sh,
		string(tm.void_character): path_void,
	}
	return TuringNode{id, route}
}
func (tm *TuringMachine) padNode() TuringNode {
	id := 3
	path_void := TuringPath{5, 'b', -1}
	route := map[string]TuringPath{
		string(tm.void_character): path_void,
	}
	return TuringNode{id, route}
}
func (tm *TuringMachine) pbdNode() TuringNode {
	id := 4
	path_void := TuringPath{5, 'a', -1}
	route := map[string]TuringPath{
		string(tm.void_character): path_void,
	}
	return TuringNode{id, route}
}
func (tm *TuringMachine) lNode() TuringNode {
	id := 5
	path_a := TuringPath{5, 'a', -1}
	path_b := TuringPath{5, 'b', -1}
	path_sh := TuringPath{5, '#', -1}
	path_o := TuringPath{0, 'o', 1}
	path_p := TuringPath{0, 'p', 1}
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
	id := 6
	path_o := TuringPath{6, 'a', -1}
	path_p := TuringPath{6, 'b', -1}
	path_void := TuringPath{7, tm.void_character, 1}
	route := map[string]TuringPath{
		"o":                       path_o,
		"p":                       path_p,
		string(tm.void_character): path_void,
	}
	return TuringNode{id, route}
}
func (tm *TuringMachine) fNode() TuringNode {
	id := 7
	return TuringNode{id, nil}
}

func OverWrite(w string, b byte, i int) string {
	result := []byte(w)
	if i < 0 {
		result = append([]byte{b}, result...)
	} else if i >= len(result) {
		result = append(result, b)
	} else {
		result[i] = b
	}
	return string(result)
}
