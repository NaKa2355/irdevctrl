package irdev

type Converter interface {
	ConvertToRawData() (RawData, error)
}
