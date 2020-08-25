---
title: Perl AddAAPoints
weight: 1
hidden: true
menuTitle: AddAAPoints
---
### AddAAPoints

**Parameter:**

points _\(uint32\)_

**Usage:**

Add the specified number of AA points to the client.

**Example**

```perl
sub EVENT_ITEM_CLICK {
    #:: Match item 123456 - Some Item
    if ($item == 123456) {
        #:: Award 20 AA points to the client who triggered the event
        $client->AddAAPoints(20);
    }
}
```