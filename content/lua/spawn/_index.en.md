---
date: 2020-08-24T16:50:16+02:00
title: Spawn
menuTitle: Spawn
weight: 25
---

## Spawn Methods (Lua)
- spawn:[CurrentNPCID](currentnpcid)(); -- uint32
- spawn:[Depop](depop)(); -- void
- spawn:[Disable](disable)(); -- void
- spawn:[Enable](enable)(); -- void
- spawn:[Enabled](enabled)(); -- bool
- spawn:[ForceDespawn](forcedespawn)(); -- void
- spawn:[GetHeading](getheading)(); -- float
- spawn:[GetID](getid)(); -- uint32
- spawn:[GetKillCount](getkillcount)(); -- uint32
- spawn:[GetSpawnCondition](getspawncondition)(); -- uint32
- spawn:[GetVariance](getvariance)(); -- uint32
- spawn:[GetX](getx)(); -- float
- spawn:[GetY](gety)(); -- float
- spawn:[GetZ](getz)(); -- float
- spawn:[LoadGrid](loadgrid)(); -- void
- spawn:[NPCPointerValid](npcpointervalid)(); -- bool
- spawn:[Repop](repop)(); -- void
- spawn:[Repop](repop)(uint32 delay); -- void
- spawn:[Reset](reset)(); -- void
- spawn:[RespawnTimer](respawntimer)(); -- uint32
- spawn:[SetCurrentNPCID](setcurrentnpcid)(uint32 nid); -- void
- spawn:[SetNPCPointer](setnpcpointer)(Lua_NPC n); -- void
- spawn:[SetRespawnTimer](setrespawntimer)(uint32 newrespawntime); -- void
- spawn:[SetTimer](settimer)(uint32 duration); -- void
- spawn:[SetVariance](setvariance)(uint32 newvariance); -- void
- spawn:[SpawnGroupID](spawngroupid)(); -- uint32