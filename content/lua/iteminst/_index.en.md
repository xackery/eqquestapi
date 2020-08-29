---
title: Iteminst
menuTitle: Iteminst
description: Lua Item Instance Class
searchTitle: Lua Item Instance Class
weight: 25
---

## Item Instance Methods
- [AddExp](addexp)({{% lua_type_number %}} exp) -- {{% lua_type_nil %}}
- [ClearTimers](cleartimers)() -- {{% lua_type_nil %}}
- [Clone](clone)() -- {{% lua_type_iteminst %}}
- [DeleteCustomData](deletecustomdata)({{% lua_type_string %}} identifier) -- {{% lua_type_nil %}}
- [GetAugmentItemID](getaugmentitemid)({{% lua_type_number %}} slot) -- {{% lua_type_number %}}
- [GetAugmentType](getaugmenttype)() -- {{% lua_type_number %}}
- [GetCharges](getcharges)() -- {{% lua_type_number %}}
- [GetColor](getcolor)() -- {{% lua_type_number %}}
- [GetCustomDataString](getcustomdatastring)() -- string
- [GetExp](getexp)() -- {{% lua_type_number %}}
- [GetID](getid)() -- {{% lua_type_number %}}
- [GetItem](getitem)({{% lua_type_number %}} slot) -- {{% lua_type_item %}}
- [GetItemID](getitemid)({{% lua_type_number %}} slot) -- {{% lua_type_number %}}
- [GetItemScriptID](getitemscriptid)() -- {{% lua_type_number %}}
- [GetKillsNeeded](getkillsneeded)({{% lua_type_number %}} current_level) -- {{% lua_type_number %}}
- [GetMaxEvolveLvl](getmaxevolvelvl)() -- {{% lua_type_number %}}
- [GetPrice](getprice)() -- {{% lua_type_number %}}
- [GetTotalItemCount](gettotalitemcount)() -- {{% lua_type_number %}}
- [GetUnscaledItem](getunscaleditem)({{% lua_type_number %}} slot) -- {{% lua_type_item %}}
- [IsAmmo](isammo)() -- {{% lua_type_boolean %}}
- [IsAugmentable](isaugmentable)() -- {{% lua_type_boolean %}}
- [IsAugmented](isaugmented)() -- {{% lua_type_boolean %}}
- [IsEquipable](isequipable)({{% lua_type_number %}} slot_id) -- {{% lua_type_boolean %}}
- [IsExpendable](isexpendable)() -- {{% lua_type_boolean %}}
- [IsInstNoDrop](isinstnodrop)() -- {{% lua_type_boolean %}}
- [IsStackable](isstackable)() -- {{% lua_type_boolean %}}
- [IsType](istype)({{% lua_type_number %}} item_class) -- {{% lua_type_boolean %}}
- [IsWeapon](isweapon)() -- {{% lua_type_boolean %}}
- [SetCharges](setcharges)({{% lua_type_number %}} charges) -- {{% lua_type_nil %}}
- [SetColor](setcolor)({{% lua_type_number %}} color) -- {{% lua_type_nil %}}
- [SetCustomData](setcustomdata)({{% lua_type_string %}} identifier, {{% lua_type_boolean %}} value) -- {{% lua_type_nil %}}
- [SetExp](setexp)({{% lua_type_number %}} exp) -- {{% lua_type_nil %}}
- [SetInstNoDrop](setinstnodrop)({{% lua_type_boolean %}} flag) -- {{% lua_type_nil %}}
- [SetPrice](setprice)({{% lua_type_number %}} price) -- {{% lua_type_nil %}}
- [SetScale](setscale)({{% lua_type_number %}} scale_factor) -- {{% lua_type_nil %}}
- [SetScaling](setscaling)({{% lua_type_boolean %}} v) -- {{% lua_type_nil %}}
- [SetTimer](settimer)({{% lua_type_string %}} name, {{% lua_type_number %}} time) -- {{% lua_type_nil %}}
- [StopTimer](stoptimer)({{% lua_type_string %}} name) -- {{% lua_type_nil %}}
