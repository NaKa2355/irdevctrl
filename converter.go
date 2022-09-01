package irdevctrl

type DataConverter interface {
	ConvertToRawData() (RawData, error)
}
