---
title: Perl Admin
weight: 1
hidden: true
menuTitle: Admin
---

## Admin

**Parameter:**

None.

**Usage:**

Returns the Admin account status of the client.
{% hint style="info" %}
Note that `$status` is exported by many events, should it prove more useful.
{% endhint %}

      **Example**

```perl
sub EVENT_SAY {
    if ($text=~/hail/i) {
        #:: Create a scalar variable to store admin status
        $GMstatus = $client->Admin();
        #:: Match if the account status is greater than one
        if ($GMstatus > 1) {
            quest::say("How ya doin', boss?");
        }
        else {
            quest::say("Hello, $name");
        }
    }
}
```
