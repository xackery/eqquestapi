---
date: 2020-08-24T16:50:16+02:00
title: Corpse
menuTitle: Corpse
weight: 25
---

## Corpse Methods (Perl)
- $corpse->[AddItem](additem)(uint32 item_id, uint16 charges, [unt16 slot = 0])
- $corpse->[AddLooter](addlooter)(mob* who)
- $corpse->[AllowMobLoot](allowmobloot)(mob* them, uint8 slot)
- $corpse->[CanMobLoot](canmobloot)(int character_id)
- $corpse->[CastRezz](castrezz)(uint16 spell_id, [mob* caster = nullptr])
- $corpse->[CompleteRezz](completerezz)()
- $corpse->[CountItems](countitems)()
- $corpse->[Delete](delete)()
- $corpse->[GetCharID](getcharid)()
- $corpse->[GetCopper](getcopper)()
- $corpse->[GetDBID](getdbid)()
- $corpse->[GetDecayTime](getdecaytime)()
- $corpse->[GetGold](getgold)()
- $corpse->[GetOwnerName](getownername)()
- $corpse->[GetPlatinum](getplatinum)()
- $corpse->[GetSilver](getsilver)()
- $corpse->[GetWornItem](getwornitem)(equipslot)
- $corpse->[IsEmpty](isempty)()
- $corpse->[IsLocked](islocked)()
- $corpse->[IsRezzed](isrezzed)()
- $corpse->[Lock](lock)()
- $corpse->[RemoveCash](removecash)()
- $corpse->[RemoveItem](removeitem)(uint16 loot_slot)
- $corpse->[ResetLooter](resetlooter)()
- $corpse->[SetCash](setcash)(uint16 copper, uint16 silver, uint16 gold, uint16 platinum)
- $corpse->[SetDecayTimer](setdecaytimer)(uint32 decay_time)
- $corpse->[Summon](summon)(client* client, bool is_spell)
- $corpse->[UnLock](unlock)()