---
title: Perl AccountName
weight: 1
hidden: true
menuTitle: AccountName
---
## AccountName

**Parameter:**

None.

**Usage:**

Returns the account name of the client.

**Example**

```perl
sub EVENT_SAY {
    if ($text=~/hail/i) {
        #:: Create a scalar variable to store the account name of the client that triggered the event
        my $AccountName = $client->AccountName();
        #:: Send a message to the client in yellow (15) text
        $client->Message(15, "Your account name is $AccountName.");
    }
}
```