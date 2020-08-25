---
date: 2020-08-24T16:50:16+02:00
title: Group
menuTitle: Group
weight: 25
---

## Group Methods (Perl)
- $group->[CastGroupSpell](castgroupspell)(mob* caster, uint16 spell_id)
- $group->[DisbandGroup](disbandgroup)()
- $group->[GetHighestLevel](gethighestlevel)()
- $group->[GetID](getid)()
- $group->[GetLeader](getleader)()
- $group->[GetLeaderName](getleadername)()
- $group->[GetMember](getmember)(int group_index)
- $group->[GetTotalGroupDamage](gettotalgroupdamage)(mob* other)
- $group->[GroupCount](groupcount)()
- $group->[GroupMessage](groupmessage)(mob* sender, uint8 language, string message)
- $group->[IsGroupMember](isgroupmember)(client)
- $group->[IsLeader](isleader)(mob* target)
- $group->[SendHPPacketsFrom](sendhppacketsfrom)(mob* new_member)
- $group->[SendHPPacketsTo](sendhppacketsto)(mob* new_member)
- $group->[SetLeader](setleader)(mob* new_leader)
- $group->[SplitExp](splitexp)(uint32 exp, mob* other)
- $group->[SplitMoney](splitmoney)(uint32 copper, uint32 silver, uint32 gold, uint32 platinum)
- $group->[TeleportGroup](teleportgroup)(mob* sender, uint32 zone_id, float x, float y, float z, float heading)