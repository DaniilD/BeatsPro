package requests

type Track struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Bpm         int    `json:"bpm"`
	IsDeleted   bool   `json:"isDeleted"`
}
