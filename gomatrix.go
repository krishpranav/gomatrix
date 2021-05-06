package main

// imports
import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/nsf/termbox-go"
)

// point struct
type Point struct {
	Head bool
	Char byte
	Age  int
}

// variables
var clear map[string]func()

var cols [][]Point
var width int
var height int

// main function
func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}

	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	width, height = termbox.Size()

	cols = make2d(width, height)

	data := make([]byte, width*height)
	rand.Read(data)

	for i := range cols {
		for j := range cols[i] {
			cols[i][j] = Point{false, (data[i*height+j] % 94) + 33, 0}
		}
	}

	data = make([]byte, 32)
	rand.Read(data)

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	create()
loop:
	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey && (ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC) {
				break loop
			}
		default:
			 
			print(cols)
			time.Sleep(time.Millisecond * 50)
			step()
			create()
		}
	}
}

// make 2d design
func make2d(width int, height int) [][]Point {
	arr := make([][]Point, width)
	for i := range arr {
		arr[i] = make([]Point, height)
	}
	return arr
}

// create rand
func create() {
	chance := 7
	q := rand.Int() % chance
	for i := 0; i < q; i++ {
		p := rand.Int() % width
		l := rand.Int()%24 + 9
		cols[p][len(cols[p])-1].Age = l
	}
}

// color head, age
func step() {
	newcols := make([][]Point, width)
	for i := range newcols {
		newcols[i] = append([]Point(nil), cols[i]...)
	}
	for i := len(cols) - 1; i >= 0; i-- {
		for j := 0; j < len(cols[i]); j++ {
			if j != len(cols[i])-1 {
				if cols[i][j].Age == 0 && cols[i][j+1].Age > 0 {
					newcols[i][j].Age = cols[i][j+1].Age
					newcols[i][j].Head = true
					newcols[i][j+1].Head = false
				}
			}
			if cols[i][j].Age > 0 {
				newcols[i][j].Age--
			}
		}
	}
	cols = newcols
}

// print matrix
func print(data [][]Point) {
	for i := range data[0] {
		for j := range data {
			toshow := ' '

			a := j
			b := height - i - 1

			point := data[a][b]

			if data[a][b].Age > 0 {
				toshow = []rune(string(point.Char))[0]
			}
			if point.Head {
				termbox.SetCell(a, height-b-1, toshow, termbox.ColorWhite, termbox.ColorBlack)
				 
			} else {
				termbox.SetCell(a, height-b-1, toshow, termbox.ColorGreen, termbox.ColorBlack)
				 
			}
		}
		if i < len(data[0])-1 {
			 
		}
	}
	termbox.Flush()
}

// initialize
func initClear() {
	clear = make(map[string]func())  
	clear["linux"] = func() {
		cmd := exec.Command("clear")  
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")  
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear")  
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// clear screen
func clearScreen() {
	fmt.Printf("%s\n", runtime.GOOS)
	value, ok := clear[runtime.GOOS]  
	if ok {                           
		value()  
	} else {  
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}