---
title: Perl AddCrystals
weight: 1
hidden: true
menuTitle: AddCrystals
---

## AddCrystals

**Parameter:**

radiant\_count _\(uint32\)_, ebon\_count _\(uint32\)_

**Usage:**

Add the specified amount of the Radiant or Ebon crystals to the client.

**Example**

```perl
sub EVENT_ITEM_CLICK {
    #:: Match item 57816 - Bag of Ebon Crystals
    if ($itemid == 57816) {
        #:: Give 40902 - Ebon Crystal x100 to the client who triggered the event
        $client->AddCrystals(0, 100);
    }
}
```