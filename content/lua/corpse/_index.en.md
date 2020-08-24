---
date: 2016-04-09T16:50:16+02:00
title: Corpse
menuTitle: Corpse
weight: 25
---

# Corpse Methods \(Lua\)

- corpse:[AddItem(uint32 itemnum, uint16 charges, int16 slot, uint32 aug1, uint32 aug2, uint32 aug3, uint32 aug4, uint32 aug5)](additem); -- void
- corpse:[AddLooter(Lua_Mob who)](addlooter); -- void
- corpse:[AllowMobLoot(Lua_Mob them, uint8 slot)](allowmobloot); -- void
- corpse:[Bury()](bury); -- void
- corpse:[CanMobLoot(int charid)](canmobloot); -- bool
- corpse:[CountItems()](countitems); -- uint32
- corpse:[Delete()](delete); -- void
- corpse:[Depop()](depop); -- void
- corpse:[GetCharID()](getcharid); -- uint32
- corpse:[GetCopper()](getcopper); -- uint32
- corpse:[GetDBID()](getdbid); -- uint32
- corpse:[GetDecayTime()](getdecaytime); -- uint32
- corpse:[GetGold()](getgold); -- uint32
- corpse:[GetPlatinum()](getplatinum); -- uint32
- corpse:[GetSilver()](getsilver); -- uint32
- corpse:[GetWornItem(int16 equipSlot)](getwornitem); -- uint32
- corpse:[IsEmpty()](isempty); -- bool
- corpse:[IsLocked()](islocked); -- bool
- corpse:[IsRezzed()](isrezzed); -- bool
- corpse:[Lock()](lock); -- void
- corpse:[RemoveCash()](removecash); -- void
- corpse:[RemoveItem(uint16 lootslot)](removeitem); -- void
- corpse:[ResetLooter()](resetlooter); -- void
- corpse:[Save()](save); -- bool
- corpse:[SetCash(uint32 copper, uint32 silver, uint32 gold, uint32 platinum)](setcash); -- void
- corpse:[SetDecayTimer(uint32 decaytime)](setdecaytimer); -- void
- corpse:[Summon(Lua_Client client, bool spell, bool checkdistance)](summon); -- bool
- corpse:[UnLock()](unlock); -- void