package viewmodel

func (model *ViewModel) SetSize(width int, height int) {
	model.Dims = ViewModelDims{
		Width:  width,
		Height: height,
	}
}

func (model *ViewModel) GetSize() ViewModelDims {
	return model.Dims
}

func (model *ViewModel) SetFocused(focused bool) bool {
	model.Focused = focused
	return model.Focused
}

func (model *ViewModel) IsFocused() bool {
	return model.Focused
}
