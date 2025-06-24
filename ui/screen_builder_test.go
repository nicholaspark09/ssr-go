package ui

import (
	"encoding/json"
	"main/model"
	"main/utils"
	"strings"
	"testing"
)

func TestBasicComponents(t *testing.T) {
	tests := []struct {
		name      string
		component model.ComponentNode
		validate  func(t *testing.T, json string)
	}{
		{
			name:      "Simple Text Component",
			component: Text("Hello World"),
			validate: func(t *testing.T, json string) {
				if !strings.Contains(json, `"type":"text"`) {
					t.Error("Missing text component type")
				}
				if !strings.Contains(json, `"Hello World"`) {
					t.Error("Missing text content")
				}
			},
		},
		{
			name:      "Styled Text Component",
			component: StyledText("Header Text", "headline1"),
			validate: func(t *testing.T, json string) {
				if !strings.Contains(json, `"style":"headline1"`) {
					t.Error("Missing text style")
				}
			},
		},
		{
			name:      "Button with Action",
			component: Button("Click Me", NavigationAction("home")),
			validate: func(t *testing.T, json string) {
				if !strings.Contains(json, `"type":"button"`) {
					t.Error("Missing button type")
				}
				if !strings.Contains(json, `"onClick"`) {
					t.Error("Missing onClick action")
				}
				if !strings.Contains(json, `"navigation"`) {
					t.Error("Missing navigation action type")
				}
			},
		},
		{
			name:      "Image Component",
			component: Image("https://example.com/image.jpg"),
			validate: func(t *testing.T, json string) {
				if !strings.Contains(json, `"type":"image"`) {
					t.Error("Missing image type")
				}
				if !strings.Contains(json, `"https://example.com/image.jpg"`) {
					t.Error("Missing image URL")
				}
			},
		},
		{
			name:      "Circle Image",
			component: CircleImage("https://example.com/avatar.jpg", 64),
			validate: func(t *testing.T, json string) {
				if !strings.Contains(json, `"shape":"circle"`) {
					t.Error("Missing circle shape")
				}
				if !strings.Contains(json, `"size":64`) {
					t.Error("Missing size property")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, err := json.Marshal(tt.component)
			if err != nil {
				t.Fatalf("Failed to marshal component: %v", err)
			}

			jsonStr := string(bytes)
			utils.PrettyPrintJSON(t, jsonStr, tt.name)

			if tt.validate != nil {
				tt.validate(t, jsonStr)
			}
		})
	}
}

func TestLayoutComponents(t *testing.T) {
	tests := []struct {
		name      string
		component model.ComponentNode
		validate  func(t *testing.T, json string)
	}{
		{
			name: "Column with Children",
			component: Column(
				Text("First item"),
				Text("Second item"),
				Button("Action", NavigationAction("next")),
			),
			validate: func(t *testing.T, json string) {
				if !strings.Contains(json, `"type":"column"`) {
					t.Error("Missing column type")
				}
				if !strings.Contains(json, `"children"`) {
					t.Error("Missing children array")
				}
			},
		},
		{
			name: "Row with Weighted Children",
			component: NewComponent("row").
				WithChildren(
					NewComponent("text").
						WithProperty("text", "Left").
						WithModifier(model.ModifierConfig{Weight: Float32Ptr(1)}).
						Build(),
					NewComponent("text").
						WithProperty("text", "Right").
						WithModifier(model.ModifierConfig{Weight: Float32Ptr(2)}).
						Build(),
				).
				Build(),
			validate: func(t *testing.T, json string) {
				if !strings.Contains(json, `"weight":1`) {
					t.Error("Missing weight property")
				}
				if !strings.Contains(json, `"weight":2`) {
					t.Error("Missing second weight property")
				}
			},
		},
		{
			name: "Card with Elevation",
			component: CardWithElevation(4.0,
				Text("Card content"),
				Spacer(16),
			),
			validate: func(t *testing.T, json string) {
				if !strings.Contains(json, `"type":"card"`) {
					t.Error("Missing card type")
				}
				if !strings.Contains(json, `"elevation":4`) {
					t.Error("Missing elevation property")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, err := json.Marshal(tt.component)
			if err != nil {
				t.Fatalf("Failed to marshal component: %v", err)
			}

			jsonStr := string(bytes)
			utils.PrettyPrintJSON(t, jsonStr, tt.name)

			if tt.validate != nil {
				tt.validate(t, jsonStr)
			}
		})
	}
}

func TestEnhancedLazyColumn(t *testing.T) {
	// Test data with mixed content types
	items := []map[string]interface{}{
		ItemWithTemplate(
			map[string]interface{}{
				"title":    "Custom Header",
				"subtitle": "Dynamic content",
			},
			Card(
				Text("{{title}}"),
				Text("{{subtitle}}"),
			),
		),
		ChartBarItem("Sample Chart", "Test data", []model.ChartDataPoint{
			{Label: "A", Value: 10},
			{Label: "B", Value: 20},
		}),
		SpacerItem(24),
		ItemWithComponentType(
			map[string]interface{}{
				"message": "Custom component type",
			},
			"custom_type",
		),
	}

	template := model.ItemTemplate{
		Type:   "default",
		Layout: Text("{{title}}"),
	}

	component := EnhancedLazyColumn(
		EnhancedStaticDataSource(items),
		template,
	)

	bytes, err := json.Marshal(component)
	if err != nil {
		t.Fatalf("Failed to marshal enhanced lazy column: %v", err)
	}

	jsonStr := string(bytes)
	utils.PrettyPrintJSON(t, jsonStr, "Enhanced Lazy Column")

	// Validate structure
	if !strings.Contains(jsonStr, `"type":"enhanced_lazy_column"`) {
		t.Error("Missing enhanced_lazy_column type")
	}
	if !strings.Contains(jsonStr, `"dataSource"`) {
		t.Error("Missing dataSource")
	}
	if !strings.Contains(jsonStr, `"itemTemplate"`) {
		t.Error("Missing itemTemplate")
	}
	if !strings.Contains(jsonStr, `"component_type":"chart_bar"`) {
		t.Error("Missing chart_bar component type")
	}
	if !strings.Contains(jsonStr, `"component_type":"spacer"`) {
		t.Error("Missing spacer component type")
	}
}

func TestCompleteScreen(t *testing.T) {
	// Build a complete screen similar to Emma's dashboard
	chartData := []model.ChartDataPoint{
		{Label: "Mon", Value: 6, Color: StringPtr("#FF9F43")},
		{Label: "Tue", Value: 7, Color: StringPtr("#FF9F43")},
		{Label: "Wed", Value: 5, Color: StringPtr("#FF9F43")},
	}

	screen := NewScreen("test_dashboard", "Test Dashboard", "1.0").
		WithLayout(
			Column(
				CenteredTopAppBar("Test Analytics"),
				NewComponent("scrollable_column").
					WithModifier(model.ModifierConfig{
						FillMaxSize: BoolPtr(true),
						Padding:     IntPtr(16),
					}).
					WithChildren(
						StyledText("Analytics Overview", "headline2"),
						Card(
							BarChart("Sample Data", chartData),
						),
						Spacer(24),
						Row(
							NewComponent("card").
								WithModifier(model.ModifierConfig{Weight: Float32Ptr(1)}).
								WithChildren(Text("Left panel")).
								Build(),
							NewComponent("card").
								WithModifier(model.ModifierConfig{Weight: Float32Ptr(1)}).
								WithChildren(Text("Right panel")).
								Build(),
						),
					).
					Build(),
			),
		).
		WithTheme(model.ThemeConfig{
			PrimaryColor:    "#3B82F6",
			SecondaryColor:  "#10B981",
			BackgroundColor: "#F9FAFB",
			TextColor:       "#1F2937",
		})

	jsonStr, err := screen.ToPrettyJSON()
	if err != nil {
		t.Fatalf("Failed to generate screen JSON: %v", err)
	}

	utils.PrettyPrintJSON(t, jsonStr, "Complete Screen")

	expectedFields := []string{"version", "screen", "theme"}
	utils.ValidateJSONStructure(t, jsonStr, expectedFields)

	if !strings.Contains(jsonStr, `"id": "test_dashboard"`) {
		t.Error("Missing screen ID")
	}
	if !strings.Contains(jsonStr, `"title": "Test Dashboard"`) {
		t.Error("Missing screen title")
	}
	if !strings.Contains(jsonStr, `"layout"`) {
		t.Error("Missing layout")
	}
}
