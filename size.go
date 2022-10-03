package ui

import tea "github.com/charmbracelet/bubbletea"

type Size struct {
	Width  int
	Height int
}

func NewSizeFromSizeMsg(msg tea.WindowSizeMsg) Size {
	return Size{
		Width:  msg.Width,
		Height: msg.Height,
	}
}

type SizeDefinition interface {
}

func FixedSize(v int) SizeDefinition {
	return fixed{
		Value: v,
	}
}

type fixed struct {
	SizeDefinition
	Value int
}

func RemainingSize() SizeDefinition {
	return remaining{}
}

type remaining struct {
	SizeDefinition
}

func RatioSize(d int) SizeDefinition {
	return ratio{
		Denominator: d,
	}
}

type ratio struct {
	SizeDefinition
	Denominator int
}

func Resolve(req []SizeDefinition, max int) []int {
	flexiNo := 0
	fixedSize := 0
	for _, r := range req {
		switch x := r.(type) {
		case fixed:
			fixedSize += x.Value
		case ratio:
			fixedSize += max / x.Denominator
		case remaining:
			flexiNo++
		}
	}

	realSizes := []int{}
	remainingSpace := max
	for _, r := range req {
		s := 0
		if remainingSpace > 0 {
			switch x := r.(type) {
			case fixed:
				s = x.Value
			case ratio:
				s = max / x.Denominator
			case remaining:
				if flexiNo == 1 {
					// the last one
					s = max - fixedSize
					flexiNo--
				} else {
					s = (max - fixedSize) / flexiNo
					fixedSize += s
					flexiNo--
				}
			}
		}
		if s > remainingSpace {
			s = remainingSpace
		}
		remainingSpace -= s
		realSizes = append(realSizes, s)

	}
	return realSizes
}
