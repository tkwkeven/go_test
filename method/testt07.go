package method

type Ssa struct {
	Name string
	Age  int
}

func (u *Ssa) Getname() string {
	return u.Name
}

var sa Ssa

func Newss(name string, age int) *Ssa {
	sa.Name = name
	sa.Age = age
	return &sa
}
