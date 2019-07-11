// Copyright 2016 - 2019 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to
// and read from XLSX files. Support reads and writes XLSX file generated by
// Microsoft Excel™ 2007 and later. Support save file without losing original
// charts of XLSX. This library needs Go version 1.8 or later.

package excelize

import "fmt"

// SheetViewOption is an option of a view of a worksheet. See SetSheetViewOptions().
type SheetViewOption interface {
	setSheetViewOption(view *xlsxSheetView)
}

// SheetViewOptionPtr is a writable SheetViewOption. See GetSheetViewOptions().
type SheetViewOptionPtr interface {
	SheetViewOption
	getSheetViewOption(view *xlsxSheetView)
}

type (
	// DefaultGridColor specified a flag indicating that the consuming
	// application should use the default grid lines color (system dependent).
	// Overrides any color specified in colorId.
	DefaultGridColor bool
	// RightToLeft specified a flag indicating whether the sheet is in 'right to
	// left' display mode. When in this mode, Column A is on the far right,
	// Column B ;is one column left of Column A, and so on. Also, information in
	// cells is displayed in the Right to Left format.
	RightToLeft bool
	// ShowFormulas specified a flag indicating whether this sheet should display
	// formulas.
	ShowFormulas bool
	// ShowGridLines specified a flag indicating whether this sheet should
	// display gridlines.
	ShowGridLines bool
	// ShowRowColHeaders specified a flag indicating whether the sheet should
	// display row and column headings.
	ShowRowColHeaders bool
	// ZoomScale specified a window zoom magnification for current view
	// representing percent values. This attribute is restricted to values
	// ranging from 10 to 400. Horizontal & Vertical scale together.
	ZoomScale float64
	// TopLeftCell specified a location of the top left visible cell Location of
	// the top left visible cell in the bottom right pane (when in Left-to-Right
	// mode).
	TopLeftCell string
	/* TODO
	// ShowWhiteSpace specified flag indicating whether page layout view shall
	// display margins. False means do not display left, right, top (header), and
	// bottom (footer) margins (even when there is data in the header or footer).
	ShowWhiteSpace bool
	// ShowZeros is a SheetViewOption.
	ShowZeros bool
	// WindowProtection is a SheetViewOption.
	WindowProtection bool
	*/
)

// Defaults for each option are described in XML schema for CT_SheetView

func (o TopLeftCell) setSheetViewOption(view *xlsxSheetView) {
	view.TopLeftCell = string(o)
}

func (o *TopLeftCell) getSheetViewOption(view *xlsxSheetView) {
	*o = TopLeftCell(string(view.TopLeftCell))
}

func (o DefaultGridColor) setSheetViewOption(view *xlsxSheetView) {
	view.DefaultGridColor = boolPtr(bool(o))
}

func (o *DefaultGridColor) getSheetViewOption(view *xlsxSheetView) {
	*o = DefaultGridColor(defaultTrue(view.DefaultGridColor)) // Excel default: true
}

func (o RightToLeft) setSheetViewOption(view *xlsxSheetView) {
	view.RightToLeft = bool(o) // Excel default: false
}

func (o *RightToLeft) getSheetViewOption(view *xlsxSheetView) {
	*o = RightToLeft(view.RightToLeft)
}

func (o ShowFormulas) setSheetViewOption(view *xlsxSheetView) {
	view.ShowFormulas = bool(o) // Excel default: false
}

func (o *ShowFormulas) getSheetViewOption(view *xlsxSheetView) {
	*o = ShowFormulas(view.ShowFormulas) // Excel default: false
}

func (o ShowGridLines) setSheetViewOption(view *xlsxSheetView) {
	view.ShowGridLines = boolPtr(bool(o))
}

func (o *ShowGridLines) getSheetViewOption(view *xlsxSheetView) {
	*o = ShowGridLines(defaultTrue(view.ShowGridLines)) // Excel default: true
}

func (o ShowRowColHeaders) setSheetViewOption(view *xlsxSheetView) {
	view.ShowRowColHeaders = boolPtr(bool(o))
}

func (o *ShowRowColHeaders) getSheetViewOption(view *xlsxSheetView) {
	*o = ShowRowColHeaders(defaultTrue(view.ShowRowColHeaders)) // Excel default: true
}

func (o ZoomScale) setSheetViewOption(view *xlsxSheetView) {
	// This attribute is restricted to values ranging from 10 to 400.
	if float64(o) >= 10 && float64(o) <= 400 {
		view.ZoomScale = float64(o)
	}
}

func (o *ZoomScale) getSheetViewOption(view *xlsxSheetView) {
	*o = ZoomScale(view.ZoomScale)
}

// getSheetView returns the SheetView object
func (f *File) getSheetView(sheetName string, viewIndex int) (*xlsxSheetView, error) {
	xlsx, err := f.workSheetReader(sheetName)
	if err != nil {
		return nil, err
	}
	if viewIndex < 0 {
		if viewIndex < -len(xlsx.SheetViews.SheetView) {
			return nil, fmt.Errorf("view index %d out of range", viewIndex)
		}
		viewIndex = len(xlsx.SheetViews.SheetView) + viewIndex
	} else if viewIndex >= len(xlsx.SheetViews.SheetView) {
		return nil, fmt.Errorf("view index %d out of range", viewIndex)
	}

	return &(xlsx.SheetViews.SheetView[viewIndex]), err
}

// SetSheetViewOptions sets sheet view options.
// The viewIndex may be negative and if so is counted backward (-1 is the last view).
//
// Available options:
//    DefaultGridColor(bool)
//    RightToLeft(bool)
//    ShowFormulas(bool)
//    ShowGridLines(bool)
//    ShowRowColHeaders(bool)
//    ZoomScale(float64)
//    TopLeftCell(string)
// Example:
//    err = f.SetSheetViewOptions("Sheet1", -1, ShowGridLines(false))
func (f *File) SetSheetViewOptions(name string, viewIndex int, opts ...SheetViewOption) error {
	view, err := f.getSheetView(name, viewIndex)
	if err != nil {
		return err
	}

	for _, opt := range opts {
		opt.setSheetViewOption(view)
	}
	return nil
}

// GetSheetViewOptions gets the value of sheet view options.
// The viewIndex may be negative and if so is counted backward (-1 is the last view).
//
// Available options:
//    DefaultGridColor(bool)
//    RightToLeft(bool)
//    ShowFormulas(bool)
//    ShowGridLines(bool)
//    ShowRowColHeaders(bool)
//    ZoomScale(float64)
//    TopLeftCell(string)
// Example:
//    var showGridLines excelize.ShowGridLines
//    err = f.GetSheetViewOptions("Sheet1", -1, &showGridLines)
func (f *File) GetSheetViewOptions(name string, viewIndex int, opts ...SheetViewOptionPtr) error {
	view, err := f.getSheetView(name, viewIndex)
	if err != nil {
		return err
	}

	for _, opt := range opts {
		opt.getSheetViewOption(view)
	}
	return nil
}
