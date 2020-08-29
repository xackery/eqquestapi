---
title: Inventory
menuTitle: Inventory
description: Lua Inventory Class
searchTitle: Lua Inventory Class
weight: 25
---

## Inventory Methods
- [GetItem](getitem)({{% lua_type_number %}} slot_id, {{% lua_type_number %}} bag_slot) -- {{% lua_type_iteminst %}}
- [PutItem](putitem)({{% lua_type_number %}} slot_id, {{% lua_type_iteminst %}} item) -- {{% lua_type_number %}}
- [PushCursor](pushcursor)({{% lua_type_iteminst %}} item) -- {{% lua_type_number %}}
- [SwapItem](swapitem)({{% lua_type_number %}} source_slot, {{% lua_type_number %}} destination_slot) -- {{% lua_type_boolean %}}
- [DeleteItem](deleteitem)({{% lua_type_number %}} slot_id, {{% lua_type_number %}} quantity) -- {{% lua_type_boolean %}}
- [CheckNoDrop](checknodrop)({{% lua_type_number %}} slot_id) -- {{% lua_type_boolean %}}
- [PopItem](popitem)({{% lua_type_number %}} slot_id) -- {{% lua_type_iteminst %}}
- [HasItem](hasitem)({{% lua_type_number %}} item_id, {{% lua_type_number %}} quantity, {{% lua_type_number %}} where) -- {{% lua_type_number %}}
- [HasSpaceForItem](hasspaceforitem)({{% lua_type_item %}} item, {{% lua_type_number %}} quantity) -- {{% lua_type_boolean %}}
- [HasItemByUse](hasitembyuse)({{% lua_type_number %}} use, {{% lua_type_number %}} quantity, {{% lua_type_number %}} where) -- {{% lua_type_number %}}
- [HasItemByLoreGroup](hasitembyloregroup)({{% lua_type_number %}} loregroup, {{% lua_type_number %}} where) -- {{% lua_type_number %}}
- [FindFreeSlot](findfreeslot)({{% lua_type_boolean %}} for_bag, {{% lua_type_boolean %}} try_cursor, {{% lua_type_number %}} min_size, {{% lua_type_boolean %}} is_arrow) -- {{% lua_type_number %}}
- [CalcSlotId](calcslotid)({{% lua_type_number %}} slot_id, {{% lua_type_number %}} bag_slot) -- {{% lua_type_number %}}
- [CalcBagIdx](calcbagidx)({{% lua_type_number %}} slot_id) -- {{% lua_type_number %}}
- [CalcSlotFromMaterial](calcslotfrommaterial)({{% lua_type_number %}} material) -- {{% lua_type_number %}}
- [CalcMaterialFromSlot](calcmaterialfromslot)({{% lua_type_number %}} equipslot) -- {{% lua_type_number %}}
- [CanItemFitInContainer](canitemfitincontainer)({{% lua_type_item %}} item, {{% lua_type_item %}} container) -- {{% lua_type_boolean %}}
- [SupportsContainers](supportscontainers)({{% lua_type_number %}} slot_id) -- {{% lua_type_boolean %}}
- [GetSlotByItemInst](getslotbyiteminst)({{% lua_type_iteminst %}} inst) -- {{% lua_type_number %}}
