Mac utility
===========

Small utility program that quickly converts mac addresses from one format to another and returns vendor.
You can pass the command at runtime, or when prompted.

* Converts single mac addresses
* Accepts multiline input, parses out mac and IP addresses (if present) and converts each one. 
    * To reduce clutter only one format is shown. 
    * Accepts IPv4 and IPv6 addresses.
    * Outputs results to csv
* Windows and Linux versions are in bin

Single input
------------
You can pass a single input at runtime, or be prompted. You must pass ONLY the mac address at this time, there is no regex for single addresses.

Passing mac at runtime:

```
$ ./mac aa-00-00-f0-ed-86
Welcome to mac. For more details visit https://github.com/Jamous/mac

DIGITAL EQUIPMENT
aa-00-00-f0-ed-86
aa00.00f0.ed86
aa:00:00:f0:ed:86
AA:00:00:F0:ED:86
```

Passing mac when prompted:
```
$ ./mac
Welcome to mac. For more details visit https://github.com/Jamous/mac

Mac address or addresses to convert (press enter twice):
aa-00-00-f0-ed-86

DIGITAL EQUIPMENT
aa-00-00-f0-ed-86
aa00.00f0.ed86
aa:00:00:f0:ed:86
AA:00:00:F0:ED:86
```

Multiple input
--------------
You can pass multiple inputs when prompted. I built this to quickly parse through an arp or neighbor table. 

IPv4 example:
```
$ ./mac
Welcome to mac. For more details visit https://github.com/Jamous/mac

Mac address or addresses to convert (press enter twice):
Internet  10.10.16.1              -   68ef.bdb5.fd3f  ARPA   Vlan500
Internet  10.10.18.5              0   0015.65bb.8fcf  ARPA   Vlan500
Internet  10.10.18.10             1   d076.8f71.9ca6  ARPA   Vlan500

10.10.16.1   68:ef:bd:b5:fd:3f   Cisco Systems
10.10.18.5   00:15:65:bb:8f:cf   XIAMEN YEALINK NETWORK TECHNOLOGY
10.10.18.10   d0:76:8f:71:9c:a6   Calix
```

IPv6 example:
```
$ ./mac
Welcome to mac. For more details visit https://github.com/Jamous/mac

Mac address or addresses to convert (press enter twice):
fe80::fa2f:5bff:fe02:64a dev eth1 lladdr f8:2f:5b:02:06:4a STALE
fe80::dab3:70ff:fe70:b0bc dev eth1 lladdr d8:b3:70:70:b0:bc STALE
fe80::ae8b:a9ff:fe65:24dd dev eth1 lladdr ac:8b:a9:65:24:dd STALE

fe80::fa2f:5bff:fe02:64a   f8:2f:5b:02:06:4a   eGauge Systems
fe80::dab3:70ff:fe70:b0bc   d8:b3:70:70:b0:bc   Ubiquiti
fe80::ae8b:a9ff:fe65:24dd   ac:8b:a9:65:24:dd   Ubiquiti
```

Installing
----------
No install is needed. Just add mac to your system path!


Version
-------
0.1 - 05/7/24