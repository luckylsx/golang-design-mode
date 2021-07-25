package filterpattern

type Person struct {
	Name          string
	Gender        string
	MaritalStatus string
}

func GetPerson(name string, gender string, marital string) Person {
	return Person{
		Name:          name,
		Gender:        gender,
		MaritalStatus: marital,
	}
}

type Criteria interface {
	MeetCriteria(persons []Person) []Person
}

type CriteriaMale struct{}

// MeetCriteria 过滤男性
func (cm CriteriaMale) MeetCriteria(persons []Person) []Person {
	var malePersons []Person
	for _, person := range persons {
		if person.Gender == "Male" {
			malePersons = append(malePersons, person)
		}
	}
	return malePersons
}

type CriteriaFemale struct{}

// MeetCriteria 过滤女性
func (cf CriteriaFemale) MeetCriteria(persons []Person) []Person {
	var FemalePersons []Person
	for _, v := range persons {
		if v.Gender == "Female" {
			FemalePersons = append(FemalePersons, v)
		}
	}
	return FemalePersons
}

type CriteriaSingle struct{}

// MeetCriteria file the single person
func (cs CriteriaSingle) MeetCriteria(persons []Person) []Person {
	var singlePersons []Person
	for _, v := range persons {
		if v.MaritalStatus == "Single" {
			singlePersons = append(singlePersons, v)
		}
	}
	return singlePersons
}

type CriteriaMarried struct{}

// MeetCriteria file the single person
func (cm CriteriaMarried) MeetCriteria(persons []Person) []Person {
	var MarriedPersons []Person
	for _, v := range persons {
		if v.MaritalStatus == "Married" {
			MarriedPersons = append(MarriedPersons, v)
		}
	}
	return MarriedPersons
}

// AndCriterial define the and logic
type AndCriterial struct {
	criterial      Criteria
	otherCriterial Criteria
}

func (s *AndCriterial) AndCriterial(criterial, otherCriterial Criteria) {
	s.criterial = criterial
	s.otherCriterial = otherCriterial
}

func (s *AndCriterial) MeetCriteria(persons []Person) []Person {
	firstCriteriaPersons := s.criterial.MeetCriteria(persons)
	return s.otherCriterial.MeetCriteria(firstCriteriaPersons)
}
