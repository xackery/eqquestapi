---
date: 2020-08-24T16:50:16+02:00
title: Iteminst
menuTitle: Iteminst
weight: 25
---

## Iteminst Methods (Lua)
- iteminst:[AddExp](addexp)(uint32 exp); -- void
- iteminst:[ClearTimers](cleartimers)(); -- void
- iteminst:[Clone](clone)(); -- Lua_ItemInst
- iteminst:[DeleteCustomData](deletecustomdata)(std::string identifier); -- void
- iteminst:[GetAugment](getaugment)(int slot); -- Lua_ItemInst
- iteminst:[GetAugmentItemID](getaugmentitemid)(int slot); -- uint32
- iteminst:[GetAugmentType](getaugmenttype)(); -- int
- iteminst:[GetCharges](getcharges)(); -- int
- iteminst:[GetColor](getcolor)(); -- uint32
- iteminst:[GetCustomData](getcustomdata)(std::string identifier); -- std::string
- iteminst:[GetCustomDataString](getcustomdatastring)(); -- std::string
- iteminst:[GetExp](getexp)(); -- uint32
- iteminst:[GetID](getid)(); -- uint32
- iteminst:[GetItem](getitem)(); -- Lua_Item
- iteminst:[GetItem](getitem)(int slot); -- Lua_ItemInst
- iteminst:[GetItemID](getitemid)(int slot); -- uint32
- iteminst:[GetItemScriptID](getitemscriptid)(); -- uint32
- iteminst:[GetKillsNeeded](getkillsneeded)(int current_level); -- uint32
- iteminst:[GetMaxEvolveLvl](getmaxevolvelvl)(); -- int
- iteminst:[GetPrice](getprice)(); -- uint32
- iteminst:[GetTotalItemCount](gettotalitemcount)(); -- int
- iteminst:[GetUnscaledItem](getunscaleditem)(int slot); -- Lua_Item
- iteminst:[IsAmmo](isammo)(); -- bool
- iteminst:[IsAugmentable](isaugmentable)(); -- bool
- iteminst:[IsAugmented](isaugmented)(); -- bool
- iteminst:[IsEquipable](isequipable)(int race, int class_); -- bool
- iteminst:[IsEquipable](isequipable)(int slot_id); -- bool
- iteminst:[IsExpendable](isexpendable)(); -- bool
- iteminst:[IsInstNoDrop](isinstnodrop)(); -- bool
- iteminst:[IsStackable](isstackable)(); -- bool
- iteminst:[IsType](istype)(int item_class); -- bool
- iteminst:[IsWeapon](isweapon)(); -- bool
- iteminst:[Lua_ItemInst](lua_iteminst)(const Lua_ItemInst& o); -- Lua_ItemInst::Lua_ItemInst(const
- iteminst:[SetCharges](setcharges)(int charges); -- void
- iteminst:[SetColor](setcolor)(uint32 color); -- void
- iteminst:[SetCustomData](setcustomdata)(std::string identifier, bool value); -- void
- iteminst:[SetCustomData](setcustomdata)(std::string identifier, float value); -- void
- iteminst:[SetCustomData](setcustomdata)(std::string identifier, int value); -- void
- iteminst:[SetCustomData](setcustomdata)(std::string identifier, std::string value); -- void
- iteminst:[SetExp](setexp)(uint32 exp); -- void
- iteminst:[SetInstNoDrop](setinstnodrop)(bool flag); -- void
- iteminst:[SetPrice](setprice)(uint32 price); -- void
- iteminst:[SetScale](setscale)(double scale_factor); -- void
- iteminst:[SetScaling](setscaling)(bool v); -- void
- iteminst:[SetTimer](settimer)(std::string name, uint32 time); -- void
- iteminst:[StopTimer](stoptimer)(std::string name); -- void
- iteminst:[operator=](operator=)(const Lua_ItemInst& o); -- Lua_ItemInst&