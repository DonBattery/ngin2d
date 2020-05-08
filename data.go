package ngin2d

// DataHolder is the interface of objects that can set get and reset an arbitrary data object
type DataHolder interface {
	GetData() interface{}
	SetData(interface{})
	ResetData()
}

// DataSet is the minimal implementation of DataSet
type DataSet struct {
	data interface{}
}

// GetData returns the DataSet's data object
func (ds *DataSet) GetData() interface{} {
	return ds.data
}

// SetData sets the DataSet's data object
func (ds *DataSet) SetData(data interface{}) {
	ds.data = data
}

// ResetData sets DataSet's data object to nil
func (ds *DataSet) ResetData() {
	ds.data = nil
}
