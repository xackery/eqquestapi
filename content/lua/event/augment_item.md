---
title: event_augment_item
searchTitle: Lua Event event_augment_item
searchDescription: Triggers when a player inserts an augment
weight: 1
hidden: true
menuTitle: event_augment_item
---

#### event_augment_item

Triggers when a player inserts an augment

**Syntax**

function {{% lua_type_functionname event_augment_item %}}({{% lua_type_class %}} e)


**Parameters**
- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_augment_item(e)
    e.self:Say("augment " .. e.item:GetName() .. " has been inserted!");
end
```

**Exported Variables**

These variables are available with event_augment_item
- {{% lua_type_item %}} **e.aug**: augment item that triggered the event
- {{% lua_type_number %}} **e.slot_id**: slot ID that the augment was inserted into