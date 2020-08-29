---
title: EVENT_SAY
searchTitle: Perl Event Say
description: Triggers when a player talks to an NPC while targetting them.
weight: 1
hidden: true
menuTitle: EVENT_SAY
---

#### EVENT_SAY

Triggers when a player talks to an NPC while targetting them.

**Syntax**
```perl
sub EVENT_SAY { }
```

**Parameters**
- None

**Example**

```perl
sub EVENT_SAY {
   quest::say("Hello, $name. I have [".quest::saylink("discovered")."] something that may interest you.");
}
```

**Exported Variables**

These variables are available with EVENT_SAY (Note: Some may not be available based on your perl_event_export_settings table settings)

- string **$text**: player message that triggered the event
- uint32 **$langid**: language id the message was in
- uint **$data**: uint object id the mess
- string **$name**: name of player that triggered the event
- string **$race**: race of player that triggered the event
- string **$class**: class of player that triggered the event
- string **$ulevel**: level of player that triggered the event
- int **$userid**: entity ID of player that triggered the event
- string **$mname**: name of target NPC 
- int **$mobid**: ID of target NPC
- int **$mlevel**: level of target NPC
- int **$hpratio**: value of 0 to 100 representing hp ratio of target NPC
- float **$x**: x location of target NPC
- float **$y**: y location of target NPC
- float **$z**: z location of target NPC
- float **$h**: heading direction of target NPC
- int **$faction**: faction of player that triggered the event
- int **$uguild_id**: guildID of player that triggered the event
- int **$uguildrank**: guild rank of player that triggered the event
- int **$status**: status of player that triggered the event
- int **$zoneid**: zone ID
- string **$zoneln**: zone long name
- string **$zonesn**: zone short name
- int **$instanceid**: zone instance ID (0 if in default) ", zone->GetInstanceID());
- int **$instanceversion**: zone instance version 
- int **$zonehour**: zone time hour
- int **$zonemin**: zone time minute
- int **$zonetime**: zone time based on formula (eqTime.hour - 1) * 100 + eqTime.minute)
- int **$zoneweather**: zone weather
- int **$$oncursor**: item ID of item on cursor on player that triggered the event

If NPC has a target, then
- int **$targetid**: target NPC's target entity ID
- string **$Targetname**: target NPC's target name

{{% global_vars_perl %}}