Inter System Communication Protocols
	- USB
	- UART
	- USART

Intra System Communication Protocols
	- I2C
	- SPI
	- CAN


USB - Universal Serial Bus
	- Two wired serial communication
	- 127 devices at any given time
	- Vcc, GND, Data+ and Data- pins
	- Fast Simple, Low Cost, Plug and Play
	- Needs a powerful master device, and specific drivers

UART - Universal Asynchronous Receiver/Transmitter 
	- Converts parallel data into serial data
	- Two wired protocol with Tx and Rx pins
	- Sends data in asynchronous fashion (i.e. no clock signal)
	- Embeds start and stop bits with data bits to define beginning and end of frame
	- Rx detects start bit begins to read data at a baud rate
	- Half duplex communication mode either tx or rx at any given time not both

USART - Universal Synchronous Asynchronous Receiver/Transmitter
	- Identical to UART just with added functional synchronous 
	- Tx generates a clock signal which is seen on the Rx end from the data stream
	- Needs not to agree on a baud rate
	- Is full duplex so you may send and receive at the same time
	- Has all the abilities of UART so you can use it in place
	- Has parity bits, no clock, cost effective, 2 wire communication
	- No multiple master slave functions, baud rates need be ~10% of each other



I2C - Inter Integrated Circuit (I2C)
	- Serial communication
	- Connect peripheral chips with microcontrollers
	- Two wire system to carry information
	- SDA (Serial Data Line), SCL (Serial Clock Line)
	- Both wires are bidirectional
	- Master slave protocol. Each slave is assigned a unique address.
	- Master establishes communication by sending the target slave address and R/W flag.
	- Master drives the SCL clock line
	- SDA line low when SCL is high (unique condition) is a start/stop sequence
		○ SDA goes low SCL stays high -> Start Sequence
		○ SCL goes high SCL stays low -> Stop Sequence
	- Slave device then moves into active mode other devices are left off
	- Once slave is ready communication begins.
	- One bit ACK is replied by receiver if transmitter transmits 1 byte of data
	- Stop condition issued at end of communication
	- Gives good communication between devices that are accessed infrequently
	- Addresses ease master slave communication
	- Cost and circuit complexity does not end up on number of devices
	- Limited in speed however.
	
SPI - Serial Peripheral Interface
	- 4 wire protocol that uses MOSI (Master Out Slave In) and MISO (Master In Slave Out), SS (Slave Select), and SCLK (Serial Clock)
	- It’s a master slave communication protocol
	- Configures clock at a frequency
	- SS line used to select the correct slave by pulling the line low (normally high)
		○ SS line per slave device from master
	- Communication is established when correct slave is selected.
	- Full duplex and doesn't limit data to 8 bit words
	- Faster than asynchronous serial communication protocol
	- Multiple slaves supported, universally accepted, cheap.
	- Requires more wires, master should control all communications (slave-slave impossible)
	- Circuit complexity with increasing number of devices.

CAN - Control Area Network
	- Serial communication.
	- Tow wire CAN High and CAN Low.
	- High and Low must agree or an error has occurred.
	- Message Oriented communication.
	- Low cost, reliable, robust, secured and fast.
	- Somewhat complex and automotive oriented.
	- All devices listen on the bus and decided at destination if the data is worth keeping
	- Can be prioritized by ID
	- Basic concept is that all devices can transfer data, bit by bit the ids are compared and sent together as long as they are the same. As soon as one drops to a lower priority that frame is dropped from the network and must try again after transfer of the most important one has completed. I.e. they go into back-off when a recessive bit is sent when another frame sends a dominant bit.
	- Automotive ethernet exists to solve the problem of low throughput
	- CSMA/CD protocol so nodes on the bus can back-off.
	
	
Ethernet 802.3
	- Use MAC Addresses and switches to move link layer datagrams
	- MAC - Multiple Access Protocol
	- CSMA/CD Carrier Sensing Multiple Access w/ Collision Detection
		○ Algorithm to listen before transmit and handle packet collisions over shared mediums.
	- ARP has each adapter with unique LAN MAC Address
	- ARP table made to have IP address, MAC Address and TTL
		○ Every 20 minutes entries are forgotten in switch table
	- Bus vs Star patterns
		○ Shared coaxial cable vs center switch one device per bus
	- Frame encapsulates IP data frame. Has destination address, source address data and the CRC to correct bits
	- Switch tables
		○ Broadcast can cause packet storms
		○ Tables constructed as needed.
		○ Mac Address, Interface and TTL 
	
Wi-Fi 802.11
	- DCF or PCF: Distributed Coordination Function, Point Coordination Function
		○ DCF has handshake mode or without handshake depending on length of packet
		○ PCF has the access point as the master and decides who transfers
	- Exposed Terminal Problem
		○ Detecting waves from out of range device don't know where it is though so cannot know if it can transfer to something where the waves will not collide. No concept of space; wasted time; no solution.
	- Hidden Terminal Problem
		○ No solution for where two out of range devices try to tx to a common in range device.

IP - Internet Protocol
	- Uses IP addresses to uniquely identify internet connected devices
	- Sender specifies sender and receiver IP addresses
	- IPv4 and IPv6 addresses
	- Public and Private IP addresses
	- Fragmented and has checksums
	- Routers are able to route the IP datagrams which hold link layer datagrams
	- IP masks used to show grouping of addresses
	- Routing tables have destination address, mask, next hop, and interface
	- TTL after exceeding threshold it drops the packet.
	- Private to Public Addresses need a Network Address Translation table
	- Routing tables constructed with algorithms such as RIP (Routing information protocol), OSPF (Open shortest path first) and BGP (Border gateway protocol).

TCP - Transmission Control Protocol
	- Safe, reliable transport layer protocol
	- Used commonly on top of IP so called TCP/IP stack lots
	- Has source port, destination port, sequence numbers, ack numbers, window sizes, checksum, options, and data.
		○ Some others as well.
		○ SYN, ACK bits used to synchronize and acknowledge
	- First computer sends SYN bit to ask to establish connection
	- Second sends back the ACK bit and SYN bit
	- First device sends back with an ACK and transfer can begin for data.
	- Each data packet can be used to send ACKs to previous packets
	- If ACK is missed sender just resends until an ACK comes
	- Timer kept to wait to see for 
	- Receiver will ignore extra packets and wait until next sequence packet comes dropping others.
	- Keeps sending same ACK for the next order packet

UDP - User Datagram Protocol 
	- Lightweight data transport layer protocol
	- Uses ports with source port, destination port the length and a checksum
	- Cannot have resends
	- Simple and fast
	- Really good for time sensitive situations where accuracy not as imperative 




