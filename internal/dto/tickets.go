package dto

type CreateTicketRequest struct{
	Title string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UpdateTicketRequest struct{
	Status string `json:"status" binding:"required"`
}