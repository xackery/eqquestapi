---
title: event_trade
searchTitle: Lua Event event_trade
description: Triggers when a player trades with a NPC
weight: 1
hidden: true
menuTitle: event_trade
---

#### event_trade

Triggers when a player trades with a NPC

**Syntax**

function {{% lua_type_functionname event_trade %}}({{% lua_type_class %}} e) -- {{% lua_type_nil %}}

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_trade(e)
    local item_lib = require("items")
    local turnins = 0
    if (e.other:GetGlobal("greatadventuresturnins") ~= "Undefined") then
        turnins = tonumber(e.other:GetGlobal("greatadventuresturnins"))
    end
    if (item_lib.check_turn_in(e.trade, {item1 = 67617})) then -- Prathun's Writings
        e.self:Emote("takes the dusty tome from " .. e.other:GetName() .. "'s hands.")
        e.self:Say("I thought I would never have to face this reality, but here it is -- the proof I asked for. While these pages detail the adventures of my love, it is you who are the great one. While I shall forever be pained by this, I am in your debt for helping me. Please leave me now and find those others who suffer the continued disappearance of their loved ones. If you have already helped everyone, please tell De'van that your task is complete and he will reward you.")
        e.other:SetGlobal("greatadventuresturnins", tostring(bit.bor(turnins, 1024)), 5, "F")
    end
    item_lib.return_items(e.self, e.other, e.trade)
end
```

**Exported Variables**

These variables are available with event_trade
- {{% type_client lua %}} **e.other**: player that triggered the event
- {{% type_npc mob %}} **e.self**: mob emitting the event
- {{% lua_type_number %}} **e.platinum**: amount in platinum traded
- {{% lua_type_number %}} **e.gold**: amount in gold traded
- {{% lua_type_number %}} **e.silver**: amount in silver traded
- {{% lua_type_number %}} **e.copper**: amount in copper traded
- {{% lua_type_number %}} **e.trade**: not yet determined
- {{% lua_type_table %}} **e.item**[number]: e.g. item0, represents item instances traded