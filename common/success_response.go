package common

type SuccessRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data interface{}, paging interface{},
	filter interface{}) *SuccessRes {
	return &SuccessRes{data, paging, filter}
}

func SimpleSuccessResponse(data interface{}) *SuccessRes {
	return &SuccessRes{data, nil, nil}
}
