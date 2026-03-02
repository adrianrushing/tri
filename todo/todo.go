package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	position int
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

type ByPri []Item

func (s ByPri) Len() int { return len(s) }

func (s ByPri) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByPri) Less(i, j int) bool {
	if s[i].Priority == s[j].Priority {
		return s[i].position < s[j].position
	}
	return s[i].Priority < s[j].Priority
}

func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "(1)"
	}

	if i.Priority == 3 {
		return "(3)"
	}
	return ""
}

func ReadItems(filename string) ([]Item, error) {

	b, err := os.ReadFile(filename)

	if err != nil {
		return []Item{}, err
	}

	var items []Item

	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	for i := range items {
		items[i].position = i + 1
	}

	return items, nil
}

func SaveItems(filename string, items []Item) error {

	b, err := json.Marshal(items)

	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}
