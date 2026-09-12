package dtos

type CategoryCreateDto struct {
	Name string `json:"name" validate:"required"`
}

func (t CategoryCreateDto) TableName() string {
	return "categorys"
}
