package ui

import "github.com/nicholaspark09/ssr-go/model"

type ComponentBuilder struct {
	component model.ComponentNode
}

func NewComponent(componentType string) *ComponentBuilder {
	return &ComponentBuilder{
		component: model.ComponentNode{
			Type:       componentType,
			Properties: make(map[string]interface{}),
		},
	}
}

func (cb *ComponentBuilder) WithID(id string) *ComponentBuilder {
	cb.component.ID = &id
	return cb
}

func (cb *ComponentBuilder) WithProperty(key string, value interface{}) *ComponentBuilder {
	if cb.component.Properties == nil {
		cb.component.Properties = make(map[string]interface{})
	}
	cb.component.Properties[key] = value
	return cb
}

func (cb *ComponentBuilder) WithModifier(modifier model.ModifierConfig) *ComponentBuilder {
	cb.component.Modifier = &modifier
	return cb
}

func (cb *ComponentBuilder) WithChildren(children ...model.ComponentNode) *ComponentBuilder {
	cb.component.Children = append(cb.component.Children, children...)
	return cb
}

func (cb *ComponentBuilder) WithAction(actionName string, action model.ActionConfig) *ComponentBuilder {
	if cb.component.Actions == nil {
		cb.component.Actions = make(map[string]model.ActionConfig)
	}
	cb.component.Actions[actionName] = action
	return cb
}

func (cb *ComponentBuilder) WithDataSource(dataSource model.DataSource) *ComponentBuilder {
	cb.component.DataSource = &dataSource
	return cb
}

func (cb *ComponentBuilder) WithItemTemplate(template model.ItemTemplate) *ComponentBuilder {
	cb.component.ItemTemplate = &template
	return cb
}

func (cb *ComponentBuilder) Build() model.ComponentNode {
	return cb.component
}

func EnhancedLazyColumn(dataSource model.DataSource, itemTemplate model.ItemTemplate) model.ComponentNode {
	return NewComponent("enhanced_lazy_column").
		WithDataSource(dataSource).
		WithItemTemplate(itemTemplate).
		Build()
}

func EnhancedStaticDataSource(items []map[string]interface{}) model.DataSource {
	return model.DataSource{
		Type:  "static",
		Items: items,
	}
}

func ItemWithTemplate(data map[string]interface{}, template model.ComponentNode) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range data {
		result[k] = v
	}
	result["template"] = template
	return result
}

func ItemWithComponentType(data map[string]interface{}, componentType string) map[string]interface{} {
	result := make(map[string]interface{})
	// Copy all data
	for k, v := range data {
		result[k] = v
	}
	// Add the component type
	result["component_type"] = componentType
	return result
}

func Text(text string) model.ComponentNode {
	return NewComponent("text").
		WithProperty("text", text).
		Build()
}

func StyledText(text, style string) model.ComponentNode {
	return NewComponent("text").
		WithProperty("text", text).
		WithProperty("style", style).
		Build()
}

func Button(text string, onClick model.ActionConfig) model.ComponentNode {
	return NewComponent("button").
		WithProperty("text", text).
		WithAction("onClick", onClick).
		Build()
}

func Image(url string) model.ComponentNode {
	return NewComponent("image").
		WithProperty("url", url).
		Build()
}

func CircleImage(url string, size int) model.ComponentNode {
	return NewComponent("image").
		WithProperty("url", url).
		WithProperty("shape", "circle").
		WithProperty("size", size).
		Build()
}

func Column(children ...model.ComponentNode) model.ComponentNode {
	return NewComponent("column").
		WithChildren(children...).
		Build()
}

func ScrollableColumn(children ...model.ComponentNode) model.ComponentNode {
	return NewComponent("scrollable_column").
		WithChildren(children...).
		Build()
}

func Row(children ...model.ComponentNode) model.ComponentNode {
	return NewComponent("row").
		WithChildren(children...).
		Build()
}

func Card(children ...model.ComponentNode) model.ComponentNode {
	return NewComponent("card").
		WithChildren(children...).
		Build()
}

func CardWithElevation(elevation float32, children ...model.ComponentNode) model.ComponentNode {
	return NewComponent("card").
		WithProperty("elevation", elevation).
		WithChildren(children...).
		Build()
}

func Spacer(height int) model.ComponentNode {
	return NewComponent("spacer").
		WithProperty("height", height).
		Build()
}

func SpacerItem(height int) map[string]interface{} {
	return ItemWithComponentType(map[string]interface{}{
		"height": height,
	}, "spacer")
}

func TopAppBar(title string) model.ComponentNode {
	return NewComponent("top_app_bar").
		WithProperty("title", title).
		Build()
}

