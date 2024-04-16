package entities

import (
	"github.com/KidPudel/snakegame_go/common"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Food struct {
	Position rl.Vector2
}

func (food *Food) Draw() {
	rl.DrawRectangleV(food.Position, rl.NewVector2(common.SnakeHeadSize, common.SnakeHeadSize), common.FoodColor)
}
