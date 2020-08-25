---
title: Perl AutoSplitEnabled
weight: 1
hidden: true
menuTitle: AutoSplitEnabled
---

## AutoSplitEnabled

**Parameter:**

None.

**Usage:**

Turns on auto-split \(coin\) for the client.

**Example**

```perl
sub EVENT_GROUP_CHANGE {
    #:: Turn on auto-split
    $client->AutoSplitEnabled()
}
```
