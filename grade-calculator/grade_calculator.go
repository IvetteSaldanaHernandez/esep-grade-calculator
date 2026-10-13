package esepunittests

type GradeCalculator struct {
	assignments []Grade
	exams       []Grade
	essays      []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		assignments: make([]Grade, 0),
		exams:       make([]Grade, 0),
		essays:      make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	switch gradeType {
	case Assignment:
		gc.assignments = append(gc.assignments, Grade{
			Name:  name,
			Grade: grade,
			Type:  Assignment,
		})
	case Exam:
		gc.exams = append(gc.exams, Grade{
			Name:  name,
			Grade: grade,
			Type:  Exam,
		})
	case Essay:
		gc.essays = append(gc.essays, Grade{
			Name:  name,
			Grade: grade,
			Type:  Essay,
		})
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	// Calculate the average score for each grade category.
	assignmentAverage := computeAverage(gc.assignments)
	examAverage := computeAverage(gc.exams)
	essayAverage := computeAverage(gc.essays)

	// Apply the required category weights.
	weightedGrade := float64(assignmentAverage)*0.50 +
		float64(examAverage)*0.35 +
		float64(essayAverage)*0.15

	return int(weightedGrade)
}

func computeAverage(grades []Grade) int {
	// Return zero if the category has no grades.
	if len(grades) == 0 {
		return 0
	}

	sum := 0

	// Add the score of every grade in the category.
	for _, grade := range grades {
		sum += grade.Grade
	}

	return sum / len(grades)
}
