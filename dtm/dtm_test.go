package dtm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOverwrite(t *testing.T) {
	w := "neko"
	t.Run("normal overwrite", func(t *testing.T) {
		w2 := "heko"
		w = OverWrite(w, 'h', 0)
		if w2 != w {
			t.Error(w)
		}
	})
	t.Run("second", func(t *testing.T) {
		w3 := "hako"
		w = OverWrite(w, 'a', 1)
		if w3 != w {
			t.Error(w)
		}

	})
}

func TestTuringMachine(t *testing.T) {
	tm := TuringMachineSet('$')

	t.Run("overwrite test", func(t *testing.T) {
		tm.word = "bb#"
		w2 := "ab#"
		tm.overWrite('a', 0)
		if w2 != tm.word {
			t.Error(tm.word)
		}
	})

	t.Run("p0Node", func(t *testing.T) {
		node := tm.p0Node()
		if node.id != 0 {
			t.Error(node)
		}
		if !reflect.DeepEqual(node.wordPath('a'), TuringPath{1, 'o', 1}) {
			t.Error(node.wordPath('a'))
		}
		if !reflect.DeepEqual(node.wordPath('b'), TuringPath{2, 'p', 1}) {
			t.Error(node.wordPath('b'))
		}
		if !reflect.DeepEqual(node.wordPath('#'), TuringPath{6, '#', -1}) {
			t.Error(node.wordPath('#'))
		}
		if !reflect.DeepEqual(node.wordPath('x'), errorTuringPath()) {
			t.Error(node.wordPath('x'))
		}
		fmt.Println(node)
	})
	t.Run("paNode", func(t *testing.T) {
		node := tm.paNode()
		if node.id != 1 {
			t.Error(node)
		}
		fmt.Println(node)
	})
	t.Run("pbNode", func(t *testing.T) {
		node := tm.pbNode()
		if node.id != 2 {
			t.Error(node)
		}
		fmt.Println(node)
	})
	t.Run("padNode", func(t *testing.T) {
		node := tm.padNode()
		if node.id != 3 {
			t.Error(node)
		}
		fmt.Println(node)
	})
	t.Run("pbdNode", func(t *testing.T) {
		node := tm.pbdNode()
		if node.id != 4 {
			t.Error(node)
		}
		fmt.Println(node)
	})
	t.Run("lNode", func(t *testing.T) {
		node := tm.lNode()
		if node.id != 5 {
			t.Error(node)
		}
		fmt.Println(node)
	})
	t.Run("cNode", func(t *testing.T) {
		node := tm.cNode()
		if node.id != 6 {
			t.Error(node)
		}
		fmt.Println(node)
	})
	t.Run("fNode", func(t *testing.T) {
		node := tm.fNode()
		if node.id != 7 {
			t.Error(node)
		}
		fmt.Println(node)
	})

	t.Run("run", func(t *testing.T) {
		wlist := []string{"a#", "b#", "ab#"}
		ans := []string{"a#ab", "b#ba", "ab#abba"}
		for i, w := range wlist {
			if tm.Run(w) != ans[i] {
				t.Error("Q:", w, "A:", ans[i])
			}
			if tm.RunWithDebug(w, 50) != ans[i] {
				t.Error("Q:", w, "A:", ans[i])
			}
		}
	})
}
