---
title: Spawn
menuTitle: Spawn
description: Lua Spawn Class
searchTitle: Lua Spawn Class
weight: 25
---

## Spawn Methods
- Spawn:[LoadGrid](loadgrid)(); -- void
- Spawn:[Enable](enable)(); -- void
- Spawn:[Disable](disable)(); -- void
- Spawn:[Enabled](enabled)(); -- bool
- Spawn:[Reset](reset)(); -- void
- Spawn:[Depop](depop)(); -- void
- Spawn:[Repop](repop)(number delay); -- void
- Spawn:[ForceDespawn](forcedespawn)(); -- void
- Spawn:[GetID](getid)(); -- number
- Spawn:[GetX](getx)(); -- number
- Spawn:[GetY](gety)(); -- number
- Spawn:[GetZ](getz)(); -- number
- Spawn:[GetHeading](getheading)(); -- number
- Spawn:[SetRespawnTimer](setrespawntimer)(number newrespawntime); -- void
- Spawn:[SetVariance](setvariance)(number newvariance); -- void
- Spawn:[GetVariance](getvariance)(); -- number
- Spawn:[SpawnGroupID](spawngroupid)(); -- number
- Spawn:[CurrentNPCID](currentnpcid)(); -- number
- Spawn:[SetCurrentNPCID](setcurrentnpcid)(number nid); -- void
- Spawn:[GetSpawnCondition](getspawncondition)(); -- number
- Spawn:[NPCPointerValid](npcpointervalid)(); -- bool
- Spawn:[SetNPCPointer](setnpcpointer)(Lua_NPC n); -- void
- Spawn:[SetTimer](settimer)(number duration); -- void
- Spawn:[GetKillCount](getkillcount)(); -- number
