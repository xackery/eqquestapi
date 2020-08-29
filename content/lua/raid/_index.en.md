---
title: Raid
menuTitle: Raid
description: Lua Raid Class
searchTitle: Lua Raid Class
weight: 25
---

## Raid Methods
- Raid:[IsRaidMember](israidmember)(const char *name); -- bool
- Raid:[CastGroupSpell](castgroupspell)(Lua_Mob caster, number spell_id, number group_id); -- void
- Raid:[GroupCount](groupcount)(number group_id); -- number
- Raid:[RaidCount](raidcount)(); -- number
- Raid:[GetGroup](getgroup)(Lua_Client c); -- number
- Raid:[SplitExp](splitexp)(number exp, Lua_Mob other); -- void
- Raid:[GetTotalRaidDamage](gettotalraiddamage)(Lua_Mob other); -- number
- Raid:[SplitMoney](splitmoney)(number gid, number copper, number silver, number gold, number platinum, Lua_Client splitter); -- void
- Raid:[BalanceHP](balancehp)(number penalty, number group_id); -- void
- Raid:[IsLeader](isleader)(Lua_Client c); -- bool
- Raid:[IsGroupLeader](isgroupleader)(const char *name); -- bool
- Raid:[GetHighestLevel](gethighestlevel)(); -- number
- Raid:[GetLowestLevel](getlowestlevel)(); -- number
- Raid:[GetClientByIndex](getclientbyindex)(number index); -- unknown - Lua_Client
- Raid:[TeleportGroup](teleportgroup)(Lua_Mob sender, number zone_id, number instance_id, float x, float y, float z, float h, number group_id); -- void
- Raid:[TeleportRaid](teleportraid)(Lua_Mob sender, number zone_id, number instance_id, float x, float y, float z, float h); -- void
- Raid:[GetID](getid)(); -- number
- Raid:[GetMember](getmember)(number index); -- unknown - Lua_Client
- Raid:[GetGroupNumber](getgroupnumber)(number index); -- number
