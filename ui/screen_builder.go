package ui

import (
	"encoding/json"
	"fmt"
	"main/model"
)

type ScreenBuilder struct {
	screen model.ComponentScreen
}

func NewScreen(id, title string, version string) *ScreenBuilder {
	return &ScreenBuilder{
		screen: model.ComponentScreen{
			Version: version,
			Screen: model.ScreenLayout{
				ID:    id,
				Title: title,
			},
		},
	}
}

func (sb *ScreenBuilder) WithLayout(layout model.ComponentNode) *ScreenBuilder {
	sb.screen.Screen.Layout = layout
	return sb
}

func (sb *ScreenBuilder) WithTheme(theme model.ThemeConfig) *ScreenBuilder {
	sb.screen.Theme = &theme
	return sb
}

func (sb *ScreenBuilder) WithData(data model.DataConfig) *ScreenBuilder {
	sb.screen.Data = &data
	return sb
}

func (sb *ScreenBuilder) ToJSON() (string, error) {
	bytes, err := json.Marshal(sb.screen)
	if err != nil {
		return "", fmt.Errorf("failed to marshal screen to JSON: %w", err)
	}
	return string(bytes), nil
}

func (sb *ScreenBuilder) ToPrettyJSON() (string, error) {
	bytes, err := json.MarshalIndent(sb.screen, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal screen to JSON: %w", err)
	}
	return string(bytes), nil
}

func (sb *ScreenBuilder) Build() model.ComponentScreen {
	return sb.screen
}

func BoolPtr(b bool) *bool {
	return &b
}

func StringPtr(s string) *string {
	return &s
}

func IntPtr(i int) *int {
	return &i
}

func Float32Ptr(f float32) *float32 {
	return &f
}
