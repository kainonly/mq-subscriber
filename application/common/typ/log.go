package typ

type Log struct {
	Identity string                 `json:"Identity"`
	Queue    string                 `json:"Queue"`
	Url      string                 `json:"Url"`
	Secret   string                 `json:"Secret"`
	Body     map[string]interface{} `json:"Body"`
	Status   bool                   `json:"Status"`
	Response map[string]interface{} `json:"Response"`
	Time     int64                  `json:"Time"`
}
