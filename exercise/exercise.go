package exercise

import (
	"fmt"
	"slices"
)

type Item struct {
	Name string
	Type string
}

type Player struct {
	Name      string
	Inventory []Item
}

// Pick up an item
func (p *Player) PickUpItem(item Item) {
	p.Inventory = append(p.Inventory, item)
}

func (p *Player) DropItem(item Item) {
	for i := range p.Inventory {
		if p.Inventory[i] == item {
			p.Inventory = slices.Delete(p.Inventory, i, i+1)
			fmt.Println("Item dropped:", item.Name)
			return
		}
	}
	fmt.Println("Item not found in inventory:", item.Name)
}

func (p *Player) UseItem(item Item) {
	for i, itemValue := range p.Inventory {
		if item.Name == itemValue.Name {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			fmt.Println("Used item:", item.Name)
			return
		}
	}
	fmt.Println("No usable item found in inventory")
}
