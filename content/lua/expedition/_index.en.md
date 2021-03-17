---
title: Expedition
menuTitle: Expedition
weight: 25
---

## Expedition Methods (Lua)
- [AddLockout](addlockout)(std::string event_name, uint32_t seconds); -- void
- [AddLockoutDuration](addlockoutduration)(std::string event_name, number seconds, bool members_only); -- void
- [AddReplayLockout](addreplaylockout)(uint32_t seconds); -- void
- [AddReplayLockoutDuration](addreplaylockoutduration)(number seconds, bool members_only); -- void
- [GetDynamicZoneID](getdynamiczoneid)(); -- unknown - uint32_t
- [GetID](getid)(); -- unknown - uint32_t
- [GetInstanceID](getinstanceid)(); -- number
- [GetLeaderName](getleadername)(); -- string
- [GetLockouts](getlockouts)(lua_State* L); -- unknown - object
- [GetLootEventByNPCTypeID](getlooteventbynpctypeid)(uint32_t npc_type_id); -- string
- [GetLootEventBySpawnID](getlooteventbyspawnid)(uint32_t spawn_id); -- string
- [GetMemberCount](getmembercount)(); -- unknown - uint32_t
- [GetMembers](getmembers)(lua_State* L); -- unknown - object
- [GetName](getname)(); -- string
- [GetSecondsRemaining](getsecondsremaining)(); -- number
- [GetUUID](getuuid)(); -- string
- [GetZoneID](getzoneid)(); -- number
- [GetZoneName](getzonename)(); -- string
- [GetZoneVersion](getzoneversion)(); -- number
- [HasLockout](haslockout)(std::string event_name); -- bool
- [HasReplayLockout](hasreplaylockout)(); -- bool
- [IsLocked](islocked)(); -- bool
- [RemoveCompass](removecompass)(); -- void
- [RemoveLockout](removelockout)(std::string event_name); -- void
- [SetCompass](setcompass)(std::string zone_name, float x, float y, float z); -- void
- [SetLocked](setlocked)(bool lock_expedition, number lock_msg, uint32_t msg_color); -- void
- [SetLootEventByNPCTypeID](setlooteventbynpctypeid)(uint32_t npc_type_id, std::string event_name); -- void
- [SetLootEventBySpawnID](setlooteventbyspawnid)(uint32_t spawn_id, std::string event_name); -- void
- [SetReplayLockoutOnMemberJoin](setreplaylockoutonmemberjoin)(bool enable); -- void
- [SetSafeReturn](setsafereturn)(std::string zone_name, float x, float y, float z, float heading); -- void
- [SetSecondsRemaining](setsecondsremaining)(uint32_t seconds_remaining); -- void
- [SetZoneInLocation](setzoneinlocation)(float x, float y, float z, float heading); -- void
- [UpdateLockoutDuration](updatelockoutduration)(std::string event_name, uint32_t duration, bool members_only); -- void
