---
title: event_aggro_say
searchTitle: Lua Event event_aggro_say
description: Triggers when an NPC aggros anything
weight: 1
hidden: true
menuTitle: event_aggro_say
---

#### event_aggro_say

Triggers when an NPC aggros anything

**Syntax**

function {{% lua_type_functionname event_aggro_say %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_aggro_say(e)
    e.self:Say("Hail " .. e.other:GetName() .. "!");        
end
```

**Exported Variables**

These variables are available with event_aggro_say
- {{% lua_type_client %}} **e.other**: player that triggered the event
- {{% lua_type_mob %}} **e.self**: mob emitting the event
- {{% lua_type_string %}} **e.message**: player message that triggered the event
- {{% lua_type_number %}} **e.language**: language id the message was in