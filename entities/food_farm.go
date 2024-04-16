package entities

import (
	"math"
	"math/rand"
	"slices"

	"github.com/KidPudel/snakegame_go/common"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type FoodFarm struct {
	Foods []Food
	// logical state
	Eaten int

	EatenSound *rl.Sound
}

func InitFoodFarm() FoodFarm {
	sound := rl.LoadSound("resources/eat.wav")
	return FoodFarm{Foods: []Food{{Position: rl.NewVector2(float32(math.Ceil(float64(rand.Intn(rl.GetScreenWidth()/int(common.SnakeHeadSize)))))*common.SnakeHeadSize, float32(math.Ceil(float64(rand.Intn(int(rl.GetScreenHeight()))/int(common.SnakeHeadSize))))*common.SnakeHeadSize)}}, EatenSound: &sound}
}

func (farm *FoodFarm) SpawnFood() {
	farm.Foods = append(farm.Foods, Food{Position: rl.NewVector2(float32(math.Ceil(float64(rand.Intn(rl.GetScreenWidth()/int(common.SnakeHeadSize)))))*common.SnakeHeadSize, float32(math.Ceil(float64(rand.Intn(int(rl.GetScreenHeight()))/int(common.SnakeHeadSize))))*common.SnakeHeadSize)})
}

// Update farm if player has eaten food
func (farm *FoodFarm) Update(playerPosition rl.Vector2) {
	farm.Foods = slices.DeleteFunc(farm.Foods, func(food Food) bool {
		return food.Position == playerPosition
	})
	if len(farm.Foods) == 0 {
		farm.Eaten++
		rl.PlaySound(*farm.EatenSound)
		farm.SpawnFood()
	}

}

func (farm *FoodFarm) Draw() {
	for _, food := range farm.Foods {
		food.Draw()
	}
}
