package utils

type SignUpBody struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

type UpdateBody struct {
	Id        int32  `json:"id" binding:"required"`
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

type Response struct {
	Response string
}

type MethodsRequest struct {
	Method   int32 `json:"method"`
	WaitTime int32 `json:"waitTime"`
}
