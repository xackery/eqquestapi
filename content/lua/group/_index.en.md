---
title: Group Class
menuTitle: Group Class
searchDescription: Lua Group Class
searchTitle: Lua Group Class
weight: 25
---

## Group Methods
- [CastGroupSpell](castgroupspell)({{% type_npc mob %}} caster, {{% lua_type_number %}} spell_id) -- {{% lua_type_nil %}}
- [DisbandGroup](disbandgroup)() -- {{% lua_type_nil %}}
- [DoesAnyMemberHaveExpeditionLockout](doesanymemberhaveexpeditionlockout)(std::string expedition_name, std::string event_name, number max_check_count); -- bool
- [GetHighestLevel](gethighestlevel)() -- {{% lua_type_number %}}
- [GetID](getid)() -- {{% lua_type_number %}}
- [GetLeader](getleader)() -- {{% type_npc mob %}}
- [GetLeaderName](getleadername)() -- {{% lua_type_string %}}
- [GetLowestLevel](getlowestlevel)() -- {{% lua_type_number %}}
- [GetMember](getmember)({{% lua_type_number %}} index) -- {{% type_npc mob %}}
- [GetTotalGroupDamage](gettotalgroupdamage)({{% type_npc mob %}} other) -- {{% lua_type_number %}}
- [GroupCount](groupcount)() -- {{% lua_type_number %}}
- [GroupMessage](groupmessage)({{% type_npc mob %}} sender, {{% lua_type_number %}} language, {{% lua_type_string %}} message) -- {{% lua_type_nil %}}
- [IsGroupMember](isgroupmember)({{% type_npc mob %}} mob) -- {{% lua_type_boolean %}}
- [IsLeader](isleader)({{% type_npc mob %}} leader) -- {{% lua_type_boolean %}}
- [SetLeader](setleader)({{% type_npc mob %}} leader) -- {{% lua_type_nil %}}
- [SplitExp](splitexp)({{% lua_type_number %}} exp, {{% type_npc mob %}} other) -- {{% lua_type_nil %}}
- [SplitMoney](splitmoney)({{% lua_type_number %}} copper, {{% lua_type_number %}} silver, {{% lua_type_number %}} gold, {{% lua_type_number %}} platinum, {{% type_client lua %}} splitter) -- {{% lua_type_nil %}}
- [TeleportGroup](teleportgroup)({{% type_npc mob %}} sender, {{% lua_type_number %}} zone_id, {{% lua_type_number %}} instance_id, {{% lua_type_number %}} x, {{% lua_type_number %}} y, {{% lua_type_number %}} z, {{% lua_type_number %}} h) -- {{% lua_type_nil %}}
