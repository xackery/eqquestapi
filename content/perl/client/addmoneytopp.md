---
title: Perl AddMoneyToPP
weight: 1
hidden: true
menuTitle: AddMoneyToPP
---

## AddMoneyToPP

**Parameter:**

copper _\(uint32\)_, silver _\(uint32\)_, gold _\(uint32\)_, platinum _\(uint32\)_, update\_client _\(bool\)_

**Usage:**

Adds the amount of currency specified to the client.

**Example**

```perl
sub EVENT_ITEM {
    #:: Match item 123456 - Some Item
    if (plugin::takeItem(123456 => 1) {
        #:: Give 1cp, 2sp, 3gp, 4pp to the client who triggered the event, but don't tell them about it
        $client->AddMoneyToPP(1, 2, 3, 4, 0);
    }
}
```
