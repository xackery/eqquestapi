---
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
- Corpse:[AddItem](additem)(number itemnum, uint16 charges, int16 slot, number aug1, number aug2, number aug3, number aug4, number aug5); -- void
- Corpse:[GetWornItem](getwornitem)(int16 equipSlot); -- number
- Corpse:[RemoveItem](removeitem)(uint16 lootslot); -- void
- Corpse:[SetCash](setcash)(number copper, number silver, number gold, number platinum); -- void
- Corpse:[RemoveCash](removecash)(); -- void
- Corpse:[IsEmpty](isempty)(); -- bool
- Corpse:[SetDecayTimer](setdecaytimer)(number decaytime); -- void
- Corpse:[CanMobLoot](canmobloot)(number charid); -- bool
- Corpse:[AllowMobLoot](allowmobloot)(Lua_Mob them, uint8 slot); -- void
- Corpse:[Summon](summon)(Lua_Client client, bool spell, bool checkdistance); -- bool
- Corpse:[GetCopper](getcopper)(); -- number
- Corpse:[GetSilver](getsilver)(); -- number
- Corpse:[GetGold](getgold)(); -- number
- Corpse:[GetPlatinum](getplatinum)(); -- number
- Corpse:[AddLooter](addlooter)(Lua_Mob who); -- void
