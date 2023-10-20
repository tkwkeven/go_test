package method

type Ssa struct {
	Name string
	Age  int
}

func Getname(s *Ssa) string {
	return s.Name
}

/*var sa *Ssa

func Newss(name string, age int) *Ssa {
	sa.name = name
	sa.age = age
	return sa
}*/
