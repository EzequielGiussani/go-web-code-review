package internal

import "errors"

var (
	ErrVehicleAlreadyExistsService = errors.New("Vehicle already exists")
	ErrVehiclesNotFoundByCriteria  = errors.New("No vehicles found with the given criteria")
	ErrVehicleNotFoundService      = errors.New("Vehicle with the provided ID not found")
	ErrNoVehiclesByBrandService    = errors.New("No vehicles found with the given brand")
)

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	Create(v Vehicle) (err error)
	GetByColorAndYear(color string, year int) (v map[int]Vehicle, err error)
	GetByBrandBetweenYears(brand string, yearStart int, yearEnd int) (v map[int]Vehicle, err error)
	GetSpeedAvgByBrand(brand string) (speedAvg float64, err error)
	CreateMultiple(v map[int]Vehicle) (err error)
	ListByWeightRange(weightMin, weightMax float64) (v map[int]Vehicle, err error)
	ListByDimensions(d map[string]float64) (v map[int]Vehicle, err error)
	Update(v *Vehicle) (err error)
	Delete(id int) (err error)
	GetAverageCapacityByBrand(brand string) (capacityAvg float64, err error)
}
