---
title: AddPVPPoints
weight: 1
hidden: true
menuTitle: AddPVPPoints
---
## AddPVPPoints

**Parameter:**

points _\(uint32\)_

**Usage:**

Adds the specified number of PVP points to the client.

**Example**

```perl
sub EVENT_ITEM {
    #:: Match item 123456 - Some Item
    if (plugin::takeItem(123456 => 1) {
        #:: Grant one pvp point to the client who triggered the event
        $client->AddPVPPoints(1);
    }
}
```
