---
date: 2020-08-24T16:50:16+02:00
title: Corpse
menuTitle: Corpse
weight: 25
---

## Corpse Methods (Lua)
- corpse:[AddItem](additem)(uint32 itemnum, uint16 charges, int16 slot, uint32 aug1, uint32 aug2, uint32 aug3, uint32 aug4, uint32 aug5); -- void
- corpse:[AddLooter](addlooter)(Lua_Mob who); -- void
- corpse:[AllowMobLoot](allowmobloot)(Lua_Mob them, uint8 slot); -- void
- corpse:[Bury](bury)(); -- void
- corpse:[CanMobLoot](canmobloot)(int charid); -- bool
- corpse:[CountItems](countitems)(); -- uint32
- corpse:[Delete](delete)(); -- void
- corpse:[Depop](depop)(); -- void
- corpse:[GetCharID](getcharid)(); -- uint32
- corpse:[GetCopper](getcopper)(); -- uint32
- corpse:[GetDBID](getdbid)(); -- uint32
- corpse:[GetDecayTime](getdecaytime)(); -- uint32
- corpse:[GetGold](getgold)(); -- uint32
- corpse:[GetPlatinum](getplatinum)(); -- uint32
- corpse:[GetSilver](getsilver)(); -- uint32
- corpse:[GetWornItem](getwornitem)(int16 equipSlot); -- uint32
- corpse:[IsEmpty](isempty)(); -- bool
- corpse:[IsLocked](islocked)(); -- bool
- corpse:[IsRezzed](isrezzed)(); -- bool
- corpse:[Lock](lock)(); -- void
- corpse:[RemoveCash](removecash)(); -- void
- corpse:[RemoveItem](removeitem)(uint16 lootslot); -- void
- corpse:[ResetLooter](resetlooter)(); -- void
- corpse:[Save](save)(); -- bool
- corpse:[SetCash](setcash)(uint32 copper, uint32 silver, uint32 gold, uint32 platinum); -- void
- corpse:[SetDecayTimer](setdecaytimer)(uint32 decaytime); -- void
- corpse:[Summon](summon)(Lua_Client client, bool spell, bool checkdistance); -- bool
- corpse:[UnLock](unlock)(); -- void