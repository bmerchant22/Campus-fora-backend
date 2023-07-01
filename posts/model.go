package posts

import (
	"time"

	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Name      string     `json:"name"`
	Questions []Question `gorm:"foreignKey:TopicID"`
	Tags      []Tag      `gorm:"foreignKey:TopicID"`
}

type Question struct {
	gorm.Model
	UUID                 string                 `json:"uuid"`
	TopicID              uint                   `json:"topic_id"`
	Title                string                 `json:"title"`
	Content              string                 `json:"content"`
	CreatedByUser        string                 `json:"created_by_user"`
	Answers              []Answer               `gorm:"foreignKey:ParentID" json:"answers"`
	Tags                 []Tag                  `gorm:"many2many:question_tags;" json:"tags"`
	UserStarredQuestions []UserStarredQuestions `gorm:"foreignKey:UserID"`
}

type Tag struct {
	gorm.Model
	Name string `json:"name"`
}

type Answer struct {
	gorm.Model
	UUID          string    `json:"uuid"`
	ParentID      uint      `json:"parent_id"`
	Content       string    `json:"content"`
	IsAnswer      bool      `json:"is_answer"`
	CreatedByUser string    `json:"created_by_user"`
	Comments      []Comment `gorm:"foreignKey:ParentID"`
}

type Comment struct {
	gorm.Model
	UUID          string `json:"uuid"`
	ParentID      uint   `json:"parent_id"`
	Content       string `json:"content"`
	CreatedByUser string `json:"created_by_user"`
}

type QuestionDetail struct {
	ID            uint      `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	CreatedByUser string    `json:"created_by_user"`
}

type UserStarredQuestions struct {
	UserID     uint           `gorm:"primaryKey"`
	QuestionId uint           `gorm:"primaryKey"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
