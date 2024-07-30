package pojo

// 请求表单
type RequestTracking struct {
	Locale           string `json:"Locale,omitempty"`
	ReturnSignature  bool   `json:"ReturnSignature,omitempty"`
	ReturnMilestones bool   `json:"ReturnMilestones,omitempty"`
	ReturnPOD        bool   `json:"ReturnPOD,omitempty"`
}

// 请求 header
type RequestHeader struct {
	Authorization  string `json:"Authorization,require" :"Authorization" :"Authorization"`
	TransId        string `json:"TransId,require" :"TransId"`
	TransactionSrc string `json:"TransactionSrc" require`
}
