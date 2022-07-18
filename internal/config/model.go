package config

type Page struct {
	Property []PageProperty `json:"property"`
}

type PageProperty struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	//notion type
	Type string `json:"type"`
}
