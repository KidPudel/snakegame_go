package main

import (
	"fmt"
	"math"

	"github.com/KidPudel/snakegame_go/common"
	"github.com/KidPudel/snakegame_go/entities"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Width  int32
	Height int32
	Title  string
}

func (game *Game) SetGame() {
	rl.SetTargetFPS(30)
	rl.SetWindowState(rl.FlagWindowResizable)
}

func (game *Game) StartMusic() *rl.Music {
	rl.InitAudioDevice()
	music := rl.LoadMusicStream("resources/BeepBox_snake.wav")
	rl.SetMusicVolume(music, 0.05)
	rl.PlayMusicStream(music)
	return &music

}

func (game *Game) Update(music *rl.Music) {
	rl.UpdateMusicStream(*music)
	if rl.IsWindowResized() {
		common.ScreenWidth = int32(math.Ceil(float64(rl.GetScreenWidth())/float64(common.SnakeHeadSize))) * int32(common.SnakeHeadSize)
		common.ScreenHeight = int32(math.Ceil(float64(rl.GetScreenHeight())/float64(common.SnakeHeadSize))) * int32(common.SnakeHeadSize)
		fmt.Println(common.ScreenWidth, common.ScreenHeight)
	}
}

func (game *Game) DrawField() {
	for i := range int32(math.Ceil(float64(common.ScreenHeight) / float64(common.SnakeHeadSize))) {
		rl.DrawLine(0, int32(i*int32(common.SnakeHeadSize)), common.ScreenWidth, int32(i*int32(common.SnakeHeadSize)), rl.Black)
	}
	for i := range int32(math.Ceil(float64(common.ScreenWidth) / float64(common.SnakeHeadSize))) {
		rl.DrawLine(int32(i*int32(common.SnakeHeadSize)), 0, int32(i*int32(common.SnakeHeadSize)), common.ScreenHeight, rl.Black)
	}
}

func (game *Game) Draw() {
	rl.ClearBackground(common.BackgroundColor)
	game.DrawField()
}

func (game *Game) Run() {
	rl.InitWindow(game.Width, game.Height, game.Title)
	game.SetGame()
	music := game.StartMusic()

	deadSound := rl.LoadSound("resources/hit.wav")
	snake := entities.Snake{DeadSound: &deadSound}
	foodFarm := entities.InitFoodFarm()

	for !rl.WindowShouldClose() {

		// handle input
		snake.HandleInput()

		// update, apply modifications
		game.Update(music)
		snake.Update(&foodFarm.Eaten)
		foodFarm.Update(snake.HeadPosition)

		// render
		rl.BeginDrawing()

		game.Draw()
		foodFarm.Draw()

		snake.Draw()

		rl.DrawRectangleV(rl.NewVector2((float32(common.ScreenWidth)/2)-(common.SnakeHeadSize*3), float32(0)), rl.NewVector2(common.SnakeHeadSize*5, common.SnakeHeadSize*2), rl.ColorAlpha(rl.Black, 0.5))
		rl.DrawText(fmt.Sprint("Score: ", foodFarm.Eaten), int32((float64(common.ScreenWidth)/2)-(float64(common.SnakeHeadSize)*3)), 10, 23, rl.White)

		// end of a frame, back buffer is complete
		rl.EndDrawing()

	}
	rl.UnloadMusicStream(*music)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func main() {
	game := Game{Width: common.ScreenWidth, Height: common.ScreenHeight, Title: "snake"}
	game.Run()
}
