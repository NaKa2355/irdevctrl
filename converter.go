package irdev

type DataConverter interface {
	ConvertToRawData() (RawData, error)
}
