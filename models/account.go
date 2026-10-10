package models

// source thầy cho Balance là string, đổi sang float64 như trong bài học
// để cộng trừ tiền được
type Account struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}
