package models

import "fmt"

//
//工厂模式（Factory Pattern）
//

type Shape interface {
	Draw()
}

type Rectangle struct{}

func (r Rectangle) Draw() {
	fmt.Println("Draw Rectangle!")
}

type Square struct{}

func (r Square) Draw() {
	fmt.Println("Draw Square!")
}

type Circle struct{}

func (r Circle) Draw() {
	fmt.Println("Draw Circle!")
}

type ShapeFactory struct{}

func (r ShapeFactory) GetFactory(shapeType string) Shape {
	if shapeType == "" {
		return nil
	}

	switch shapeType {
	case "Rectangle":
		return Rectangle{}
	case "Square":
		return Square{}
	case "Circle":
		return Circle{}
	}
	return nil
}

func TestShapeFactory() {
	factory := ShapeFactory{}
	factory.GetFactory("Rectangle").Draw()
	factory.GetFactory("Square").Draw()
	factory.GetFactory("Circle").Draw()
}
