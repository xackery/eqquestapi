---
date: 2020-08-24T16:50:16+02:00
title: Inventory
menuTitle: Inventory
weight: 25
---

## Inventory Methods (Lua)
- Inventory:[GetItem](getitem)(number slot_id, number bag_slot); -- unknown - Lua_ItemInst
- Inventory:[PutItem](putitem)(number slot_id, Lua_ItemInst item); -- number
- Inventory:[PushCursor](pushcursor)(Lua_ItemInst item); -- number
- Inventory:[SwapItem](swapitem)(number source_slot, number destination_slot); -- bool
- Inventory:[DeleteItem](deleteitem)(number slot_id, number quantity); -- bool
- Inventory:[CheckNoDrop](checknodrop)(number slot_id); -- bool
- Inventory:[PopItem](popitem)(number slot_id); -- unknown - Lua_ItemInst
- Inventory:[HasItem](hasitem)(number item_id, number quantity, number where); -- number
- Inventory:[HasSpaceForItem](hasspaceforitem)(Lua_Item item, number quantity); -- bool
- Inventory:[HasItemByUse](hasitembyuse)(number use, uint8 quantity, uint8 where); -- number
- Inventory:[HasItemByLoreGroup](hasitembyloregroup)(uint32 loregroup, number where); -- number
- Inventory:[FindFreeSlot](findfreeslot)(bool for_bag, bool try_cursor, number min_size, bool is_arrow); -- number
- Inventory:[CalcSlotId](calcslotid)(number slot_id, number bag_slot); -- number
- Inventory:[CalcBagIdx](calcbagidx)(number slot_id); -- number
- Inventory:[CalcSlotFromMaterial](calcslotfrommaterial)(number material); -- number
- Inventory:[CalcMaterialFromSlot](calcmaterialfromslot)(number equipslot); -- number
- Inventory:[CanItemFitInContainer](canitemfitincontainer)(Lua_Item item, Lua_Item container); -- bool
- Inventory:[SupportsContainers](supportscontainers)(number slot_id); -- bool
- Inventory:[GetSlotByItemInst](getslotbyiteminst)(Lua_ItemInst inst); -- number
