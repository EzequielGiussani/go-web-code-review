package service

import (
	"app/internal"
	"errors"
	"fmt"
	"math"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) Create(v internal.Vehicle) (err error) {
	// create vehicle in repository
	if err = s.rp.Create(v); err != nil {
		// check error type
		switch {
		case errors.Is(err, internal.ErrVehicleAlreadyExistsRepo):
			return fmt.Errorf("%w: %w", internal.ErrVehicleAlreadyExistsService, err)
		}
	}
	// return nil error
	return nil
}

func (s *VehicleDefault) GetByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	// get vehicles by color and year from repository
	v, err = s.rp.GetByColorAndYear(color, year)

	// check error type
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrNoVehiclesByColorYearRepo):
			return nil, fmt.Errorf("%w: %w", internal.ErrVehiclesNotFoundByCriteria, err)
		}
	}

	// return vehicles
	return v, nil
}

func (s *VehicleDefault) GetByBrandBetweenYears(brand string, yearStart int, yearEnd int) (v map[int]internal.Vehicle, err error) {
	// get vehicles by brand and years from repository
	v, err = s.rp.GetByBrandBetweenYears(brand, yearStart, yearEnd)

	// check error type
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrNoVehiclesByBrandYearsRepo):
			return nil, fmt.Errorf("%w: %w", internal.ErrVehiclesNotFoundByCriteria, err)
		}
	}

	// return vehicles
	return v, nil
}

func (s *VehicleDefault) GetSpeedAvgByBrand(brand string) (speedAvg float64, err error) {
	speedAvg, err = s.rp.GetSpeedAvgByBrand(brand)

	if err != nil {
		switch {
		case errors.Is(err, internal.ErrNoVehiclesByBrandRepo):
			return 0, fmt.Errorf("%w: %w", internal.ErrVehiclesNotFoundByCriteria, err)
		}
	}

	return speedAvg, nil
}

func (s *VehicleDefault) CreateMultiple(v map[int]internal.Vehicle) (err error) {
	err = s.rp.CreateMultiple(v)

	if err != nil {
		switch {
		case errors.Is(err, internal.ErrVehicleAlreadyExistsRepo):
			return fmt.Errorf("%w: %w", internal.ErrVehicleAlreadyExistsService, err)
		}
	}

	return nil
}

func (s *VehicleDefault) ListByWeightRange(weightMin, weightMax float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.ListByWeightRange(weightMin, weightMax)

	if err != nil {
		switch {
		case errors.Is(err, internal.ErrNoVehiclesByWeightRangeRepo):
			return nil, fmt.Errorf("%w: %w", internal.ErrVehiclesNotFoundByCriteria, err)
		}
	}

	return v, nil
}

func (s *VehicleDefault) ListByDimensions(d map[string]float64) (v map[int]internal.Vehicle, err error) {

	//Dimensions modifier
	var minLength = 0.0
	var maxLength = math.MaxFloat64
	var minWidth = 0.0
	var maxWidth = math.MaxFloat64

	if _, ok := d["min_length"]; ok {
		minLength = d["min_length"]
	}

	if _, ok := d["max_length"]; ok {
		maxLength = d["max_length"]
	}

	if _, ok := d["min_width"]; ok {
		minWidth = d["min_width"]
	}

	if _, ok := d["max_width"]; ok {
		maxWidth = d["max_width"]
	}

	v, err = s.rp.ListByDimensions(minLength, maxLength, minWidth, maxWidth)

	if err != nil {
		switch {
		case errors.Is(err, internal.ErrNoVehiclesByDimensionsRepo):
			return nil, fmt.Errorf("%w: %w", internal.ErrVehiclesNotFoundByCriteria, err)
		}
	}

	return v, nil
}

func (s *VehicleDefault) Update(v *internal.Vehicle) (err error) {

	err = s.rp.Update(v)

	if err != nil {
		switch {
		case errors.Is(err, internal.ErrVehicleNotFoundRepo):
			return fmt.Errorf("%w: %w", internal.ErrVehicleNotFoundService, err)
		}
	}

	return nil

}

func (s *VehicleDefault) Delete(id int) (err error) {

	err = s.rp.Delete(id)

	if err != nil {
		switch {
		case errors.Is(err, internal.ErrVehicleNotFoundRepo):
			return fmt.Errorf("%w: %w", internal.ErrVehicleNotFoundService, err)
		}
	}

	return nil

}

func (s *VehicleDefault) GetAverageCapacityByBrand(brand string) (capacityAvg float64, err error) {
	capacityAvg, err = s.rp.GetAverageCapacityByBrand(brand)

	if err != nil {
		switch {
		case errors.Is(err, internal.ErrNoVehiclesByBrandRepo):
			return 0, fmt.Errorf("%w: %w", internal.ErrVehiclesNotFoundByCriteria, err)
		}
	}

	return capacityAvg, nil
}
