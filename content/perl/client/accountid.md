---
title: Perl AccountID
weight: 1
hidden: true
menuTitle: AccountID
---
## AccountID

**Parameter:**

None.

**Usage:**

Returns the account ID of the client.

**Example**

```perl
sub EVENT_SAY {
    if ($text=~/hail/i) {
        #:: Create a scalar variable to store the account ID of the client that triggered the event
        my $AccountID = $client->AccountID();
        #:: Send a message to the client in yellow (15) text
        $client->Message(15, "Your account ID is $AccountID.");
    }
}