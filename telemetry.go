package vainglory

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

// unmarshalEvent takes in the raw json and determines which event to unmarshal
// the json into. Additionally, it adds a pointer to the event to some helper
// slices to allow users to quickly parse specific events. Everything is kept
// in chronological order.
func (tr *TelemetryResponse) unmarshalEvent(js []byte, t string) {
	switch t {
	case "HeroBan":
		v := HeroBanEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.HeroBanEvents = append(tr.HeroBanEvents, &v)
	case "HeroSelect":
		v := HeroSelectEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.HeroSelectEvents = append(tr.HeroSelectEvents, &v)
	case "HeroSkinSelect":
		v := HeroSkinSelectEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.HeroSkinSelectEvents = append(tr.HeroSkinSelectEvents, &v)
	case "PlayerFirstSpawn":
		v := PlayerFirstSpawnEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.PlayerFirstSpawnEvents = append(tr.PlayerFirstSpawnEvents, &v)
	case "LevelUp":
		v := LevelUpEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.LevelUpEvents = append(tr.LevelUpEvents, &v)
	case "BuyItem":
		v := BuyItemEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.BuyItemEvents = append(tr.BuyItemEvents, &v)
	case "LearnAbility":
		v := LearnAbilityEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.LearnAbilityEvents = append(tr.LearnAbilityEvents, &v)
	case "UseAbility":
		v := UseAbilityEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.UseAbilityEvents = append(tr.UseAbilityEvents, &v)
	case "UseItemAbility":
		v := UseItemAbilityEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.UseItemAbilityEvents = append(tr.UseItemAbilityEvents, &v)
	case "DealDamage":
		v := DealDamageEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.DealDamageEvents = append(tr.DealDamageEvents, &v)
	case "HealTarget":
		v := HealTargetEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.HealTargetEvents = append(tr.HealTargetEvents, &v)
	case "EarnXP":
		v := EarnXPEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.EarnXPEvents = append(tr.EarnXPEvents, &v)
	case "KillActor":
		v := KillActorEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.KillActorEvents = append(tr.KillActorEvents, &v)
	case "Vampirism":
		v := VampirismEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.VampirismEvents = append(tr.VampirismEvents, &v)
	case "GoldFromTowerKill":
		v := GoldFromTowerKillEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.GoldFromTowerKillEvents = append(tr.GoldFromTowerKillEvents, &v)
	case "SellItem":
		v := SellItemEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.SellItemEvents = append(tr.SellItemEvents, &v)
	case "NPCkillNPC":
		v := NPCkillNPCEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.NPCkillNPCEvents = append(tr.NPCkillNPCEvents, &v)
	}
	return
}

// ToFile will save a TelemetryResponse to the file at a specified location.
// These data are always static and so it makes sense to cache/save this somewhere
// locally to prevent from having to request the large file multiple times
func (tr *TelemetryResponse) ToFile(path string) (err error) {
	var b []byte
	b, err = json.Marshal(tr.Events)
	if err != nil {
		return
	}
	var pretty bytes.Buffer
	err = json.Indent(&pretty, b, "", "\t")
	if err != nil {
		return
	}
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	w := bufio.NewWriter(f)
	_, err = w.Write(pretty.Bytes())
	if err != nil {
		fmt.Println(err.Error())
	}
	pretty.Reset()
	err = w.Flush()
	if err != nil {
		fmt.Println(err.Error())
	}
	f.Close()
	return
}
