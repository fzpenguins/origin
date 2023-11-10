package serialize

type Response struct {
	//Id     uint        `json:"id"`
	Status uint        `json:"status"`
	Msg    string      `json:"msg"`
	Error  error       `json:"error"`
	Data   interface{} `json:"data"`
}

type ResponseWithoutData struct {
	Status uint   `json:"status"`
	Msg    string `json:"msg"`
	Error  error  `json:"error"`
}

func WithoutData(response Response) ResponseWithoutData {
	return ResponseWithoutData{
		Status: response.Status,
		Msg:    response.Msg,
		Error:  response.Error,
	}
}

type TokenData struct {
	Data interface{} `json:"data"`
	Id   uint        `json:"id"`
}
