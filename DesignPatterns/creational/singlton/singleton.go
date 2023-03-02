package singlton

type SingleTon interface {
	AddOne() int
}

type singlton struct {
	count int
}

var instance *singlton

func GetInstance() *singlton {
	if instance == nil {
		instance = new(singlton)
	}
	return instance
}

func (s *singlton) AddOne() int {
	s.count += 1
	return s.count

}
