package module

import (
	"fmt"
	"slot-server/internal/slot/api/proto"
	"slot-server/internal/slot/model"
)

type Strip struct {
	Symbol []int `toml:"symbol"`
	Weight []int `toml:"weight"`
}

func (s *Strip) StakingWeight() {
	for i := 1; i < len(s.Weight); i++ {
		s.Weight[i] += s.Weight[i-1]
	}
}

func (s *Strip) getSymbolStripByIndex(start int, count int) []int {

	strip := make([]int, count)
	for i := 0; i < count; i++ {
		if start >= len(s.Symbol) {
			start -= len(s.Symbol)
		} else if start < 0 {
			start += len(s.Symbol)
		}

		strip[i] = s.Symbol[start]
		start++

	}
	return strip
}

// GenRandomGrid Make Spin Result
func GenRandomGrid(strip []Strip, column int) [][]int {
	rnd := rollReel(strip)
	return makeGrid(strip, column, rnd)
}
func rollReel(strip []Strip) []int {
	rnd := make([]int, len(strip))
	for i := 0; i < len(strip); i++ {
		rnd[i] = rollStrip(strip[i])
	}
	return rnd
}
func rollStrip(strip Strip) int {
	rand := GenRandom()
	max := strip.Weight[len(strip.Weight)-1]
	weightRand := int(rand * float64(max))

	for index, weight := range strip.Weight {
		if weight >= weightRand {
			return index
		}
	}
	panic(fmt.Sprintf("Gen Random error!!!!! %d, %f", weightRand, rand))
}
func makeGrid(strip []Strip, column int, randoms []int) [][]int {
	grid := make([][]int, len(strip))
	for i := 0; i < len(strip); i++ {
		start := -column/2 + randoms[i]
		grid[i] = strip[i].getSymbolStripByIndex(start, column)
	}
	return grid
}

// LineWins Calculate LinePay for NormalLinePay Slot
//func _(grid [][]int, lines [][]int, pays []Line, bet float32) []models.LineWin {
//	var wins []models.LineWin
//	for _, pay := range pays {
//		for _, line := range lines {
//			lineToCheck := getLine(grid, line)
//			matched := getWinsByLine(lineToCheck, pay.Symbol)
//			if matched > 2 {
//				win := models.LineWin{
//					Win:      pay.Payout[matched-1] * bet,
//					Position: line[0:matched],
//				}
//				wins = append(wins, win)
//			}
//		}
//	}
//	return wins
//}
//func getLine(grid [][]int, line []int) []int {
//	lineToCheck := make([]int, len(line))
//	for i, pos := range line {
//		lineToCheck[i] = grid[i][pos]
//	}
//	return lineToCheck
//}
//func getWinsByLine(line []int, symbol int) int {
//	for i, s := range line {
//		if s != symbol {
//			return i
//		}
//	}
//	return len(line)
//}

func AllLineWins(grid [][]int, pays []model.Line, bet float32) []*proto.AllLineWin {
	var wins []*proto.AllLineWin
	checked := map[int]bool{}

	for _, initSymbol := range grid[0] {
		if checked[initSymbol] == true {
			continue
		}
		checked[initSymbol] = true

		var winPos []int32
		var matchedCount int32 = 1
		for i := 0; i < len(grid); i++ {
			bitmask, count := findSymbol(grid[i], initSymbol, i > 0)
			if count > 0 {
				winPos = append(winPos, bitmask)
				matchedCount *= count
			} else {
				break
			}
		}

		winCount := len(winPos)
		if winCount > 2 {
			for _, pay := range pays {
				if pay.Symbol == initSymbol {
					winCash := pay.Payout[winCount-1] * float32(matchedCount) * bet
					win := proto.AllLineWin{
						Win:      winCash,
						Position: winPos,
					}
					wins = append(wins, &win)
				}
			}

		}
	}

	return wins
}
func findSymbol(strip []int, symbolToFind int, checkWild bool) (int32, int32) {
	var bitmask int32
	var count int32 = 0
	for i, symbol := range strip {
		if symbol == symbolToFind || (checkWild && symbol == model.WildSymbol) {
			bitmask |= 1 << uint(i)
			count++
		}
	}
	return bitmask, count
}

func ScatterWin(grid [][]int, scatter model.Scatter, totalBet float32) *proto.ScatterWin {

	symbolToFind := scatter.Symbol
	var winPos []int32
	var matchedCount int32 = 0
	for i := 0; i < len(grid); i++ {
		bitmask, count := findSymbol(grid[i], symbolToFind, false)
		winPos = append(winPos, bitmask)
		matchedCount += count
	}

	if matchedCount == 0 {
		return nil
	}

	winCash := scatter.Payout[matchedCount-1] * totalBet
	bonusParam := scatter.Bonus[matchedCount-1]

	win := proto.ScatterWin{
		Win:        winCash,
		Position:   winPos,
		Bonus:      int32(scatter.BonusType),
		BonusParam: float32(bonusParam),
	}

	if winCash == 0 && bonusParam == 0 {
		return nil
	}

	return &win
}
