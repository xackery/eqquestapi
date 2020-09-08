---
title: event_task_fail
searchTitle: Lua Event event_task_fail
searchDescription: Triggers when a player talks to a NPC
weight: 1
hidden: true
menuTitle: event_task_fail
---

#### event_task_fail

Triggers when a player gets a fail on a task

**Syntax**

function {{% lua_type_functionname event_task_fail %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_task_fail(e)
    e.self:task_fail("Hail " .. e.other:GetName() .. "!");        
end
```

**Exported Variables**

These variables are available with event_task_fail
- {{% type_client lua %}} **e.other**: player that triggered the event
- {{% type_npc mob %}} **e.self**: mob emitting the event
- {{% lua_type_string %}} **e.message**: player message that triggered the event
- {{% lua_type_number %}} **e.language**: language id the message was in