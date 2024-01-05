package main

import (
	"errors"
	"fmt"
)

type Mouse struct {
	name  string
	price int64
}

type Keyboard struct {
	name         string
	price        int64
	keyboardType string
}

type GamingStore interface {
	GetMouse(segment string) Mouse
	GetKeyboard(segment string) Keyboard
}

type LogitechStore struct {
}

func (LogitechStore) GetMouse(segment string) Mouse {
	switch segment {
	case "low":
		return Mouse{
			name:  "Logitech G102",
			price: 360000,
		}
	case "medium":
		return Mouse{
			name:  "Logitech G304",
			price: 660000,
		}
	case "high":
		return Mouse{
			name:  "Logitech G903",
			price: 1600000,
		}
	default:
		return Mouse{}
	}
}

func (LogitechStore) GetKeyboard(segment string) Keyboard {
	switch segment {
	case "low":
		return Keyboard{
			name:         "Logitech G413",
			price:        1290000,
			keyboardType: "TKL",
		}
	case "medium":
		return Keyboard{
			name:         "Logitech Pro X GX",
			price:        2090000,
			keyboardType: "TKL",
		}
	case "high":
		return Keyboard{
			name:         "Logitech G913",
			price:        3490000,
			keyboardType: "TKL",
		}
	default:
		return Keyboard{}
	}
}

type RazerStore struct {
}

func (RazerStore) GetMouse(segment string) Mouse {
	switch segment {
	case "low":
		return Mouse{
			name:  "Razer Cobra",
			price: 890000,
		}
	case "medium":
		return Mouse{
			name:  "Razer Viper V3",
			price: 1650000,
		}
	case "high":
		return Mouse{
			name:  "Razer Cobra Pro",
			price: 2850000,
		}
	default:
		return Mouse{}
	}
}

func (RazerStore) GetKeyboard(segment string) Keyboard {
	switch segment {
	case "low":
		return Keyboard{
			name:         "Razer BlackWindow V3",
			price:        1490000,
			keyboardType: "TKL",
		}
	case "medium":
		return Keyboard{
			name:         "Razer BlackWindow V3 Mini Hyper Speed",
			price:        2390000,
			keyboardType: "65%",
		}
	case "high":
		return Keyboard{
			name:         "Razer BlackWindow V3 Pro",
			price:        5190000,
			keyboardType: "Full size",
		}
	default:
		return Keyboard{}
	}
}

func CreateGamingStore(storeName string) (GamingStore, error) {
	switch storeName {
	case "logitech":
		return new(LogitechStore), nil
	case "razer":
		return new(RazerStore), nil
	default:
		return nil, errors.New("store not found")
	}
}

func main() {
	store, err := CreateGamingStore("logitech")

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(store.GetMouse("high"))
	fmt.Println(store.GetKeyboard("low"))
}
