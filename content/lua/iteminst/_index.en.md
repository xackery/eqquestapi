---
date: 2020-08-24T16:50:16+02:00
title: Iteminst
menuTitle: Iteminst
weight: 25
---

## Iteminst Methods (Lua)
- ItemInst:[IsType](istype)(number item_class); -- bool
- ItemInst:[IsStackable](isstackable)(); -- bool
- ItemInst:[IsEquipable](isequipable)(number slot_id); -- bool
- ItemInst:[IsAugmentable](isaugmentable)(); -- bool
- ItemInst:[GetAugmentType](getaugmenttype)(); -- number
- ItemInst:[IsExpendable](isexpendable)(); -- bool
- ItemInst:[GetItem](getitem)(number slot); -- unknown - Lua_Item
- ItemInst:[GetUnscaledItem](getunscaleditem)(number slot); -- unknown - Lua_Item
- ItemInst:[GetItemID](getitemid)(number slot); -- number
- ItemInst:[GetTotalItemCount](gettotalitemcount)(); -- number
- ItemInst:[GetAugmentItemID](getaugmentitemid)(number slot); -- number
- ItemInst:[IsAugmented](isaugmented)(); -- bool
- ItemInst:[IsWeapon](isweapon)(); -- bool
- ItemInst:[IsAmmo](isammo)(); -- bool
- ItemInst:[GetID](getid)(); -- number
- ItemInst:[GetItemScriptID](getitemscriptid)(); -- number
- ItemInst:[GetCharges](getcharges)(); -- number
- ItemInst:[SetCharges](setcharges)(number charges); -- void
- ItemInst:[GetPrice](getprice)(); -- number
- ItemInst:[SetPrice](setprice)(number price); -- void
- ItemInst:[SetColor](setcolor)(number color); -- void
- ItemInst:[GetColor](getcolor)(); -- number
- ItemInst:[IsInstNoDrop](isinstnodrop)(); -- bool
- ItemInst:[SetInstNoDrop](setinstnodrop)(bool flag); -- void
- ItemInst:[GetCustomDataString](getcustomdatastring)(); -- string
- ItemInst:[SetCustomData](setcustomdata)(std::string identifier, bool value); -- void
- ItemInst:[DeleteCustomData](deletecustomdata)(std::string identifier); -- void
- ItemInst:[SetScaling](setscaling)(bool v); -- void
- ItemInst:[SetScale](setscale)(double scale_factor); -- void
- ItemInst:[GetExp](getexp)(); -- number
- ItemInst:[SetExp](setexp)(number exp); -- void
- ItemInst:[AddExp](addexp)(number exp); -- void
- ItemInst:[GetMaxEvolveLvl](getmaxevolvelvl)(); -- number
- ItemInst:[GetKillsNeeded](getkillsneeded)(number current_level); -- number
- ItemInst:[Clone](clone)(); -- unknown - Lua_ItemInst
- ItemInst:[SetTimer](settimer)(std::string name, number time); -- void
- ItemInst:[StopTimer](stoptimer)(std::string name); -- void
- ItemInst:[ClearTimers](cleartimers)(); -- void
