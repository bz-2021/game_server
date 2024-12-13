package main

import (
	"fmt"
	"image/color"
	"log"
	"strconv"

	"github.com/ByteArena/box2d"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	resize             = 20 // 渲染放大系数
	ppm                = 1  // 像素到物理单位的比例
	obstacleWidth      = 6.0
	obstacleHeight     = 6.0
	tankWidth          = 3.0
	tankHeight         = 3.0
	shellWidth         = 10.0
	shellHeight        = 10.0
	screenWidth        = 800
	screenHeight       = 800
	timeStep           = 1.0 / 60
	velocityIterations = 6
	positionIterations = 0
)

type Map struct {
	World         *box2d.B2World
	TankBody      *box2d.B2Body
	Obstacles     []*box2d.B2Body
	Shells        []*box2d.B2Body
	Width, Height int
}

func NewMap(width, height int) *Map {
	gravity := box2d.MakeB2Vec2(0, 0)
	world := box2d.MakeB2World(gravity)

	obstacleBodyDef := box2d.MakeB2BodyDef()
	obstacleBodyDef.Type = box2d.B2BodyType.B2_staticBody
	obstacleBodyDef.Position = box2d.MakeB2Vec2(6, 6)
	obstacleBody := world.CreateBody(&obstacleBodyDef)

	obstacleShape := createRectangleShape(obstacleWidth, obstacleHeight)
	obstacleFixtureDef := box2d.MakeB2FixtureDef()
	obstacleFixtureDef.Shape = &obstacleShape
	obstacleFixtureDef.Density = 1.0
	obstacleBody.CreateFixtureFromDef(&obstacleFixtureDef)

	tankBodyDef := box2d.MakeB2BodyDef()
	tankBodyDef.Type = box2d.B2BodyType.B2_dynamicBody
	tankBodyDef.Position = box2d.MakeB2Vec2(12, 12)
	tankBody := world.CreateBody(&tankBodyDef)

	tankShape := createRectangleShape(tankWidth, tankHeight)
	tankFixtureDef := box2d.MakeB2FixtureDef()
	tankFixtureDef.Shape = &tankShape
	tankFixtureDef.Density = 0.5
	tankFixtureDef.Friction = 0.4
	tankBody.CreateFixtureFromDef(&tankFixtureDef)

	return &Map{
		World:     &world,
		TankBody:  tankBody,
		Obstacles: []*box2d.B2Body{obstacleBody},
		Width:     width,
		Height:    height,
	}
}

func createRectangleShape(width, height float64) box2d.B2PolygonShape {
	shape := box2d.MakeB2PolygonShape()
	shape.SetAsBox(width/2, height/2)
	return shape
}

func (m *Map) Update() error {
	m.handleInput()
	m.World.Step(timeStep, velocityIterations, positionIterations)
	return nil
}

func (m *Map) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	ebitenutil.DebugPrint(screen, "FPS: "+strconv.FormatFloat(ebiten.ActualFPS(), 'f', 1, 64))
	ebitenutil.DebugPrint(screen, "\nTPS: "+strconv.FormatFloat(ebiten.ActualTPS(), 'f', 1, 64))
	for _, obstacle := range m.Obstacles {
		pos := obstacle.GetPosition()
		vector.DrawFilledRect(screen, float32((pos.X-obstacleWidth/2)*ppm*resize), float32((pos.Y-obstacleHeight/2)*ppm*resize), float32(obstacleWidth)*ppm*resize, float32(obstacleHeight)*ppm*resize, color.Black, false)
	}
	//vectorvector.DrawFilledRect(screen, 0, 0, float32(obstacleWidth)*ppm*resize, float32(obstacleHeight)*ppm*resize, color.Black, false)

	tankPos := m.TankBody.GetPosition()
	fmt.Println(tankPos.X, tankPos.Y)
	vector.DrawFilledRect(screen, float32((tankPos.X-tankWidth/2)*ppm*resize), float32((tankPos.Y-tankHeight/2)*ppm*resize), float32(tankWidth)*ppm*resize, float32(tankHeight)*ppm*resize, color.RGBA{255, 0, 0, 255}, false)
}

func (m *Map) Layout(outsideWidth, outsideHeight int) (int, int) {
	return m.Width, m.Height
}

func (m *Map) handleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		force := box2d.MakeB2Vec2(0, -5)
		m.TankBody.SetLinearVelocity(force)
		fmt.Println("up")
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		force := box2d.MakeB2Vec2(0, 5)
		m.TankBody.SetLinearVelocity(force)
		fmt.Println("down")
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		force := box2d.MakeB2Vec2(-5, 0)
		m.TankBody.SetLinearVelocity(force)
		fmt.Println("left")
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		force := box2d.MakeB2Vec2(5, 0)
		m.TankBody.SetLinearVelocity(force)
		fmt.Println("right")
	}
}

func main() {
	mp := NewMap(screenWidth, screenHeight)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Tank Game")
	if err := ebiten.RunGame(mp); err != nil {
		log.Fatal(err)
	}
}
