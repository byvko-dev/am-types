package glossary

type AchievementDetails struct {
	ID          string            `json:"achievement_id" bson:"achievement_id"`
	Name        map[string]string `json:"realm" bson:"name"`
	Section     string            `json:"section" bson:"section"`
	ImageURL    string            `json:"image" bson:"image"`
	Description string            `json:"description" bson:"description"`
	Condition   string            `json:"condition" bson:"condition"`
}
