---
title: Perl AssignToInstance
weight: 1
hidden: true
menuTitle: AssignToInstance
---

## AssignToInstance

**Parameter:**

instance\_id _\(uint16\)_

**Usage:**

Assigns the specified instance to the client.

**Example**

```perl
#:: Create a scalar variable to store the instance id
$InstanceID = quest::CreateInstance("soldungb", 0, 86400); #:: Instance of Sol B, based on version 0, for 86400 seconds (1 day)
#:: Assign the client to the instance
$client->AssignToInstance($InstanceID);
```
