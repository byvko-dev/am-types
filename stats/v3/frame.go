package stats

import "github.com/byvko-dev/am-types/wargaming/v2/statistics"

type Frame struct {
	Total        statistics.StatsFrame        `json:"total" bson:"total"`
	Achievements statistics.AchievementsFrame `json:"achievements" bson:"achievements"`
	Ratings      map[string]int               `json:"ratings" bson:"ratings"`
	Vehicles     map[int]VehicleStats         `json:"vehicles" bson:"vehicles"`
}

type CompleteFrame struct {
	Regular Frame `json:"regular" bson:"regular"`
	Rating  Frame `json:"rating" bson:"rating"`
}

func (f *Frame) Add(b *Frame) {
	f.Total.Add(&b.Total)
	f.Achievements.Add(&b.Achievements)
	for key, value := range b.Ratings {
		f.Ratings[key] += value
	}
	for key, value := range b.Vehicles {
		if _, ok := f.Vehicles[key]; !ok {
			f.Vehicles[key] = value
		} else {
			v := f.Vehicles[key]
			v.Add(&value)
			f.Vehicles[key] = v
		}
	}
}

func (f *Frame) Subtract(b *Frame) {
	f.Total.Subtract(&b.Total)
	f.Achievements.Subtract(&b.Achievements)
	for key, value := range b.Ratings {
		f.Ratings[key] -= value
	}
	for key, value := range b.Vehicles {
		if _, ok := f.Vehicles[key]; ok {
			v := f.Vehicles[key]
			v.Subtract(&value)
			f.Vehicles[key] = v
		}
	}
}
