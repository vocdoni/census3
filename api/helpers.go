package api

import (
	"fmt"
	"strconv"

	"go.vocdoni.io/dvote/httprouter"
)

// paginationFromCtx extracts from the request and returns the page size,
// the database page size, the current cursor and the direction of that cursor.
// The page size is the number of elements of the page, the database page size
// is the number of elements of the page plus one, to get the previous and next
// cursors. The cursor and the direction are extracted from the request. If
// both cursors are provided, it returns an error.
func paginationFromCtx(ctx *httprouter.HTTPContext) (int32, int32, string, bool, error) {
	// define the initial page size by increasing the probvided value to get
	// the previous and next cursors
	pageSize := defaultPageSize
	dbPageSize := defaultPageSize + 1
	// if the page size is provided, use the provided value instead, increasing
	// it by 2 to get the previous and next cursors
	if strPageSize := ctx.Request.URL.Query().Get("pageSize"); strPageSize != "" {
		if intPageSize, err := strconv.Atoi(strPageSize); err == nil && intPageSize > 1 {
			pageSize = int32(intPageSize)
			dbPageSize = int32(intPageSize) + 1
		}
	}
	// get posible previous and next cursors
	prevCursor := ctx.Request.URL.Query().Get("prevCursor")
	nextCursor := ctx.Request.URL.Query().Get("nextCursor")
	// if both cursors are provided, return an error
	if prevCursor != "" && nextCursor != "" {
		return 0, 0, "", false, fmt.Errorf("both cursors provided, next and previous")
	}
	// by default go forward, if the previous cursor is provided, go backwards
	goForward := prevCursor == ""
	cursor := nextCursor
	if nextCursor == "" {
		cursor = prevCursor
	}
	// return the page size, the cursor and the direction
	return pageSize, dbPageSize, cursor, goForward, nil
}

// paginationToRequest returns the rows of the page, the next cursor and the
// previous cursor. If the rows size is the same as the database page size, the
// last element of the page is the next cursor, so it has to be removed from the
// rows. If the current page is the first one, the previous cursor is nil, and
// the rows are empty, because the first element is the cursor and there is
// include it in the following page. It uses generics to support any type of
// rows. The cursors will alwways be strings.
func paginationToRequest[T any](rows []T, dbPageSize int32, cursor string, goForward bool) ([]T, *T, *T) {
	// if the rows are empty there is no results or next and previous cursor
	if len(rows) == 0 {
		return rows, nil, nil
	}
	// by default, the next cursor is the last element of the page, and the
	// previous cursor is the first element of the page
	nextCursor := &rows[len(rows)-1]
	prevCursor := &rows[0]
	// if the length of the rows is less than the maximun page size, there is
	// no next cursor, and all the rows are part of the page
	if len(rows) < int(dbPageSize)-1 {
		if len(rows) > 1 {
			return rows, nil, prevCursor
		}
		// if the rows has just one element, there is no next or previous cursor, so
		// if the direction is forward, the next cursor is nil, and if the direction
		// is backwards, the previous cursor is nil and the rows are empty, because
		// the first element is the cursor and there is include it in the following
		// page.
		if len(rows) == 1 {
			if goForward {
				nextCursor = nil
			} else {
				prevCursor = nil
				rows = []T{}
			}
		}
	}
	// if the page size is the same as the database page size, the last element
	// of the page is the next cursor, so it has to be removed from the rows
	if len(rows) == int(dbPageSize) {
		rows = rows[:len(rows)-1]
	}
	return rows, nextCursor, prevCursor
}
