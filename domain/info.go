package domain

type ServiceInfo struct {
	student Student
	labWork LaboratoryWork
}

type ServiceChannels struct {
	LabsQueue          chan ServiceInfo
	LabsToCorrectQueue chan ServiceInfo
	AcceptedLabsQueue  chan ServiceInfo
	StudentDoneChan    chan struct{}
	LabsDoneChan       chan struct{}
}

func NewServiceChannels(students []Student, labsCount int) ServiceChannels {
	return ServiceChannels{
		LabsQueue:          make(chan ServiceInfo, 2*len(students)*labsCount),
		LabsToCorrectQueue: make(chan ServiceInfo, 2*len(students)*labsCount),
		AcceptedLabsQueue:  make(chan ServiceInfo, 2*len(students)*labsCount),
		StudentDoneChan:    make(chan struct{}, len(students)),
		LabsDoneChan:       make(chan struct{}, len(students)),
	}
}
