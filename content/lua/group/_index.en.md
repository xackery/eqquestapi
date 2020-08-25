---
date: 2020-08-24T16:50:16+02:00
title: Group
menuTitle: Group
weight: 25
---

## Group Methods (Lua)
- group:[CastGroupSpell](castgroupspell)(Lua_Mob caster, int spell_id); -- void
- group:[DisbandGroup](disbandgroup)(); -- void
- group:[GetHighestLevel](gethighestlevel)(); -- int
- group:[GetID](getid)(); -- int
- group:[GetLeader](getleader)(); -- Lua_Mob
- group:[GetLowestLevel](getlowestlevel)(); -- int
- group:[GetMember](getmember)(int index); -- Lua_Mob
- group:[GetTotalGroupDamage](gettotalgroupdamage)(Lua_Mob other); -- uint32
- group:[GroupCount](groupcount)(); -- int
- group:[GroupMessage](groupmessage)(Lua_Mob sender, int language, const char *message); -- void
- group:[IsGroupMember](isgroupmember)(Lua_Mob mob); -- bool
- group:[IsLeader](isleader)(Lua_Mob leader); -- bool
- group:[SetLeader](setleader)(Lua_Mob leader); -- void
- group:[SplitExp](splitexp)(uint32 exp, Lua_Mob other); -- void
- group:[SplitMoney](splitmoney)(uint32 copper, uint32 silver, uint32 gold, uint32 platinum); -- void
- group:[SplitMoney](splitmoney)(uint32 copper, uint32 silver, uint32 gold, uint32 platinum, Lua_Client splitter); -- void
- group:[TeleportGroup](teleportgroup)(Lua_Mob sender, uint32 zone_id, uint32 instance_id, float x, float y, float z, float h); -- void