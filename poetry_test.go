package poetry

import "testing"

var (

	p3 = Poem{{"The mortal fruit upon the bough",
		"Hands above the nuptial bed.",
		"The cat-bird in the tree returns",
		"The forfeit of his mutual vow.",
	}, {"The hard, untimely apple of",
		"The branch that feeds on watered rain,",
		"Takes the place upon her lips",
		"Of her late lamented love."},
		{"Many hands together press,",
			"Shaped within a static prayer",
			"Recall to one the chorister",
			"Docile in his sexless dress.",
		}}
)

func TestPoem_NumStanzas(t *testing.T) {
	p := Poem{{"The mortal fruit upon the bough",
		"Hands above the nuptial bed.",
		"The cat-bird in the tree returns",
		"The forfeit of his mutual vow.",
	}, {"The hard, untimely apple of",
		"The branch that feeds on watered rain,",
		"Takes the place upon her lips",
		"Of her late lamented love."},
		{"Many hands together press,",
			"Shaped within a static prayer",
			"Recall to one the chorister",
			"Docile in his sexless dress.",
		}}
//Stanzaz: 3,  Lines: 12

	if p.NumStanzas() != 3 {
		t.Fatalf("Unexpected stanza count %d", p.NumStanzas())
	}

	p2 := Poem{}
	if p2.NumStanzas() != 0 {
		t.Fatalf("Empty poem not equal to 0")
	}

	if p3.NumStanzas() != 3 {
		t.Fatalf("Unexpected stanza count %d", p.NumStanzas())
	}



}

func TestPoem_NumLines(t *testing.T) {
	p := Poem{{"The mortal fruit upon the bough",
		"Hands above the nuptial bed.",
		"The cat-bird in the tree returns",
		"The forfeit of his mutual vow.",
	}, {"The hard, untimely apple of",
		"The branch that feeds on watered rain,",
		"Takes the place upon her lips",
		"Of her late lamented love."},
		{"Many hands together press,",
			"Shaped within a static prayer",
			"Recall to one the chorister",
			"Docile in his sexless dress.",
		}}
	//Stanzaz: 3,  Lines: 12

	if p.NumLines() != 12 {
		t.Fatalf("Unexpected line count %d", p.NumLines())
	}

	if p3.NumLines() != 12 {
		t.Fatalf("Unexpected line count %d", p3.NumLines())
	}

	p2 := Poem{}
	if p2.NumLines() != 0 {
		t.Fatalf("Empty poem not equal to 0")
	}
}

func TestStats(t *testing.T)  {
	p := Poem{}
	v,c,q := p.Stats()
	if v != 0 || c != 0 || q != 0 {
		t.Fatalf("Bad number of vowels, constanants or punctations of empty Poem")
	}

	p = Poem{{"Hello"}}
	v,c,q = p.Stats()
	if v != 2 || c != 3 {
		t.Fatalf("Bad number of vowels, constanants or punctations of Poem {{\"Hello\"}}")
	}

	p = Poem{{"Hello, World!"}}
	v,c,q = p.Stats()
	if v != 3 || c != 7 || q != 2 {
		t.Fatalf("Bad number of vowels, constanants or punctations of Poem {{\"Hello, World!\"}} (%d , %d , %d)",
			v , c, q)
	}

}