package settings

import (
	"b2broker-task/domain"
	"time"
)

func MakeStudents() []domain.Student {
	timeForLab := 1 * time.Second
	timeForLabCorrect := 1 * time.Second
	return []domain.Student{
		{Name: "Nikita", Lastname: "Tatarskikh", TimeForLab: timeForLab, TimeForLabCorrect: timeForLabCorrect},
		{Name: "Ivan", Lastname: "Virolainen", TimeForLab: timeForLab, TimeForLabCorrect: timeForLabCorrect},
		{Name: "Maxim", Lastname: "Stolyarov", TimeForLab: timeForLab, TimeForLabCorrect: timeForLabCorrect},
	}
}

func MakeTeachers() []domain.Teacher {
	timeForCheck := 1 * time.Second

	return []domain.Teacher{
		{Name: "Vitaly", Lastname: "Koshelev", TimeForCheck: timeForCheck},
		{Name: "Ivan", Lastname: "Holopov", TimeForCheck: timeForCheck},
	}
}

func MakeExaminers() []domain.Examiner {
	timeForExam := 5 * time.Second

	return []domain.Examiner{
		{ExaminerId: 1, TimeForExam: timeForExam},
		{ExaminerId: 2, TimeForExam: timeForExam},
	}
}
