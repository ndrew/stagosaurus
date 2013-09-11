package blog

import "testing"

func TestMetaRegex(t *testing.T) {
	data := [...]string{"<!--!\n{}\n-->", "  <!--!\n{}\n-->", "testo  <!--!\n{}\n--> "}

	var m []string
	for _, s := range data {
		m = metadataRE.FindStringSubmatch(s)
		if len(m) < 2 || m[3] != "{}" {
			t.Errorf("Can't extract metadata from %s: %q", "<!--!\n{}\n-->", m)
		}
	}
}

func TestNewMeta(t *testing.T) {
	meta := new(Meta)

	err := meta.FromString("<!--!\n{\"Ready\":true}\n-->")

	if err != nil {
		t.Error(err)
	}

	if meta.Ready == false {
		t.Error("incorrectly parsed json")
	}

	// time testing
	err = meta.FromString("<!--!\n{\"Date\":\"2013-01-22T10:30:55Z\"}\n-->")

	if err != nil {
		t.Error(err)
	}

	if "2013-01-22 10:30:55 +0000 UTC" != meta.Date.String() {
		t.Error("incorrecly interpretated date")
	}

}
