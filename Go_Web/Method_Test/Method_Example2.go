package main

import "fmt"

const (
	WHITE   =   iota        // iota 常量计数器
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth    float64
	color                   Color
}

type BoxList []Box          // a slice of boxes

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {           // 传值只是复制，所以最后输出没有真的改变颜色，在这里接收者必须作为一个指针
	b.color = c
}

func (bl BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITE)

	for _, b := range bl{
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {          // Go可以自动识别，Go知道receiver是指针可以自动转换
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (bl BoxList) PaintItBlack2() {          // 这么写也对
	for i, _ := range bl {
		(&bl[i]).SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{    4,  4,  4,  RED     },
		Box{    10, 10, 1,  YELLOW  },
		Box{    1,  1,  20, BLACK   },
		Box{    10, 10, 1,  BLUE    },
		Box{    10, 30, 1,  WHITE   },
		Box{    20, 20, 20, YELLOW  },
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volumn of the first one is", boxes[0].Volume())
	fmt.Println("The color of the last one is", boxes[len(boxes) - 1].color)
	fmt.Println("The biggest one is", boxes.BiggestsColor().String())
	fmt.Println("Paint them all black")
	boxes.PaintItBlack2()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("The biggest one is", boxes.BiggestsColor().String())
	fmt.Println("The biggest one is", boxes.BiggestsColor())
}
