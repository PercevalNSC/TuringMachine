package dtm

import (
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
	w := "bb#"
	tm := TuringMachineSet(w, '$')
	t.Run("overwrite test", func(t *testing.T) {
		w2 := "ab#"
		tm.overWrite('a', 0)
		if w2 != tm.words {
			t.Error(tm.words)
		}
	})
}
