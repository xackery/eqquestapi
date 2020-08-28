---
title: AddEXP
weight: 1
hidden: true
menuTitle: AddEXP
---
## AddEXP

**Parameter:**

experience\_points _\(uint32\)_

**Usage:**

Add the specified amount of experience to the client.

**Example**

```perl
sub EVENT_ITEM {
    #:: Match item 123456 - Some Item
    if (plugin::takeItem(123456 => 1) {
        #:: Grant 100 experience to the client who triggered the event
        $client->AddEXP(100);
    }
}
```
