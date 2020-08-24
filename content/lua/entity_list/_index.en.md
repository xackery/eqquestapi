---
date: 2016-04-09T16:50:16+02:00
title: Entity List
menuTitle: Entity List
weight: 25
---

# Entity List Methods \(Lua\)

entity_list:CanAddHateForMob(Lua_Mob p); -- bool
entity_list:ChannelMessage(Lua_Mob from, int channel_num, int language, const char *message); -- void
entity_list:ClearClientPetitionQueue(); -- void
entity_list:ClearFeignAggro(Lua_Mob who); -- void
entity_list:DeleteNPCCorpses(); -- int
entity_list:DeletePlayerCorpses(); -- int
entity_list:DoubleAggro(Lua_Mob who); -- void
entity_list:Fighting(Lua_Mob who); -- bool
entity_list:FilteredMessageClose(Lua_Mob sender, bool skip_sender, float dist, uint32 type, int filter, const char *message); -- void
entity_list:FindDoor(uint32 id); -- Lua_Door
entity_list:GetClientByAccID(uint32 acct_id); -- Lua_Client
entity_list:GetClientByCharID(uint32 char_id); -- Lua_Client
entity_list:GetClientByID(int id); -- Lua_Client
entity_list:GetClientByName(const char *name); -- Lua_Client
entity_list:GetClientByWID(uint32 wid); -- Lua_Client
entity_list:GetClientList(); -- Lua_Client_List
entity_list:GetCorpseByID(int id); -- Lua_Corpse
entity_list:GetCorpseByName(const char *name); -- Lua_Corpse
entity_list:GetCorpseByOwner(Lua_Client client); -- Lua_Corpse
entity_list:GetCorpseList(); -- Lua_Corpse_List
entity_list:GetDoorsByDBID(uint32 db_id); -- Lua_Door
entity_list:GetDoorsByDoorID(uint32 door_id); -- Lua_Door
entity_list:GetDoorsByID(int id); -- Lua_Door
entity_list:GetDoorsList(); -- Lua_Doors_List
entity_list:GetGroupByClient(Lua_Client client); -- Lua_Group
entity_list:GetGroupByID(int id); -- Lua_Group
entity_list:GetGroupByLeaderName(const char *name); -- Lua_Group
entity_list:GetGroupByMob(Lua_Mob mob); -- Lua_Group
entity_list:GetMob(const char *name); -- Lua_Mob
entity_list:GetMob(int id); -- Lua_Mob
entity_list:GetMobByNpcTypeID(int npc_type); -- Lua_Mob
entity_list:GetMobID(int id); -- Lua_Mob
entity_list:GetMobList(); -- Lua_Mob_List
entity_list:GetNPCByID(int id); -- Lua_NPC
entity_list:GetNPCByNPCTypeID(int npc_type); -- Lua_NPC
entity_list:GetNPCBySpawnID(int spawn_id); -- Lua_NPC
entity_list:GetNPCList(); -- Lua_NPC_List
entity_list:GetObjectByDBID(uint32 db_id); -- Lua_Object
entity_list:GetObjectByID(int id); -- Lua_Object
entity_list:GetObjectList(); -- Lua_Object_List
entity_list:GetRaidByClient(Lua_Client client); -- Lua_Raid
entity_list:GetRaidByID(int id); -- Lua_Raid
entity_list:GetRandomClient(float x, float y, float z, float dist); -- Lua_Client
entity_list:GetRandomClient(float x, float y, float z, float dist, Lua_Client exclude); -- Lua_Client
entity_list:GetShuffledClientList(); -- Lua_Client_List
entity_list:GetSpawnByID(uint32 id); -- Lua_Spawn
entity_list:GetSpawnList(); -- Lua_Spawn_List
entity_list:HalveAggro(Lua_Mob who); -- void
entity_list:IsMobSpawnedByNpcTypeID(int npc_type); -- bool
entity_list:MakeNameUnique(const char *name); -- std::string
entity_list:Message(uint32 guild_dbid, uint32 type, const char *message); -- void
entity_list:MessageClose(Lua_Mob sender, bool skip_sender, float dist, uint32 type, const char *message); -- void
entity_list:MessageGroup(Lua_Mob who, bool skip_close, uint32 type, const char *message); -- void
entity_list:MessageStatus(uint32 guild_dbid, int min_status, uint32 type, const char *message); -- void
entity_list:OpenDoorsNear(Lua_Mob opener); -- void
entity_list:RemoveFromHateLists(Lua_Mob who); -- void
entity_list:RemoveFromHateLists(Lua_Mob who, bool set_to_one); -- void
entity_list:RemoveFromTargets(Lua_Mob mob); -- void
entity_list:RemoveFromTargets(Lua_Mob mob, bool RemoveFromXTargets); -- void
entity_list:RemoveNumbers(const char *name); -- std::string
entity_list:ReplaceWithTarget(Lua_Mob target, Lua_Mob new_target); -- void
entity_list:SignalAllClients(int signal); -- void
entity_list:SignalMobsByNPCID(uint32 npc_id, int signal); -- void