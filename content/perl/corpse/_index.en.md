---
searchTitle: Perl Corpse Class
title: Corpse Class
menuTitle: Corpse Class
weight: 25
---

## Corpse Methods
- [AddItem](additem)(uint32 item_id, uint16 charges, [unt16 slot = 0])
- [AddLooter](addlooter)(mob* who)
- [AllowMobLoot](allowmobloot)(mob* them, uint8 slot)
- [CanMobLoot](canmobloot)(int character_id)
- [CastRezz](castrezz)(uint16 spell_id, [mob* caster = nullptr])
- [CompleteRezz](completerezz)()
- [CountItems](countitems)()
- [Delete](delete)()
- [GetCharID](getcharid)()
- [GetCopper](getcopper)()
- [GetDBID](getdbid)()
- [GetDecayTime](getdecaytime)()
- [GetGold](getgold)()
- [GetOwnerName](getownername)()
- [GetPlatinum](getplatinum)()
- [GetSilver](getsilver)()
- [GetWornItem](getwornitem)(equipslot)
- [IsEmpty](isempty)()
- [IsLocked](islocked)()
- [IsRezzed](isrezzed)()
- [Lock](lock)()
- [RemoveCash](removecash)()
- [RemoveItem](removeitem)(uint16 loot_slot)
- [ResetLooter](resetlooter)()
- [SetCash](setcash)(uint16 copper, uint16 silver, uint16 gold, uint16 platinum)
- [SetDecayTimer](setdecaytimer)(uint32 decay_time)
- [Summon](summon)(client* client, bool is_spell)
- [UnLock](unlock)()