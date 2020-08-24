---
date: 2016-04-09T16:50:16+02:00
title: NPC
menuTitle: NPC
weight: 25
---

# NPC Methods \(Lua\)

npc:AI_SetRoambox(float dist, float max_x, float min_x, float max_y, float min_y); -- void
npc:AI_SetRoambox(float dist, float max_x, float min_x, float max_y, float min_y, uint32 delay, uint32 mindelay); -- void
npc:AddAISpell(int priority, int spell_id, int type, int mana_cost, int recast_delay, int resist_adjust); -- void
npc:AddAISpell(int priority, int spell_id, int type, int mana_cost, int recast_delay, int resist_adjust, int min_hp, int max_hp); -- void
npc:AddCash(int copper, int silver, int gold, int platinum); -- void
npc:AddItem(int item_id, int charges); -- void
npc:AddItem(int item_id, int charges, bool equip); -- void
npc:AddItem(int item_id, int charges, bool equip, int aug1); -- void
npc:AddItem(int item_id, int charges, bool equip, int aug1, int aug2); -- void
npc:AddItem(int item_id, int charges, bool equip, int aug1, int aug2, int aug3); -- void
npc:AddItem(int item_id, int charges, bool equip, int aug1, int aug2, int aug3, int aug4); -- void
npc:AddItem(int item_id, int charges, bool equip, int aug1, int aug2, int aug3, int aug4, int aug5); -- void
npc:AddItem(int item_id, int charges, bool equip, int aug1, int aug2, int aug3, int aug4, int aug5, int aug6); -- void
npc:AddLootTable(); -- void
npc:AddLootTable(int id); -- void
npc:AssignWaypoints(int grid); -- void
npc:CalculateNewWaypoint(); -- void
npc:CheckNPCFactionAlly(int faction); -- int
npc:ClearItemList(); -- void
npc:CountLoot(); -- int
npc:DisplayWaypointInfo(Lua_Client to); -- void
npc:DoClassAttacks(Lua_Mob target); -- void
npc:GetAccuracyRating(); -- int
npc:GetAttackDelay(); -- int
npc:GetAttackSpeed(); -- float
npc:GetAvoidanceRating(); -- int
npc:GetCopper(); -- uint32
npc:GetFollowCanRun(); -- bool
npc:GetFollowDistance(); -- int
npc:GetFollowID(); -- int
npc:GetGold(); -- uint32
npc:GetGrid(); -- int
npc:GetGuardPointX(); -- float
npc:GetGuardPointY(); -- float
npc:GetGuardPointZ(); -- float
npc:GetLoottableID(); -- int
npc:GetMaxDMG(); -- uint32
npc:GetMaxDamage(int level); -- uint32
npc:GetMaxWp(); -- int
npc:GetMinDMG(); -- uint32
npc:GetNPCFactionID(); -- int
npc:GetNPCHate(Lua_Mob ent); -- int
npc:GetNPCSpellsID(); -- int
npc:GetPetSpellID(); -- int
npc:GetPlatinum(); -- uint32
npc:GetPrimSkill(); -- int
npc:GetPrimaryFaction(); -- int
npc:GetRawAC(); -- int
npc:GetScore(); -- int
npc:GetSecSkill(); -- int
npc:GetSilver(); -- uint32
npc:GetSlowMitigation(); -- float
npc:GetSp2(); -- uint32
npc:GetSpawnKillCount(); -- int
npc:GetSpawnPointH(); -- float
npc:GetSpawnPointID(); -- int
npc:GetSpawnPointX(); -- float
npc:GetSpawnPointY(); -- float
npc:GetSpawnPointZ(); -- float
npc:GetSpellFocusDMG(); -- int
npc:GetSpellFocusHeal(); -- int
npc:GetSwarmOwner(); -- int
npc:GetSwarmTarget(); -- int
npc:GetWaypointMax(); -- int
npc:IsAnimal(); -- bool
npc:IsGuarding(); -- bool
npc:IsOnHatelist(Lua_Mob ent); -- bool
npc:MerchantCloseShop(); -- void
npc:MerchantOpenShop(); -- void
npc:ModifyNPCStat(const char *stat, const char *value); -- void
npc:MoveTo(float x, float y, float z, float h, bool save); -- void
npc:NextGuardPosition(); -- void
npc:PauseWandering(int pause_time); -- void
npc:PickPocket(Lua_Client thief); -- void
npc:RecalculateSkills(); -- void
npc:RemoveAISpell(int spell_id); -- void
npc:RemoveCash(); -- void
npc:RemoveItem(int item_id); -- void
npc:RemoveItem(int item_id, int quantity); -- void
npc:RemoveItem(int item_id, int quantity, int slot); -- void
npc:ResumeWandering(); -- void
npc:SaveGuardSpot(float x, float y, float z, float heading); -- void
npc:SetCopper(uint32 amt); -- void
npc:SetFollowCanRun(bool v); -- void
npc:SetFollowDistance(int dist); -- void
npc:SetFollowID(int id); -- void
npc:SetGold(uint32 amt); -- void
npc:SetGrid(int grid); -- void
npc:SetNPCFactionID(int id); -- void
npc:SetPetSpellID(int id); -- void
npc:SetPlatinum(uint32 amt); -- void
npc:SetPrimSkill(int skill_id); -- void
npc:SetSaveWaypoint(int wp); -- void
npc:SetSecSkill(int skill_id); -- void
npc:SetSilver(uint32 amt); -- void
npc:SetSimpleRoamBox(float box_size); -- void
npc:SetSimpleRoamBox(float box_size, float move_distance); -- void
npc:SetSimpleRoamBox(float box_size, float move_distance, int move_delay); -- void
npc:SetSp2(int sg2); -- void
npc:SetSpellFocusDMG(int focus); -- void
npc:SetSpellFocusHeal(int focus); -- void
npc:SetSwarmTarget(int target); -- void
npc:SetTaunting(bool t); -- void
npc:SetWaypointPause(); -- void
npc:Signal(int id); -- void
npc:StartSwarmTimer(uint32 duration); -- void
npc:StopWandering(); -- void
npc:UpdateWaypoint(int wp); -- void