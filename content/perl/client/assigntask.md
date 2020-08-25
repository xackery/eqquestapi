---
title: Perl AssignTask
weight: 1
hidden: true
menuTitle: AssignTask
---

## AssignTask

**Parameter:**

task\_id _\(int\)_, npc\_id _\(int\)_, enforce\_level\_requirement _\(bool\) \[Default = false\]_

**Usage:**

Assigns the specified task, credited to the specified NPC, to the client.  Enforcing the level requirement of the task is optional, and defaults to false.

**Example**

```perl
#:: Assign Task 105, credit the npc by id, with level requirement enforced
$client->AssignTask(105, $npc->GetID(), 1);
#:: Assign Task 104, credit the npc by id, ignore level requirement
$client->AssignTask(104, $npc->GetID());
```
