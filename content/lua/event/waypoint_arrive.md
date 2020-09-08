---
title: event_waypoint_arrive
searchTitle: Lua Event event_waypoint_arrive
searchDescription: Triggers when a NPC aggros anything
weight: 1
hidden: true
menuTitle: event_waypoint_arrive
---

#### event_waypoint_arrive

Triggers when a NPC arrives at a waypoint

**Syntax**

function {{% lua_type_functionname event_waypoint_arrive %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_waypoint_arrive(e)
    e.self:Say("I have arrived at a waypoint");
end
```

**Exported Variables**

These variables are available with event_waypoint_arrive
- {{% type_npc mob %}} **e.self**: mob emitting the event