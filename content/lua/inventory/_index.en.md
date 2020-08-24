---
date: 2016-04-09T16:50:16+02:00
title: Inventory
menuTitle: Inventory
weight: 25
---

# Inventory Methods \(Lua\)

- inventory:[CalcBagIdx(int slot_id)](inventory) -- int
- inventory:[CalcMaterialFromSlot(int equipslot)](calcmaterialfromslot) -- int
- inventory:[CalcSlotFromMaterial(int material)](calcslotfrommaterial) -- int
- inventory:[CalcSlotId(int slot_id)](calcslotid) -- int
- inventory:[CalcSlotId(int slot_id, int bag_slot)](calcslotid) -- int
- inventory:[CanItemFitInContainer(Lua_Item item, Lua_Item container)](canitemfitincontainer) -- bool
- inventory:[CheckNoDrop(int slot_id)](checknodrop) -- bool
- inventory:[DeleteItem(int slot_id)](deleteitem) -- bool
- inventory:[DeleteItem(int slot_id, int quantity)](deleteitem) -- bool
- inventory:[FindFreeSlot(bool for_bag, bool try_cursor)](findfreeslot) -- int
- inventory:[FindFreeSlot(bool for_bag, bool try_cursor, int min_size)](findfreeslot) -- int
- inventory:[FindFreeSlot(bool for_bag, bool try_cursor, int min_size, bool is_arrow)](findfreeslot) -- int
- inventory:[GetItem(int slot_id)](getitem) -- Lua_ItemInst
- inventory:[GetItem(int slot_id, int bag_slot)](getitem) -- Lua_ItemInst
- inventory:[GetSlotByItemInst(Lua_ItemInst inst)](getslotbyiteminst) -- int
- inventory:[HasItem(int item_id)](hasitem) -- int
- inventory:[HasItem(int item_id, int quantity)](hasitem) -- int
- inventory:[HasItem(int item_id, int quantity, int where)](hasitem) -- int
- inventory:[HasItemByLoreGroup(uint32 loregroup)](hasitembyloregroup) -- int
- inventory:[HasItemByLoreGroup(uint32 loregroup, int where)](hasitembyloregroup) -- int
- inventory:[HasItemByUse(int use)](hasitembyuse) -- int
- inventory:[HasItemByUse(int use, uint8 quantity)](hasitembyuse) -- int
- inventory:[HasItemByUse(int use, uint8 quantity, uint8 where)](hasitembyuse) -- int
- inventory:[HasSpaceForItem(Lua_Item item, int quantity)](hasspaceforitem) -- bool
- inventory:[PopItem(int slot_id)](popitem) -- Lua_ItemInst
- inventory:[PushCursor(Lua_ItemInst item)](pushcursor) -- int
- inventory:[PutItem(int slot_id, Lua_ItemInst item)](putitem) -- int
- inventory:[SupportsContainers(int slot_id)](supportscontainers) -- bool
- inventory:[SwapItem(int source_slot, int destination_slot)](swapitem) -- bool