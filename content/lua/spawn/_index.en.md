---
title: Spawn Class
menuTitle: Spawn Class
searchDescription: Lua Spawn Class
searchTitle: Lua Spawn Class
weight: 25
---

## Spawn Methods
- [LoadGrid](loadgrid)(); -- void
- [Enable](enable)(); -- void
- [Disable](disable)(); -- void
- [Enabled](enabled)(); -- bool
- [Reset](reset)(); -- void
- [Depop](depop)(); -- void
- [Repop](repop)(number delay); -- void
- [ForceDespawn](forcedespawn)(); -- void
- [GetID](getid)(); -- number
- [GetX](getx)(); -- number
- [GetY](gety)(); -- number
- [GetZ](getz)(); -- number
- [GetHeading](getheading)(); -- number
- [SetRespawnTimer](setrespawntimer)(number newrespawntime); -- void
- [SetVariance](setvariance)(number newvariance); -- void
- [GetVariance](getvariance)(); -- number
- [SpawnGroupID](spawngroupid)(); -- number
- [CurrentNPCID](currentnpcid)(); -- number
- [SetCurrentNPCID](setcurrentnpcid)(number nid); -- void
- [GetSpawnCondition](getspawncondition)(); -- number
- [NPCPointerValid](npcpointervalid)(); -- bool
- [SetNPCPointer](setnpcpointer)(Lua_NPC n); -- void
- [SetTimer](settimer)(number duration); -- void
- [GetKillCount](getkillcount)(); -- number
