---
title: event_task_accepted
searchTitle: Lua Event event_task_accepted
searchDescription: Triggers when a player talks to a NPC
weight: 1
hidden: true
menuTitle: event_task_accepted
---

#### event_task_accepted

Triggers when a player accepts a task

**Syntax**

function {{% lua_type_functionname event_task_accepted %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_task_accepted(e)
    e.self:task_accepted("Hail " .. e.other:GetName() .. "!");        
end
```

**Exported Variables**

These variables are available with event_task_accepted
- {{% type_client lua %}} **e.other**: player that triggered the event
- {{% type_npc mob %}} **e.self**: mob emitting the event
- {{% lua_type_string %}} **e.message**: player message that triggered the event
- {{% lua_type_number %}} **e.language**: language id the message was in