package dynamicgen

type Base struct {
	Wh                   [2]int        `json:"wh"`
	DoNotGenerateContext bool          `json:"doNotGenerateContext"`
	Entities             []interface{} `json:"entities"`
}

type ImageEntity struct {
	Type string
	Xy   [2]int
  Wh [2]int
	Url  string
}

type TextEntity struct {
	Type      string
	Content   string
	Xy        [2]float64
	MaxWidth  float64
	MaxHeight float64
	Spacing   float64
	Color     string
	Font      FontEntity
}

type FontEntity struct {
	Face string
	Size float64
}
