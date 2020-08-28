---
title: AddAlternateCurrencyValue
searchTitle: Perl Client AddAlternateCurrencyValue
weight: 1
hidden: true
menuTitle: AddAlternateCurrencyValue
---
## AddAlternateCurrencyValue

**Parameter:**

currency\_id _\(uint32\)_ amount _\(int32\)_

**Usage:**

Add the specified amount of the specified currency to the client.

**Example**

```perl
sub EVENT_ITEM_CLICK {
    #:: Match item 57810 - Bag of Doubloons
    if ($itemid == 57810) {
        #:: Give 79910 - Doubloon x100 to the client who triggered the event
        $client->AddAlternateCurrencyValue(79910, 100);
    }
}
```
