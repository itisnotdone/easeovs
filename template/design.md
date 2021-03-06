# Public Network Block
- Use 172.16.0.0/12, 172.16.0.0 - 172.31.255.255
- 17 bits network block per a region and 18 bits network block per a fabric
- aregion.dev: 172.30.0.0/17
  - 1st fabric in aregion: 172.30.0.0/18
  - 2nd fabric in aregion: 172.30.64.0/18
- bregion.dev: 172.30.128.0/17
  - 1st fabric in bregion: 172.30.128.0/18
  - 2nd fabric in bregion: 172.30.192.0/18
- default.dev: 172.31.0.0/24

# Private Network Block
- Use 192.168.0.0/16, 192.168.0.0 - 192.168.255.255

# VID definition for public VLANs
- For the 1st fabric in the 1st region, 10 will be added to the 3rd octet of their network ID. i.e. 172.30.0.0 will have VID 10.
- Adding 10 is to avoid using pre-reserved VIDs such as 0 and 1.
  - https://www.cisco.com/c/en/us/td/docs/switches/lan/catalyst6500/ios/12-2SX/configuration/guide/book/vlans.html
  - From 2 to 1001 will be used as VID

# VID definition for private(internal only) VLANs
For internal VLANs, 800 will be added to 3rd octet of the base CIDR. i.e., 800 from 172.31.0.0, 992 from 172.31.192.0
