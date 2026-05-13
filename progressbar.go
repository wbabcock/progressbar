// Package progressbar provides a simple customizable terminal progress bar.
package progressbar

import (
	"fmt"

	"github.com/wbabcock/progressbar/colors"
	"github.com/wbabcock/progressbar/indicator"
)

// ProgressBar represents a terminal progress bar.
type ProgressBar struct {
	length    uint
	percent   uint
	indicator indicator.Indicator
	color     colors.Color
}

// NewProgressBar creates a new progress bar with default settings.
//
// Defaults:
//   - Length: 50
//   - Indicator: "#"
//   - Color: Reset
func NewProgressBar() *ProgressBar {
	return &ProgressBar{
		length:    50,
		percent:   0,
		indicator: indicator.Hash,
		color:     colors.Reset,
	}
}

// WithLength returns a copy of the progress bar with the given length.
func (p ProgressBar) WithLength(length uint) *ProgressBar {
	return &ProgressBar{
		length:    length,
		percent:   p.percent,
		indicator: p.indicator,
		color:     p.color,
	}
}

// WithIndicator returns a copy of the progress bar with the given indicator character.
func (p ProgressBar) WithIndicator(indicator indicator.Indicator) *ProgressBar {
	return &ProgressBar{
		length:    p.length,
		percent:   p.percent,
		indicator: indicator,
		color:     p.color,
	}
}

// WithColor returns a copy of the progress bar with the given color.
func (p ProgressBar) WithColor(color colors.Color) *ProgressBar {
	return &ProgressBar{
		length:    p.length,
		percent:   p.percent,
		indicator: p.indicator,
		color:     color,
	}
}

// SetLength updates the progress bar length.
func (p *ProgressBar) SetLength(length uint) {
	p.length = length
}

// SetIndicator updates the progress bar indicator character.
func (p *ProgressBar) SetIndicator(indicator indicator.Indicator) {
	p.indicator = indicator
}

// SetColor updates the progress bar color.
func (p *ProgressBar) SetColor(color colors.Color) {
	p.color = color
}

// Update redraws the progress bar using the provided percentage.
//
// The percentage should be between 0 and 100.
func (p *ProgressBar) Update(percent uint) {
	if percent > 100 {
		percent = 100
	}

	p.percent = percent // Set the progress bar percentage

	numChars := p.percent * p.length / 100
	emptyChars := p.length - numChars

	fmt.Printf("\r%s[", p.color)

	for range numChars {
		fmt.Printf("%s", p.indicator)
	}

	for range emptyChars {
		fmt.Print(" ")
	}

	fmt.Printf("] %d%%\033[0m", p.percent)
}
