package model

type DataSource struct {
	Type       string                   `json:"type"` // "api", "static", "database"
	URL        *string                  `json:"url"`
	Method     *string                  `json:"method"`
	Headers    map[string]string        `json:"headers"`
	Items      []map[string]interface{} `json:"items"`
	Pagination *PaginationConfig        `json:"pagination"`
}

type PaginationConfig struct {
	Type        string  `json:"type"` // "offset", "cursor", "page"
	PageSize    int     `json:"pageSize"`
	OffsetParam *string `json:"offsetParam"`
	LimitParam  *string `json:"limitParam"`
	CursorParam *string `json:"cursorParam"`
	PageParam   *string `json:"pageParam"`
}

type ItemTemplate struct {
	Type    string                  `json:"type"`
	Layout  ComponentNode           `json:"layout"`
	Actions map[string]ActionConfig `json:"actions"`
}

type ActionConfig struct {
	Type        string            `json:"type"`
	Destination *string           `json:"destination"`
	Params      map[string]string `json:"params"`
}

type LoadingState struct {
	Type  string `json:"type"`
	Count *int   `json:"count"`
}

type ModifierConfig struct {
	Padding       *int            `json:"padding"`
	FillMaxSize   *bool           `json:"fillMaxSize"`
	FillMaxWidth  *bool           `json:"fillMaxWidth"`
	Width         *int            `json:"width"`
	Height        *int            `json:"height"`
	Weight        *float32        `json:"weight"`
	PaddingStart  *int            `json:"paddingStart"`
	PaddingTop    *int            `json:"paddingTop"`
	PaddingEnd    *int            `json:"paddingEnd"`
	PaddingBottom *int            `json:"paddingBottom"`
	Gradient      *GradientConfig `json:"gradient"`
}

type GradientConfig struct {
	Type   string   `json:"type"`
	Colors []string `json:"colors"`
	Angle  *float32 `json:"angle"`
	StartX *float32 `json:"startX"`
	StartY *float32 `json:"startY"`
	EndX   *float32 `json:"endX"`
	EndY   *float32 `json:"endY"`
	Radius *float32 `json:"radius"`
}

type ComponentNode struct {
	Type            string                  `json:"type"`
	ID              *string                 `json:"id"`
	Properties      map[string]interface{}  `json:"properties"`
	Modifier        *ModifierConfig         `json:"modifier"`
	Children        []ComponentNode         `json:"children"`
	Actions         map[string]ActionConfig `json:"actions"`
	Arrangement     *string                 `json:"arrangement"`
	Columns         *int                    `json:"columns"`
	DataSource      *DataSource             `json:"dataSource"`
	ItemTemplate    *ItemTemplate           `json:"itemTemplate"`
	LoadingTemplate *LoadingState           `json:"loadingTemplate"`
	EmptyTemplate   *ComponentNode          `json:"emptyTemplate"`
	ErrorTemplate   *ComponentNode          `json:"errorTemplate"`
}

type ScreenLayout struct {
	ID     string        `json:"id"`
	Title  string        `json:"title"`
	Layout ComponentNode `json:"layout"`
}

type ThemeConfig struct {
	PrimaryColor    string `json:"primaryColor"`
	SecondaryColor  string `json:"secondaryColor"`
	BackgroundColor string `json:"backgroundColor"`
	TextColor       string `json:"textColor"`
}

type DataConfig struct {
	APIEndpoints map[string]string      `json:"apiEndpoints"`
	StaticData   map[string]interface{} `json:"staticData"`
}

type ComponentScreen struct {
	Version string       `json:"version"`
	Screen  ScreenLayout `json:"screen"`
	Theme   *ThemeConfig `json:"theme"`
	Data    *DataConfig  `json:"data,omitempty"`
}

type ChartDataPoint struct {
	Label    string                 `json:"label"`
	Value    float64                `json:"value"`
	Color    *string                `json:"color"`
	Metadata map[string]interface{} `json:"metadata"`
}

type ChartSeries struct {
	Name  string           `json:"name"`
	Data  []ChartDataPoint `json:"data"`
	Color *string          `json:"color"`
}

type ChartConfig struct {
	Title      *string  `json:"title"`
	Subtitle   *string  `json:"subtitle"`
	ShowLegend *bool    `json:"showLegend"`
	ShowGrid   *bool    `json:"showGrid"`
	ShowLabels *bool    `json:"showLabels"`
	ShowValues *bool    `json:"showValues"`
	Animated   *bool    `json:"animated"`
	Colors     []string `json:"colors"`
	Height     *int     `json:"height"`
	Width      *int     `json:"width"`
}
