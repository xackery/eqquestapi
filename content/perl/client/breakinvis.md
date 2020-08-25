---
title: Perl BreakInvis
weight: 1
hidden: true
menuTitle: BreakInvis
---

## BreakInvis

**Parameter:**

None.

**Usage:**

Turns off invisibility effects for the client.

**Example**

```perl
sub EVENT_CLICKDOOR {
    #:: Turn off invis for any client that clicks the door
    $client->BreakInvis()
}
```