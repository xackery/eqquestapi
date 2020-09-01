---
title: EntityList Class
menuTitle: EntityList Class
description: Lua EntityList Class
searchTitle: Lua EntityList Class
weight: 25
---

## Entity List Methods
- [CanAddHateForMob](canaddhateformob)({{% type_npc mob %}} p) -- {{% lua_type_boolean %}}
- [ChannelMessage](channelmessage)({{% type_npc mob %}} from, {{% lua_type_number %}} channel_num, {{% lua_type_number %}} language, {{% lua_type_string %}} message) -- {{% lua_type_nil %}}
- [ClearClientPetitionQueue](clearclientpetitionqueue)() -- {{% lua_type_nil %}}
- [ClearFeignAggro](clearfeignaggro)({{% type_npc mob %}} who) -- {{% lua_type_nil %}}
- [DeleteNPCCorpses](deletenpccorpses)() -- {{% lua_type_number %}}
- [DeletePlayerCorpses](deleteplayercorpses)() -- {{% lua_type_number %}}
- [DoubleAggro](doubleaggro)({{% type_npc mob %}} who) -- {{% lua_type_nil %}}
- [Fighting](fighting)({{% type_npc mob %}} who) -- {{% lua_type_boolean %}}
- [FilteredMessageClose](filteredmessageclose)({{% type_npc mob %}} sender, {{% lua_type_boolean %}} skip_sender, {{% lua_type_number %}} dist, {{% lua_type_number %}} type, {{% lua_type_number %}} filter, {{% lua_type_string %}} message) -- {{% lua_type_nil %}}
- [FindDoor](finddoor)({{% lua_type_number %}} id) -- {{% lua_type_door %}}
- [GetClientByAccID](getclientbyaccid)({{% lua_type_number %}} acct_id) -- {{% type_client lua %}}
- [GetClientByCharID](getclientbycharid)({{% lua_type_number %}} char_id) -- {{% type_client lua %}}
- [GetClientByID](getclientbyid)({{% lua_type_number %}} id) -- {{% type_client lua %}}
- [GetClientByName](getclientbyname)({{% lua_type_string %}} name) -- {{% type_client lua %}}
- [GetClientByWID](getclientbywid)({{% lua_type_number %}} wid) -- {{% type_client lua %}}
- [GetClientList](getclientlist)() -- {{% type_client lua_list %}}
- [GetCorpseByID](getcorpsebyid)({{% lua_type_number %}} id) -- {{% lua_type_corpse %}}
- [GetCorpseByName](getcorpsebyname)({{% lua_type_string %}} name) -- {{% lua_type_corpse %}}
- [GetCorpseByOwner](getcorpsebyowner)({{% type_client lua %}} client) -- {{% lua_type_corpse %}}
- [GetCorpseList](getcorpselist)() -- {{% lua_type_corpse_list %}}
- [GetDoorsByDBID](getdoorsbydbid)({{% lua_type_number %}} db_id) -- {{% lua_type_door %}}
- [GetDoorsByDoorID](getdoorsbydoorid)({{% lua_type_number %}} door_id) -- {{% lua_type_door %}}
- [GetDoorsByID](getdoorsbyid)({{% lua_type_number %}} id) -- {{% lua_type_door %}}
- [GetDoorsList](getdoorslist)() -- {{% lua_type_door %}}s_List
- [GetGroupByClient](getgroupbyclient)({{% type_client lua %}} client) -- {{% lua_type_group %}}
- [GetGroupByID](getgroupbyid)({{% lua_type_number %}} id) -- {{% lua_type_group %}}
- [GetGroupByLeaderName](getgroupbyleadername)({{% lua_type_string %}} name) -- {{% lua_type_group %}}
- [GetGroupByMob](getgroupbymob)({{% type_npc mob %}} mob) -- {{% lua_type_group %}}
- [GetMobByNpcTypeID](getmobbynpctypeid)({{% lua_type_number %}} npc_type) -- {{% type_npc mob %}}
- [GetMobID](getmobid)({{% lua_type_number %}} id) -- {{% type_npc mob %}}
- [GetMobList](getmoblist)() -- {{% type_npc mob_list %}}
- [GetNPCByID](getnpcbyid)({{% lua_type_number %}} id) -- {{% type_npc lua %}}
- [GetNPCByNPCTypeID](getnpcbynpctypeid)({{% lua_type_number %}} npc_type) -- {{% type_npc lua %}}
- [GetNPCBySpawnID](getnpcbyspawnid)({{% lua_type_number %}} spawn_id) -- {{% type_npc lua %}}
- [GetNPCList](getnpclist)() -- {{% type_npc lua_list %}}
- [GetObjectByDBID](getobjectbydbid)({{% lua_type_number %}} db_id) -- {{% lua_type_object %}}
- [GetObjectByID](getobjectbyid)({{% lua_type_number %}} id) -- {{% lua_type_object %}}
- [GetObjectList](getobjectlist)() -- {{% lua_type_object_list %}}
- [GetRaidByClient](getraidbyclient)({{% type_client lua %}} client) -- {{% type_raid lua %}}
- [GetRaidByID](getraidbyid)({{% lua_type_number %}} id) -- {{% type_raid lua %}}
- [GetRandomClient](getrandomclient)({{% lua_type_number %}} x, {{% lua_type_number %}} y, {{% lua_type_number %}} z, {{% lua_type_number %}} dist, {{% type_client lua %}} exclude) -- {{% type_client lua %}}
- [GetShuffledClientList](getshuffledclientlist)() -- {{% type_client lua_list %}}
- [GetSpawnByID](getspawnbyid)({{% lua_type_number %}} id) -- {{% lua_type_spawn %}}
- [GetSpawnList](getspawnlist)() -- {{% lua_type_spawn_list %}}
- [HalveAggro](halveaggro)({{% type_npc mob %}} who) -- {{% lua_type_nil %}}
- [IsMobSpawnedByNpcTypeID](ismobspawnedbynpctypeid)({{% lua_type_number %}} npc_type) -- {{% lua_type_boolean %}}
- [MakeNameUnique](makenameunique)({{% lua_type_string %}} name) -- {{% lua_type_string %}}
- [Message](message)({{% lua_type_number %}} guild_dbid, {{% lua_type_number %}} type, {{% lua_type_string %}} message) -- {{% lua_type_nil %}}
- [MessageClose](messageclose)({{% type_npc mob %}} sender, {{% lua_type_boolean %}} skip_sender, {{% lua_type_number %}} dist, {{% lua_type_number %}} type, {{% lua_type_string %}} message) -- {{% lua_type_nil %}}
- [MessageGroup](messagegroup)({{% type_npc mob %}} who, {{% lua_type_boolean %}} skip_close, {{% lua_type_number %}} type, {{% lua_type_string %}} message) -- {{% lua_type_nil %}}
- [MessageStatus](messagestatus)({{% lua_type_number %}} guild_dbid, {{% lua_type_number %}} min_status, {{% lua_type_number %}} type, {{% lua_type_string %}} message) -- {{% lua_type_nil %}}
- [OpenDoorsNear](opendoorsnear)({{% type_npc mob %}} opener) -- {{% lua_type_nil %}}
- [RemoveFromHateLists](removefromhatelists)({{% type_npc mob %}} who, {{% lua_type_boolean %}} set_to_one) -- {{% lua_type_nil %}}
- [RemoveFromTargets](removefromtargets)({{% type_npc mob %}} mob, {{% lua_type_boolean %}} RemoveFromXTargets) -- {{% lua_type_nil %}}
- [RemoveNumbers](removenumbers)({{% lua_type_string %}} name) -- {{% lua_type_string %}}
- [ReplaceWithTarget](replacewithtarget)({{% type_npc mob %}} target, {{% type_npc mob %}} new_target) -- {{% lua_type_nil %}}
- [SignalAllClients](signalallclients)({{% lua_type_number %}} signal) -- {{% lua_type_nil %}}
- [SignalMobsByNPCID](signalmobsbynpcid)({{% lua_type_number %}} npc_id, {{% lua_type_number %}} signal) -- {{% lua_type_nil %}}
