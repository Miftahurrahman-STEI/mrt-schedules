package station

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/Miftahurrahman-STEI/mrt-schedules/modules/common/client"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
	CheckScheduleByStation(id string) (response []StationScheduleResponse, err error)
}

func convertDataToResponses(schedules []StationSchedule) ([]StationScheduleResponse, error) {
	var responses []StationScheduleResponse
	for _, schedule := range schedules {
		response := StationScheduleResponse{
			Id:        schedule.Id,
			StationId: schedule.StationId,
			Time:      schedule.Time,
		}
		responses = append(responses, response)
	}
	return responses, nil
}

type Schedule struct {
	// Define the fields for Schedule
	Id        string    `json:"id"`
	StationId string    `json:"station_id"`
	Time      time.Time `json:"time"`
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service) GetAllStation() (response []StationResponse, err error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	// hit url
	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)

	// kita keluarin response
	for _, item := range stations {
		// response = append(response, StationResponse{
		// 	Id: item.Id,
		// 	Name: item.Name,
		// })
		response = append(response, StationResponse(item))
	}

	return
}

// CheckScheduleByStation implements Service.
func (s *service) CheckScheduleByStation(id string) (response []StationScheduleResponse, err error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var schedule []StationSchedule
	err = json.Unmarshal(byteResponse, &schedule)

	if err != nil {
		return
	}

	response, err = convertDataToResponses(scheduleSelected)

	if err != nil {
		return
	}

	var scheduleSelected StationSchedule

	if scheduleSelected.StationId == "" {
		err = errors.New("station not found")
		return
	}

	response, err = ConvertDataToResponses(scheduleSelected)

	if err != nil {
		return
	}

	return
}
