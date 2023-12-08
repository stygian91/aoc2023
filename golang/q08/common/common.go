package common

import "strings"

type LR = int

type Nodes struct {
	Keys    []string
	Mapping map[string][]string
}

const (
	Left = iota
	Right
)

func ParseLR(line string) []LR {
	result := []LR{}

	for _, ch := range line {
		switch ch {
		case 'L':
			result = append(result, Left)
		case 'R':
			result = append(result, Right)
		default:
			panic("Invalid LR instruction")
		}
	}

	return result
}

func ParseNodes(lines []string) Nodes {
	mapping := map[string][]string{}
	keys := []string{}

	for _, line := range lines {
		key, values := parseNode(line)
		keys = append(keys, key)
		mapping[key] = values
	}

	return Nodes{
		Keys: keys,
		Mapping: mapping,
	}
}

func parseNode(line string) (string, []string) {
	parts := strings.Split(line, " = ")
	values := strings.Split(parts[1][1:len(parts[1])-1], ", ")

	return parts[0], values
}
