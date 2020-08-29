---
title: Group
menuTitle: Group
description: Lua Group Class
searchTitle: Lua Group Class
weight: 25
---

## Group Methods
- [DisbandGroup](disbandgroup)() -- {{% lua_type_nil %}}
- [IsGroupMember](isgroupmember)({{% lua_type_mob %}} mob) -- {{% lua_type_boolean %}}
- [CastGroupSpell](castgroupspell)({{% lua_type_mob %}} caster, {{% lua_type_number %}} spell_id) -- {{% lua_type_nil %}}
- [SplitExp](splitexp)({{% lua_type_number %}} exp, {{% lua_type_mob %}} other) -- {{% lua_type_nil %}}
- [GroupMessage](groupmessage)({{% lua_type_mob %}} sender, {{% lua_type_number %}} language, {{% lua_type_string %}} message) -- {{% lua_type_nil %}}
- [GetTotalGroupDamage](gettotalgroupdamage)({{% lua_type_mob %}} other) -- {{% lua_type_number %}}
- [SplitMoney](splitmoney)({{% lua_type_number %}} copper, {{% lua_type_number %}} silver, {{% lua_type_number %}} gold, {{% lua_type_number %}} platinum, {{% lua_type_client %}} splitter) -- {{% lua_type_nil %}}
- [SetLeader](setleader)({{% lua_type_mob %}} leader) -- {{% lua_type_nil %}}
- [GetLeader](getleader)() -- {{% lua_type_mob %}}
- [GetLeaderName](getleadername)() -- {{% lua_type_string %}}
- [IsLeader](isleader)({{% lua_type_mob %}} leader) -- {{% lua_type_boolean %}}
- [GroupCount](groupcount)() -- {{% lua_type_number %}}
- [GetHighestLevel](gethighestlevel)() -- {{% lua_type_number %}}
- [GetLowestLevel](getlowestlevel)() -- {{% lua_type_number %}}
- [TeleportGroup](teleportgroup)({{% lua_type_mob %}} sender, {{% lua_type_number %}} zone_id, {{% lua_type_number %}} instance_id, {{% lua_type_number %}} x, {{% lua_type_number %}} y, {{% lua_type_number %}} z, {{% lua_type_number %}} h) -- {{% lua_type_nil %}}
- [GetID](getid)() -- {{% lua_type_number %}}
- [GetMember](getmember)({{% lua_type_number %}} index) -- {{% lua_type_mob %}}
