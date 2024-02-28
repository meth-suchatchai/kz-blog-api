package calculateservices

type defaultService struct {
}

func NewService() Service {
	return &defaultService{}
}
