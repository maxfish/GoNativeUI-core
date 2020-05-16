![Go](https://github.com/maxfish/GoNativeUI-core/workflows/Go/badge.svg)
[![Maintainability](https://api.codeclimate.com/v1/badges/1f9d6d2be93564c60962/maintainability)](https://codeclimate.com/github/maxfish/GoNativeUI-core/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/1f9d6d2be93564c60962/test_coverage)](https://codeclimate.com/github/maxfish/GoNativeUI-core/test_coverage)

*Heavily WIP / Not yet ready for use*

# GoNativeUI
Pure Go, no dependecies, cross-platform, barebone, retained-mode GUI.
The primary goal of the library is to provide a simple-to-use GUI for cross-platform desktop app/tools.

These are not goals of this project:
* Create a GUI which can be set up with a couple of lines of code
* Target Javascript, Electron, mobile phones, ...
* Provide an extensive set of Layouts or a complete widgets set similar to: QT, GTK, ...

## Core
The core repository contains the logic code of the GUI and it does **not** contain any rendering code.

The only assumption made is that the embedded font (Roboto Regular) will be rendered using the embedded [Distance Field](https://github.com/libgdx/libgdx/wiki/Distance-field-fonts) PNG. By implementing an interface is possible to replace the base font and to use different data and rendering techniques.

## Widgets currently available
* Label
* Button, ToggleButton, Checkbox

## Layout
The _Box_ container allows to create complex layouts by applying very simple rules.

The layout is inspired by the [The Box Model](https://developer.mozilla.org/en-US/docs/Archive/Mozilla/XUL/Tutorial/The_Box_Model) from the, now defunct, [Mozilla XUL](https://developer.mozilla.org/en-US/docs/Archive/Mozilla/XUL).


## To do
- [ ] Button action callbacks
- [ ] Widget content alignment
- [ ] Widget alignment within containers (only _stretch_ is supported right now)
- [ ] Containers padding
- [ ] Radio buttons
- [ ] Input fields
- [ ] Grid container / Property grid
