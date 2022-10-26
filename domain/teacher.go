package domain

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type Teacher struct {
	Name         string
	Lastname     string
	TimeForCheck time.Duration
	Probability  float64
}

const (
	accepted = "accepted"
	rejected = "rejected"
)

// CheckLabs - метод преподавателя для проверки лабораторных работ
func (t *Teacher) CheckLabs(srvChannels ServiceChannels, teacherWorkerPoolWaitGroup *sync.WaitGroup) {
	defer teacherWorkerPoolWaitGroup.Done()

	checkLabsWG := new(sync.WaitGroup)
	checkLabsWG.Add(1)

	// Мониторит очередь сдачи, либо принимает, либо отравляет лабораторную работу обратно как не принятую.
	go func() {
		defer checkLabsWG.Done()
		for {
			select {
			case <-srvChannels.StudentDoneChan:
				log.Println("[DEBUG] One of the students has all labs accepted")
				return
			case lab := <-srvChannels.LabsQueue:
				log.Printf(
					"Teacher %s %s started checking lab № %d sent by %s %s",
					t.Name, t.Lastname, lab.labWork.SerialNumber, lab.student.Name, lab.student.Lastname,
				)

				time.Sleep(t.TimeForCheck)

				source := rand.NewSource(time.Now().Unix())
				rnd := rand.New(source)
				potentialResults := []string{accepted, rejected}

				if potentialResults[rnd.Intn(len(potentialResults))] == rejected {
					srvChannels.LabsToCorrectQueue <- ServiceInfo{
						labWork: LaboratoryWork{SerialNumber: lab.labWork.SerialNumber},
						student: Student{
							Name:     lab.student.Name,
							Lastname: lab.student.Lastname,
						},
					}

					log.Printf(
						"Teacher %s %s rejected lab № %d sent by %s %s",
						t.Name, t.Lastname, lab.labWork.SerialNumber, lab.student.Name, lab.student.Lastname,
					)
				} else {
					srvChannels.AcceptedLabsQueue <- lab
					log.Printf(
						"Teacher %s %s accepted lab № %d sent by student %s %s",
						t.Name, t.Lastname, lab.labWork.SerialNumber, lab.student.Name, lab.student.Lastname,
					)
				}

			}
		}
	}()
}

type TeacherWorkerPool struct {
	teachers []Teacher
}

func NewTeacherWorkerPool(teachers []Teacher) TeacherWorkerPool {
	return TeacherWorkerPool{teachers: teachers}
}

// DoTasks - метод воркекра-преподавателя, запускает воркера и начинает проверку лабораторных работ i-тым преподавателем
func (twp TeacherWorkerPool) DoTasks(srvChannels ServiceChannels, group *sync.WaitGroup) {
	wg := new(sync.WaitGroup)
	wg.Add(len(twp.teachers))

	for i := 0; i < len(twp.teachers); i++ {
		go twp.teachers[i].CheckLabs(srvChannels, wg)
	}

	group.Done()
}
