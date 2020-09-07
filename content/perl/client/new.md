---
title: new
searchTitle: Perl Client new
searchDescription: new creates a new instance reference to the client calling the script
weight: 1
hidden: true
menuTitle: new
---

#### new

new creates a new instance reference to the client calling the script

**Syntax**

$client->{{% lua_type_functionname new %}}() -- {{% type_client lua %}}

**Parameters**

**Example**

```perl
sub EVENT_SAY {
    my $newClient = $client->new();
    quest::say("new client is "+$newClient->Name());
}
```