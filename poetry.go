package poetry

import "unicode"

type Line string
type Stanza []Line
type Poem []Stanza

func NewPoem() Poem  {
	return Poem{}
}

func (p Poem) NumStanzas() int {
	return len(p)
}

func (s Stanza) NumLines() int   {
	return len(s)
}

func (p Poem) NumLines() (count int) {
	for _,s := range p {
		count += s.NumLines()
	}
	return
}

func (p Poem) Stats() (numVowels, numConstanants, numPuncuations int) {

	numVowels = 0        // don't have to imitialize to 0
	numConstanants = 0   // it is done by default as returned
						 // variable default initializers
	zeds := 0
	//fmt.Printf("Starting loop\n\n")
	for _, s := range p {
		for _, l := range s {
			for _, c := range l {
				//fmt.Printf("%c ", c)
				if unicode.IsPunct(c) {
					numPuncuations++
				}
				switch c {
				case ',', ' ', '!':
					//numPuncuations++
				case 'a', 'e', 'i', 'o', 'u':
					numVowels++
				case 'z':
					zeds++
					fallthrough
				default:
					numConstanants++
				}
			}
		}
	}
	//fmt.Printf("Ending loop\n\n")

	//return numVowels, numConstanants
	return     ///  default return variables
}


