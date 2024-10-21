package formvalidate

type LoginParams struct {
	UserName string `json:"userName" xml:"userName" form:"userName" valid:"Required"`
	Password string `json:"password" xml:"password" form:"password" valid:"Required"`
}
