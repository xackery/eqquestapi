---
title: event_augment_insert
searchTitle: Lua Event event_augment_insert
description: Triggers when a player inserts an augment
weight: 1
hidden: true
menuTitle: event_augment_insert
---

#### event_augment_insert

Triggers when a player inserts an augment

**Syntax**

function {{% lua_type_functionname event_augment_insert %}}({{% lua_type_class %}} e)


**Parameters**
- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_augment_insert(e)
    e.self:Say("item " .. e.item:GetName() .. " has an augment!");
end
```

**Exported Variables**

These variables are available with event_augment_insert
- {{% lua_type_item %}} **e.item**: item that triggered the event
- {{% lua_type_number %}} **e.slot_id**: slot ID that the augment was inserted into