package config

type SlotGame struct {
	ID       int    `json:"id"`
	ParSheet string `json:"parSheet"`
}

type Config struct {
	SlotGames []SlotGame `json:"SlotGames"`
}
