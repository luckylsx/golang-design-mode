package filterpattern

import "testing"

func TestFilterPattern(t *testing.T) {
	var persons []Person
	persons = append(persons, GetPerson("mike", "Male", "Single"))
	persons = append(persons, GetPerson("jack", "Female", "Single"))
	persons = append(persons, GetPerson("Tom", "Male", "Married"))
	persons = append(persons, GetPerson("mice", "Male", "Married"))
	maleFilter := new(CriteriaMale)
	femaleFilter := new(CriteriaFemale)
	singleFilter := new(CriteriaSingle)
	merriefFilter := new(CriteriaMarried)
	andFilter := new(AndCriterial)
	andFilter.AndCriterial(maleFilter, merriefFilter)
	mps := andFilter.MeetCriteria(persons)
	t.Logf("male and merried person : %+v", mps)
	andFilter.AndCriterial(femaleFilter, singleFilter)
	ms := andFilter.MeetCriteria(persons)
	t.Logf("female and single person: %+v", ms)
}
