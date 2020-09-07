---
title: Real
searchTitle: Lua EQ Real
searchDescription: Rolls a random real number. It is normally suggested to use math.Random(low, high) if whole numbers are acceptable.
weight: 1
hidden: true
menuTitle: Real
---

## Real

Rolls a random real number. It is normally suggested to use math.Random(low, high) if whole numbers are acceptable.

**Syntax**

Random.Real({{% lua_type_number %}} low, {{% lua_type_number %}} high) -- {{% lua_type_number %}}

**Parameters**

- {{% lua_type_number %}} **low**: low number inclusive
- {{% lua_type_number %}} **high**: high number inclusive

**Example**

```lua
function event_say(e)
   e.self:Say("random number from 0 to 100 (floating precision): " .. Random.Real(0, 100));
end
```