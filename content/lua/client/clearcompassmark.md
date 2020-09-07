---
title: ClearCompassMark
searchTitle: Lua EQ ClearCompassMark
searchDescription: Clears compass's current mark.
weight: 1
hidden: true
menuTitle: ClearCompassMark
---

#### ClearCompassMark

Clears compass's current mark.

**Syntax**

client:{{% lua_type_functionname ClearCompassMark %}}() -- {{% lua_type_nil %}}

**Parameters**

**Example**

```lua
function event_say(e)
    e.other:ClearCompassMark();
    e.self:Say("compass cleared");
end
```