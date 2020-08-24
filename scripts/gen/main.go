package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("failed", err)
		os.Exit(1)
	}
	fmt.Println("done")
}

func run() error {
	methods := strings.Split(methodDump, "\n")
	language := "Lua"
	category := "Client"

	indexFuncs := []string{}
	path := fmt.Sprintf("../../content/%s/%s/", strings.ToLower(language), strings.ToLower(category))
	for _, method := range methods {

		funcName := method
		if language == "Perl" {
			funcName = funcName[strings.Index(funcName, "->")+2:]
			funcName = funcName[0:strings.Index(funcName, "(")]
		}

		if language == "Lua" {
			funcName = funcName[strings.Index(funcName, ":")+1:]
			funcName = funcName[0:strings.Index(funcName, "(")]
		}

		filename := path + strings.ToLower(funcName) + ".md"
		f, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("create: %w", err)
		}
		funcPage := fmt.Sprintf(`---
title: %s %s
weight: 1
hidden: true
menuTitle: %s
---
## %s
`, language, funcName, funcName, funcName)
		funcPage += fmt.Sprintf("```%s\n%s\n```", strings.ToLower(language), method)
		f.WriteString(funcPage)
		f.Close()

		indexMethod := "- " + method[0:strings.Index(method, funcName)] + "[" + funcName + "](" + strings.ToLower(funcName) + ")" + method[strings.Index(method, funcName)+len(funcName):]

		indexFuncs = append(indexFuncs, indexMethod)
	}

	indexPage := fmt.Sprintf(`---
date: 2020-08-24T16:50:16+02:00
title: %s
menuTitle: %s
weight: 25
---

## %s Methods (%s)`, category, category, category, language)

	indexPage += strings.Join(indexFuncs, "\n")
	f, err := os.Create(path + "_index.en.md")
	if err != nil {
		return fmt.Errorf("create index: %w", err)
	}
	f.WriteString(indexPage)
	f.Close()

	return nil
}

