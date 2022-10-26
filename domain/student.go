package domain

import (
	"log"
	"sync"
	"time"
)

type Student struct {
	Name              string
	Lastname          string
	TimeForLab        time.Duration
	TimeForLabCorrect time.Duration
	AcceptedLabs      int
}

// DoLab - метод студента для выполнеиня лабораторных работ.
func (s *Student) DoLab(srvChannels ServiceChannels, doTaskWG *sync.WaitGroup) {
	doLabWG := new(sync.WaitGroup)
	doLabWG.Add(3)

	// выполняет лабораторные работы по списку и отправляет их в очередь сдачи
	go func() {
		defer doLabWG.Done()
		for i := 1; i <= 8; i++ {
			log.Printf("Lab № %d was started by student: %s %s", i, s.Name, s.Lastname)
			time.Sleep(s.TimeForLab)
			srvChannels.LabsQueue <- ServiceInfo{
				student: Student{
					Name:     s.Name,
					Lastname: s.Lastname,
				},
				labWork: LaboratoryWork{SerialNumber: i},
			}
			log.Printf("Lab № %d completed by student: %s %s", i, s.Name, s.Lastname)
		}
	}()

	// считает количество лабораторныех работ
	go func() {
		defer doLabWG.Done()
		for range srvChannels.AcceptedLabsQueue {
			s.AcceptedLabs++
			log.Printf(
				"Student: %s %s has %d accepted labs",
				s.Name, s.Lastname, s.AcceptedLabs,
			)
			if s.AcceptedLabs == 8 {
				log.Printf(
					"[DEBUG] Student: %s %s completed all labs %d",
					s.Name, s.Lastname, s.AcceptedLabs,
				)
				srvChannels.StudentDoneChan <- struct{}{}
				srvChannels.LabsDoneChan <- struct{}{}

				return
			}
		}
	}()

	// исправлет непринятые лабораторные работы и отправляет их обратно в очечедь сдачи
	go func() {
		defer doLabWG.Done()
		for {
			select {
			case <-srvChannels.LabsDoneChan:
				log.Println("[DEBUG] One of the students notifying that all labs are accepted]")
				return
			case cl := <-srvChannels.LabsToCorrectQueue:
				log.Printf(
					"Student: %s %s started correcting Lab № %d",
					s.Name, s.Lastname, cl.labWork.SerialNumber,
				)

				time.Sleep(s.TimeForLabCorrect)

				srvChannels.LabsQueue <- ServiceInfo{
					student: Student{
						Name:     s.Name,
						Lastname: s.Lastname,
					},
					labWork: LaboratoryWork{SerialNumber: cl.labWork.SerialNumber},
				}

				log.Printf(
					"Student: %s %s sent corrected Lab № %d",
					s.Name, s.Lastname, cl.labWork.SerialNumber,
				)
			}
		}
	}()

	doLabWG.Wait()

	doTaskWG.Done()
}

type StudentWorkerPool struct {
	students []Student
}

func NewStudentsWorkerPool(students []Student) StudentWorkerPool {
	return StudentWorkerPool{students: students}
}

// DoTasks - метод воркекра-преподавателя, запускает воркера и начинает выполнение лабораторных работ i-тым студентом
func (swp StudentWorkerPool) DoTasks(srvChannels ServiceChannels, group *sync.WaitGroup) {
	doTasksWG := new(sync.WaitGroup)
	doTasksWG.Add(len(swp.students))

	for i := 0; i < len(swp.students); i++ {
		go swp.students[i].DoLab(srvChannels, doTasksWG)
	}

	doTasksWG.Wait()

	group.Done()
}
