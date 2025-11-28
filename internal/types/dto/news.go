package dto

// Web - Create/Update News
type NewsRequest struct {
	Title       string   `json:"title" validate:"required"`
	Excerpt     string   `json:"excerpt"`
	Content     string   `json:"content" validate:"required"`
	Thumbnail   string   `json:"thumbnail"`
	Category    string   `json:"category" validate:"required,oneof=announcement event program_update success_story tips regulation general"`
	IsPublished bool     `json:"is_published"`
	Tags        []string `json:"tags"`
}

// Web - News Response
type NewsResponse struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Slug        string   `json:"slug"`
	Excerpt     string   `json:"excerpt"`
	Content     string   `json:"content"`
	Thumbnail   string   `json:"thumbnail"`
	Category    string   `json:"category"`
	AuthorID    int      `json:"author_id"`
	AuthorName  string   `json:"author_name"`
	IsPublished bool     `json:"is_published"`
	PublishedAt *string  `json:"published_at"`
	ViewsCount  int      `json:"views_count"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	Tags        []string `json:"tags"`
}

// Web - News List Response
type NewsListResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Slug        string  `json:"slug"`
	Excerpt     string  `json:"excerpt"`
	Thumbnail   string  `json:"thumbnail"`
	Category    string  `json:"category"`
	AuthorName  string  `json:"author_name"`
	IsPublished bool    `json:"is_published"`
	PublishedAt *string `json:"published_at"`
	ViewsCount  int     `json:"views_count"`
	CreatedAt   string  `json:"created_at"`
}

// Mobile - News List Response
type NewsListMobile struct {
	ID         int     `json:"id"`
	Title      string  `json:"title"`
	Slug       string  `json:"slug"`
	Excerpt    string  `json:"excerpt"`
	Thumbnail  string  `json:"thumbnail"`
	Category   string  `json:"category"`
	AuthorName string  `json:"author_name"`
	ViewsCount int     `json:"views_count"`
	CreatedAt  string  `json:"created_at"`
}

// Mobile - News Detail Response
type NewsDetailMobile struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	Slug       string   `json:"slug"`
	Content    string   `json:"content"`
	Thumbnail  string   `json:"thumbnail"`
	Category   string   `json:"category"`
	AuthorName string   `json:"author_name"`
	ViewsCount int      `json:"views_count"`
	CreatedAt  string   `json:"created_at"`
	Tags       []string `json:"tags"`
}

// Query Parameters
type NewsQueryParams struct {
	Page       int    `query:"page"`
	Limit      int    `query:"limit"`
	Category   string `query:"category"`
	Search     string `query:"search"`
	Tag        string `query:"tag"`
	IsPublished *bool  `query:"is_published"`
}