# Lösung
137846528820

# Lösungsbeschreibung
## Lösungsart


# Dauer


## Algorithmus
Wenn man alle Varianten hat, die mit "Rechts" starten, kann man diese mal zwei nehmen, da es genauso viele Varianten gibt, die mit "Runter" beginnen.

Für 2x2:
a,b,c,d

1: Rechts, Rechts, runter, runter
2: Rechts, runter, Rechts, runter
3: Rechts, runter, runter, Rechts
4: runter, Rechts, Rechts, runter
5: runter, Rechts, runter, Rechts
6: runter, Runter, rechts, rechts

Binär wäre das
1, 1, 0, 0
1, 0, 1, 0
1, 0, 0, 1
0, 1, 1, 0
0, 1, 0, 1

Das erinnert an Permutation mit Wiederholungen (Urne mit zweifarbigen Kugeln)
X! / (Y! × Z!)
Wobei X die Gesamtanzahl der Felder sind, wärend X die Schritte runter und Y die Schritte nach Rechts sind
4! / (2! x 2!) = 24 / (2 * 2) = 24 / 4 = 6
Für 4x4 würde also gelten
8! / (4! * 4!) = 40320 / (24 * 24) = 70
Kann man das kürzen ?
8! / (4! * 4!) = ( 8*7*6*5 ) *4! / (4!*4!) =  ( 8*7*6*5 ) / 4! = 70 
( 8*7*6*5 ) / 4! = ( 8*7*6*5 ) / (4 * 3 * 2 * 1) = ( 8 / 4*2 ) * (7*6*5)/(3*1) = 7*5*2 = 70

Für 20x20 würde das bedeuten:
X = 40
Y und Z = 20
40! / (20! x 20!) = ( 40 * 39 * ... * 22 * 21 ) *20! / (20!*20! ==> ( 40 * 39 * ... * 22 * 21 ) / 20!
=> (40*39*38*37*36*35*34*33*32*31*30*29*28*27*26*25*24*23*22*21) / (20*19*18*17*16*15*14*13*12*11*10*9*8*7*6*5*4*3*2*1)
=> (40/(20*2))*(39/(13*3))*(36/(9*4))*(35/(7*5))* (38*37*34*33*32*31*30*29*28*27*26*25*24*23*22*21)/(19*18*17*16*15*14*12*11*10*8*6)
=>(38*37*34*33*32*31*30*29*28*27*26*25*24*23*22*21)/(19*18*17*16*15*14*12*11*10*8*6)
=>137846528820