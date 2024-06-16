package Day12

import (
	"maps"
	"slices"
	"strings"
	"testing"
)

type Paths struct {
	ToFrom map[string][]string
}

type Route struct {
	Path    []string
	Current string
	Seen    map[string]int
	Twice   bool
}

func ParseInput(input string) *Paths {
	toFrom := make(map[string][]string)
	paths := &Paths{
		ToFrom: toFrom,
	}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "-")
		paths.ToFrom[parts[0]] = append(paths.ToFrom[parts[0]], parts[1])
		paths.ToFrom[parts[1]] = append(paths.ToFrom[parts[1]], parts[0])
	}
	return paths
}

func Part1(input string) int {
	paths := ParseInput(input)
	var queue []*Route
	for _, next := range paths.ToFrom["start"] {
		queue = append(queue, &Route{
			Path:    []string{"start", next},
			Current: next,
			Seen:    map[string]int{"start": 1, next: 1},
		})
	}
	var finishRoutes []*Route
	for len(queue) > 0 {
		route := queue[0]
		queue = queue[1:]
		if route.Current == "end" {
			finishRoutes = append(finishRoutes, route)
			continue
		}
		if route.Current == "start" {
			continue
		}
		for _, next := range paths.ToFrom[route.Current] {
			if i, ok := route.Seen[next]; ok && i > 0 {
				// skip lowercase caverns
				if next[0] >= 'a' && next[0] <= 'z' {
					continue
				}
			}
			newSeen := maps.Clone(route.Seen)
			newSeen[next] = 1
			newPath := append(route.Path, next)
			queue = append(queue, &Route{
				Path:    newPath,
				Current: next,
				Seen:    newSeen,
			})
		}
	}
	return len(finishRoutes)
}

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "Example1",
			input: `start-A
start-b
A-c
A-b
b-d
A-end
b-end`,
			want: 10,
		},
		{
			name: "Example2",
			input: `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`,
			want: 19,
		},
		{
			name: "Example3",
			input: `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`,
			want: 226,
		},
		{
			name: "real",
			input: `xq-XZ
zo-yr
CT-zo
yr-xq
yr-LD
xq-ra
np-zo
end-LD
np-LD
xq-kq
start-ra
np-kq
LO-end
start-xq
zo-ra
LO-np
XZ-start
zo-kq
LO-yr
kq-XZ
zo-LD
kq-ra
XZ-yr
LD-ws
np-end
kq-yr`,
			want: 5457,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Part1(test.input); got != test.want {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}

func Part2(input string) int {
	paths := ParseInput(input)
	var queue []*Route
	for _, next := range paths.ToFrom["start"] {
		queue = append(queue, &Route{
			Path:    []string{"start", next},
			Current: next,
			Seen:    map[string]int{"start": 1, next: 1},
		})
	}
	var finishRoutes []*Route
	for len(queue) > 0 {
		route := queue[0]
		queue = queue[1:]
		if route.Current == "end" {
			finishRoutes = append(finishRoutes, route)
			continue
		}
		if route.Current == "start" {
			continue
		}
		for _, next := range paths.ToFrom[route.Current] {
			// never backtrack to start
			if next == "start" {
				continue
			}
			if i, ok := route.Seen[next]; ok && i > 0 {
				// skip lowercase caverns if we've already been there twice or twice to another one
				if next[0] >= 'a' && next[0] <= 'z' && route.Twice {
					continue
				}
			}
			newSeen := maps.Clone(route.Seen)
			twice := route.Twice
			if _, ok := newSeen[next]; ok {
				newSeen[next]++
			} else {
				newSeen[next] = 1
			}
			if next[0] >= 'a' && next[0] <= 'z' {
				if newSeen[next] >= 2 {
					twice = true
				}
			}
			newPath := append(slices.Clip(route.Path), next)
			queue = append(queue, &Route{
				Path:    newPath,
				Current: next,
				Seen:    newSeen,
				Twice:   twice,
			})
		}
	}
	return len(finishRoutes)
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "Example1",
			input: `start-A
start-b
A-c
A-b
b-d
A-end
b-end`,
			want: 36,
		},
		{
			name: "Example2",
			input: `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`,
			want: 103,
		},
		{
			name: "Example3",
			input: `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`,
			want: 3509,
		},
		{
			name: "real",
			input: `xq-XZ
zo-yr
CT-zo
yr-xq
yr-LD
xq-ra
np-zo
end-LD
np-LD
xq-kq
start-ra
np-kq
LO-end
start-xq
zo-ra
LO-np
XZ-start
zo-kq
LO-yr
kq-XZ
zo-LD
kq-ra
XZ-yr
LD-ws
np-end
kq-yr`,
			want: 128506,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Part2(test.input); got != test.want {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}
