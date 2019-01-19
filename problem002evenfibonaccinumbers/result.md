# Lösung
4613732

# Lösungsbeschreibung
## Lösungsart
Bruteforce

# Dauer
38 Nanosekunden

## Algorithmus
tmp1 = 1
tmp2 = 1
tmp = 0
Solange die tmp kleiner als das Limit ( 4 Millionen ) ist
  tmp = tmp1 + tmp2
  tmp1 = tmp2
  tmp2 = tmp
  Prüfe ob Zahl gerade ist
  if tmp%2 == 0 {
    result += tmp2
  }
