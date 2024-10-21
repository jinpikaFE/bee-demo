package formvalidate

type Test struct {
	Id   int    `json:"id" xml:"id" form:"id"`
	Name string `json:"name" xml:"name" form:"name" valid:"Required"`
	Age  int    `json:"age" xml:"age" form:"age" valid:"Required"`
}
