# Plan for oppgave (Test-dreven utvikling) RiverCrossing Prosjekt


Prosjektet bør benytte "Test-Dreven" utvikling, der koden testes kontinuerlig før implementering. Dette vil gi store fordeler, ettersom koden er testet til perfeksjon før den implementeres. Dette vil gjøre prosessen enklere. 

CASE: Problemdefinisjon (eller definisjon av "verden"): En Homo Sapiens (HS) skal over Elven med en Kylling, en sekk med Korn og en Rev (kan også være Ulv, Sau og Kålhode) og har en Båt til disposisjon. Båten kan kun ta HS og en av de andre ting/dyr. Når HS passer på ting/dyr, som er på samme sted, kan ikke disse gjøre noe ulovlig. Når HS er i båten, så kan den ikke passe på dyrene, som er på land og de kan gjennomføre ulovlige handlinger. En mulig utfordring kan være å finne en måte å få alt over elven uten at ting/dyr blir spist av hverandre, og eventuelt med færrest mulig krysninger


# Funksjoner koden bør inneholde:

1) Viewstate funksjon
2) Grafisk fremstilling av systemet
3) PutInBoat-funksjon som gjør at ting kan "puttes" i båten. 
4) CrossRiver - Funksjon som (Flytter båten til den andre siden av elven), og som returnerer en ny tilstand.
