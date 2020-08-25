---
date: 2020-08-24T16:50:16+02:00
title: Inventory
menuTitle: Inventory
weight: 25
---

## Inventory Methods (Lua)
- inventory:[CalcBagIdx](calcbagidx)(int slot_id); -- int
- inventory:[CalcMaterialFromSlot](calcmaterialfromslot)(int equipslot); -- int
- inventory:[CalcSlotFromMaterial](calcslotfrommaterial)(int material); -- int
- inventory:[CalcSlotId](calcslotid)(int slot_id); -- int
- inventory:[CalcSlotId](calcslotid)(int slot_id, int bag_slot); -- int
- inventory:[CanItemFitInContainer](canitemfitincontainer)(Lua_Item item, Lua_Item container); -- bool
- inventory:[CheckNoDrop](checknodrop)(int slot_id); -- bool
- inventory:[DeleteItem](deleteitem)(int slot_id); -- bool
- inventory:[DeleteItem](deleteitem)(int slot_id, int quantity); -- bool
- inventory:[FindFreeSlot](findfreeslot)(bool for_bag, bool try_cursor); -- int
- inventory:[FindFreeSlot](findfreeslot)(bool for_bag, bool try_cursor, int min_size); -- int
- inventory:[FindFreeSlot](findfreeslot)(bool for_bag, bool try_cursor, int min_size, bool is_arrow); -- int
- inventory:[GetItem](getitem)(int slot_id); -- Lua_ItemInst
- inventory:[GetItem](getitem)(int slot_id, int bag_slot); -- Lua_ItemInst
- inventory:[GetSlotByItemInst](getslotbyiteminst)(Lua_ItemInst inst); -- int
- inventory:[HasItem](hasitem)(int item_id); -- int
- inventory:[HasItem](hasitem)(int item_id, int quantity); -- int
- inventory:[HasItem](hasitem)(int item_id, int quantity, int where); -- int
- inventory:[HasItemByLoreGroup](hasitembyloregroup)(uint32 loregroup); -- int
- inventory:[HasItemByLoreGroup](hasitembyloregroup)(uint32 loregroup, int where); -- int
- inventory:[HasItemByUse](hasitembyuse)(int use); -- int
- inventory:[HasItemByUse](hasitembyuse)(int use, uint8 quantity); -- int
- inventory:[HasItemByUse](hasitembyuse)(int use, uint8 quantity, uint8 where); -- int
- inventory:[HasSpaceForItem](hasspaceforitem)(Lua_Item item, int quantity); -- bool
- inventory:[PopItem](popitem)(int slot_id); -- Lua_ItemInst
- inventory:[PushCursor](pushcursor)(Lua_ItemInst item); -- int
- inventory:[PutItem](putitem)(int slot_id, Lua_ItemInst item); -- int
- inventory:[SupportsContainers](supportscontainers)(int slot_id); -- bool
- inventory:[SwapItem](swapitem)(int source_slot, int destination_slot); -- bool