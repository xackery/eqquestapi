---
title: Message
weight: 1
hidden: true
menuTitle: Message
---
## Message

**Parameter:**

[emote_color_type]({{<ref "#Emote Colors">}}) _uint32_ message _string_

**Usage:**

Sends a message from a mob

**Example**

```perl
sub EVENT_SAY {
    if ($text=~/hail/i) {
        #:: Send a message in yellow (15) text
        $mob->Message(15, "Test");
    }
}
```
{{% emote_colors %}}