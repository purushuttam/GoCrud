package request

type CreateNoteRequest struct {
	Content string `validate:"required,min=2,max=100" json:"content"`
}

type UpdateNoteResponse struct {
	Id      int    `validate:"required"`
	Content string `validate:"required,max=200" json:"content"`
}
