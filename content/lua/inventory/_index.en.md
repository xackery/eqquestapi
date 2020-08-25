---
date: 2020-08-24T16:50:16+02:00
title: Inventory
menuTitle: Inventory
weight: 25
---

## Inventory Methods (Lua)
- Inventory:[GetItem](getitem)(int slot_id, int bag_slot); -- unknown - Lua_ItemInst
- Inventory:[PutItem](putitem)(int slot_id, Lua_ItemInst item); -- number
- Inventory:[PushCursor](pushcursor)(Lua_ItemInst item); -- number
- Inventory:[SwapItem](swapitem)(int source_slot, int destination_slot); -- bool
- Inventory:[DeleteItem](deleteitem)(int slot_id, int quantity); -- bool
- Inventory:[CheckNoDrop](checknodrop)(int slot_id); -- bool
- Inventory:[PopItem](popitem)(int slot_id); -- unknown - Lua_ItemInst
- Inventory:[HasItem](hasitem)(int item_id, int quantity, int where); -- number
- Inventory:[HasSpaceForItem](hasspaceforitem)(Lua_Item item, int quantity); -- bool
- Inventory:[HasItemByUse](hasitembyuse)(int use, uint8 quantity, uint8 where); -- number
- Inventory:[HasItemByLoreGroup](hasitembyloregroup)(uint32 loregroup, int where); -- number
- Inventory:[FindFreeSlot](findfreeslot)(bool for_bag, bool try_cursor, int min_size, bool is_arrow); -- number
- Inventory:[CalcSlotId](calcslotid)(int slot_id, int bag_slot); -- number
- Inventory:[CalcBagIdx](calcbagidx)(int slot_id); -- number
- Inventory:[CalcSlotFromMaterial](calcslotfrommaterial)(int material); -- number
- Inventory:[CalcMaterialFromSlot](calcmaterialfromslot)(int equipslot); -- number
- Inventory:[CanItemFitInContainer](canitemfitincontainer)(Lua_Item item, Lua_Item container); -- bool
- Inventory:[SupportsContainers](supportscontainers)(int slot_id); -- bool
- Inventory:[GetSlotByItemInst](getslotbyiteminst)(Lua_ItemInst inst); -- number
