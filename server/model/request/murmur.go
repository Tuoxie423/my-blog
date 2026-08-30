package request

type MurmurCreate struct {
	Content string `json:"content" binding:"required"`
}

type MurmurDelete struct {
	Ids []uint `json:"ids" binding:"required"`
}
