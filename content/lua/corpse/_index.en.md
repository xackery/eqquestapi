---
title: Corpse
menuTitle: Corpse
description: Lua Corpse Method List
weight: 25
---

## Corpse Methods (Lua)
- [AddItem](additem)({{% lua_type_number %}} itemnum, {{% lua_type_number %}} charges, {{% lua_type_number %}} slot, {{% lua_type_number %}} aug1, {{% lua_type_number %}} aug2, {{% lua_type_number %}} aug3, {{% lua_type_number %}} aug4, {{% lua_type_number %}} aug5) -- {{% lua_type_nil %}}
- [AddLooter](addlooter)({{% lua_type_mob %}} who) -- {{% lua_type_nil %}}
- [AllowMobLoot](allowmobloot)({{% lua_type_mob %}} them, {{% lua_type_number %}} slot) -- {{% lua_type_nil %}}
- [Bury](bury)() -- {{% lua_type_nil %}}
- [CanMobLoot](canmobloot)({{% lua_type_number %}} charid) -- {{% lua_type_boolean %}}
- [CountItems](countitems)() -- {{% lua_type_number %}}
- [Delete](delete)() -- {{% lua_type_nil %}}
- [Depop](depop)() -- {{% lua_type_nil %}}
- [GetCharID](getcharid)() -- {{% lua_type_number %}}
- [GetCopper](getcopper)() -- {{% lua_type_number %}}
- [GetDBID](getdbid)() -- {{% lua_type_number %}}
- [GetDecayTime](getdecaytime)() -- {{% lua_type_number %}}
- [GetGold](getgold)() -- {{% lua_type_number %}}
- [GetOwnerName](getownername) -- {{% lua_type_string %}}
- [GetPlatinum](getplatinum)() -- {{% lua_type_number %}}
- [GetSilver](getsilver)() -- {{% lua_type_number %}}
- [GetWornItem](getwornitem)({{% lua_type_number %}} equipSlot) -- {{% lua_type_number %}}
- [IsEmpty](isempty)() -- {{% lua_type_boolean %}}
- [IsLocked](islocked)() -- {{% lua_type_boolean %}}
- [IsRezzed](isrezzed)() -- {{% lua_type_boolean %}}
- [Lock](lock)() -- {{% lua_type_nil %}}
- [RemoveCash](removecash)() -- {{% lua_type_nil %}}
- [RemoveItem](removeitem)({{% lua_type_number %}} lootslot) -- {{% lua_type_nil %}}
- [ResetLooter](resetlooter)() -- {{% lua_type_nil %}}
- [Save](save)() -- {{% lua_type_boolean %}}
- [SetCash](setcash)({{% lua_type_number %}} copper, {{% lua_type_number %}} silver, {{% lua_type_number %}} gold, {{% lua_type_number %}} platinum) -- {{% lua_type_nil %}}
- [SetDecayTimer](setdecaytimer)({{% lua_type_number %}} decaytime) -- {{% lua_type_nil %}}
- [Summon](summon)({{% lua_type_client %}} client, {{% lua_type_boolean %}} spell, {{% lua_type_boolean %}} checkdistance) -- {{% lua_type_boolean %}}
- [UnLock](unlock)() -- {{% lua_type_nil %}}
