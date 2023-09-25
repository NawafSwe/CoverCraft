package valueobjects

type job struct {
	Company        string `json:"company"`
	Location       string `json:"location"`
	Description    string `json:"description"`
	Qualifications string `json:"qualifications"`
	Name           string `json:"name"`
}

type resume struct {
	Text string `json:"text"`
}

type OptimizeResumeRequest struct {
	Job    job    `json:"job"`
	Resume resume `json:"resume"`
}
type OptimizationResult struct {
	Result string `json:"result"`
}
