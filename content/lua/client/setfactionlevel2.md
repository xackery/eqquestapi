---
title: SetFactionLevel2
searchTitle: Lua EQ SetFactionLevel2
searchDescription: Clears compass's current mark.
weight: 1
hidden: true
menuTitle: SetFactionLevel2
---

#### SetFactionLevel2

Sets a faction level

**Syntax**

client:{{% lua_type_functionname SetFactionLevel2 %}}{{% lua_type_number %}} char_id, {{% lua_type_number %}} npc_id, {{% lua_type_number %}} char_class, {{% lua_type_number %}} char_race, {{% lua_type_number %}} char_deity, {{% lua_type_number %}} value, {{% lua_type_number %}} temp) -- {{% lua_type_nil %}}

**Parameters**

- **char_id**: character id to match
- **npc_id**: npc type id to match 
- **char_class**: class id to match
- **char_race**: race id to match
- **char_deity**: deity id to match
- **value**: value to set faction 2
- **temp**: is it temporary or not? 0 or 1

**Example**

```lua
function event_say(e)
    e.other:SetFactionLevel2(e.other:GetID(), e.self:GetNpcTypeID(), e.other:GetClass(), e.other:GetRace(), e.other:GetDeity(), 1000, 1);
    e.self:Say("setting faction level 2 to 1000");
end
```