// Package viewmodel: baseclass for generic properties
package viewmodel

type ViewModel struct {
	Focused bool
	Dims    ViewModelDims
}

type ViewModelDims struct {
	Width  int
	Height int
}
