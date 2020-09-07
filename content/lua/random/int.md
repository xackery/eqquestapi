---
title: Int
searchTitle: Lua EQ Int
searchDescription: Rolls a random whole number. Alternatively can use math.random(low, high) for built in lua randomizing.
weight: 1
hidden: true
menuTitle: Int
---

## Int

Rolls a random whole number. Alternatively can use math.random(low, high) for built in lua randomizing.

**Syntax**

Random.Int({{% lua_type_number %}} low, {{% lua_type_number %}} high) -- {{% lua_type_number %}}

**Parameters**

- {{% lua_type_number %}} **low**: low number inclusive
- {{% lua_type_number %}} **high**: high number inclusive

**Example**

```lua
function event_say(e)
   e.self:Say("random number from 0 to 100: " .. Random.Int(0, 100));
end
```