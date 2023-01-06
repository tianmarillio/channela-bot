package models

type Channel struct {
	ID               string `json:"id" gorm:"primaryKey"`
	YoutubeChannelId string `json:"youtubeChannelId" gorm:"uniqueIndex"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	CustomUrl        string `json:"customUrl"`
	CreatedBy        string `json:"createdBy"`
	UpdatedBy        string `json:"updatedBy"`

	// FIXME:
	// Creator User `gorm:"references:ID,foreignKey:CreatedBy"`

	Timestamp
}
