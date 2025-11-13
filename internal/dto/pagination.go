package dto

// PaginatedResponse - универсальная структура для ответа с пагинацией
type PaginatedResponse[T any] struct {
	// Данные
	Data []T `json:"data"`

	// Мета-информация о пагинации
	Meta PaginationMeta `json:"meta"`
}

func NewPaginatedResponse[T any](
	data []T,
	currentPage int,
	perPage int,
	total int,
) PaginatedResponse[T] {

	count := len(data)
	totalPages := (total + perPage - 1) / perPage // округление вверх

	return PaginatedResponse[T]{
		Data: data,
		Meta: PaginationMeta{
			CurrentPage: currentPage,
			PerPage:     perPage,
			Total:       total,
			TotalPages:  totalPages,
			Count:       count,
		},
	}
}

// PaginationMeta - мета-информация о пагинации
type PaginationMeta struct {
	// Текущая страница
	CurrentPage int `json:"current_page"`
	// Количество элементов на странице
	PerPage int `json:"per_page"`
	// Общее количество элементов
	Total int `json:"total"`
	// Общее количество страниц
	TotalPages int `json:"total_pages"`
	// Количество элементов на текущей странице
	Count int `json:"count"`
}
