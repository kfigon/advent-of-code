package day8

import (
	"io"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// https://adventofcode.com/2019/day/8

func readFile(t *testing.T) []int {
	file, err := os.Open("data.txt")
	require.NoError(t, err)
	defer file.Close()

	content, err := io.ReadAll(file)
	require.NoError(t, err)

	strContent := string(content)
	out := make([]int, 0)
	for i := 0; i < len(strContent); i++ {
		v, err := strconv.Atoi(string(strContent[i]))
		require.NoError(t, err)
		
		out = append(out, v)
	}
	return out
	
}
func TestParsingLayers(t *testing.T) {
	in := []int{1,2,3,4,5,6,7,8,9,0,1,2}
	width,height := 3,2

	image := parseInput(in, width, height)

	assert.Equal(t, 2, image.numberOfLayers())
	assert.Equal(t, []int{1,2,3,4,5,6}, image.extractLayer(0))
	assert.Equal(t, []int{7,8,9,0,1,2}, image.extractLayer(1))
}

func TestParsingLayers2(t *testing.T) {
	in := []int{1,2,3,4,5,6,7,8,9,0,1,2}
	width,height := 3,2

	image := parseInput(in, width, height)

	assert.Equal(t, 2, image.numberOfLayers())
	assert.Equal(t, []int{1,2,3,4,5,6}, image.extractLayer(0))
	assert.Equal(t, []int{7,8,9,0,1,2}, image.extractLayer(1))
}

func TestPart1(t *testing.T) {
	file := readFile(t)
	const width int = 25
	const height int = 6

	image := parseInput(file, width, height)
	layer := image.findLayerWithSmallestNumberOf(0)
	num := findOccurences(layer,1)*findOccurences(layer,2)
	assert.Equal(t, 1806,num)
}

func TestLayeringPart2(t *testing.T) {
	input := []int{0,2,2,2,1,1,2,2,2,2,1,2,0,0,0,0}
	img := parseInput(input,2,2)

	assert.Equal(t, 4, img.numberOfLayers())
	assert.Equal(t, 4, img.layerSize())
	assert.Equal(t, []int{0,2,2,2}, img.extractLayer(0))
	assert.Equal(t, []int{1,1,2,2}, img.extractLayer(1))
	assert.Equal(t, []int{2,2,1,2}, img.extractLayer(2))
	assert.Equal(t, []int{0,0,0,0}, img.extractLayer(3))
}

func TestPart2Example(t *testing.T) {
	input := []int{0,2,2,2,1,1,2,2,2,2,1,2,0,0,0,0}
	img := parseInput(input,2,2)

	assert.Equal(t, []int{0,1,1,0}, img.getFinalImage())
}

func TestPart2(t *testing.T) {
	file := readFile(t)
	const width int = 25
	const height int = 6

	img := parseInput(file,width,height)

	got := img.getFinalImage()
	assert.Equal(t, []int{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 
		1, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 
		1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 
		1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 
		0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 
		0, 1, 0, 0, 1, 0}, got)

	formattedMsg := ""
	for i := 0; i < len(got); i++ {
		v := got[i]
		if i != 0 && i % width == 0 {
			formattedMsg += "\n"
		}
		if v == 1 {
			formattedMsg+="8"
		} else {
			formattedMsg+=" "
		}
	}
	// JAFRA
	// assert.Equal(t,"asd", formattedMsg)
}

type size struct {
	width int
	height int
}

func parseInput(data []int, width, height int) *image {
	return &image {
		in: data,
		s: size{width: width,height: height},
	}
}

type image struct {
	in []int
	s size
}

func (i *image) numberOfLayers() int { return len(i.in)/i.layerSize() }
func (i *image) layerSize() int { return i.s.width*i.s.height }

func (i *image) extractLayer(layerNum int) []int {
	if layerNum >= i.numberOfLayers(){
		return []int{}
	}

	start := layerNum * i.layerSize()
	stop := start+i.layerSize()
	return i.in[start:stop]
}

func findOccurences(arr []int, toFind int) int {
	occurences := 0
	for i := 0; i < len(arr); i++ {
		if toFind == arr[i]{
			occurences++
		}
	}
	return occurences
}

func (i *image) findLayerWithSmallestNumberOf(numToFind int) []int {

	var min *int
	layerIdx := 0
	for layerNum := 0; layerNum < i.numberOfLayers(); layerNum++ {
		layerData := i.extractLayer(layerNum)
		occurences := findOccurences(layerData, numToFind)
		
		if min == nil || occurences < *min {
			min = &occurences
			layerIdx = layerNum
		}
	}

	if min == nil {
		return []int{}
	}
	return i.extractLayer(layerIdx)
}

func (i *image) getFinalImage() []int {
	out := make([]int,0)
	for pixel := 0; pixel < i.layerSize(); pixel++ {
		out = append(out, 2)
	}

	for layerIdx := 0; layerIdx < i.numberOfLayers(); layerIdx++ {
		layer := i.extractLayer(layerIdx)

		for pixel := 0; pixel < len(layer); pixel++ {
			out[pixel] = determineTopVisiblePixel(out[pixel], layer[pixel])
		}
	}
	return out
}

const (
	BLACK = 0
	WHITE = 1
	TRANSPARENT = 2
)

func determineTopVisiblePixel(currentPixel, nextLayerPixel int) int {
	if currentPixel == TRANSPARENT {
		return nextLayerPixel
	}
	return currentPixel
}