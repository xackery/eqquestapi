---
date: 2020-08-24T16:50:16+02:00
title: Event Functions
menuTitle: Event Functions
searchTitle: Event Functions
weight: 25
---

## Event Functions

Note that every event has an enum represented with e.g. Event.event_cast syntax

- [event_aggro_say](aggro_say)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_augment_insert](augment_insert)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_augment_item](augment_item)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_augment_remove](augment_remove)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_cast_begin](cast_begin)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_cast_on](cast_on)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_cast](cast)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_click_door](click_door)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_click_object](click_object)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_combat](combat)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_combine_failure](combine_failure)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_combine_success](combine_success)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_command](command)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_connect](connect)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_death_complete](death_complete)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_death_zone](death_zone)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_death](death)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_destroy_item](destroy_item)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_disconnect](disconnect)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_discover_item](discover_item)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_drop_item](drop_item)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_duel_lose](duel_lose)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_duel_win](duel_win)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_encounter_load](encounter_load)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_encounter_unload](encounter_unload)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_enter_area](enter_area)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_enter_zone](enter_zone)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_enter](enter)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_environmental_damage](environmental_damage)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_equip_item](equip_item)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_exit](exit)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_feign_death](feign_death)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_fish_failure](fish_failure)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_fish_start](fish_start)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_fish_success](fish_success)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_forage_failure](forage_failure)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_forage_success](forage_success)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_group_change](group_change)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_hate_list](hate_list)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_hp](hp)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_item_click_cast](item_click_cast)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_item_click](item_click)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_item_enter_zone](item_enter_zone)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_item_tick](item_tick)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_killed_merit](killed_merit)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_leave_area](leave_area)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_level_up](level_up)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_loot](loot)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_player_pickup](player_pickup)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_popup_response](popup_response)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_proximity_say](proximity_say)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_say](say)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_scale_calc](scale_calc)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_signal](signal)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_slay](slay)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_spawn_zone](spawn_zone)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_spawn](spawn)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_spell_buff_tic_client](spell_buff_tic_client)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_spell_effect_client](spell_effect_client)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_spell_fade](spell_fade)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_target_change](target_change)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_task_accepted](task_accepted)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_task_complete](task_complete)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_task_fail](task_fail)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_task_stage_complete](task_stage_complete)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_task_update](task_update)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_tick](tick)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_timer](timer)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_trade](trade)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_unaugment_item](unaugment_item)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_unequip_item](unequip_item)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_unhandled_opcode](unhandled_opcode)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_use_skill](use_skill)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_waypoint_arrive](waypoint_arrive)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_waypoint_depart](waypoint_depart)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_weapon_proc](weapon_proc)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}
- [event_zone](zone)({{% lua_type_class %}} e) -- {{% lua_type_nil %}}