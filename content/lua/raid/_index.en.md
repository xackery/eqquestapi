---
date: 2020-08-24T16:50:16+02:00
title: Raid
menuTitle: Raid
weight: 25
---

## Raid Methods (Lua)
- Raid:[IsRaidMember](israidmember)(const char *name); -- bool
- Raid:[CastGroupSpell](castgroupspell)(Lua_Mob caster, number spell_id, uint32 group_id); -- void
- Raid:[GroupCount](groupcount)(uint32 group_id); -- number
- Raid:[RaidCount](raidcount)(); -- number
- Raid:[GetGroup](getgroup)(Lua_Client c); -- number
- Raid:[SplitExp](splitexp)(uint32 exp, Lua_Mob other); -- void
- Raid:[GetTotalRaidDamage](gettotalraiddamage)(Lua_Mob other); -- number
- Raid:[SplitMoney](splitmoney)(uint32 gid, uint32 copper, uint32 silver, uint32 gold, uint32 platinum, Lua_Client splitter); -- void
- Raid:[BalanceHP](balancehp)(number penalty, uint32 group_id); -- void
- Raid:[IsLeader](isleader)(Lua_Client c); -- bool
- Raid:[IsGroupLeader](isgroupleader)(const char *name); -- bool
- Raid:[GetHighestLevel](gethighestlevel)(); -- number
- Raid:[GetLowestLevel](getlowestlevel)(); -- number
- Raid:[GetClientByIndex](getclientbyindex)(number index); -- unknown - Lua_Client
- Raid:[TeleportGroup](teleportgroup)(Lua_Mob sender, uint32 zone_id, uint32 instance_id, float x, float y, float z, float h, uint32 group_id); -- void
- Raid:[TeleportRaid](teleportraid)(Lua_Mob sender, uint32 zone_id, uint32 instance_id, float x, float y, float z, float h); -- void
- Raid:[GetID](getid)(); -- number
- Raid:[GetMember](getmember)(number index); -- unknown - Lua_Client
- Raid:[GetGroupNumber](getgroupnumber)(number index); -- number
