package internal

import (
	"errors"
)

var (
	ErrVehicleAlreadyExistsRepo    = errors.New("Vehicle ID already present")
	ErrNoVehiclesByColorYearRepo   = errors.New("No vehicles found with the given color and year")
	ErrNoVehiclesByBrandYearsRepo  = errors.New("No vehicles found with the given brand and years")
	ErrNoVehiclesByBrandRepo       = errors.New("No vehicles found with the given brand")
	ErrNoVehiclesByWeightRangeRepo = errors.New("No vehicles found with the given weight range")
	ErrNoVehiclesByDimensionsRepo  = errors.New("No vehicles found with the given dimensions")
	ErrVehicleNotFoundRepo         = errors.New("Vehicle with the provided ID not found")
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	Create(v Vehicle) (err error)
	GetByColorAndYear(color string, year int) (v map[int]Vehicle, err error)
	GetByBrandBetweenYears(brand string, yearStart int, yearEnd int) (v map[int]Vehicle, err error)
	GetSpeedAvgByBrand(brand string) (speedAvg float64, err error)
	CreateMultiple(v map[int]Vehicle) (err error)
	ListByWeightRange(weightMin, weightMax float64) (v map[int]Vehicle, err error)
	ListByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v map[int]Vehicle, err error)
	Update(v *Vehicle) (err error)
	Delete(id int) (err error)
	GetAverageCapacityByBrand(brand string) (capacityAvg float64, err error)
}
