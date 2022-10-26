package main

import (
	"b2broker-task/domain"
	"b2broker-task/settings"
	"log"
	"sync"
	"time"
)

const LabsCount = 8

func main() {
	start := time.Now()
	log.Println("app started")

	students := settings.MakeStudents()
	teachers := settings.MakeTeachers()
	//examiners := settings.MakeExaminers()

	studentsWorkerPool := domain.NewStudentsWorkerPool(students)
	teachersWorkerPool := domain.NewTeacherWorkerPool(teachers)
	//examinerWorkerPool := domain.NewExaminerWorkerPool(examiners)

	serviceChannels := domain.NewServiceChannels(students, LabsCount)

	mainWG := new(sync.WaitGroup)
	mainWG.Add(2)

	// Запуск воркер пула студентов
	go studentsWorkerPool.DoTasks(serviceChannels, mainWG)
	// запуск воркер пула преподавателей
	go teachersWorkerPool.DoTasks(serviceChannels, mainWG)

	// запуск воркер пула экзаменаторов
	//go examinerWorkerPool.DoTasks(examChan, mainWG)

	mainWG.Wait()
	log.Println(time.Since(start))

	for i := range students {
		log.Printf(
			"Student %s %s has %d accepted labs",
			students[i].Name, students[i].Lastname, students[i].AcceptedLabs,
		)
	}
	log.Println("app finished")
}
