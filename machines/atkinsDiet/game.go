package atkinsDiet

import (
	"math/rand"
	"time"
)

//go:generate easyjson -all .

type Spin struct {
	Type  SpinKind
	Total uint
	Stops [5]uint
}

func Run(bet uint, lines uint8) []Spin {
	rr := make([]Spin, 0, 1)

	// Rounded to smaller
	betPerLine := bet / uint(lines)

	spins := 1
	var bonus uint16 = 1
	var totalMul uint
	for i := 0; i < spins; i++ {
		res := Spin{
			Type: Main,
		}
		if i != 0 {
			res.Type = Free
		}

		stops := generateStops()

		symbols, scatters := getSymbols(stops)
		if scatters >= 3 {
			bonus = 3
			spins += 10 // Adding free spins
		}

		//symbols = [5][3]Symbol{
		//	{Steak, BuffaloWings, Butter},
		//	{Mayonnaise, BuffaloWings, Steak},
		//	{Steak, BuffaloWings, Butter},
		//	{Butter, BuffaloWings, Bacon},
		//	{Mayonnaise, Eggs, Sausage},
		//} Total: 107

		totalMul += uint((payTable[Scatter][scatters] + calculateWin(symbols, lines)) * bonus)

		res.Total = betPerLine * totalMul
		for i, stop := range stops {
			res.Stops[i] = uint(stop) + 1
		}
		rr = append(rr, res)
	}

	return rr
}

func calculateWin(symbols [5][3]Symbol, activeLines uint8) (totalMul uint16) {
	for _, line := range lines[:activeLines-1] {
		lineSym := symbols[0][line[0]]
		var inRow uint8
		var wildInRow uint8
		for reel, elem := range line {
			curSym := symbols[reel][elem]

			if curSym == Scatter {
				break
			}

			if lineSym == Wild {
				if curSym == Wild {
					wildInRow++
				} else {
					lineSym = curSym
				}
			}

			if lineSym != Wild && curSym != Wild && curSym != lineSym {
				break
			}

			inRow++
		}

		var multiplier uint16
		// Calculating multiplier for the Wilds in a row
		multiplier = payTable[Wild][wildInRow]

		// Checking if multiplier from regular symbols is bigger than current
		n := inRow + wildInRow
		if n > 5 {
			n = 5
		}
		if payTable[lineSym][n] > multiplier {
			multiplier = payTable[lineSym][n]
		}

		totalMul += multiplier
		//fmt.Println("line:", lineN+1, "m:", multiplier)
	}
	return totalMul
}

func generateStops() [5]uint8 {
	//r := rand.NewSource(time.Now().UnixNano())

	result := [5]uint8{}
	for reel := 0; reel < 5; reel++ {
		//result[reel] = uint8(r.Int63() % 32)
		result[reel] = uint8(rand.NewSource(time.Now().UnixNano()).Int63() % 32)
	}
	return result
}

// getSymbols returns matrix of all symbols and number of scatters on the screen
func getSymbols(reels [5]uint8) ([5][3]Symbol, uint8) {
	result := [5][3]Symbol{}

	var scatters uint8
	var top, bot uint8
	for reel, mid := range reels {
		top = mid - 1
		bot = mid + 1

		if mid == 0 {
			top = 31
		} else if mid == 31 {
			bot = 0
		}

		result[reel] = [3]Symbol{
			Reals[reel][top],
			Reals[reel][mid],
			Reals[reel][bot],
		}

		for _, s := range result[reel] {
			if s == Scatter {
				scatters++
			}
		}
	}
	return result, scatters
}
