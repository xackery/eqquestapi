---
title: event_augment_remove
searchTitle: Lua Event event_augment_remove
description: Triggers when a player removes an augment
weight: 1
hidden: true
menuTitle: event_augment_remove
---

#### event_augment_remove

Triggers when a player removes an augment

**Syntax**

function {{% lua_type_functionname event_augment_remove %}}({{% lua_type_class %}} e)


**Parameters**
- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_augment_remove(e)
    e.self:Say("augment " .. e.item:GetName() .. " was removed!");
end
```

**Exported Variables**

These variables are available with event_augment_remove
- {{% lua_type_item %}} **e.item**: item that triggered the event
- {{% lua_type_number %}} **e.slot_id**: slot ID that the augment was inserted into
- {{% lua_type_boolean %}} **e.destroyed**: returns true if augment was destroyed while being removed