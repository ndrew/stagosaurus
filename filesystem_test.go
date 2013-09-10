package blog

import "testing"

func TestFilePathConcatenation(t *testing.T) {

	if "testo/pesto" != concatFilePaths("testo", "pesto") {
		t.Error("incorrect concatenation result")
	}
}
