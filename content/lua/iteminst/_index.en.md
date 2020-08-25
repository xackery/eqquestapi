---
date: 2020-08-24T16:50:16+02:00
title: Iteminst
menuTitle: Iteminst
weight: 25
---

## Iteminst Methods (Lua)
- ItemInst:[IsType](istype)(int item_class); -- bool
- ItemInst:[IsStackable](isstackable)(); -- bool
- ItemInst:[IsEquipable](isequipable)(int slot_id); -- bool
- ItemInst:[IsAugmentable](isaugmentable)(); -- bool
- ItemInst:[GetAugmentType](getaugmenttype)(); -- number
- ItemInst:[IsExpendable](isexpendable)(); -- bool
- ItemInst:[GetItem](getitem)(int slot); -- unknown - Lua_Item
- ItemInst:[GetUnscaledItem](getunscaleditem)(int slot); -- unknown - Lua_Item
- ItemInst:[GetItemID](getitemid)(int slot); -- number
- ItemInst:[GetTotalItemCount](gettotalitemcount)(); -- number
- ItemInst:[GetAugmentItemID](getaugmentitemid)(int slot); -- number
- ItemInst:[IsAugmented](isaugmented)(); -- bool
- ItemInst:[IsWeapon](isweapon)(); -- bool
- ItemInst:[IsAmmo](isammo)(); -- bool
- ItemInst:[GetID](getid)(); -- number
- ItemInst:[GetItemScriptID](getitemscriptid)(); -- number
- ItemInst:[GetCharges](getcharges)(); -- number
- ItemInst:[SetCharges](setcharges)(int charges); -- void
- ItemInst:[GetPrice](getprice)(); -- number
- ItemInst:[SetPrice](setprice)(uint32 price); -- void
- ItemInst:[SetColor](setcolor)(uint32 color); -- void
- ItemInst:[GetColor](getcolor)(); -- number
- ItemInst:[IsInstNoDrop](isinstnodrop)(); -- bool
- ItemInst:[SetInstNoDrop](setinstnodrop)(bool flag); -- void
- ItemInst:[GetCustomDataString](getcustomdatastring)(); -- string
- ItemInst:[SetCustomData](setcustomdata)(std::string identifier, bool value); -- void
- ItemInst:[DeleteCustomData](deletecustomdata)(std::string identifier); -- void
- ItemInst:[SetScaling](setscaling)(bool v); -- void
- ItemInst:[SetScale](setscale)(double scale_factor); -- void
- ItemInst:[GetExp](getexp)(); -- number
- ItemInst:[SetExp](setexp)(uint32 exp); -- void
- ItemInst:[AddExp](addexp)(uint32 exp); -- void
- ItemInst:[GetMaxEvolveLvl](getmaxevolvelvl)(); -- number
- ItemInst:[GetKillsNeeded](getkillsneeded)(int current_level); -- number
- ItemInst:[Clone](clone)(); -- unknown - Lua_ItemInst
- ItemInst:[SetTimer](settimer)(std::string name, uint32 time); -- void
- ItemInst:[StopTimer](stoptimer)(std::string name); -- void
- ItemInst:[ClearTimers](cleartimers)(); -- void
