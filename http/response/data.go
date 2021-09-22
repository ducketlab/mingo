package response

// Data to be used by controllers.
type Data struct {
	Code      *int        `json:"code"`                // Custom return code 0: indicates normal
	Type      string      `json:"type,omitempty"`      // Data type, can be default
	Namespace string      `json:"namespace,omitempty"` // The scope of the exception
	Reason    string      `json:"reason,omitempty"`    // Abnormal reason
	Message   string      `json:"message,omitempty"`   // Descriptive information about this response
	Data      interface{} `json:"data,omitempty"`      // Specific data returned
	Meta      interface{} `json:"meta,omitempty"`      // Data meta
}

func NewData(data interface{}) *Data {
	code := -1
	return &Data{
		Code: &code,
		Data: data,
	}
}