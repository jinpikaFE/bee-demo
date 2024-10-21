package formvalidate

type User struct {
	Id       int    `json:"id" xml:"id" form:"id"`
	UserName string `json:"userName" xml:"userName" form:"userName" valid:"Required"`
	Password string `json:"password" xml:"password" form:"password" valid:"Required"`
}
