---
date: 2020-08-24T16:50:16+02:00
title: Raid
menuTitle: Raid
weight: 25
---

## Raid Methods (Lua)
- raid:[BalanceHP](balancehp)(int penalty, uint32 group_id); -- void
- raid:[CastGroupSpell](castgroupspell)(Lua_Mob caster, int spell_id, uint32 group_id); -- void
- raid:[GetClientByIndex](getclientbyindex)(int index); -- Lua_Client
- raid:[GetGroup](getgroup)(Lua_Client c); -- int
- raid:[GetGroup](getgroup)(const char *c); -- int
- raid:[GetGroupNumber](getgroupnumber)(int index); -- int
- raid:[GetHighestLevel](gethighestlevel)(); -- int
- raid:[GetID](getid)(); -- int
- raid:[GetLowestLevel](getlowestlevel)(); -- int
- raid:[GetMember](getmember)(int index); -- Lua_Client
- raid:[GetTotalRaidDamage](gettotalraiddamage)(Lua_Mob other); -- uint32
- raid:[GroupCount](groupcount)(uint32 group_id); -- int
- raid:[IsGroupLeader](isgroupleader)(const char *name); -- bool
- raid:[IsLeader](isleader)(Lua_Client c); -- bool
- raid:[IsLeader](isleader)(const char *c); -- bool
- raid:[IsRaidMember](israidmember)(const char *name); -- bool
- raid:[RaidCount](raidcount)(); -- int
- raid:[SplitExp](splitexp)(uint32 exp, Lua_Mob other); -- void
- raid:[SplitMoney](splitmoney)(uint32 copper, uint32 silver, uint32 gold, uint32 platinum); -- void
- raid:[SplitMoney](splitmoney)(uint32 copper, uint32 silver, uint32 gold, uint32 platinum, Lua_Client splitter); -- void
- raid:[TeleportGroup](teleportgroup)(Lua_Mob sender, uint32 zone_id, uint32 instance_id, float x, float y, float z, float h, uint32 group_id); -- void
- raid:[TeleportRaid](teleportraid)(Lua_Mob sender, uint32 zone_id, uint32 instance_id, float x, float y, float z, float h); -- void