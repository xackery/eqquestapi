---
title: Raid Class
menuTitle: Raid Class
searchTitle: Perl Raid Class
weight: 25
---

## Raid Methods
- $raid->[BalanceHP](balancehp)(int32 penalty, uint32 group_id)
- $raid->[CastGroupSpell](castgroupspell)(mob* caster, uint16 spell_id, uint32 group_id)
- $raid->[GetClientByIndex](getclientbyindex)(uint16 raid_indez)
- $raid->[GetGroup](getgroup)(string name)
- $raid->[GetHighestLevel](gethighestlevel)()
- $raid->[GetID](getid)()
- $raid->[GetLowestLevel](getlowestlevel)()
- $raid->[GetMember](getmember)(int raid_index)
- $raid->[GetTotalRaidDamage](gettotalraiddamage)([mob* other = nullptr])
- $raid->[GroupCount](groupcount)(uint32 group_id)
- $raid->[IsGroupLeader](isgroupleader)(string name)
- $raid->[IsLeader](isleader)(string name)
- $raid->[IsRaidMember](israidmember)(string name)
- $raid->[RaidCount](raidcount)()
- $raid->[SplitExp](splitexp)(uint32 experience, [mob* other = nullptr])
- $raid->[SplitMoney](splitmoney)(uint32 copper, uint32 silver, uint32 gold, uint32 platinum)
- $raid->[TeleportGroup](teleportgroup)(mob* sender, uint32 zone_id, float X, float Y, float Z, float heading, uint32 group_id)
- $raid->[TeleportRaid](teleportraid)(mob* sender, uint32 zone_id, float X, float Y, float Z, float heading)