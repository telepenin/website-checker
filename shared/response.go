package shared

type Response struct {
	Website  Website
	Code     int      `json:"code"`
	Duration float64  `json:"duration"`
	Content  [][]byte `json:"content"`
}
