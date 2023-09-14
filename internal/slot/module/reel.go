package module

import (
	"fmt"
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

type Pay struct {
	Symbol int       `toml:"symbol,omitempty"`
	Payout []float64 `toml:"payout,omitempty"`
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
func _(grid [][]int, lines [][]int, pays []Pay, bet float64) []model.LineWin {
	var wins []model.LineWin
	for _, pay := range pays {
		for _, line := range lines {
			lineToCheck := getLine(grid, line)
			matched := getWinsByLine(lineToCheck, pay.Symbol)
			if matched > 2 {
				win := model.LineWin{
					Win:      pay.Payout[matched-1] * bet,
					Position: line[0:matched],
				}
				wins = append(wins, win)
			}
		}
	}
	return wins
}
func getLine(grid [][]int, line []int) []int {
	lineToCheck := make([]int, len(line))
	for i, pos := range line {
		lineToCheck[i] = grid[i][pos]
	}
	return lineToCheck
}
func getWinsByLine(line []int, symbol int) int {
	for i, s := range line {
		if s != symbol {
			return i
		}
	}
	return len(line)
}

func AllLineWins(grid [][]int, pays []Pay, bet float64) []model.AllLineWin {
	var wins []model.AllLineWin

	for _, initSymbol := range grid[0] {
		var matchedGrid [][]int
		for i := 0; i < len(grid); i++ {
			matchedIndex := findSymbol(grid[i], initSymbol, i > 0)
			if len(matchedIndex) > 0 {
				matchedGrid = append(matchedGrid, matchedIndex)
			} else {
				break
			}
		}

		winCount := len(matchedGrid)
		if winCount > 2 {
			matchedCount := 1
			for _, match := range matchedGrid {
				matchedCount *= len(match)
			}

			for _, pay := range pays {
				if pay.Symbol == initSymbol {
					win := model.AllLineWin{
						Win:      pay.Payout[winCount-1] * float64(matchedCount) * bet,
						Position: matchedGrid,
					}
					wins = append(wins, win)
				}
			}

		}
	}

	return wins
}
func findSymbol(strip []int, symbolToFind int, checkWild bool) []int {
	var matchIndex []int
	for i, symbol := range strip {
		if symbol == symbolToFind || (checkWild && symbol == WildSymbol) {
			matchIndex = append(matchIndex, i)
		}
	}
	return matchIndex
}
