---
title: Group
menuTitle: Group
weight: 25
---

## Group Methods (Lua)
- Group:[DisbandGroup](disbandgroup)(); -- void
- Group:[IsGroupMember](isgroupmember)(Lua_Mob mob); -- bool
- Group:[CastGroupSpell](castgroupspell)(Lua_Mob caster, number spell_id); -- void
- Group:[SplitExp](splitexp)(number exp, Lua_Mob other); -- void
- Group:[GroupMessage](groupmessage)(Lua_Mob sender, number language, const char *message); -- void
- Group:[GetTotalGroupDamage](gettotalgroupdamage)(Lua_Mob other); -- number
- Group:[SplitMoney](splitmoney)(number copper, number silver, number gold, number platinum, Lua_Client splitter); -- void
- Group:[SetLeader](setleader)(Lua_Mob leader); -- void
- Group:[GetLeader](getleader)(); -- unknown - Lua_Mob
- Group:[GetLeaderName](getleadername)(); -- string
- Group:[IsLeader](isleader)(Lua_Mob leader); -- bool
- Group:[GroupCount](groupcount)(); -- number
- Group:[GetHighestLevel](gethighestlevel)(); -- number
- Group:[GetLowestLevel](getlowestlevel)(); -- number
- Group:[TeleportGroup](teleportgroup)(Lua_Mob sender, number zone_id, number instance_id, float x, float y, float z, float h); -- void
- Group:[GetID](getid)(); -- number
- Group:[GetMember](getmember)(number index); -- unknown - Lua_Mob
