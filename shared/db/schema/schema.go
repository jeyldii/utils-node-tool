package schema

type Endpoints struct {
	Currency  string   `json:"currency"`
	Addresses []string `json:"addresses"`
	Reserve   string   `json:"reserve"`
}