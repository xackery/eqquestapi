---
date: 2020-08-24T16:50:16+02:00
title: Corpse
menuTitle: Corpse
weight: 25
---

## Corpse Methods (Lua)
- Corpse:[GetCharID](getcharid)(); -- number
- Corpse:[GetDecayTime](getdecaytime)(); -- number
- Corpse:[Lock](lock)(); -- void
- Corpse:[UnLock](unlock)(); -- void
- Corpse:[IsLocked](islocked)(); -- bool
- Corpse:[ResetLooter](resetlooter)(); -- void
- Corpse:[GetDBID](getdbid)(); -- number
- Corpse:[IsRezzed](isrezzed)(); -- bool
- Corpse:[GetOwnerName](getownername); -- 
- Corpse:[Save](save)(); -- bool
- Corpse:[Delete](delete)(); -- void
- Corpse:[Bury](bury)(); -- void
- Corpse:[Depop](depop)(); -- void
- Corpse:[CountItems](countitems)(); -- number
- Corpse:[AddItem](additem)(uint32 itemnum, uint16 charges, int16 slot, uint32 aug1, uint32 aug2, uint32 aug3, uint32 aug4, uint32 aug5); -- void
- Corpse:[GetWornItem](getwornitem)(int16 equipSlot); -- number
- Corpse:[RemoveItem](removeitem)(uint16 lootslot); -- void
- Corpse:[SetCash](setcash)(uint32 copper, uint32 silver, uint32 gold, uint32 platinum); -- void
- Corpse:[RemoveCash](removecash)(); -- void
- Corpse:[IsEmpty](isempty)(); -- bool
- Corpse:[SetDecayTimer](setdecaytimer)(uint32 decaytime); -- void
- Corpse:[CanMobLoot](canmobloot)(int charid); -- bool
- Corpse:[AllowMobLoot](allowmobloot)(Lua_Mob them, uint8 slot); -- void
- Corpse:[Summon](summon)(Lua_Client client, bool spell, bool checkdistance); -- bool
- Corpse:[GetCopper](getcopper)(); -- number
- Corpse:[GetSilver](getsilver)(); -- number
- Corpse:[GetGold](getgold)(); -- number
- Corpse:[GetPlatinum](getplatinum)(); -- number
- Corpse:[AddLooter](addlooter)(Lua_Mob who); -- void
