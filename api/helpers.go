package api

import (
	"fmt"
	"strconv"

	"go.vocdoni.io/dvote/httprouter"
	"go.vocdoni.io/dvote/log"
)

func paginationFromCtx(ctx *httprouter.HTTPContext) (int32, string, bool, error) {
	// define the initial page size by increasing the probvided value to get
	// the previous and next cursors
	pageSize := defaultPageSize + 2
	// if the page size is provided, use the provided value instead, increasing
	// it by 2 to get the previous and next cursors
	if strPageSize := ctx.Request.URL.Query().Get("pageSize"); strPageSize != "" {
		if intPageSize, err := strconv.Atoi(strPageSize); err == nil && intPageSize > 1 {
			pageSize = int32(intPageSize) + 2
		}
	}
	// get posible previous and next cursors
	prevCursor := ctx.Request.URL.Query().Get("prevCursor")
	nextCursor := ctx.Request.URL.Query().Get("nextCursor")
	// if both cursors are provided, return an error
	if prevCursor != "" && nextCursor != "" {
		return 0, "", false, fmt.Errorf("both cursors provided, next and previous")
	}
	// by default go forward, if the previous cursor is provided, go backwards
	goForward := prevCursor == ""
	cursor := nextCursor
	if nextCursor == "" {
		cursor = prevCursor
	}
	// return the page size, the cursor and the direction
	return pageSize, cursor, goForward, nil
}

func paginationToRequest[T any](rows []T, pageSize int32, cursor string, goForward bool) ([]T, *T, *T) {
	// define vars to get correct the rows for the current page
	var rowsStart, rowsEnd int = 0, len(rows)
	// by default, the next cursor is the last element of the slice, and the
	// previous cursor is the first element of the slice
	prevCursorIdx, nextCursorIdx := 0, len(rows)-1
	// This represetents the edge case of first or intermediate pages with or
	// without cursor. If the rows slice has the same size as the page size,
	// and a cursor is not provided, the next cursor is the second to last
	// element of the slice, no previous cursor is needed and the correct rows
	// are from the first element to the second to last element of the slice.
	// If the cursor is provided, the boundary rows will not be included in the
	// rows slice.
	log.Info(len(rows), pageSize)
	if len(rows) == int(pageSize) {
		if cursor == "" {
			nextCursorIdx = len(rows) - 2
			prevCursorIdx = -1
			rowsEnd = len(rows) - 2
		} else {
			rowsStart = 1
			rowsEnd = len(rows) - 1
		}
	} else if len(rows) > int(pageSize)-2 {
		// This represetents the edge case of last pages in both directions.
		// If the rows slice has more elements than the page size but less than
		// the page size subtracted by 2 and cursor is provided or direction is
		// forward, the correct rows are from the first element to the second to
		// last element of the slice. If no cursor is provided or direction is
		// backwards, the correct rows are from the second element to the last.
		if cursor != "" || goForward {
			rowsEnd = len(rows) - 1
		} else {
			rowsStart = 1
		}
	}
	// If the next cursor index is greater than 0, the next cursor is the
	// element in the next cursor index.
	var nextCursor *T
	if nextCursorIdx >= 0 {
		nextCursor = &rows[nextCursorIdx]
	}
	// If the previous cursor index is greater than 0, the previous cursor is
	// the element in the previous cursor index.
	var prevCursor *T
	if prevCursorIdx >= 0 {
		prevCursor = &rows[prevCursorIdx]
	}
	// If the next and previous cursor indexes are the same, and the direction
	// is forward, the next cursor is nil. If the direction is backwards, the
	// previous cursor is nil. This is to avoid returning the same cursor for
	// prev and next.
	if nextCursorIdx == prevCursorIdx {
		if goForward {
			nextCursor = nil
		} else {
			prevCursor = nil
		}
	}
	return rows[rowsStart:rowsEnd], nextCursor, prevCursor
}
