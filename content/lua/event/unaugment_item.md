---
title: event_item_unaugment_item
searchTitle: Lua Event event_item_unaugment_item
searchDescription: Triggers when a player equips an item
weight: 1
hidden: true
menuTitle: event_item_unaugment_item
---

#### event_item_unaugment_item

Triggers when a player equips an item

Only usable in the following files:
```
/quests/[global|zone.name]/items/script_[item.scriptfileid].[lua|perl]
/quests/[global|zone.name]/items/[item.charmfile].[lua|perl]
/quests/[global|zone.name]/items/[item.id].[lua|perl]
```

**Syntax**

function {{% lua_type_functionname event_item_unaugment_item %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_item_unaugment_item(e)
    eq.debug(e.item.GetID() .. "triggered an unaugment_item");
end
```

**Exported Variables**

These variables are available with event_item_unaugment_item
- {{% type_client lua %}} **e.owner**: player that triggered the event/owns the item
- {{% lua_type_iteminst %}} **e.self**: item instance that triggered the event