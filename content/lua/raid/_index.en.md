---
title: Raid Class
menuTitle: Raid Class
searchDescription: Lua Raid Class
searchTitle: Lua Raid Class
weight: 25
---

## Raid Methods
- [BalanceHP](balancehp)(number penalty, number group_id); -- void
- [CastGroupSpell](castgroupspell)(Lua_Mob caster, number spell_id, number group_id); -- void
- [DoesAnyMemberHaveExpeditionLockout](doesanymemberhaveexpeditionlockout)(std::string expedition_name, std::string event_name, number max_check_count); -- bool
- [GetClientByIndex](getclientbyindex)(number index); -- unknown - Lua_Client
- [GetGroup](getgroup)(Lua_Client c); -- number
- [GetGroupNumber](getgroupnumber)(number index); -- number
- [GetHighestLevel](gethighestlevel)(); -- number
- [GetID](getid)(); -- number
- [GetLowestLevel](getlowestlevel)(); -- number
- [GetMember](getmember)(number index); -- unknown - Lua_Client
- [GetTotalRaidDamage](gettotalraiddamage)(Lua_Mob other); -- number
- [GroupCount](groupcount)(number group_id); -- number
- [IsGroupLeader](isgroupleader)(const char *name); -- bool
- [IsLeader](isleader)(Lua_Client c); -- bool
- [IsRaidMember](israidmember)(const char *name); -- bool
- [RaidCount](raidcount)(); -- number
- [SplitExp](splitexp)(number exp, Lua_Mob other); -- void
- [SplitMoney](splitmoney)(number gid, number copper, number silver, number gold, number platinum, Lua_Client splitter); -- void
- [TeleportGroup](teleportgroup)(Lua_Mob sender, number zone_id, number instance_id, float x, float y, float z, float h, number group_id); -- void
- [TeleportRaid](teleportraid)(Lua_Mob sender, number zone_id, number instance_id, float x, float y, float z, float h); -- void
