package vainglory

import "time"

// TelemetryEvent is an interface for TelemetryEvent's.
type TelemetryEvent interface {
	GetType() string         // returns the event type.
	GetTimestamp() time.Time // returns the Timestamp of the event.
}

// Base is the base of all telemetery event types.
// This information will always be returned.
type Base struct {
	Timestamp time.Time `json:"time"`
	Type      string    `json:"type"`
}

// GetType returns the event type.
func (b Base) GetType() string {
	return b.Type
}

// GetTimestamp returns the event timestamp.
func (b Base) GetTimestamp() time.Time {
	return b.Timestamp
}

type HeroBanEvent struct {
	Base
	Payload struct {
		Hero string `json:"Hero"`
		Team string `json:"Team"`
	} `json:"payload"`
}

type HeroSelectEvent struct {
	Base
	Payload struct {
		Hero   string `json:"Hero"`
		Team   string `json:"Team"`
		Player string `json:"Player"`
		Handle string `json:"Handle"`
	} `json:"payload"`
}

type HeroSkinSelectEvent struct {
	Base
	Payload struct {
		Hero string `json:"Hero"`
		Skin string `json:"Skin"`
	} `json:"payload"`
}

type PlayerFirstSpawnEvent struct {
	Base
	Payload struct {
		Team  string `json:"Team"`
		Actor string `json:"Actor"`
	} `json:"payload"`
}

type LevelUpEvent struct {
	Base
	Payload struct {
		Team         string `json:"Team"`
		Actor        string `json:"Actor"`
		Level        int    `json:"Level"`
		LifetimeGold int    `json:"LifetimeGold"`
	} `json:"payload"`
}

type BuyItemEvent struct {
	Base
	Payload struct {
		Team          string    `json:"Team"`
		Actor         string    `json:"Actor"`
		Item          string    `json:"Item"`
		Cost          int       `json:"Cost"`
		RemainingGold int       `json:"RemainingGold"`
		Position      []float64 `json:"Position"`
	} `json:"payload"`
}

type LearnAbilityEvent struct {
	Base
	Payload struct {
		Team    string `json:"Team"`
		Actor   string `json:"Actor"`
		Ability string `json:"Ability"`
		Level   int    `json:"Level"`
	} `json:"payload"`
}

type UseAbilityEvent struct {
	Base
	Payload struct {
		Team           string    `json:"Team"`
		Actor          string    `json:"Actor"`
		Ability        string    `json:"Ability"`
		Position       []float64 `json:"Position"`
		TargetActor    string    `json:"TargetActor"`
		TargetPosition []float64 `json:"TargetPosition"`
	} `json:"payload"`
}

type UseItemAbilityEvent struct {
	Base
	Payload struct {
		Team           string    `json:"Team"`
		Actor          string    `json:"Actor"`
		Ability        string    `json:"Ability"`
		Position       []float64 `json:"Position"`
		TargetActor    string    `json:"TargetActor"`
		TargetPosition []float64 `json:"TargetPosition"`
	} `json:"payload"`
}

type DealDamageEvent struct {
	Base
	Payload struct {
		Team         string `json:"Team"`
		Actor        string `json:"Actor"`
		Target       string `json:"Target"`
		Source       string `json:"Source"`
		Damage       int    `json:"Damage"`
		Dealt        int    `json:"Dealt"`
		IsHero       int    `json:"IsHero"`
		TargetIsHero int    `json:"TargetIsHero"`
	} `json:"payload"`
}

type HealTargetEvent struct {
	Base
	Payload struct {
		Team         string `json:"Team"`
		Actor        string `json:"Actor"`
		TargetActor  string `json:"TargetActor"`
		TargetTeam   string `json:"TargetTeam"`
		Source       string `json:"Source"`
		Heal         int    `json:"Heal"`
		Healed       int    `json:"Healed"`
		IsHero       int    `json:"IsHero"`
		TargetIsHero int    `json:"TargetIsHero"`
	} `json:"payload"`
}

type EarnXPEvent struct {
	Base
	Payload struct {
		Team       string `json:"Team"`
		Actor      string `json:"Actor"`
		Source     string `json:"Source"`
		Amount     int    `json:"Amount"`
		SharedWith int    `json:"SharedWith"`
	} `json:"payload"`
}

type KillActorEvent struct {
	Base
	Payload struct {
		Team         string    `json:"Team"`
		Actor        string    `json:"Actor"`
		Killed       string    `json:"Killed"`
		KilledTeam   string    `json:"KilledTeam"`
		Gold         string    `json:"Gold"`
		IsHero       int       `json:"IsHero"`
		TargetIsHero int       `json:"TargetIsHero"`
		Position     []float64 `json:"Position"`
	} `json:"payload"`
}

type VampirismEvent struct {
	Base
	Payload struct {
		Actor        string `json:"Actor"`
		Team         string `json:"Team"`
		TargetActor  string `json:"TargetActor"`
		TargetTeam   string `json:"TargetTeam"`
		Source       string `json:"Source"`
		Vamp         string `json:"Vamp"`
		IsHero       int    `json:"IsHero"`
		TargetIsHero int    `json:"TargetIsHero"`
	} `json:"payload"`
}

type GoldFromTowerKillEvent struct {
	Base
	Payload struct {
		Team   string `json:"Team"`
		Actor  string `json:"Actor"`
		Amount int    `json:"Amount"`
	} `json:"payload"`
}

type SellItemEvent struct {
	Base
	Payload struct {
		Team  string `json:"Team"`
		Actor string `json:"Actor"`
		Item  string `json:"Item"`
		Cost  int    `json:"Cost"`
	} `json:"payload"`
}

type NPCkillNPCEvent struct {
	Base
	Payload struct {
		Team         string    `json:"Team"`
		Actor        string    `json:"Actor"`
		Killed       string    `json:"Killed"`
		KilledTeam   string    `json:"KilledTeam"`
		Gold         string    `json:"Gold"`
		IsHero       int       `json:"IsHero"`
		TargetIsHero int       `json:"TargetIsHero"`
		Position     []float64 `json:"Position"`
	} `json:"payload"`
}
