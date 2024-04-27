package jobs

type StatusJobsResponse struct {
	Count  int64  `json:"count"`
	Status string `json:"status"`
}
