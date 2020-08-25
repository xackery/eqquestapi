---
date: 2020-08-24T16:50:16+02:00
title: Object
menuTitle: Object
weight: 25
---

## Object Methods (Lua)
- Object:[Depop](depop)(); -- void
- Object:[Repop](repop)(); -- void
- Object:[SetModelName](setmodelname)(const char *name); -- void
- Object:[GetModelName](getmodelname)(); -- string
- Object:[GetX](getx)(); -- number
- Object:[GetY](gety)(); -- number
- Object:[GetZ](getz)(); -- number
- Object:[GetHeading](getheading)(); -- number
- Object:[SetX](setx)(float x); -- void
- Object:[SetY](sety)(float y); -- void
- Object:[SetZ](setz)(float z); -- void
- Object:[SetHeading](setheading)(float h); -- void
- Object:[SetLocation](setlocation)(float x, float y, float z); -- void
- Object:[SetItemID](setitemid)(uint32 item_id); -- void
- Object:[GetItemID](getitemid)(); -- number
- Object:[SetIcon](seticon)(uint32 icon); -- void
- Object:[GetIcon](geticon)(); -- number
- Object:[SetType](settype)(uint32 type); -- void
- Object:[GetType](gettype)(); -- number
- Object:[GetDBID](getdbid)(); -- number
- Object:[ClearUser](clearuser)(); -- void
- Object:[SetID](setid)(number user); -- void
- Object:[GetID](getid)(); -- number
- Object:[Save](save)(); -- bool
- Object:[VarSave](varsave)(); -- number
- Object:[DeleteItem](deleteitem)(number index); -- void
- Object:[StartDecay](startdecay)(); -- void
- Object:[IsGroundSpawn](isgroundspawn)(); -- bool
- Object:[Close](close)(); -- void
- Object:[GetEntityVariable](getentityvariable)(const char *name); -- string
- Object:[SetEntityVariable](setentityvariable)(const char *name, const char *value); -- void
- Object:[EntityVariableExists](entityvariableexists)(const char *name); -- bool
