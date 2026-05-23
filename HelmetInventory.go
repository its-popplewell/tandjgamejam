package main

type HelmetInventory struct {
	capacity    int32
	used_volume int32
	modifiers   []Modifier
}

func NewHelmetInventory(capacity int32) HelmetInventory {
	return HelmetInventory{
		capacity:    capacity,
		used_volume: 0,
		modifiers:   []Modifier{},
	}
}

func (hi *HelmetInventory) addModifier(toAdd Modifier) bool {
	if hi.getRemainingCapacity() < toAdd.size {
		return false
	}

	hi.modifiers = append(hi.modifiers, toAdd)
	return true
}

func (hi *HelmetInventory) removeModifier(toRemove Modifier) {
	DeleteFunc(&hi.modifiers, func(m Modifier) bool {
		return (m.getId() == toRemove.getId())
	})
}

func (hi HelmetInventory) getRemainingCapacity() int32 {
	return hi.capacity - hi.used_volume
}

func (hi HelmetInventory) getAttackModifier() [2]int32 {
	outp := [2]int32{0, 0}
	for _, m := range hi.modifiers {
		modifier := m.getAttackModifier()
		outp[0] += modifier[0]
		outp[1] += modifier[1]
	}

	return outp
}
