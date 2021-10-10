package main

import "fmt"

type renderer interface {
	renderCircle(radius int)
}

type vectorRenderer struct{}

func (v *vectorRenderer) renderCircle(radius int) {
	fmt.Println("drawing circle with radius ", radius)
}

type circle struct {
	renderer renderer
	radius   int
}

func (c *circle) draw() {
	c.renderer.renderCircle(c.radius)
}

func main() {
	vectorRenderer := &vectorRenderer{}
	circle := &circle{
		renderer: vectorRenderer,
		radius:   10,
	}

	circle.draw()
}
