---
title: event_scale_calc
searchTitle: Lua Event event_scale_calc
searchDescription: Triggers when an item triggers a scale calculation
weight: 1
hidden: true
menuTitle: event_scale_calc
---

#### event_scale_calc

Triggers when an item triggers a scale calculation. This can happen when equipped(?), or an aug is inserted/removed

Only usable in the following files:
```
/quests/[global|zone.name]/items/script_[item.scriptfileid].[lua|perl]
/quests/[global|zone.name]/items/[item.charmfile].[lua|perl]
/quests/[global|zone.name]/items/[item.id].[lua|perl]
```

**Syntax**

function {{% lua_type_functionname event_scale_calc %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_scale_calc(e)
    eq.debug(e.self.GetID() .. " triggered a scale event");
end
```

**Exported Variables**

These variables are available with event_scale_calc
- {{% type_client lua %}} **e.owner**: player that triggered the event/owns the item
- {{% lua_type_iteminst %}} **e.self**: item instance that triggered the event
