---
title: event_say
searchTitle: Lua Event Say
description: Triggers when a player talks to an NPC while targetting them.
weight: 1
hidden: true
menuTitle: event_say
---

#### event_say

Triggers when a player talks to an NPC while targetting them.

**Syntax**
```lua
function event_say(e)
end
```

**Parameters**
- **e**: Event packed argument

**Example**

```lua
function event_say(e)
    e.self:Say("Hail " .. e.other:GetName() .. "!");        
end
```

**Exported Variables**

These variables are available with event_say
- [Client](/en/lua/client) **e.other**: player that triggered the event
- [Mob](/en/lua/mob) **e.self**: mob the message was said to
- string **e.message**: player message that triggered the event
- number **e.language**: language id the message was in