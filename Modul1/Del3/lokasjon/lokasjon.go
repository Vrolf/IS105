package lokasjon 

import "fmt"


var HomoSapiensFlytt = [16]string { // String syntax
	"[Vest Siden - 🐔 & 🦊 & 🌽 & 🧑][_________________][ Ingen på Øst Siden- Øst Siden]", // Utformingen er delt i 3, der den første er vest siden, den i midten er "elven", mens den tredje er Øst Siden
	"[Vest Siden - 🦊 & 🌽][__________[🐔 & 🧑]________][ Ingen på Øst Siden- Øst Siden]",
	"[Vest Siden - 🦊 & 🌽][____________________________][🐔 & 🧑 - Øst Siden]",
	"[Vest Siden - 🦊 & 🌽][___________[🧑]____________][🐔 - Øst Siden]",
	"[Vest Siden - 🦊 & 🌽 & 🧑][______________________][🐔 - Øst Siden]",
	"[Vest Siden - 🦊][__________[🧑 & 🌽]_____________][🐔 - Øst Siden]",
	"[Vest Siden - 🦊][_________________________________][🐔 & 🌽 & 🧑 - Øst Siden]",
	"[Vest Siden - 🦊][____________[🧑 & 🐔]___________][🌽 - Øst Siden]",
	"[Vest Siden - 🦊 & 🐔 & 🧑][______________________][🌽 - Øst Siden]",
	"[Vest Siden - 🐔][___________[🧑 & 🦊]____________][🌽 - Øst Siden]",
	"[Vest Siden - 🐔][_________________________________][🦊 🌽 🧑 - Øst Siden]",
	"[Vest Siden - 🐔][______________[🧑]_______________][🦊 🌽 - Øst Siden]",
	"[Vest Siden - 🐔 🧑][______________________________][🦊 🌽 - Øst Siden]",
	"[Vest Siden - Ingen på Vest Siden][___[🐔 & 🧑]____][🦊 & 🌽 - Øst Siden]",
	"[Vest Siden - Ingen på Vest Siden][_________________][🐔 & 🌽 & 🧑 & 🦊 - Øst Siden, Alle På Øst Siden]"}

var Stadie = HomoSapiensFlytt[0] // Function som flytter Homosapiens til 0

