package gateway

type SampleUser struct {
}

func NewSampleUser() *SampleUser {
	return &SampleUser{}
}

func (s *SampleUser) Create() error {
	return nil
}
