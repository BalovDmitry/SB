package task

type Solution struct {
	UpdateTime string `json:"time"`
	Value      string `json:"solution"`
}

type Solutions map[string]Solution