var methodDump = `client:AccountID(); -- uint32
client:AddAAPoints(int points); -- void
client:AddAlternateCurrencyValue(uint32 currency, int amount); -- void
client:AddCrystals(uint32 radiant, uint32 ebon); -- void
client:AddEXP(uint32 add_exp); -- void
client:AddEXP(uint32 add_exp, int conlevel); -- void
client:AddEXP(uint32 add_exp, int conlevel, bool resexp); -- void
client:AddLevelBasedExp(int exp_pct); -- void
client:AddLevelBasedExp(int exp_pct, int max_level); -- void
client:AddLevelBasedExp(int exp_pct, int max_level, bool ignore_mods); -- void
client:AddMoneyToPP(uint32 copper, uint32 silver, uint32 gold, uint32 platinum, bool update_client); -- void
client:AddPVPPoints(uint32 points); -- void
client:AddSkill(int skill_id, int value); -- void
client:Admin(); -- int
client:AssignTask(int task, int npc_id); -- void
client:AssignTask(int task, int npc_id, bool enforce_level_requirement); -- void
client:AssignToInstance(int instance_id); -- void
client:AutoSplitEnabled(); -- bool
client:BreakInvis(); -- void
client:CalcATK(); -- int
client:CalcCurrentWeight(); -- int
client:CalcPriceMod(Lua_Mob other, bool reverse); -- float
client:CanHaveSkill(int skill_id); -- bool
client:ChangeLastName(const char *in); -- void
client:CharacterID(); -- uint32
client:CheckIncreaseSkill(int skill_id, Lua_Mob target); -- void
client:CheckIncreaseSkill(int skill_id, Lua_Mob target, int chance_mod); -- void
client:CheckSpecializeIncrease(int spell_id); -- void
client:ClearCompassMark(); -- void
client:ClearZoneFlag(int zone_id); -- void
client:Connected(); -- bool
client:DecreaseByID(uint32 type, int amt); -- bool
client:DeleteItemInInventory(int slot_id, int quantity); -- void
client:DeleteItemInInventory(int slot_id, int quantity, bool update_client); -- void
client:DisableAreaEndRegen(); -- void
client:DisableAreaHPRegen(); -- void
client:DisableAreaManaRegen(); -- void
client:DisableAreaRegens(); -- void
client:Disconnect(); -- void
client:DropItem(int slot_id); -- void
client:Duck(); -- void
client:DyeArmorBySlot(uint8 slot, uint8 red, uint8 green, uint8 blue); -- void
client:DyeArmorBySlot(uint8 slot, uint8 red, uint8 green, uint8 blue, uint8 use_tint); -- void
client:EnableAreaEndRegen(int value); -- void
client:EnableAreaHPRegen(int value); -- void
client:EnableAreaManaRegen(int value); -- void
client:EnableAreaRegens(int value); -- void
client:Escape(); -- void
client:FailTask(int task); -- void
client:FilteredMessage(Mob *sender, uint32 type, int filter, const char *message); -- void
client:FindMemmedSpellBySlot(int slot); -- uint16
client:FindSpellBookSlotBySpellID(int spell_id); -- int
client:ForageItem(); -- void
client:ForageItem(bool guarantee); -- void
client:Freeze(); -- void
client:GetAAExp(); -- uint32
client:GetAAPercent(); -- uint32
client:GetAAPoints(); -- int
client:GetAccountAge(); -- int
client:GetAccountFlag(std::string flag); -- std::string
client:GetAggroCount(); -- int
client:GetAllMoney(); -- uint64
client:GetAlternateCurrencyValue(uint32 currency); -- int
client:GetAnon(); -- bool
client:GetAugmentIDAt(int slot_id, int aug_slot); -- int
client:GetBaseAGI(); -- int
client:GetBaseCHA(); -- int
client:GetBaseDEX(); -- int
client:GetBaseFace(); -- int
client:GetBaseINT(); -- int
client:GetBaseRace(); -- int
client:GetBaseSTA(); -- int
client:GetBaseSTR(); -- int
client:GetBaseWIS(); -- int
client:GetBindHeading(); -- float
client:GetBindHeading(int index); -- float
client:GetBindX(); -- float
client:GetBindX(int index); -- float
client:GetBindY(); -- float
client:GetBindY(int index); -- float
client:GetBindZ(); -- float
client:GetBindZ(int index); -- float
client:GetBindZoneID(); -- uint32
client:GetBindZoneID(int index); -- uint32
client:GetCarriedMoney(); -- uint64
client:GetCharacterFactionLevel(int faction_id); -- int
client:GetClientVersion(); -- int
client:GetClientVersionBit(); -- uint32
client:GetCorpseCount(); -- int
client:GetCorpseID(int corpse); -- int
client:GetCorpseItemAt(int corpse, int slot); -- int
client:GetDisciplineTimer(uint32 timer_id); -- uint32
client:GetDiscSlotBySpellID(int32 spell_id); -- int
client:GetDisplayAC()
client:GetDuelTarget(); -- int
client:GetEXP(); -- uint32
client:GetEbonCrystals(); -- uint32
client:GetEndurance(); -- int
client:GetEndurancePercent(); -- int
client:GetFace(); -- int
client:GetFactionLevel(uint32 char_id, uint32 npc_id, uint32 race, uint32 class_, uint32 deity, uint32 faction, Lua_NPC npc); -- int
client:GetFeigned(); -- bool
client:GetGM(); -- bool
client:GetGroup(); -- Lua_Group
client:GetGroupPoints(); -- uint32
client:GetHorseId(); -- int
client:GetHunger(); -- int
client:GetIP(); -- uint32
client:GetInstrumentMod(int spell_id); -- int
client:GetInventory(); -- Lua_Inventory
client:GetItemIDAt(int slot_id); -- int
client:GetLDoNLosses(); -- int
client:GetLDoNLossesTheme(int theme); -- int
client:GetLDoNPointsTheme(int theme); -- int
client:GetLDoNWins(); -- int
client:GetLDoNWinsTheme(int theme); -- int
client:GetLanguageSkill(int skill_id); -- int
client:GetMaxEndurance(); -- int
client:GetModCharacterFactionLevel(int faction); -- int
client:GetMoney(uint8 type, uint8 subtype); -- uint32
client:GetNextAvailableSpellBookSlot(); -- int
client:GetNextAvailableSpellBookSlot(int start); -- int
client:GetPVP(); -- bool
client:GetPVPPoints(); -- uint32
client:GetRadiantCrystals(); -- uint32
client:GetRaid(); -- Lua_Raid
client:GetRaidPoints(); -- uint32
client:GetRawItemAC(); -- int
client:GetRawSkill(int skill_id); -- int
client:GetSkillPoints(); -- int
client:GetSpentAA(); -- int
client:GetStartZone(); -- int
client:GetThirst(); -- int
client:GetTotalSecondsPlayed(); -- uint32
client:GetWeight(); -- int
client:GoFish(); -- void
client:GrantAlternateAdvancementAbility(int aa_id, int points); -- bool
client:GrantAlternateAdvancementAbility(int aa_id, int points, bool ignore_cost); -- bool
client:GuildID(); -- uint32
client:GuildRank(); -- int
client:HasSkill(int skill_id); -- bool
client:HasSpellScribed(int spell_id); -- bool
client:HasZoneFlag(int zone_id); -- bool
client:Hungry(); -- bool
client:InZone(); -- bool
client:IncStats(int type, int value); -- void
client:IncreaseLanguageSkill(int skill_id); -- void
client:IncreaseLanguageSkill(int skill_id, int value); -- void
client:IncreaseSkill(int skill_id); -- void
client:IncreaseSkill(int skill_id, int value); -- void
client:IncrementAA(int aa); -- void
client:IsCrouching(); -- bool
client:IsDead(); -- bool
client:IsDueling(); -- bool
client:IsGrouped(); -- bool
client:IsLD(); -- bool
client:IsMedding(); -- bool
client:IsRaidGrouped(); -- bool
client:IsSitting(); -- bool
client:IsStanding(); -- bool
client:IsTaskActive(int task); -- bool
client:IsTaskActivityActive(int task, int activity); -- bool
client:IsTaskCompleted(int task); -- bool
client:KeyRingAdd(uint32 item); -- void
client:KeyRingCheck(uint32 item); -- bool
client:Kick(); -- void
client:LearnRecipe(uint32 recipe); -- void
client:LeaveGroup(); -- void
client:MarkSingleCompassLoc(float in_x, float in_y, float in_z); -- void
client:MarkSingleCompassLoc(float in_x, float in_y, float in_z, int count); -- void
client:MaxSkill(int skill_id); -- int
client:MemSpell(int spell_id, int slot); -- void
client:MemSpell(int spell_id, int slot, bool update_client); -- void
client:MemmedCount(); -- int
client:MovePC(int zone, float x, float y, float z, float heading); -- void
client:MovePCInstance(int zone, int instance, float x, float y, float z, float heading); -- void
client:MoveZone(const char *zone_short_name); -- void
client:MoveZoneGroup(const char *zone_short_name); -- void
client:MoveZoneRaid(const char *zone_short_name); -- void
client:MoveZoneInstance(uint16 instance_id); -- void
client:MoveZoneInstanceGroup(uint16 instance_id); -- void
client:MoveZoneInstanceRaid(uint16 instance_id); -- void
client:NukeItem(uint32 item_num); -- void
client:NukeItem(uint32 item_num, int where_to_check); -- void
client:OpenLFGuildWindow(); -- void
client:PlayMP3(std::string file); -- void
client:PushItemOnCursor(Lua_ItemInst inst); -- bool
client:PutItemInInventory(int slot_id, Lua_ItemInst inst); -- bool
client:QuestReadBook(const char *text, int type); -- void
client:QuestReward(Lua_Mob target); -- void
client:QuestReward(Lua_Mob target, uint32 copper); -- void
client:QuestReward(Lua_Mob target, uint32 copper, uint32 silver); -- void
client:QuestReward(Lua_Mob target, uint32 copper, uint32 silver, uint32 gold); -- void
client:QuestReward(Lua_Mob target, uint32 copper, uint32 silver, uint32 gold, uint32 platinum); -- void
client:QuestReward(Lua_Mob target, uint32 copper, uint32 silver, uint32 gold, uint32 platinum, uint32 itemid); -- void
client:QuestReward(Lua_Mob target, uint32 copper, uint32 silver, uint32 gold, uint32 platinum, uint32 itemid, uint32 exp); -- void
client:QuestReward(Lua_Mob target, uint32 copper, uint32 silver, uint32 gold, uint32 platinum, uint32 itemid, uint32 exp, bool faction); -- void
client:QueuePacket(Lua_Packet app); -- void
client:QueuePacket(Lua_Packet app, bool ack_req); -- void
client:QueuePacket(Lua_Packet app, bool ack_req, int client_connection_status); -- void
client:QueuePacket(Lua_Packet app, bool ack_req, int client_connection_status, int filter); -- void
client:RefundAA(); -- void
client:ResetAA(); -- void
client:ResetDisciplineTimer(uint32 timer_id); -- void
client:ResetTrade(); -- void
client:Save(); -- void
client:Save(int commit_now); -- void
client:SaveBackup(); -- void
client:ScribeSpell(int spell_id, int slot); -- void
client:ScribeSpell(int spell_id, int slot, bool update_client); -- void
client:SendColoredText(uint32 type, std::string msg); -- void
client:SendItemScale(Lua_ItemInst inst); -- void
client:SendMarqueeMessage(uint32 type, uint32 priority, uint32 fade_in, uint32 fade_out, uint32 duration, std::string msg); -- void
client:SendOPTranslocateConfirm(Lua_Mob caster, int spell_id); -- void
client:SendSound(); -- void
client:SendWebLink(const char *site); -- void
client:SendZoneFlagInfo(Lua_Client to); -- void
client:SetAAPoints(int points); -- void
client:SetAATitle(const char *title); -- void
client:SetAccountFlag(std::string flag, std::string val); -- void
client:SetAlternateCurrencyValue(uint32 currency, int amount); -- void
client:SetBaseClass(int v); -- void
client:SetBaseGender(int v); -- void
client:SetBaseRace(int v); -- void
client:SetBindPoint(); -- void
client:SetBindPoint(int to_zone); -- void
client:SetBindPoint(int to_zone, int to_instance); -- void
client:SetBindPoint(int to_zone, int to_instance, float new_x); -- void
client:SetBindPoint(int to_zone, int to_instance, float new_x, float new_y); -- void
client:SetBindPoint(int to_zone, int to_instance, float new_x, float new_y, float new_z); -- void
client:SetConsumption(int in_hunger, int in_thirst); -- void
client:SetDeity(int v); -- void
client:SetDuelTarget(int c); -- void
client:SetDueling(bool v); -- void
client:SetEXP(uint32 set_exp, uint32 set_aaxp); -- void
client:SetEXP(uint32 set_exp, uint32 set_aaxp, bool resexp); -- void
client:SetEndurance(int endur); -- void
client:SetFactionLevel(uint32 char_id, uint32 npc_id, int char_class, int char_race, int char_deity); -- void
client:SetFactionLevel2(uint32 char_id, int faction_id, int char_class, int char_race, int char_deity, int value, int temp); -- void
client:SetFeigned(bool v); -- void
client:SetGM(bool v); -- void
client:SetHorseId(int id); -- void
client:SetHunger(int in_hunger); -- void
client:SetLanguageSkill(int language, int value); -- void
client:SetMaterial(int slot_id, uint32 item_id); -- void
client:SetPVP(bool v); -- void
client:SetPrimaryWeaponOrnamentation(uint32 model_id); -- void
client:SetSecondaryWeaponOrnamentation(uint32 model_id); -- void
client:SetSkill(int skill_id, int value); -- void
client:SetSkillPoints(int skill); -- void
client:SetStartZone(int zone_id); -- void
client:SetStartZone(int zone_id, float x); -- void
client:SetStartZone(int zone_id, float x, float y); -- void
client:SetStartZone(int zone_id, float x, float y, float z); -- void
client:SetStats(int type, int value); -- void
client:SetThirst(int in_thirst); -- void
client:SetTint(int slot_id, uint32 color); -- void
client:SetTitleSuffix(const char *text); -- void
client:SetZoneFlag(int zone_id); -- void
client:Signal(uint32 id); -- void
client:Stand(); -- void
client:SummonItem(uint32 item_id); -- void
client:SummonItem(uint32 item_id, int charges); -- void
client:SummonItem(uint32 item_id, int charges, uint32 aug1); -- void
client:SummonItem(uint32 item_id, int charges, uint32 aug1, uint32 aug2); -- void
client:SummonItem(uint32 item_id, int charges, uint32 aug1, uint32 aug2, uint32 aug3); -- void
client:SummonItem(uint32 item_id, int charges, uint32 aug1, uint32 aug2, uint32 aug3, uint32 aug4); -- void
client:SummonItem(uint32 item_id, int charges, uint32 aug1, uint32 aug2, uint32 aug3, uint32 aug4, uint32 aug5); -- void
client:SummonItem(uint32 item_id, int charges, uint32 aug1, uint32 aug2, uint32 aug3, uint32 aug4, uint32 aug5, bool attuned); -- void
client:SummonItem(uint32 item_id, int charges, uint32 aug1, uint32 aug2, uint32 aug3, uint32 aug4, uint32 aug5, bool attuned, int to_slot); -- void
client:TGB(); -- bool
client:TakeMoneyFromPP(uint64 copper); -- bool
client:TakeMoneyFromPP(uint64 copper, bool update_client); -- bool
client:Thirsty(); -- bool
client:TrainDisc(int itemid); -- void
client:TrainDiscBySpellID(int32 spell_id); -- void
client:UnFreeze(); -- void
client:Undye(); -- void
client:UnmemSpell(int slot); -- void
client:UnmemSpell(int slot, bool update_client); -- void
client:UnmemSpellAll(); -- void
client:UnmemSpellAll(bool update_client); -- void
client:UnmemSpellBySpellID(int32 spell_id); -- void
client:UnscribeSpell(int slot); -- void
client:UnscribeSpell(int slot, bool update_client); -- void
client:UnscribeSpellAll(); -- void
client:UnscribeSpellAll(bool update_client); -- void
client:UntrainDisc(int slot); -- void
client:UntrainDisc(int slot, bool update_client); -- void
client:UntrainDiscAll(); -- void
client:UntrainDiscAll(bool update_client); -- void
client:UpdateGroupAAs(int points, uint32 type); -- void
client:UpdateLDoNPoints(int points, uint32 theme); -- void
client:UpdateTaskActivity(int task, int activity, int count); -- void
client:UseDiscipline(int spell_id, int target_id); -- bool
client:WorldKick(); -- void`
