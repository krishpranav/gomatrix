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

// type
type Point struct {
	Head 	bool
	Char	byte
	Age		int 
}

// variables
var clear map[string]func()

var cols [][]Point
var width int
var height int

// main fucntion
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
	

}