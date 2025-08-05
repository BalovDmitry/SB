package task

type Solution struct {
	UpdateTime string `json:"time"`
	Value      string `json:"solution"`
}

type Solutions struct {
	Id                 int    `json:"id" validate:"required,min=0"`
	Name               string `json:"name" validate:"required"`
	Description        string `json:"description" validate:"required"`
	LanguageToSolution map[string]Solution
}

type SolutionView struct {
	IsSingle       bool
	Solutions      []Solutions
	SolutionSingle Solutions
}

//type Solutions map[string]Solution
