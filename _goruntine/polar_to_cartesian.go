package _goruntine

import (
	"fmt"
	"math"
)

type Polar struct {
	Radius float64
	Angle  float64
}

type Cartesian struct {
	X float64
	Y float64
}

func TestConvertion() {
	var radius, angle float64
	var p = new(Polar)
	in := make(chan *Polar)
	defer close(in)
	out := make(chan *Cartesian)
	defer close(out)
	convert(in, out)

	for {
		fmt.Printf("Please enter the polar coordinate(radius, angle): ")
		fmt.Scanf("%f %f", &radius, &angle)

		p.Radius = radius
		p.Angle = angle
		in <- p
		c := <-out
		fmt.Printf("convert result: X=%.02f, Y=%.02f\n", c.X, c.Y)
	}
}

func convert(in chan *Polar, out chan *Cartesian) {
	go func() {
		for {
			p := <-in
			angle := p.Angle / 180.0 * math.Pi
			c := new(Cartesian)
			c.X = p.Radius * math.Cos(angle)
			c.Y = p.Radius * math.Sin(angle)
			out <- c
		}
	}()
}

// ============================================================================================================
//const result = "Polar: radius=%.02f angle=%.02f degrees -- Cartesian: x=%.02f y=%.02f\n"
//
//var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " + "or %s to quit."
//
//func init() {
//	if runtime.GOOS == "windows" {
//		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
//	} else { // Unix-like
//		prompt = fmt.Sprintf(prompt, "Ctrl+D")
//	}
//}
//
//func main() {
//	questions := make(chan polar)
//	defer close(questions)
//	answers := createSolver(questions)
//	defer close(answers)
//	interact(questions, answers)
//}
//
//func createSolver(questions chan polar) chan cartesian {
//	answers := make(chan cartesian)
//	go func() {
//		for {
//			polarCoord := <-questions
//			Θ := polarCoord.Θ * math.Pi / 180.0 // degrees to radians
//			x := polarCoord.radius * math.Cos(Θ)
//			y := polarCoord.radius * math.Sin(Θ)
//			answers <- cartesian{x, y}
//		}
//	}()
//	return answers
//}
//
//func interact(questions chan polar, answers chan cartesian) {
//	reader := bufio.NewReader(os.Stdin)
//	fmt.Println(prompt)
//	for {
//		fmt.Printf("Radius and angle: ")
//		line, err := reader.ReadString('\n')
//		if err != nil {
//			break
//		}
//		line = line[:len(line)-1] // chop of newline character
//		if numbers := strings.Fields(line); len(numbers) == 2 {
//			polars, err := floatsForStrings(numbers)
//			if err != nil {
//				fmt.Fprintln(os.Stderr, "invalid number")
//				continue
//			}
//			questions <- polar{polars[0], polars[1]}
//			coord := <-answers
//			fmt.Printf(result, polars[0], polars[1], coord.x, coord.y)
//		} else {
//			fmt.Fprintln(os.Stderr, "invalid input")
//		}
//	}
//	fmt.Println()
//}
//
//func floatsForStrings(numbers []string) ([]float64, error) {
//	var floats []float64
//	for _, number := range numbers {
//		if x, err := strconv.ParseFloat(number, 64); err != nil {
//			return nil, err
//		} else {
//			floats = append(floats, x)
//		}
//	}
//	return floats, nil
//}
