---
title: Perl setskill
weight: 1
hidden: true
menuTitle: setskill
---
## setskill
```perl
$quest->setskill(int skill_id, int value)
```
**Parameters:**

skill\_id _\(int\)_, value _\(int\)_

**Usage:**

Used to set the specified [Skill](https://github.com/EQEmu/Server/wiki/Skills) to the specified value.

**Example:**
```perl
#:: Set baking to 100
$quest->setskill(60, 100);
```