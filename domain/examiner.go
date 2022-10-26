package domain

import (
	"sync"
	"time"
)

type Examiner struct {
	TimeForExam time.Duration
	ExaminerId  int
}

type ExaminersWorkerPool struct {
	examiners []Examiner
}

func NewExaminerWorkerPool(examiners []Examiner) ExaminersWorkerPool {
	return ExaminersWorkerPool{examiners: examiners}
}

func (ewp ExaminersWorkerPool) DoTasks(examChan chan ServiceInfo, group *sync.WaitGroup) {
	wg := sync.WaitGroup{}
	wg.Add(len(ewp.examiners))

	for i := 0; i < len(ewp.examiners); i++ {
		go ewp.examiners[i].DoExam(examChan, &wg)
	}
	group.Done()
}

func (ex Examiner) DoExam(examChan chan ServiceInfo, g *sync.WaitGroup) {
	//	doExamWG := new(sync.WaitGroup)
	//	doExamWG.Add(1)
	//	go func() {
	//		//TODO: implement me
	//		}
	//	}()
	//	doExamWG.Wait()
	//	g.Done()
	//
}
