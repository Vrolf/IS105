package lokasjon 

import "fmt"


var HomoSapiensFlytt = [16]string { // String syntax
	"[Vest Siden - π & π¦ & π½ & π§][_________________][ Ingen pΓ₯ Γst Siden- Γst Siden]", // Utformingen er delt i 3, der den fΓΈrste er vest siden, den i midten er "elven", mens den tredje er Γst Siden
	"[Vest Siden - π¦ & π½][__________[π & π§]________][ Ingen pΓ₯ Γst Siden- Γst Siden]",
	"[Vest Siden - π¦ & π½][____________________________][π & π§ - Γst Siden]",
	"[Vest Siden - π¦ & π½][___________[π§]____________][π - Γst Siden]",
	"[Vest Siden - π¦ & π½ & π§][______________________][π - Γst Siden]",
	"[Vest Siden - π¦][__________[π§ & π½]_____________][π - Γst Siden]",
	"[Vest Siden - π¦][_________________________________][π & π½ & π§ - Γst Siden]",
	"[Vest Siden - π¦][____________[π§ & π]___________][π½ - Γst Siden]",
	"[Vest Siden - π¦ & π & π§][______________________][π½ - Γst Siden]",
	"[Vest Siden - π][___________[π§ & π¦]____________][π½ - Γst Siden]",
	"[Vest Siden - π][_________________________________][π¦ π½ π§ - Γst Siden]",
	"[Vest Siden - π][______________[π§]_______________][π¦ π½ - Γst Siden]",
	"[Vest Siden - π π§][______________________________][π¦ π½ - Γst Siden]",
	"[Vest Siden - Ingen pΓ₯ Vest Siden][___[π & π§]____][π¦ & π½ - Γst Siden]",
	"[Vest Siden - Ingen pΓ₯ Vest Siden][_________________][π & π½ & π§ & π¦ - Γst Siden, Alle PΓ₯ Γst Siden]"}

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

if Flytt == "π§ og π til π€" && Stadie == HomoSapiensFlytt[0] {
	Stadie = HomoSapiensFlytt[1] // Flytter Homosapiens og kylling til bΓ₯t
	fmt.Println(HomoSapiensFlytt[1]) // fmt.Println([]) trenger jeg fΓ₯r Γ₯ vise den visuelle framvisningen, hadde jeg ikke hatt denne ville outputten i terminalen bare vΓ¦rt "Ferdig" 
	return HomoSapiensFlytt[1]
}

if Flytt == "π§ og π til π€" && Stadie == HomoSapiensFlytt[2] {
	Stadie = HomoSapiensFlytt[1]
	fmt.Println(HomoSapiensFlytt[1])
	return HomoSapiensFlytt[1]
}

if Flytt == "π§ og π til Γst Siden" && Stadie == HomoSapiensFlytt[1] {
	Stadie = HomoSapiensFlytt[2]
	fmt.Println(HomoSapiensFlytt[2])
	return HomoSapiensFlytt[2]
}

if Flytt == "π§ og π til Γst Siden" && Stadie == HomoSapiensFlytt[3] {
	Stadie = HomoSapiensFlytt [2]
	fmt.Println(HomoSapiensFlytt[2])
	return HomoSapiensFlytt[2]

}
if Flytt == "π§ til π€" && Stadie == HomoSapiensFlytt[2] {
	Stadie = HomoSapiensFlytt[3]
	fmt.Println(HomoSapiensFlytt[3])
	return HomoSapiensFlytt[3]
} 

if Flytt == "π§ til π€" && Stadie == HomoSapiensFlytt[4] {
	Stadie = HomoSapiensFlytt[3]
	fmt.Println(HomoSapiensFlytt[3])
	return HomoSapiensFlytt[3]
}
if Flytt == "π§ til Vest Siden" && Stadie == HomoSapiensFlytt[3] {
	Stadie = HomoSapiensFlytt[4]
	fmt.Println(HomoSapiensFlytt[4])
	return HomoSapiensFlytt[4]
}

if Flytt == "π§ til Vest Siden" && Stadie == HomoSapiensFlytt[5] {
	Stadie = HomoSapiensFlytt[4]
	fmt.Println(HomoSapiensFlytt[4])
	return HomoSapiensFlytt[4]

}

if Flytt == "π§ og π½ til π€" && Stadie == HomoSapiensFlytt[4] {
	Stadie = HomoSapiensFlytt[5]
	fmt.Println(HomoSapiensFlytt[5])
	return HomoSapiensFlytt[5]
}

