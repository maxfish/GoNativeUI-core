![Go](https://github.com/maxfish/GoNativeUI-core/workflows/Go/badge.svg)
[![Maintainability](https://api.codeclimate.com/v1/badges/1f9d6d2be93564c60962/maintainability)](https://codeclimate.com/github/maxfish/GoNativeUI-core/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/1f9d6d2be93564c60962/test_coverage)](https://codeclimate.com/github/maxfish/GoNativeUI-core/test_coverage)

*Heavily WIP / Not yet ready for use*

# GoNativeUI
Pure Go, no dependencies, cross-platform, bare bone, retained-mode GUI.
The primary goal of the library is to provide a simple-to-use GUI for cross-platform
 desktop app/tools.

These are not goals of this project:
* Create a GUI which can be set up with a couple of lines of code
* Target Javascript, Electron, mobile phones, ...
* Provide an extensive set of Layouts, or a complete widgets set similar to QT, GTK, ...

## Core
The core repository contains the logic code of the GUI and it does not
 contain **any** rendering code.

## Widgets
- [x] Label
- [x] Button
- [x] Toggle button
- [x] Checkbox
- [ ] Radio button
- [X] Input fields

## Layout
The only container available is _Box_. It allows creating complex layouts by applying
 very simple rules.

The layout is inspired by the [The Box Model](https://developer.mozilla.org/en-US/docs/Archive/Mozilla/XUL/Tutorial/The_Box_Model)
 from [Mozilla XUL](https://developer.mozilla.org/en-US/docs/Archive/Mozilla/XUL) (now defunct).

## Other ToDo
- [X] Button action callbacks
- [X] Widget content alignment
- [ ] Widget alignment within containers (only _stretch_ is supported right now)
- [x] Containers padding
- [ ] Grid container / Property grid
