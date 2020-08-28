---
title: AddLevelBasedExp
weight: 1
hidden: true
menuTitle: AddLevelBasedExp
---

## AddLevelBasedExp

**Parameter:**

exp\_percentage _\(uint8\)_, max\_level _\(uint8\) \[default = 0\],_ ignore\_mods _\(bool\) \[default = false\]_

**Usage:**

Adds the specified percentage of experience at each level, up until the maximum level specified, to the client.  If the player's level exceeds the maximum level specified, the player would only be rewarded the percentage of the experience at the specified level.

**Example**

```perl
sub EVENT_ITEM {
    #:: Match item 123456 - Some Item
    if (plugin::takeItem(123456 => 1) {
        #:: Grant 4 percent experience to the client who triggered the event up until level 15
        $client->AddLevelBasedExp(4, 15);
    }
}
```
