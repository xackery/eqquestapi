---
title: Door
menuTitle: Door
weight: 25
---

## Door Methods (Lua)
- Door:[SetDoorName](setdoorname)(const char *name); -- void
- Door:[GetDoorName](getdoorname)(); -- string
- Door:[GetX](getx)(); -- number
- Door:[GetY](gety)(); -- number
- Door:[GetZ](getz)(); -- number
- Door:[GetHeading](getheading)(); -- number
- Door:[SetX](setx)(float x); -- void
- Door:[SetY](sety)(float y); -- void
- Door:[SetZ](setz)(float z); -- void
- Door:[SetHeading](setheading)(float h); -- void
- Door:[SetLocation](setlocation)(float x, float y, float z); -- void
- Door:[GetDoorDBID](getdoordbid)(); -- number
- Door:[GetDoorID](getdoorid)(); -- number
- Door:[SetSize](setsize)(number sz); -- void
- Door:[GetSize](getsize)(); -- number
- Door:[SetIncline](setincline)(number incline); -- void
- Door:[GetIncline](getincline)(); -- number
- Door:[SetOpenType](setopentype)(number type); -- void
- Door:[GetOpenType](getopentype)(); -- number
- Door:[SetDisableTimer](setdisabletimer)(bool flag); -- void
- Door:[GetDisableTimer](getdisabletimer)(); -- bool
- Door:[SetLockPick](setlockpick)(number pick); -- void
- Door:[GetLockPick](getlockpick)(); -- number
- Door:[SetKeyItem](setkeyitem)(number key); -- void
- Door:[GetKeyItem](getkeyitem)(); -- number
- Door:[SetNoKeyring](setnokeyring)(number type); -- void
- Door:[GetNoKeyring](getnokeyring)(); -- number
- Door:[CreateDatabaseEntry](createdatabaseentry)(); -- void
- Door:[ForceOpen](forceopen)(Lua_Mob sender, bool alt_mode); -- void
- Door:[ForceClose](forceclose)(Lua_Mob sender, bool alt_mode); -- void
