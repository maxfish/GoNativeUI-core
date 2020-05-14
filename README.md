![Go](https://github.com/maxfish/GoNativeUI-core/workflows/Go/badge.svg)
<a href="https://codeclimate.com/github/codeclimate/codeclimate/maintainability"><img src="https://api.codeclimate.com/v1/badges/a99a88d28ad37a79dbf6/maintainability" /></a> [![Test Coverage](https://api.codeclimate.com/v1/badges/a99a88d28ad37a79dbf6/test_coverage)](https://codeclimate.com/github/codeclimate/codeclimate/test_coverage)

*Heavily WIP / Not yet ready for use*

# GoNativeUI
Pure Go, no dependecies, cross-platform, barebone, retained-mode GUI.
The primary goal of the library is to provide a simple-to-use GUI for cross-platform desktop app/tools.

These are not goals of this project:
* Create a GUI which can be set up with a couple of lines of code
* Target Javascript, Electron, mobile phones, ...
* Provide an extensive set of Layouts or a complete widgets set similar to: QT, GTK, ...

## Core
The core repository contains the logic code of the GUI and it does **not** contain rendering code. A backend for GLFW/OpenGL is WIP.

Right now, the only assumption made is that the default font (Roboto Regular) will be rendered using the embedded [Distance Field](https://github.com/libgdx/libgdx/wiki/Distance-field-fonts) PNG.

## Widgets currently available
* Label
* Button
