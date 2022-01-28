package lokasjon

import "testing"

func Test3(t *testing.T) {
    // Hva forventer jeg?
    wanted := "Alle Kom over elven"
    Stadie := Flytt("Øst Siden") // Hva fikk jeg?
    if Stadie != wanted {
             t.Errorf("Feil, fikk %q, ønsket %q.", Stadie, wanted)
    }
}

func test4(t *testing.T){
	wanted := "[Vest Siden - 🐔][_________________________________][🦊 🌽 🧑 - Øst Siden]"
	Stadie := Flytt("Båt")

if Stadie != wanted {
	t.Errorf("Feil, fikk %q, ønsket %q.", Stadie, wanted)
}
}

func test5(t *testing.T){
	wanted := "[Vest Siden - 🐔][______________[🧑]_______________][🦊 🌽 - Øst Siden]"
	Stadie := Flytt("Øst")

if Stadie != wanted {
	t.Errorf("Feil, fikk %q, ønsket %q.", Stadie, wanted)
}
}

func test6(t *testing.T){
	wanted := "[Vest Siden - 🐔 🧑][______________________________][🦊 🌽 - Øst Siden]"
	Stadie := Flytt("Øst")

if Stadie != wanted {
	t.Errorf("Feil, fikk %q, ønsket %q.", Stadie, wanted)
}
}