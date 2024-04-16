package common

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	SnakeHeadSize float32 = 20
)

var (
	ScreenWidth      int32 = 600
	ScreenHeight     int32 = 400
	SnakeColor             = rl.NewColor(182, 208, 148, 255)
	SnakeBorderColor       = rl.NewColor(159, 182, 127, 255)
	FoodColor              = rl.NewColor(106, 46, 53, 255)
	BackgroundColor        = rl.NewColor(225, 170, 125, 255)
)
