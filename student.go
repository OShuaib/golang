package Students

 var StudentsList = make(map[int]Student)

type Student struct{
	Name string
	Class string
	Age int
	Sex string
}

type Applicants struct{
	Name string
	grades []int
	age int
}


func (a Applicants) getAge() int{
	return a.age
}

//method  for average grade of applicant
func (a Applicants) getAverageGrade() int{
	sum := 0
	for _, v := range a.grades {
		sum += v
	}
	return sum / (len(a.grades))
}


