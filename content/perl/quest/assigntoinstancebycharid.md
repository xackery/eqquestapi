---
title: AssignToInstanceByCharID
searchTitle: Perl Quest AssignToInstanceByCharID
weight: 1
hidden: true
menuTitle: AssignToInstanceByCharID
searchDescription: Assigns a single player to an instance via character ID
---

#### AssignToInstanceByCharID

Assigns a single player to an instance via character ID

**Syntax**
```perl
quest::AssignToInstanceByCharID(int $instanceID, int $characterID); # void
```

**Parameters**
- **instanceID**: which instance to assign player to
- **characterID**: which character ID to assign to instance


**Example**

```perl
sub EVENT_SAY {
    # zone short name. mmca is an example of a LDON zone with different versions
    my $zoneName = "mmca"
    # all zones default to version 0 as their static normal version.
    my $version = 1;
    # time in seconds for the instance to time out
    my $duration = 5400;

    # check if character is already assigned to an instance
    my $instanceID = quest::GetInstanceIDByCharID($zoneName, $version, $client->CharacterID());
    if (!$instanceID) { 
        # if not, create an instance
        $instanceID = quest::CreateInstance($zoneName, $version, $duration);
    }

    quest::AssignToInstanceByCharID($instanceID, $client->CharacterID());
    quest::say("you are part of instance # $instanceID");
}
```