func CenteredTopAppBar(title string) model.ComponentNode {
	return NewComponent("top_app_bar").
		WithProperty("title", title).
		WithProperty("centerTitle", true).
		Build()
}

func BarChart(title string, data []model.ChartDataPoint) model.ComponentNode {
	return NewComponent("bar_chart").
		WithProperty("title", title).
		WithProperty("data", data).
		WithProperty("showLegend", true).
		WithProperty("showGrid", true).
		WithProperty("showValues", true).
		Build()
}

func LineChart(title string, series []model.ChartSeries) model.ComponentNode {
	return NewComponent("line_chart").
		WithProperty("title", title).
		WithProperty("series", series).
		WithProperty("showLegend", true).
		WithProperty("showGrid", true).
		Build()
}

func PieChart(title string, data []model.ChartDataPoint) model.ComponentNode {
	return NewComponent("pie_chart").
		WithProperty("title", title).
		WithProperty("data", data).
		WithProperty("showLegend", true).
		WithProperty("showValues", true).
		Build()
}

func RadarChart(title string, data []model.ChartDataPoint) model.ComponentNode {
	return NewComponent("radar_chart").
		WithProperty("title", title).
		WithProperty("data", data).
		WithProperty("showLabels", true).
		Build()
}

func ChartBarItem(title, subtitle string, data []model.ChartDataPoint) map[string]interface{} {
	return ItemWithComponentType(map[string]interface{}{
		"title":    title,
		"subtitle": subtitle,
		"data":     data,
	}, "chart_bar")
}

func ChartLineItem(title, subtitle string, series []model.ChartSeries) map[string]interface{} {
	return ItemWithComponentType(map[string]interface{}{
		"title":    title,
		"subtitle": subtitle,
		"series":   series,
	}, "chart_line")
}

func ChartPieItem(title, subtitle string, data []model.ChartDataPoint) map[string]interface{} {
	return ItemWithComponentType(map[string]interface{}{
		"title":    title,
		"subtitle": subtitle,
		"data":     data,
	}, "chart_pie")
}

func ChartRadarItem(title, subtitle string, data []model.ChartDataPoint) map[string]interface{} {
	return ItemWithComponentType(map[string]interface{}{
		"title":    title,
		"subtitle": subtitle,
		"data":     data,
	}, "chart_radar")
}

func StaticDataSource(items []map[string]interface{}) model.DataSource {
	return model.DataSource{
		Type:  "static",
		Items: items,
	}
}

func APIDataSource(url, method string) model.DataSource {
	return model.DataSource{
		Type:   "api",
		URL:    &url,
		Method: &method,
	}
}

func APIDataSourceWithPagination(url, method string, pageSize int) model.DataSource {
	return model.DataSource{
		Type:   "api",
		URL:    &url,
		Method: &method,
		Pagination: &model.PaginationConfig{
			Type:     "page",
			PageSize: pageSize,
		},
	}
}

func LazyColumn(dataSource model.DataSource, itemTemplate model.ItemTemplate) model.ComponentNode {
	return NewComponent("lazy_column").
		WithDataSource(dataSource).
		WithItemTemplate(itemTemplate).
		Build()
}

func LazyRow(dataSource model.DataSource, itemTemplate model.ItemTemplate) model.ComponentNode {
	return NewComponent("lazy_row").
		WithDataSource(dataSource).
		WithItemTemplate(itemTemplate).
		Build()
}

func NavigationAction(destination string) model.ActionConfig {
	return model.ActionConfig{
		Type:        "navigation",
		Destination: &destination,
	}
}

func NavigationActionWithParams(destination string, params map[string]string) model.ActionConfig {
	return model.ActionConfig{
		Type:        "navigation",
		Destination: &destination,
		Params:      params,
	}
}

func APICallAction() model.ActionConfig {
	return model.ActionConfig{
		Type: "api_call",
	}
}

func PaddingModifier(padding int) model.ModifierConfig {
	return model.ModifierConfig{
		Padding: &padding,
	}
}

func FillMaxWidthModifier() model.ModifierConfig {
	fillMaxWidth := true
	return model.ModifierConfig{
		FillMaxWidth: &fillMaxWidth,
	}
}

func FillMaxSizeModifier() model.ModifierConfig {
	fillMaxSize := true
	return model.ModifierConfig{
		FillMaxSize: &fillMaxSize,
	}
}

func SizeModifier(width, height int) model.ModifierConfig {
	return model.ModifierConfig{
		Width:  &width,
		Height: &height,
	}
}

func WeightModifier(weight float32) model.ModifierConfig {
	return model.ModifierConfig{
		Weight: &weight,
	}
}
