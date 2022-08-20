package statistics

// AchievementsFrame - Supports all player achievements, only used for leadeboards at the moment
type AchievementsFrame struct {
	MarkOfMastery    int `json:"markOfMastery,omitempty" gorm:"column:mark_of_mastery"`
	MarkOfMasteryI   int `json:"markOfMasteryI,omitempty" gorm:"column:mark_of_mastery_1"`
	MarkOfMasteryII  int `json:"markOfMasteryII,omitempty" gorm:"column:mark_of_mastery_2"`
	MarkOfMasteryIII int `json:"markOfMasteryIII,omitempty" gorm:"column:mark_of_mastery_3"`
}

// Add achievements from b to a
func (a *AchievementsFrame) Add(b *AchievementsFrame) {
	// Achievements
	a.MarkOfMastery += b.MarkOfMastery
	a.MarkOfMasteryI += b.MarkOfMasteryI
	a.MarkOfMasteryII += b.MarkOfMasteryII
	a.MarkOfMasteryIII += b.MarkOfMasteryIII
}

// Subtracts b from a
func (a *AchievementsFrame) Subtracts(b *AchievementsFrame) {
	// Achievements
	a.MarkOfMastery -= b.MarkOfMastery
	a.MarkOfMasteryI -= b.MarkOfMasteryI
	a.MarkOfMasteryII -= b.MarkOfMasteryII
	a.MarkOfMasteryIII -= b.MarkOfMasteryIII

}
