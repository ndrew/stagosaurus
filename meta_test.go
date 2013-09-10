package blog

import "testing"

func TestMetaRegex(t *testing.T) {
	m := metadataRE.FindStringSubmatch("<!--!\n{}\n-->")
	if len(m) < 2 || m[3] != "{}" {
		t.Errorf("Can't extract metadata from %s: %q", "<!--!\n{}\n-->", m)
	}

	// test leading whitespace
	m = metadataRE.FindStringSubmatch("  <!--!\n{}\n-->")
	if len(m) < 2 || m[3] != "{}" {
		t.Errorf("Can't extract metadata from %s: %q", "<!--!\n{}\n-->", m)
	}

	m = metadataRE.FindStringSubmatch("testo  <!--!\n{}\n--> ")
	if len(m) < 2 || m[3] != "{}" {

		t.Errorf("Can't extract metadata from %s: %q", "<!--!\n{}\n-->", m)
	}

}

func TestNewMeta(t *testing.T) {

}
