package d7

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	data := parse(strings.Split(testInput, "\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 95437, p1(data))
	})

	t.Run("p2", func(t *testing.T) {

	})
}

func TestFile(t *testing.T) {
	rawData, err := os.ReadFile("data.txt")
	require.NoError(t, err)

	data := parse(strings.Split(string(rawData), "\r\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Greater(t, 1062058, p1(data))
	})

	t.Run("p2", func(t *testing.T) {

	})
}

type directory struct {
	name           string
	subDirectories []*directory
	files          []file
}

func (dir directory) size() int {
	out := 0
	for _, f := range dir.files {
		out += f.size
	}
	for _, d := range dir.subDirectories {
		out += d.size()
	}
	return out
}

type file struct {
	name string
	size int
}

type stack[T any] struct {
	vs []T
}

func (s *stack[T]) push(v T) {
	s.vs = append(s.vs, v)
}

func (s *stack[T]) pop() (T, bool) {
	if len(s.vs) == 0 {
		var out T
		return out, false
	}
	out := s.vs[len(s.vs)-1]
	s.vs = s.vs[:len(s.vs)-1]
	return out, true
}

func parse(lines []string) directory {
	dirStack := &stack[*directory]{vs: []*directory{}}
	root := directory{
		name: "/",
	}
	var currentDir *directory = &root
	dirStack.push(currentDir)

	for _, line := range lines {
		if line == "$ ls" || line == "$ cd /" {
			continue
		} else if line == "$ cd .." {
			parent, ok := dirStack.pop()
			if !ok {
				continue
			}
			currentDir = parent
		} else if strings.HasPrefix(line, "$ cd ") {
			dirStack.push(currentDir)

			dirName := strings.TrimPrefix(line, "$ cd ")
			for _, d := range currentDir.subDirectories {
				if d.name == dirName {
					currentDir = d
					break
				}
			}
		} else if strings.HasPrefix(line, "dir") {
			newDir := directory{
				name: strings.TrimPrefix(line, "dir "),
			}
			(*currentDir).subDirectories = append((*currentDir).subDirectories, &newDir)
		} else if f, ok := parseFileLine(line); ok {
			fdata := file{size: f.a, name: f.b}
			(*currentDir).files = append((*currentDir).files, fdata)
		} else {
			return directory{} // unknown command
		}
	}
	return root
}

type pair[T any, V any] struct {
	a T
	b V
}

func parseFileLine(line string) (pair[int, string], bool) {
	ok, err := regexp.MatchString(`\d+ \w+`, line)
	if !ok || err != nil {
		return pair[int, string]{}, false
	}
	parts := strings.Split(line, " ")
	size, _ := strconv.Atoi(parts[0])
	return pair[int, string]{a: size, b: parts[1]}, true
}

func p1(fs directory) int {
	dirs := map[string]int{}
	var foo func(*directory)
	foo = func(d *directory) {
		if d == nil {
			return
		}
		dirs[d.name] = d.size()
		for _, v := range d.subDirectories {
			foo(v)
		}
	}

	for _, v := range fs.subDirectories {
		foo(v)
	}
	
	out := 0
	for _,v := range dirs {
		if v <= 100000 {
			out += v
		}
	}
	return out
}

const testInput = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`