func Flytt(Flytt string) string { //Function som bruker stringene
	
	if Flytt == "til 0"{
		Stadie = HomoSapiensFlytt[0] 
		fmt.Println(HomoSapiensFlytt[0])
		return HomoSapiensFlytt[0]
	}

if Flytt == "til 0"{
	Stadie = HomoSapiensFlytt[0]
	fmt.Println(HomoSapiensFlytt[0])
	return HomoSapiensFlytt[0]

}

if Flytt == "🧑 og 🐔 til 🚤" && Stadie == HomoSapiensFlytt[0] {
	Stadie = HomoSapiensFlytt[1] // Flytter Homosapiens og kylling til båt
	fmt.Println(HomoSapiensFlytt[1]) // fmt.Println([]) trenger jeg får å vise den visuelle framvisningen, hadde jeg ikke hatt denne ville outputten i terminalen bare vært "Ferdig" 
	return HomoSapiensFlytt[1]
}

if Flytt == "🧑 og 🐔 til 🚤" && Stadie == HomoSapiensFlytt[2] {
	Stadie = HomoSapiensFlytt[1]
	fmt.Println(HomoSapiensFlytt[1])
	return HomoSapiensFlytt[1]
}

if Flytt == "🧑 og 🐔 til Øst Siden" && Stadie == HomoSapiensFlytt[1] {
	Stadie = HomoSapiensFlytt[2]
	fmt.Println(HomoSapiensFlytt[2])
	return HomoSapiensFlytt[2]
}

if Flytt == "🧑 og 🐔 til Øst Siden" && Stadie == HomoSapiensFlytt[3] {
	Stadie = HomoSapiensFlytt [2]
	fmt.Println(HomoSapiensFlytt[2])
	return HomoSapiensFlytt[2]

}
if Flytt == "🧑 til 🚤" && Stadie == HomoSapiensFlytt[2] {
	Stadie = HomoSapiensFlytt[3]
	fmt.Println(HomoSapiensFlytt[3])
	return HomoSapiensFlytt[3]
} 

if Flytt == "🧑 til 🚤" && Stadie == HomoSapiensFlytt[4] {
	Stadie = HomoSapiensFlytt[3]
	fmt.Println(HomoSapiensFlytt[3])
	return HomoSapiensFlytt[3]
}
if Flytt == "🧑 til Vest Siden" && Stadie == HomoSapiensFlytt[3] {
	Stadie = HomoSapiensFlytt[4]
	fmt.Println(HomoSapiensFlytt[4])
	return HomoSapiensFlytt[4]
}

if Flytt == "🧑 til Vest Siden" && Stadie == HomoSapiensFlytt[5] {
	Stadie = HomoSapiensFlytt[4]
	fmt.Println(HomoSapiensFlytt[4])
	return HomoSapiensFlytt[4]

}

if Flytt == "🧑 og 🌽 til 🚤" && Stadie == HomoSapiensFlytt[4] {
	Stadie = HomoSapiensFlytt[5]
	fmt.Println(HomoSapiensFlytt[5])
	return HomoSapiensFlytt[5]
}

if Flytt == "🧑 og 🌽 til 🚤" && Stadie == HomoSapiensFlytt[6] {
	Stadie = HomoSapiensFlytt[5]
	fmt.Println(HomoSapiensFlytt[5])
	return HomoSapiensFlytt[5]
}
if Flytt == "🧑 og 🌽 til Øst Siden" && Stadie == HomoSapiensFlytt[5] {
	Stadie = HomoSapiensFlytt[6]
	fmt.Println(HomoSapiensFlytt[6])
	return HomoSapiensFlytt[6]
}
if Flytt == "🧑 og 🌽 til Øst Siden" && Stadie == HomoSapiensFlytt[7] {
	Stadie = HomoSapiensFlytt[6]
	fmt.Println(HomoSapiensFlytt[6])
	return HomoSapiensFlytt[6]

}
if Flytt == "🧑 og 🐔 til 🚤" && Stadie == HomoSapiensFlytt[6] {
	Stadie = HomoSapiensFlytt[7]
	fmt.Println(HomoSapiensFlytt[7])
	return HomoSapiensFlytt[7]
}

if Flytt == "🧑 og 🐔 til 🚤" && Stadie == HomoSapiensFlytt[8] {
	Stadie = HomoSapiensFlytt[7]
	fmt.Println(HomoSapiensFlytt[7])
	return HomoSapiensFlytt[7]
}

if Flytt == "🧑 og 🐔 til Vest Siden" && Stadie == HomoSapiensFlytt[7] {
	Stadie = HomoSapiensFlytt[8]
	fmt.Println(HomoSapiensFlytt[8])
	return HomoSapiensFlytt[8]
}

if Flytt == "🧑 og 🐔 til Vest Siden" && Stadie == HomoSapiensFlytt[9] {
	Stadie = HomoSapiensFlytt[8]
	fmt.Println(HomoSapiensFlytt[8])
	return HomoSapiensFlytt[8]

}

if Flytt == "🧑 og 🦊 til 🚤" && Stadie == HomoSapiensFlytt[8] {
	Stadie = HomoSapiensFlytt[9]
	fmt.Println(HomoSapiensFlytt[9])
	return HomoSapiensFlytt[9]
}

if Flytt == "🧑 og 🦊 til 🚤" && Stadie == HomoSapiensFlytt[10] {
	Stadie = HomoSapiensFlytt[9]
	fmt.Println(HomoSapiensFlytt[9])
	return HomoSapiensFlytt[9]
}

if Flytt == "🧑 og 🦊 til Øst Siden" && Stadie == HomoSapiensFlytt[9] {
	Stadie = HomoSapiensFlytt[10]
	fmt.Println(HomoSapiensFlytt[10])
	return HomoSapiensFlytt[10]
}

if Flytt == "🧑 og 🦊 til Øst Siden" && Stadie == HomoSapiensFlytt[11] {
	Stadie = HomoSapiensFlytt[10]
	fmt.Println(HomoSapiensFlytt[10])
	return HomoSapiensFlytt[10]
}

if Flytt == "🧑 til 🚤" && Stadie == HomoSapiensFlytt[10] {
	Stadie = HomoSapiensFlytt[11]
	fmt.Println(HomoSapiensFlytt[11])
	return HomoSapiensFlytt[11]
}

if Flytt == "🧑 til 🚤" && Stadie == HomoSapiensFlytt[12] {
	Stadie = HomoSapiensFlytt[11]
	fmt.Println(HomoSapiensFlytt[11])
	return HomoSapiensFlytt[11]
}

if Flytt == "🧑 til Vest Siden" && Stadie == HomoSapiensFlytt[11] {
	Stadie = HomoSapiensFlytt[12]
	fmt.Println(HomoSapiensFlytt[12])
	return HomoSapiensFlytt[12]
}

if Flytt == "🧑 til Vest Siden" && Stadie == HomoSapiensFlytt[13] {
	Stadie = HomoSapiensFlytt[12]
	fmt.Println(HomoSapiensFlytt[12])
	return HomoSapiensFlytt[12]

}

if Flytt == "🧑 og 🐔 til 🚤" && Stadie == HomoSapiensFlytt[12] {
	Stadie = HomoSapiensFlytt[13]
	fmt.Println(HomoSapiensFlytt[13])
	return HomoSapiensFlytt[13]
}

if Flytt == "🧑 og 🐔 til 🚤" && Stadie == HomoSapiensFlytt[14] {
	Stadie = HomoSapiensFlytt[13]
	fmt.Println(HomoSapiensFlytt[13])
	return HomoSapiensFlytt[13]
}

if Flytt == "🧑 og 🐔 til Øst Siden" && Stadie == HomoSapiensFlytt[13] {
	Stadie = HomoSapiensFlytt[14]
	fmt.Println(HomoSapiensFlytt[14])
	return HomoSapiensFlytt[14]
}

if Flytt == "🧑 og 🐔 til Øst Siden" && Stadie == HomoSapiensFlytt[15] {
	Stadie = HomoSapiensFlytt[14]
	fmt.Println(HomoSapiensFlytt[14])
	return HomoSapiensFlytt[14]
}

return "Alle Kom over elven"

}