if Flytt == "π§ og π½ til π€" && Stadie == HomoSapiensFlytt[6] {
	Stadie = HomoSapiensFlytt[5]
	fmt.Println(HomoSapiensFlytt[5])
	return HomoSapiensFlytt[5]
}
if Flytt == "π§ og π½ til Γst Siden" && Stadie == HomoSapiensFlytt[5] {
	Stadie = HomoSapiensFlytt[6]
	fmt.Println(HomoSapiensFlytt[6])
	return HomoSapiensFlytt[6]
}
if Flytt == "π§ og π½ til Γst Siden" && Stadie == HomoSapiensFlytt[7] {
	Stadie = HomoSapiensFlytt[6]
	fmt.Println(HomoSapiensFlytt[6])
	return HomoSapiensFlytt[6]

}
if Flytt == "π§ og π til π€" && Stadie == HomoSapiensFlytt[6] {
	Stadie = HomoSapiensFlytt[7]
	fmt.Println(HomoSapiensFlytt[7])
	return HomoSapiensFlytt[7]
}

if Flytt == "π§ og π til π€" && Stadie == HomoSapiensFlytt[8] {
	Stadie = HomoSapiensFlytt[7]
	fmt.Println(HomoSapiensFlytt[7])
	return HomoSapiensFlytt[7]
}

if Flytt == "π§ og π til Vest Siden" && Stadie == HomoSapiensFlytt[7] {
	Stadie = HomoSapiensFlytt[8]
	fmt.Println(HomoSapiensFlytt[8])
	return HomoSapiensFlytt[8]
}

if Flytt == "π§ og π til Vest Siden" && Stadie == HomoSapiensFlytt[9] {
	Stadie = HomoSapiensFlytt[8]
	fmt.Println(HomoSapiensFlytt[8])
	return HomoSapiensFlytt[8]

}

if Flytt == "π§ og π¦ til π€" && Stadie == HomoSapiensFlytt[8] {
	Stadie = HomoSapiensFlytt[9]
	fmt.Println(HomoSapiensFlytt[9])
	return HomoSapiensFlytt[9]
}

if Flytt == "π§ og π¦ til π€" && Stadie == HomoSapiensFlytt[10] {
	Stadie = HomoSapiensFlytt[9]
	fmt.Println(HomoSapiensFlytt[9])
	return HomoSapiensFlytt[9]
}

if Flytt == "π§ og π¦ til Γst Siden" && Stadie == HomoSapiensFlytt[9] {
	Stadie = HomoSapiensFlytt[10]
	fmt.Println(HomoSapiensFlytt[10])
	return HomoSapiensFlytt[10]
}

if Flytt == "π§ og π¦ til Γst Siden" && Stadie == HomoSapiensFlytt[11] {
	Stadie = HomoSapiensFlytt[10]
	fmt.Println(HomoSapiensFlytt[10])
	return HomoSapiensFlytt[10]
}

if Flytt == "π§ til π€" && Stadie == HomoSapiensFlytt[10] {
	Stadie = HomoSapiensFlytt[11]
	fmt.Println(HomoSapiensFlytt[11])
	return HomoSapiensFlytt[11]
}

if Flytt == "π§ til π€" && Stadie == HomoSapiensFlytt[12] {
	Stadie = HomoSapiensFlytt[11]
	fmt.Println(HomoSapiensFlytt[11])
	return HomoSapiensFlytt[11]
}

if Flytt == "π§ til Vest Siden" && Stadie == HomoSapiensFlytt[11] {
	Stadie = HomoSapiensFlytt[12]
	fmt.Println(HomoSapiensFlytt[12])
	return HomoSapiensFlytt[12]
}

if Flytt == "π§ til Vest Siden" && Stadie == HomoSapiensFlytt[13] {
	Stadie = HomoSapiensFlytt[12]
	fmt.Println(HomoSapiensFlytt[12])
	return HomoSapiensFlytt[12]

}

if Flytt == "π§ og π til π€" && Stadie == HomoSapiensFlytt[12] {
	Stadie = HomoSapiensFlytt[13]
	fmt.Println(HomoSapiensFlytt[13])
	return HomoSapiensFlytt[13]
}

if Flytt == "π§ og π til π€" && Stadie == HomoSapiensFlytt[14] {
	Stadie = HomoSapiensFlytt[13]
	fmt.Println(HomoSapiensFlytt[13])
	return HomoSapiensFlytt[13]
}

if Flytt == "π§ og π til Γst Siden" && Stadie == HomoSapiensFlytt[13] {
	Stadie = HomoSapiensFlytt[14]
	fmt.Println(HomoSapiensFlytt[14])
	return HomoSapiensFlytt[14]
}

if Flytt == "π§ og π til Γst Siden" && Stadie == HomoSapiensFlytt[15] {
	Stadie = HomoSapiensFlytt[14]
	fmt.Println(HomoSapiensFlytt[14])
	return HomoSapiensFlytt[14]
}

return "Alle Kom over elven"

}
