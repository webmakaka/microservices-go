package service

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
	tripTypes "ride-sharing/services/trip-service/pkg/types"
	"ride-sharing/shared/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	repo domain.TripRepository
}

func NewService(repo domain.TripRepository) *service {
	return &service{repo: repo}
}

func (s *service) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error) {
	t := &domain.TripModel{
		ID:       primitive.NewObjectID(),
		UserID:   fare.UserID,
		Status:   "pending",
		RideFare: fare,
	}
	return s.repo.CreateTrip(ctx, t)

}

func (s *service) GetRoute(ctx context.Context, pickup, destination *types.Coordinate) (*tripTypes.OsrmApiResponse, error) {

	return &tripTypes.OsrmApiResponse{
		Routes: []struct {
			Distance float64 `json:"distance"`
			Duration float64 `json:"duration"`
			Geometry struct {
				Coordinates [][]float64 `json:"coordinates"`
			} `json:"geometry"`
		}{
			{
				Distance: 5.0, // 5km
				Duration: 600, // 10 minutes
				Geometry: struct {
					Coordinates [][]float64 `json:"coordinates"`
				}{
					Coordinates: [][]float64{
						{pickup.Latitude, pickup.Longitude},
						{destination.Latitude, destination.Longitude},
					},
				},
			},
		},
	}, nil
}

// func (s *service) GetRoute(ctx context.Context, pickup, destination *types.Coordinate) (*types.OsrmApiResponse, error) {

// 	// http://router.project-osrm.org
// 	baseURL := "https://osrm.selfmadeengineer.com"

// 	url := fmt.Sprintf("%s/route/v1/driving/%f,%f;%f,%f?overview=full&geometries=geojson", baseURL, pickup.Longitude, pickup.Latitude, destination.Longitude, destination.Latitude)

// 	resp, err := http.Get(url)

// 	if err != nil {
// 		return nil, fmt.Errorf("[ME] failed to fetch route from OSRM API: %v", err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)

// 	if err != nil {
// 		return nil, fmt.Errorf("[ME] failed to read the response: %v", err)
// 	}

// 	var routeResp types.OsrmApiResponse
// 	if err := json.Unmarshal(body, &routeResp); err != nil {
// 		return nil, fmt.Errorf("[ME] failed to parse response: %v", err)
// 	}

// 	return &routeResp, nil
// }
