---
title: Raid Class
menuTitle: Raid Class
description: Lua Raid Class
searchTitle: Lua Raid Class
weight: 25
---

## Raid Methods
- [IsRaidMember](israidmember)(const char *name); -- bool
- [CastGroupSpell](castgroupspell)(Lua_Mob caster, number spell_id, number group_id); -- void
- [GroupCount](groupcount)(number group_id); -- number
- [RaidCount](raidcount)(); -- number
- [GetGroup](getgroup)(Lua_Client c); -- number
- [SplitExp](splitexp)(number exp, Lua_Mob other); -- void
- [GetTotalRaidDamage](gettotalraiddamage)(Lua_Mob other); -- number
- [SplitMoney](splitmoney)(number gid, number copper, number silver, number gold, number platinum, Lua_Client splitter); -- void
- [BalanceHP](balancehp)(number penalty, number group_id); -- void
- [IsLeader](isleader)(Lua_Client c); -- bool
- [IsGroupLeader](isgroupleader)(const char *name); -- bool
- [GetHighestLevel](gethighestlevel)(); -- number
- [GetLowestLevel](getlowestlevel)(); -- number
- [GetClientByIndex](getclientbyindex)(number index); -- unknown - Lua_Client
- [TeleportGroup](teleportgroup)(Lua_Mob sender, number zone_id, number instance_id, float x, float y, float z, float h, number group_id); -- void
- [TeleportRaid](teleportraid)(Lua_Mob sender, number zone_id, number instance_id, float x, float y, float z, float h); -- void
- [GetID](getid)(); -- number
- [GetMember](getmember)(number index); -- unknown - Lua_Client
- [GetGroupNumber](getgroupnumber)(number index); -- number
