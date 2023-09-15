package model

type LineWin struct {
	Win      float64 `json:"win"`
	Position []int   `json:"pos"`
}

type AllLineWin struct {
	Win      float64 `json:"win"`
	Position [][]int `json:"pos"`
}

//
//type BonusWin interface{}
