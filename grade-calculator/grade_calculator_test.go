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

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 73, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 78, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 61, Assignment)
	gradeCalculator.AddGrade("exam 1", 66, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 68, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}


func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 10, Assignment)
	gradeCalculator.AddGrade("exam 1", 35, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 50, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestPassing(t *testing.T) {
// Just adding in a test to ensure that a passing grade is a C or above. //

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 73, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 78, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if (actual_value != "A") && (actual_value != "B") && (actual_value != "C") {
		t.Errorf("Grade is not passing.")
	}
}

func TestGradeTypeName(t *testing.T) {
// Coverage file returns a 0 for line 24 in grade_calculator.go, so creating a test to hopefully get coverage. //
// Given the information in the .go file, takes in a GradeType followed by String() and returns the given gradeTypeName for that grade type object. //
	if Assignment.String() != "assignment" {
		t.Errorf("Assignment string is incorrect.")
	}
	if Essay.String() != "essay" {
		t.Errorf("Essay string is incorrect.")
	}
	if Exam.String() != "exam"{
		t.Errorf("Exam string is incorrect.")
	}
}