package domain

import "gorm.io/gorm"

type Pageable struct {
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
}

// Pagination
// 分页封装
func Pagination(pageable Pageable) func(db *gorm.DB) *gorm.DB {
	var page = pageable.Page
	var pageSize = pageable.PageSize
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		maxSize := 10000
		switch {
		case pageSize > maxSize:
			pageSize = maxSize
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

const maxPageSize = 10000

func (p Pageable) Offset() int {
	page, size := p.Page, p.PageSize
	if page == 0 {
		page = 1
	}
	switch {
	case size > maxPageSize:
		size = maxPageSize
	case size <= 0:
		size = 10
	}
	return (page - 1) * size
}
