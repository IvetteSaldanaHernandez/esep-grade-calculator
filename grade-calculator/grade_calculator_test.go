package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	calculator := NewGradeCalculator()

	calculator.AddGrade("Assignment 1", 75, Assignment)
	calculator.AddGrade("Exam 1", 75, Exam)
	calculator.AddGrade("Essay 1", 75, Essay)

	actual := calculator.GetFinalGrade()

	if actual != "C" {
		t.Errorf("Expected C, got %s", actual)
	}
}

func TestGetGradeD(t *testing.T) {
	calculator := NewGradeCalculator()

	calculator.AddGrade("Assignment 1", 65, Assignment)
	calculator.AddGrade("Exam 1", 65, Exam)
	calculator.AddGrade("Essay 1", 65, Essay)

	actual := calculator.GetFinalGrade()

	if actual != "D" {
		t.Errorf("Expected D, got %s", actual)
	}
}

func TestMultipleGradesPerCategory(t *testing.T) {
	calculator := NewGradeCalculator()

	// Assignment average: 90
	calculator.AddGrade("Assignment 1", 100, Assignment)
	calculator.AddGrade("Assignment 2", 80, Assignment)

	// Exam average: 80
	calculator.AddGrade("Exam 1", 90, Exam)
	calculator.AddGrade("Exam 2", 70, Exam)

	// Essay average: 100
	calculator.AddGrade("Essay 1", 100, Essay)

	// Weighted result: 45 + 28 + 15 = 88 (B)
	actual := calculator.GetFinalGrade()

	if actual != "B" {
		t.Errorf("Expected B, got %s", actual)
	}
}

func TestEmptyGradeCategory(t *testing.T) {
	calculator := NewGradeCalculator()

	// No grades have been added.
	actual := calculator.GetFinalGrade()

	if actual != "F" {
		t.Errorf("Expected F for an empty calculator, got %s", actual)
	}
}

func TestGradeTypeString(t *testing.T) {
	tests := []struct {
		gradeType GradeType
		expected  string
	}{
		{Assignment, "assignment"},
		{Exam, "exam"},
		{Essay, "essay"},
	}

	for _, test := range tests {
		actual := test.gradeType.String()

		if actual != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, actual)
		}
	}
}


func TestPassFailPass(t *testing.T) {
	calculator := NewGradeCalculator()

	// Enable Pass/Fail evaluation.
	calculator.passFail = true

	calculator.AddGrade("Assignment", 85, Assignment)
	calculator.AddGrade("Exam", 80, Exam)
	calculator.AddGrade("Essay", 75, Essay)

	actual := calculator.GetFinalGrade()

	if actual != "P" {
		t.Errorf("Expected P, got %s", actual)
	}
}

func TestPassFailFail(t *testing.T) {
	calculator := NewGradeCalculator()
	calculator.passFail = true

	calculator.AddGrade("Assignment", 50, Assignment)
	calculator.AddGrade("Exam", 55, Exam)
	calculator.AddGrade("Essay", 60, Essay)

	actual := calculator.GetFinalGrade()

	if actual != "F" {
		t.Errorf("Expected F, got %s", actual)
	}
}
