package xhttp

func FactoryGetUserInfo() HandleInterface {
	return &GetUserInfo{}
}

type GetUserInfo struct {
	Uuid string `json:"uuid"`
}

func (m *GetUserInfo) Handle() Response {
	// todo
	var resp Response

	return resp
}

func FactoryGetCondList() HandleInterface {
	return &GetCondList{}
}

type GetCondList struct {
	T string `json:"t"`
}

func (m *GetCondList) Handle() Response {
	return Response{}
}
