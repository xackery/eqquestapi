---
title: NPC Class
menuTitle: NPC Class
searchDescription: Lua NPC Class
searchTitle: Lua NPC Class
weight: 25
---

## NPC Methods
- [AddAISpell](addaispell)({{% lua_type_number %}} priority, {{% lua_type_number %}} spell_id, {{% lua_type_number %}} type, {{% lua_type_number %}} mana_cost, {{% lua_type_number %}} recast_delay, {{% lua_type_number %}} resist_adjust, {{% lua_type_number %}} min_hp, {{% lua_type_number %}} max_hp) -- {{% lua_type_nil %}}
- [AddCash](addcash)({{% lua_type_number %}} copper, {{% lua_type_number %}} silver, {{% lua_type_number %}} gold, {{% lua_type_number %}} platinum) -- {{% lua_type_nil %}}
- [AddItem](additem)({{% lua_type_number %}} item_id, {{% lua_type_number %}} charges, {{% lua_type_boolean %}} equip, {{% lua_type_number %}} aug1, {{% lua_type_number %}} aug2, {{% lua_type_number %}} aug3, {{% lua_type_number %}} aug4, {{% lua_type_number %}} aug5, {{% lua_type_number %}} aug6) -- {{% lua_type_nil %}}
- [AddLootTable](addloottable)({{% lua_type_number %}} id) -- {{% lua_type_nil %}}
- [AssignWaypoints](assignwaypoints)({{% lua_type_number %}} grid) -- {{% lua_type_nil %}}
- [CalculateNewWaypoint](calculatenewwaypoint)() -- {{% lua_type_nil %}}
- [CheckNPCFactionAlly](checknpcfactionally)({{% lua_type_number %}} faction) -- {{% lua_type_number %}}
- [ClearItemList](clearitemlist)() -- {{% lua_type_nil %}}
- [CountLoot](countloot)() -- {{% lua_type_number %}}
- [DisplayWaypointInfo](displaywaypointinfo)({{% type_client lua %}} to) -- {{% lua_type_nil %}}
- [DoClassAttacks](doclassattacks)({{% type_npc mob %}} target) -- {{% lua_type_nil %}}
- [GetAccuracyRating](getaccuracyrating)() -- {{% lua_type_number %}}
- [GetAttackDelay](getattackdelay)() -- {{% lua_type_number %}}
- [GetAttackSpeed](getattackspeed)() -- {{% lua_type_number %}}
- [GetAvoidanceRating](getavoidancerating)() -- {{% lua_type_number %}}
- [GetCopper](getcopper)() -- {{% lua_type_number %}}
- [GetFollowCanRun](getfollowcanrun)() -- {{% lua_type_boolean %}}
- [GetFollowDistance](getfollowdistance)() -- {{% lua_type_number %}}
- [GetFollowID](getfollowid)() -- {{% lua_type_number %}}
- [GetGold](getgold)() -- {{% lua_type_number %}}
- [GetGrid](getgrid)() -- {{% lua_type_number %}}
- [GetGuardPointX](getguardpointx)() -- {{% lua_type_number %}}
- [GetGuardPointY](getguardpointy)() -- {{% lua_type_number %}}
- [GetGuardPointZ](getguardpointz)() -- {{% lua_type_number %}}
- [GetLoottableID](getloottableid)() -- {{% lua_type_number %}}
- [GetMaxDamage](getmaxdamage)({{% lua_type_number %}} level) -- {{% lua_type_number %}}
- [GetMaxDMG](getmaxdmg)() -- {{% lua_type_number %}}
- [GetMaxWp](getmaxwp)() -- {{% lua_type_number %}}
- [GetMinDMG](getmindmg)() -- {{% lua_type_number %}}
- [GetNPCFactionID](getnpcfactionid)() -- {{% lua_type_number %}}
- [GetNPCHate](getnpchate)({{% type_npc mob %}} ent) -- {{% lua_type_number %}}
- [GetNPCSpellsID](getnpcspellsid)() -- {{% lua_type_number %}}
- [GetPetSpellID](getpetspellid)() -- {{% lua_type_number %}}
- [GetPlatinum](getplatinum)() -- {{% lua_type_number %}}
- [GetPrimaryFaction](getprimaryfaction)() -- {{% lua_type_number %}}
- [GetPrimSkill](getprimskill)() -- {{% lua_type_number %}}
- [GetRawAC](getrawac)() -- {{% lua_type_number %}}
- [GetScore](getscore)() -- {{% lua_type_number %}}
- [GetSecSkill](getsecskill)() -- {{% lua_type_number %}}
- [GetSilver](getsilver)() -- {{% lua_type_number %}}
- [GetSlowMitigation](getslowmitigation)() -- {{% lua_type_number %}}
- [GetSpawnKillCount](getspawnkillcount)() -- {{% lua_type_number %}}
- [GetSpawnPointH](getspawnpointh)() -- {{% lua_type_number %}}
- [GetSpawnPointID](getspawnpointid)() -- {{% lua_type_number %}}
- [GetSpawnPointX](getspawnpointx)() -- {{% lua_type_number %}}
- [GetSpawnPointY](getspawnpointy)() -- {{% lua_type_number %}}
- [GetSpawnPointZ](getspawnpointz)() -- {{% lua_type_number %}}
- [GetSpellFocusDMG](getspellfocusdmg)() -- {{% lua_type_number %}}
- [GetSpellFocusHeal](getspellfocusheal)() -- {{% lua_type_number %}}
- [GetSwarmOwner](getswarmowner)() -- {{% lua_type_number %}}
- [GetSwarmTarget](getswarmtarget)() -- {{% lua_type_number %}}
- [GetWaypointMax](getwaypointmax)() -- {{% lua_type_number %}}
- [IsAnimal](isanimal)() -- {{% lua_type_boolean %}}
- [IsGuarding](isguarding)() -- {{% lua_type_boolean %}}
- [IsOnHatelist](isonhatelist)({{% type_npc mob %}} ent) -- {{% lua_type_boolean %}}
- [IsTaunting](istaunting)() -- {{% lua_type_boolean %}}
- [MerchantCloseShop](merchantcloseshop)() -- {{% lua_type_nil %}}
- [MerchantOpenShop](merchantopenshop)() -- {{% lua_type_nil %}}
- [ModifyNPCStat](modifynpcstat)({{% lua_type_string %}} stat, {{% lua_type_string %}} value) -- {{% lua_type_nil %}}
- [MoveTo](moveto)({{% lua_type_number %}} x, {{% lua_type_number %}} y, {{% lua_type_number %}} z, {{% lua_type_number %}} h, {{% lua_type_boolean %}} save) -- {{% lua_type_nil %}}
- [NextGuardPosition](nextguardposition)() -- {{% lua_type_nil %}}
- [PauseWandering](pausewandering)({{% lua_type_number %}} pause_time) -- {{% lua_type_nil %}}
- [PickPocket](pickpocket)({{% type_client lua %}} thief) -- {{% lua_type_nil %}}
- [RecalculateSkills](recalculateskills)() -- {{% lua_type_nil %}}
- [RemoveAISpell](removeaispell)({{% lua_type_number %}} spell_id) -- {{% lua_type_nil %}}
- [RemoveCash](removecash)() -- {{% lua_type_nil %}}
- [RemoveItem](removeitem)({{% lua_type_number %}} item_id, {{% lua_type_number %}} quantity, {{% lua_type_number %}} slot) -- {{% lua_type_nil %}}
- [ResumeWandering](resumewandering)() -- {{% lua_type_nil %}}
- [SaveGuardSpot](saveguardspot)({{% lua_type_number %}} x, {{% lua_type_number %}} y, {{% lua_type_number %}} z, {{% lua_type_number %}} heading) -- {{% lua_type_nil %}}
- [ScaleNPC](scalenpc)(uint8 npc_level); -- void
- [SetCopper](setcopper)({{% lua_type_number %}} amt) -- {{% lua_type_nil %}}
- [SetFollowCanRun](setfollowcanrun)({{% lua_type_boolean %}} v) -- {{% lua_type_nil %}}
- [SetFollowDistance](setfollowdistance)({{% lua_type_number %}} dist) -- {{% lua_type_nil %}}
- [SetFollowID](setfollowid)({{% lua_type_number %}} id) -- {{% lua_type_nil %}}
- [SetGold](setgold)({{% lua_type_number %}} amt) -- {{% lua_type_nil %}}
- [SetGrid](setgrid)({{% lua_type_number %}} grid) -- {{% lua_type_nil %}}
- [SetNPCFactionID](setnpcfactionid)({{% lua_type_number %}} id) -- {{% lua_type_nil %}}
- [SetPetSpellID](setpetspellid)({{% lua_type_number %}} id) -- {{% lua_type_nil %}}
- [SetPlatinum](setplatinum)({{% lua_type_number %}} amt) -- {{% lua_type_nil %}}
- [SetPrimSkill](setprimskill)({{% lua_type_number %}} skill_id) -- {{% lua_type_nil %}}
- [SetSaveWaypoint](setsavewaypoint)({{% lua_type_number %}} wp) -- {{% lua_type_nil %}}
- [SetSecSkill](setsecskill)({{% lua_type_number %}} skill_id) -- {{% lua_type_nil %}}
- [SetSilver](setsilver)({{% lua_type_number %}} amt) -- {{% lua_type_nil %}}
- [SetSimpleRoamBox](setsimpleroambox)({{% lua_type_number %}} box_size, {{% lua_type_number %}} move_distance, {{% lua_type_number %}} move_delay) -- {{% lua_type_nil %}}
- [SetSpellFocusDMG](setspellfocusdmg)({{% lua_type_number %}} focus) -- {{% lua_type_nil %}}
- [SetSpellFocusHeal](setspellfocusheal)({{% lua_type_number %}} focus) -- {{% lua_type_nil %}}
- [SetSwarmTarget](setswarmtarget)({{% lua_type_number %}} target) -- {{% lua_type_nil %}}
- [SetTaunting](settaunting)({{% lua_type_boolean %}} t) -- {{% lua_type_nil %}}
- [SetWaypointPause](setwaypointpause)() -- {{% lua_type_nil %}}
- [Signal](signal)({{% lua_type_number %}} id) -- {{% lua_type_nil %}}
- [StartSwarmTimer](startswarmtimer)({{% lua_type_number %}} duration) -- {{% lua_type_nil %}}
- [StopWandering](stopwandering)() -- {{% lua_type_nil %}}
- [UpdateWaypoint](updatewaypoint)({{% lua_type_number %}} wp) -- {{% lua_type_nil %}}
