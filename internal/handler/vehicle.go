package handler

import (
	"app/internal"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

const (
	maxSpeed = 400.0
	minSpeed = 0.0
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

type VehicleJSONBatch struct {
	Vehicles []VehicleJSON `json:"vehicles"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body VehicleJSON

		if err := request.JSON(r, &body); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid JSON Body")
			return
		}

		vehicle := internal.Vehicle{
			Id: body.ID,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           body.Brand,
				Model:           body.Model,
				Registration:    body.Registration,
				Color:           body.Color,
				FabricationYear: body.FabricationYear,
				Capacity:        body.Capacity,
				MaxSpeed:        body.MaxSpeed,
				FuelType:        body.FuelType,
				Transmission:    body.Transmission,
				Weight:          body.Weight,
				Dimensions: internal.Dimensions{
					Height: body.Height,
					Length: body.Length,
					Width:  body.Width,
				},
			},
		}

		if err := h.sv.Create(vehicle); err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleAlreadyExistsService):
				fmt.Print(err.Error())
				response.Error(w, http.StatusConflict, "Vehicle already exists")
			}
			return
		}

		data := VehicleJSON{
			ID:              vehicle.Id,
			Brand:           vehicle.Brand,
			Model:           vehicle.Model,
			Registration:    vehicle.Registration,
			Color:           vehicle.Color,
			FabricationYear: vehicle.FabricationYear,
			Capacity:        vehicle.Capacity,
			MaxSpeed:        vehicle.MaxSpeed,
			FuelType:        vehicle.FuelType,
			Transmission:    vehicle.Transmission,
			Weight:          vehicle.Weight,
			Height:          vehicle.Height,
			Length:          vehicle.Length,
			Width:           vehicle.Width,
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"Message": "successful vehicle creation",
			"Data":    data,
		})

	}

}

func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		color := chi.URLParam(r, "color")
		yearString := chi.URLParam(r, "year")

		if color == "" {
			response.Error(w, http.StatusBadRequest, "Color cannot be empty")
			return
		}

		if yearString == "" {
			response.Error(w, http.StatusBadRequest, "Year cannot be empty")
			return
		}

		year, err := strconv.Atoi(yearString)

		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid year provided")
			return
		}

		vehicles, err := h.sv.GetByColorAndYear(color, year)

		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehiclesNotFoundByCriteria):
				fmt.Println(err.Error())
				response.Error(w, http.StatusNotFound, "No vehicles found with the given criteria")
			}
			return
		}

		// response
		data := make(map[int]VehicleJSON)

		for key, value := range vehicles {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success, returning vehicles by color and year",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) GetByBrandBetweenYears() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		brand := chi.URLParam(r, "brand")
		yearStartString := chi.URLParam(r, "start_year")
		yearEndString := chi.URLParam(r, "end_year")

		if brand == "" {
			response.Error(w, http.StatusBadRequest, "Brand cannot be empty")
			return
		}

		if yearStartString == "" {
			response.Error(w, http.StatusBadRequest, "Start year cannot be empty")
			return
		}

		if yearEndString == "" {
			response.Error(w, http.StatusBadRequest, "End year cannot be empty")
			return
		}

		yearStart, err := strconv.Atoi(yearStartString)

		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid start year provided")
			return
		}

		yearEnd, err := strconv.Atoi(yearEndString)

		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid end year provided")
			return
		}

		vehicles, err := h.sv.GetByBrandBetweenYears(brand, yearStart, yearEnd)

		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehiclesNotFoundByCriteria):
				fmt.Println(err.Error())
				response.Error(w, http.StatusNotFound, "No vehicles found with the given criteria")
			}
			return
		}

		// response
		data := make(map[int]VehicleJSON)

		for key, value := range vehicles {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success, returning vehicles by brand between years",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) GetSpeedAvgByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		brand := chi.URLParam(r, "brand")

		if brand == "" {
			response.Error(w, http.StatusBadRequest, "Brand cannot be empty")
			return
		}

		avg, err := h.sv.GetSpeedAvgByBrand(brand)

		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehiclesNotFoundByCriteria):
				fmt.Println(err.Error())
				response.Error(w, http.StatusNotFound, "No vehicles found with the given criteria")
			}
			return
		}

		// response

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success, returning vehicle's speed average by brand",
			"Average": avg,
		})
	}
}

func (h *VehicleDefault) CreateMultiple() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body VehicleJSONBatch

		if err := request.JSON(r, &body); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid JSON Body")
			return
		}

		var vehicles = make(map[int]internal.Vehicle)

		for _, value := range body.Vehicles {
			vehicle := internal.Vehicle{
				Id: value.ID,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           value.Brand,
					Model:           value.Model,
					Registration:    value.Registration,
					Color:           value.Color,
					FabricationYear: value.FabricationYear,
					Capacity:        value.Capacity,
					MaxSpeed:        value.MaxSpeed,
					FuelType:        value.FuelType,
					Transmission:    value.Transmission,
					Weight:          value.Weight,
					Dimensions: internal.Dimensions{
						Height: value.Height,
						Length: value.Length,
						Width:  value.Width,
					},
				},
			}

			vehicles[vehicle.Id] = vehicle
		}

		if err := h.sv.CreateMultiple(vehicles); err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleAlreadyExistsService):
				fmt.Print(err.Error())
				response.Error(w, http.StatusConflict, err.Error())
			}
			return
		}

		var data = make(map[int]VehicleJSON)

		for key, value := range vehicles {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"Message": "successful multiple vehicle creation",
			"Data":    data,
		})

	}

}

func (h *VehicleDefault) ListByWeightRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		minWeightStr := r.URL.Query().Get("weight_min")
		maxWeightStr := r.URL.Query().Get("weight_max")

		var minWeight float64
		var maxWeight float64

		if minWeightStr != "" {
			var err error
			minWeight, err = strconv.ParseFloat(minWeightStr, 64)

			if err != nil {
				response.Error(w, http.StatusBadRequest, "Invalid min weight provided")
				return
			}
		}

		if maxWeightStr != "" {
			var err error
			maxWeight, err = strconv.ParseFloat(maxWeightStr, 64)

			if err != nil {
				response.Error(w, http.StatusBadRequest, "Invalid max weight provided")
				return
			}
		}

		vehicles, err := h.sv.ListByWeightRange(minWeight, maxWeight)

		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehiclesNotFoundByCriteria):
				fmt.Println(err.Error())
				response.Error(w, http.StatusNotFound, "No vehicles found with the given criteria")
			}
			return
		}

		// response
		data := make(map[int]VehicleJSON)

		for key, value := range vehicles {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success, returning vehicles by weight range",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) ListByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		lengthStr := r.URL.Query().Get("length")
		widthStr := r.URL.Query().Get("width")

		var minLength float64
		var maxLength float64
		var minWidth float64
		var maxWidth float64

		var minLengthStr string
		var maxLengthStr string
		var minWidthStr string
		var maxWidthStr string

		if lengthStr != "" {
			minLengthStr = strings.Split(lengthStr, "-")[0]
			maxLengthStr = strings.Split(lengthStr, "-")[1]
		}

		if widthStr != "" {
			minWidthStr = strings.Split(widthStr, "-")[0]
			maxWidthStr = strings.Split(widthStr, "-")[1]
		}

		if minLengthStr != "" {
			var err error
			minLength, err = strconv.ParseFloat(minLengthStr, 64)

			if err != nil {
				response.Error(w, http.StatusBadRequest, "Invalid min length provided")
				return
			}
		}

		if maxLengthStr != "" {
			var err error
			maxLength, err = strconv.ParseFloat(maxLengthStr, 64)

			if err != nil {
				response.Error(w, http.StatusBadRequest, "Invalid max length provided")
				return
			}
		}

		if minWidthStr != "" {
			var err error
			minWidth, err = strconv.ParseFloat(minWidthStr, 64)

			if err != nil {
				response.Error(w, http.StatusBadRequest, "Invalid min width provided")
				return
			}
		}

		if maxWidthStr != "" {
			var err error
			maxWidth, err = strconv.ParseFloat(maxWidthStr, 64)

			if err != nil {
				response.Error(w, http.StatusBadRequest, "Invalid max width provided")
				return
			}
		}

		dimensions := make(map[string]float64)

		dimensions["min_length"] = minLength
		dimensions["max_length"] = maxLength
		dimensions["min_width"] = minWidth
		dimensions["max_width"] = maxWidth

		vehicles, err := h.sv.ListByDimensions(dimensions)

		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehiclesNotFoundByCriteria):
				fmt.Println(err.Error())
				response.Error(w, http.StatusNotFound, "No vehicles found with the given criteria")
			}
			return
		}

		// response
		data := make(map[int]VehicleJSON)

		for key, value := range vehicles {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success, returning vehicles by given dimensions",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")

		if idStr == "" {
			response.Error(w, http.StatusBadRequest, "id cannot be empty")
			return
		}

		id, err := strconv.Atoi(idStr)

		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid id provided")
			return
		}

		var body VehicleJSON

		if err := request.JSON(r, &body); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid JSON Body")
			return
		}

		if body.MaxSpeed <= minSpeed || body.MaxSpeed > maxSpeed {
			response.Error(w, http.StatusBadRequest, "Invalid max speed provided")
			return
		}

		vehicle := internal.Vehicle{
			Id: id,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           body.Brand,
				Model:           body.Model,
				Registration:    body.Registration,
				Color:           body.Color,
				FabricationYear: body.FabricationYear,
				Capacity:        body.Capacity,
				MaxSpeed:        body.MaxSpeed,
				FuelType:        body.FuelType,
				Transmission:    body.Transmission,
				Weight:          body.Weight,
				Dimensions: internal.Dimensions{
					Height: body.Height,
					Length: body.Length,
					Width:  body.Width,
				},
			},
		}

		if err := h.sv.Update(&vehicle); err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFoundService):
				fmt.Print(err.Error())
				response.Error(w, http.StatusNotFound, "Vehicle not found")
			}
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"Message": "successful vehicle speed update",
		})

	}

}

func (h *VehicleDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid id provided")
			return
		}

		err = h.sv.Delete(id)

		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFoundService):
				fmt.Println(err.Error())
				response.Error(w, http.StatusNotFound, "Vehicle not found")
			}
			return
		}

		// response
		response.JSON(w, http.StatusNoContent, map[string]any{
			"message": "successful deletion",
		})
	}
}

func (h *VehicleDefault) GetAverageCapacityByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		brand := chi.URLParam(r, "brand")

		if brand == "" {
			response.Error(w, http.StatusBadRequest, "Brand cannot be empty")
			return
		}

		average, err := h.sv.GetAverageCapacityByBrand(brand)

		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehiclesNotFoundByCriteria):
				fmt.Println(err.Error())
				response.Error(w, http.StatusNotFound, "No vehicles found with the given criteria")
			}
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success, returning vehicle average by brand",
			"Average": average,
		})
	}
}
