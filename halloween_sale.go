package main
import "fmt"

func main() {

	var p,d,m,s int
	fmt.Scanf("%d %d %d %d", &p, &d, &m, &s)
	
	if s < p {
		fmt.Println(0)
		return
	}
	var totGames int

	for p >= m && s > p { 
		// Compra jogos na PA decrescente
		s -= p
		p -= d
		totGames++
	}

	if p < m && totGames > 0 {
		for s >= m {
			// Compra jogos com mesmo preço M
			s -= m
			totGames++
		}
	}

	fmt.Println(totGames)
	

}