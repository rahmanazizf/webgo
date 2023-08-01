package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int16
	FloatMap  map[string]float32
	OtherData map[string]interface{}
	Massage   string
	Warning   string
	Error     string
}
