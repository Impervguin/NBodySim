package buffers

import "image/color"

type ScreenFunction func(x, y, w, h int) color.Color
