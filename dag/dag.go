package dag

const TREE = 0
const LIST = 1
const BLOB = 2

// dir
type Tree struct {
	// tree can link to 'tree', 'list' or 'blob'
	data []int
	
	links []Link
}

// file
type List struct {
	// list can link to 'blob' or 'list'(hard-link or soft-link).
	data []int
	
	links []Link
}

// an unit of data. When file is small, it can also represent file.
type Blob struct {
	data []byte
}

// edge of dag
type Link struct {
	hash []byte

	name string

	size int
}