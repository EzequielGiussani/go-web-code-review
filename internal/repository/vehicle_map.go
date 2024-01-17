package repository

import (
	"app/internal"
	"fmt"
	"math"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *VehicleMap) Create(v internal.Vehicle) (err error) {

	//Check if vehicle already exists
	if _, ok := r.db[v.Id]; ok {
		return internal.ErrVehicleAlreadyExistsRepo
	}

	// add vehicle to db
	r.db[v.Id] = v

	// return nil error
	return nil
}

func (r *VehicleMap) GetByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		if value.Color == color && value.FabricationYear == year {
			v[key] = value
		}
	}

	if len(v) == 0 {
		return nil, internal.ErrNoVehiclesByColorYearRepo
	}

	return
}

func (r *VehicleMap) GetByBrandBetweenYears(brand string, yearStart int, yearEnd int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		if value.Brand == brand && value.FabricationYear >= yearStart && value.FabricationYear <= yearEnd {
			v[key] = value
		}
	}

	if len(v) == 0 {
		return nil, internal.ErrNoVehiclesByBrandYearsRepo
	}

	return
}

func (r *VehicleMap) GetSpeedAvgByBrand(brand string) (speedAvg float64, err error) {
	speedAvg = 0
	count := 0

	// copy db
	for _, value := range r.db {
		if value.Brand == brand {
			speedAvg += value.MaxSpeed
			count++
		}
	}

	if count == 0 {
		return 0, internal.ErrNoVehiclesByBrandRepo
	}

	return speedAvg / float64(count), nil
}

func (r *VehicleMap) CreateMultiple(v map[int]internal.Vehicle) (err error) {
	// add vehicles to db
	for key, value := range v {
		if _, ok := r.db[key]; ok {
			return fmt.Errorf("%w key: %d", internal.ErrVehicleAlreadyExistsRepo, key)
		}

		r.db[key] = value
	}

	// return nil error
	return nil
}

func (r *VehicleMap) ListByWeightRange(weightMin, weightMax float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	if weightMax == 0 {
		weightMax = math.MaxFloat64
	}

	// copy db
	for key, value := range r.db {
		if value.Weight >= weightMin && value.Weight <= weightMax {
			v[key] = value
		}
	}

	if len(v) == 0 {
		return nil, internal.ErrNoVehiclesByWeightRangeRepo
	}

	return
}

func (r *VehicleMap) ListByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v map[int]internal.Vehicle, err error) {

	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		if value.Length >= minLength && value.Length <= maxLength && value.Width >= minWidth && value.Width <= maxWidth {
			v[key] = value
		}
	}

	if len(v) == 0 {
		return nil, internal.ErrNoVehiclesByDimensionsRepo
	}

	return
}

func (r *VehicleMap) Update(v *internal.Vehicle) (err error) {
	// check if vehicle exists
	if _, ok := r.db[v.Id]; !ok {
		return internal.ErrVehicleNotFoundRepo
	}

	// update vehicle
	r.db[v.Id] = *v

	// return nil error
	return nil
}

func (r *VehicleMap) Delete(id int) (err error) {
	// check if vehicle exists
	if _, ok := r.db[id]; !ok {
		return internal.ErrVehicleNotFoundRepo
	}

	// delete vehicle
	delete(r.db, id)

	// return nil error
	return nil
}

func (r *VehicleMap) GetAverageCapacityByBrand(brand string) (capacityAvg float64, err error) {
	capacityAvg = 0
	count := 0

	// copy db
	for _, value := range r.db {
		if value.Brand == brand {
			capacityAvg += float64(value.Capacity)
			count++
		}
	}

	if count == 0 {
		return 0, internal.ErrNoVehiclesByBrandRepo
	}

	return capacityAvg / float64(count), nil
}
