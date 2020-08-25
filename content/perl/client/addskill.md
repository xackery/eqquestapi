---
title: Perl AddSkill
weight: 1
hidden: true
menuTitle: AddSkill
---

## AddSkill

**Parameter:**

skill\_id _\(int\)_, value _\(uint16\)_

**Usage:**

Adds the specified skill level value to the specified skill for the client.
{% hint style="warning" %}
Note the difference between $client-&gt;AddSkill and $client-&gt;SetSkill
{% endhint %}

      **Example**

```perl
sub EVENT_COMBINE_SUCCESS {
    #:: Match Recipe 2686 - Hand Made Backpack by ID
    if ($recipe_id == 2686) {
        #:: Send the client a message in color 15 (yellow)
        $client->Message(15,"Wow, you did it--you won tailoring!");
        #:: Add 300 points to the Tailoring skill for the client that triggered the event
        $client->AddSkill(61, 300);
    }
}
```
