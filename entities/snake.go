package entities

import (
	"slices"

	"github.com/KidPudel/snakegame_go/common"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Snake struct {
	// head position
	HeadPosition rl.Vector2
	Direction    rl.Vector2

	TailPositions []rl.Vector2
	// time trigger to jump to the next cell, based on ellaped time
	EllapsedSinceLastJump float32

	DeadSound *rl.Sound
}

// separate to own functions, so main function wouldn't be a mess, do not overcomplecate it
func (snake *Snake) HandleInput() {
	// move by its size, so even
	if rl.IsKeyPressed(rl.KeyD) {
		snake.Direction = rl.NewVector2(common.SnakeHeadSize, 0)
	} else if rl.IsKeyPressed(rl.KeyA) {
		snake.Direction = rl.NewVector2(-common.SnakeHeadSize, 0)
	} else if rl.IsKeyPressed(rl.KeyW) {
		snake.Direction = rl.NewVector2(0, -common.SnakeHeadSize)
	} else if rl.IsKeyPressed(rl.KeyS) {
		snake.Direction = rl.NewVector2(0, common.SnakeHeadSize)
	}
}

func (snake *Snake) Update(eaten *int) {
	if snake.HeadPosition.X < 0 {
		snake.HeadPosition.X = float32(common.ScreenWidth) - common.SnakeHeadSize
	} else if snake.HeadPosition.X >= float32(common.ScreenWidth) {
		snake.HeadPosition.X = 0
	}
	if snake.HeadPosition.Y < 0 {
		snake.HeadPosition.Y = float32(common.ScreenHeight) - common.SnakeHeadSize
	} else if snake.HeadPosition.Y >= float32(common.ScreenHeight) {
		snake.HeadPosition.Y = 0
	}

	snake.EllapsedSinceLastJump += rl.GetFrameTime()

	if snake.EllapsedSinceLastJump >= rl.GetFrameTime()*2 {
		// set tail position
		newTailPositions := make([]rl.Vector2, *eaten)
		for i := range *eaten {
			if i == 0 {
				newTailPositions[i] = snake.HeadPosition
			} else {
				newTailPositions[i] = snake.TailPositions[i-1]
			}
		}
		snake.TailPositions = newTailPositions

		// move to to the direction
		snake.HeadPosition = rl.Vector2Add(snake.HeadPosition, snake.Direction)
		snake.EllapsedSinceLastJump = 0

		if slices.ContainsFunc(snake.TailPositions, func(tailPosition rl.Vector2) bool {
			return tailPosition == snake.HeadPosition
		}) {
			*eaten = 0
			snake.TailPositions = []rl.Vector2{}
			rl.PlaySound(*snake.DeadSound)
		}
	}
}

// draw frame
func (snake *Snake) Draw() {
	rl.DrawRectangleV(snake.HeadPosition, rl.NewVector2(common.SnakeHeadSize, common.SnakeHeadSize), common.SnakeColor)
	rl.DrawRectangleLinesEx(rl.NewRectangle(snake.HeadPosition.X, snake.HeadPosition.Y, common.SnakeHeadSize, common.SnakeHeadSize), 2, common.SnakeBorderColor)
	for _, followingPosition := range snake.TailPositions {
		rl.DrawRectangleV(followingPosition, rl.NewVector2(common.SnakeHeadSize, common.SnakeHeadSize), common.SnakeColor)
		rl.DrawRectangleLinesEx(rl.NewRectangle(followingPosition.X, followingPosition.Y, common.SnakeHeadSize, common.SnakeHeadSize), 2, common.SnakeBorderColor)
	}
}
