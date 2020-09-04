package usecase

type HealthCheckUseCase interface {
	HealthCheck() error
}

type healthCheckUseCase struct {
}

func NewHealthCheckUseCase() HealthCheckUseCase {
	return &healthCheckUseCase{}
}

func (u *healthCheckUseCase) HealthCheck() error {
	return nil
}
