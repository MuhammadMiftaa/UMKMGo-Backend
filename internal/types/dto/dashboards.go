package dto

type UMKMByCardType struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

type ApplicationStatusSummary struct {
	TotalApplications int64 `json:"total_applications,omitempty"`
	InProcess         int64 `json:"in_process,omitempty"`
	Approved          int64 `json:"approved,omitempty"`
	Rejected          int64 `json:"rejected,omitempty"`
}

type ApplicationStatusDetail struct {
	Screening int64 `json:"screening,omitempty"`
	Revised   int64 `json:"revised,omitempty"`
	Final     int64 `json:"final,omitempty"`
	Approved  int64 `json:"approved,omitempty"`
	Rejected  int64 `json:"rejected,omitempty"`
}

type ApplicationByType struct {
	Funding       int64 `json:"funding,omitempty"`
	Certification int64 `json:"certification,omitempty"`
	Training      int64 `json:"training,omitempty"`
}
