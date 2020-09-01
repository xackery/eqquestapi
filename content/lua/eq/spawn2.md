---
title: spawn2
searchTitle: Lua EQ spawn2
description: Spawns an NPC Type ID with provided details
weight: 1
hidden: true
menuTitle: spawn2
---

## spawn2

Spawns an NPC Type ID with provided details

**Syntax**

eq.spawn2({{% lua_type_number %}} npc_type, {{% lua_type_number %}} grid, {{% lua_type_number %}} unused, {{% lua_type_number %}} x, {{% lua_type_number %}} y, {{% lua_type_number %}} z) -- {{% lua_type_mob %}}

eq.spawn2({{% lua_type_number %}} npc_type, {{% lua_type_number %}} grid, {{% lua_type_number %}} unused, {{% lua_type_number %}} x, {{% lua_type_number %}} y, {{% lua_type_number %}} z, {{% lua_type_number %}} heading) -- {{% lua_type_mob %}}

**Parameters**

- {{% lua_type_number %}} **npc_type**: NPC Type ID to spawn
- {{% lua_type_number %}} **grid**: grid ID to walk. Default is 0 for none
- {{% lua_type_number %}} **unused**: Not used at this time
- {{% lua_type_number %}} **x**: X position to spawn at
- {{% lua_type_number %}} **y**: Y position to spawn at
- {{% lua_type_number %}} **z**: Z position to spawn at
- {{% lua_type_number %}} **heading**: direction to face when spawned

**Example**

```lua
function event_say(e)
   eq.spawn2(154129, 0, 0, e.self:GetX(), e.self:GetY(), e.self:GetZ(), 0);
   e.self:Say("spawned a grimling")
end
```