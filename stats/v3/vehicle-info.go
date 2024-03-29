package stats

type VehicleInfo struct {
	TankID    int               `json:"tank_id,omitempty" bson:"tank_id,omitempty"`
	Name      map[string]string `json:"name,omitempty" bson:"name,omitempty"`
	Nation    string            `json:"nation,omitempty" bson:"nation,omitempty"`
	Tier      int               `json:"tier,omitempty" bson:"tier,omitempty"`
	Type      string            `json:"type,omitempty" bson:"type,omitempty"`
	IsPremium bool              `json:"is_premium,omitempty" bson:"is_premium,omitempty"`
}
